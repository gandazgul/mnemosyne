#!/usr/bin/env python3
"""Run Mnemosyne on LongMemEval-style memory retrieval.

LongMemEval is not a single global corpus like BEIR. Each question has its own
haystack of conversation sessions, so this harness builds an isolated Mnemosyne
database per question, searches it, and scores whether the ranked results map
back to the answer session IDs.
"""

from __future__ import annotations

import argparse
import json
import math
import os
import shutil
import subprocess
import sys
import time
import urllib.request
from collections import OrderedDict, defaultdict
from datetime import datetime, timezone
from pathlib import Path
from typing import Any


DATASET_URL = "https://huggingface.co/datasets/xiaowu0162/longmemeval-cleaned/resolve/main/longmemeval_s_cleaned.json"
FORMAT_VERSION = 1
METRIC_KEYS = ["recall_any@5", "recall_any@10", "recall_all@5", "recall_all@10", "mrr@10", "ndcg@10"]


def parse_args() -> argparse.Namespace:
    repo_root = Path(__file__).resolve().parents[2]
    parser = argparse.ArgumentParser(
        description="Run Mnemosyne against the cleaned LongMemEval session-retrieval benchmark.",
    )
    parser.add_argument(
        "--binary",
        type=Path,
        default=repo_root / "mnemosyne",
        help="Path to the mnemosyne binary. Defaults to ./mnemosyne.",
    )
    parser.add_argument(
        "--data-file",
        type=Path,
        default=repo_root / "benchmarks" / "data" / "longmemeval" / "longmemeval_s_cleaned.json",
        help="Path to longmemeval_s_cleaned.json. Downloaded if missing unless --skip-download is set.",
    )
    parser.add_argument(
        "--work-dir",
        type=Path,
        default=repo_root / "benchmarks" / "work" / "longmemeval",
        help="Directory for per-question imports and SQLite DBs.",
    )
    parser.add_argument(
        "--results-dir",
        type=Path,
        default=repo_root / "benchmarks" / "results",
        help="Directory where JSON and Markdown result summaries are written.",
    )
    parser.add_argument(
        "--collection",
        default="bench_longmemeval",
        help="Mnemosyne collection name inside each per-question DB.",
    )
    parser.add_argument(
        "--config",
        type=Path,
        default=None,
        help="Path to a mnemosyne config.yaml. Passed as MNEMOSYNE_CONFIG to import/search commands.",
    )
    parser.add_argument(
        "--run-label",
        default=None,
        help="Label for this model/config run. Defaults to the config filename stem.",
    )
    parser.add_argument(
        "--doc-mode",
        choices=["session", "user-turn", "message"],
        default="session",
        help=(
            "How to turn each haystack session into Mnemosyne documents. "
            "session joins user turns like MemPalace raw; user-turn stores one user turn per doc; "
            "message stores every user/assistant message as a short doc."
        ),
    )
    parser.add_argument(
        "--limit",
        type=int,
        default=50,
        help="Search result limit. Defaults to 50 to match MemPalace raw retrieval depth.",
    )
    parser.add_argument(
        "--rerank-candidates",
        type=int,
        default=50,
        help="Candidates passed to Mnemosyne fusion/reranking. Defaults to 50.",
    )
    parser.add_argument(
        "--max-queries",
        type=int,
        default=None,
        help="Run only the first N questions. Useful for smoke tests.",
    )
    parser.add_argument(
        "--max-sessions",
        type=int,
        default=None,
        help="Import only the first N haystack sessions plus all answer sessions. Useful for smoke tests.",
    )
    parser.add_argument(
        "--reuse-db",
        action="store_true",
        help="Reuse existing per-question DBs instead of rebuilding/importing them.",
    )
    parser.add_argument(
        "--skip-download",
        action="store_true",
        help="Require --data-file to already exist.",
    )
    parser.add_argument(
        "--no-rerank",
        action="store_true",
        help="Disable cross-encoder reranking during search.",
    )
    parser.add_argument(
        "--fusion",
        choices=["rrf", "vector-bm25"],
        default=None,
        help="Search fusion strategy. Defaults to vector-bm25 for hybrid runs and rrf for source-only runs.",
    )
    parser.add_argument(
        "--bm25-weight",
        type=float,
        default=0.10,
        help="BM25 lexical weight for --fusion vector-bm25. Defaults to 0.10.",
    )
    parser.add_argument(
        "--vector-only",
        action="store_true",
        help="Use only vector search candidates.",
    )
    parser.add_argument(
        "--fts-only",
        action="store_true",
        help="Use only full-text search candidates.",
    )
    return parser.parse_args()


def sanitize_label(label: str) -> str:
    out = "".join(ch if ch.isalnum() or ch in "._-" else "-" for ch in label.strip()).strip("-")
    if not out:
        raise ValueError("label cannot be empty after sanitization")
    return out


def weight_label(value: float) -> str:
    return f"{value:g}".replace(".", "p")


def download_dataset(path: Path, skip_download: bool) -> None:
    if path.exists():
        return
    if skip_download:
        raise FileNotFoundError(f"{path} does not exist and --skip-download was provided")
    path.parent.mkdir(parents=True, exist_ok=True)
    print(f"Downloading LongMemEval cleaned dataset from {DATASET_URL}", flush=True)
    urllib.request.urlretrieve(DATASET_URL, path)  # noqa: S310 - benchmark downloads a public dataset URL.


def load_dataset(path: Path) -> list[dict[str, Any]]:
    with path.open("r", encoding="utf-8") as f:
        data = json.load(f)
    if not isinstance(data, list):
        raise ValueError("expected LongMemEval JSON root to be a list")
    return data


def subset_entry(entry: dict[str, Any], max_sessions: int | None) -> dict[str, Any]:
    if max_sessions is None or max_sessions >= len(entry["haystack_session_ids"]):
        return entry

    answer_ids = set(entry["answer_session_ids"])
    selected: list[int] = []
    seen: set[int] = set()

    for i, sess_id in enumerate(entry["haystack_session_ids"]):
        if sess_id in answer_ids:
            selected.append(i)
            seen.add(i)

    for i in range(len(entry["haystack_session_ids"])):
        if len(selected) >= max_sessions:
            break
        if i not in seen:
            selected.append(i)
            seen.add(i)

    selected.sort()
    copied = dict(entry)
    for key in ["haystack_dates", "haystack_session_ids", "haystack_sessions"]:
        copied[key] = [entry[key][i] for i in selected]
    return copied


def session_text(session: list[dict[str, str]]) -> str:
    return "\n".join(turn["content"] for turn in session if turn.get("role") == "user").strip()


def documents_for_entry(entry: dict[str, Any], doc_mode: str) -> list[dict[str, str]]:
    docs: list[dict[str, str]] = []
    for session, session_id, date in zip(
        entry["haystack_sessions"],
        entry["haystack_session_ids"],
        entry["haystack_dates"],
    ):
        if doc_mode == "session":
            content = session_text(session)
            if content:
                docs.append(
                    {
                        "content": content,
                        "corpus_id": session_id,
                        "session_id": session_id,
                        "date": date,
                        "role": "user-session",
                    }
                )
            continue

        turn_index = 0
        for turn in session:
            role = str(turn.get("role") or "")
            content = str(turn.get("content") or "").strip()
            if not content:
                continue
            if doc_mode == "user-turn" and role != "user":
                continue

            docs.append(
                {
                    "content": f"{role}: {content}" if doc_mode == "message" else content,
                    "corpus_id": f"{session_id}_turn_{turn_index}",
                    "session_id": session_id,
                    "date": date,
                    "role": role,
                }
            )
            turn_index += 1
    return docs


def write_import_file(path: Path, collection: str, entry: dict[str, Any], doc_mode: str) -> list[dict[str, str]]:
    docs = documents_for_entry(entry, doc_mode)
    path.parent.mkdir(parents=True, exist_ok=True)
    with path.open("w", encoding="utf-8") as f:
        header = {
            "version": FORMAT_VERSION,
            "exported_at": datetime.now(timezone.utc).replace(microsecond=0).isoformat().replace("+00:00", "Z"),
            "collection": collection,
            "doc_count": len(docs),
        }
        f.write(json.dumps(header, ensure_ascii=False) + "\n")
        for i, doc in enumerate(docs):
            metadata = {
                "longmemeval_doc_id": doc["corpus_id"],
                "longmemeval_session_id": doc["session_id"],
                "longmemeval_date": doc["date"],
                "longmemeval_role": doc["role"],
                "longmemeval_doc_mode": doc_mode,
                "longmemeval_index": i,
            }
            record = {
                "content": doc["content"],
                "metadata": json.dumps(metadata, ensure_ascii=False),
            }
            f.write(json.dumps(record, ensure_ascii=False) + "\n")
    return docs


def remove_db_files(db_path: Path) -> None:
    for candidate in [db_path, Path(f"{db_path}-wal"), Path(f"{db_path}-shm")]:
        if candidate.exists():
            candidate.unlink()


def mnemosyne_env(db_path: Path, config_path: Path | None) -> dict[str, str]:
    env = os.environ.copy()
    env["MNEMOSYNE_DB_PATH"] = str(db_path)
    if config_path is not None:
        env["MNEMOSYNE_CONFIG"] = str(config_path)
    return env


def run(cmd: list[str], env: dict[str, str], cwd: Path) -> subprocess.CompletedProcess[str]:
    proc = subprocess.run(
        cmd,
        cwd=cwd,
        env=env,
        text=True,
        stdout=subprocess.PIPE,
        stderr=subprocess.PIPE,
    )
    if proc.returncode != 0:
        raise RuntimeError(
            f"command failed with exit code {proc.returncode}: {' '.join(cmd)}\n"
            f"stdout:\n{proc.stdout}\n"
            f"stderr:\n{proc.stderr}"
        )
    return proc


def import_question_db(
    binary: Path,
    import_file: Path,
    collection: str,
    db_path: Path,
    config_path: Path | None,
    repo_root: Path,
) -> None:
    proc = run(
        [str(binary), "import", str(import_file), "--name", collection],
        mnemosyne_env(db_path, config_path),
        repo_root,
    )
    if proc.stdout.strip():
        print(proc.stdout.strip(), flush=True)


def query_mnemosyne(
    binary: Path,
    collection: str,
    query: str,
    limit: int,
    rerank_candidates: int,
    no_rerank: bool,
    fusion: str,
    bm25_weight: float,
    vector_only: bool,
    fts_only: bool,
    db_path: Path,
    config_path: Path | None,
    repo_root: Path,
) -> list[dict[str, Any]]:
    cmd = [
        str(binary),
        "search",
        "--name",
        collection,
        "--format",
        "json",
        "--limit",
        str(limit),
        "--rerank-candidates",
        str(rerank_candidates),
        "--fusion",
        fusion,
        "--bm25-weight",
        str(bm25_weight),
        "--no-threshold",
    ]
    if no_rerank:
        cmd.append("--no-rerank")
    if vector_only:
        cmd.append("--vector-only")
    if fts_only:
        cmd.append("--fts-only")
    cmd.append(query)

    proc = run(cmd, mnemosyne_env(db_path, config_path), repo_root)
    payload = json.loads(proc.stdout)
    return list(payload.get("results", []))


def ranked_session_ids(results: list[dict[str, Any]]) -> list[str]:
    ranked: list[str] = []
    seen: set[str] = set()
    for result in results:
        metadata = result.get("metadata") or {}
        session_id = metadata.get("longmemeval_session_id")
        if not session_id:
            continue
        session_id = str(session_id)
        if session_id in seen:
            continue
        seen.add(session_id)
        ranked.append(session_id)
    return ranked


def dcg(relevances: list[int]) -> float:
    return sum(rel / math.log2(rank + 1) for rank, rel in enumerate(relevances, start=1))


def metrics_for_query(ranked_ids: list[str], correct_ids: list[str]) -> dict[str, Any]:
    correct = set(correct_ids)
    out: dict[str, Any] = {
        "answer_session_ids": correct_ids,
        "top_10_session_ids": ranked_ids[:10],
        "first_relevant_rank": None,
    }
    for k in [5, 10]:
        top = ranked_ids[:k]
        hits = set(top) & correct
        out[f"recall_any@{k}"] = 1.0 if hits else 0.0
        out[f"recall_all@{k}"] = 1.0 if correct and correct.issubset(set(top)) else 0.0

    for rank, session_id in enumerate(ranked_ids[:10], start=1):
        if session_id in correct:
            out["first_relevant_rank"] = rank
            break
    out["mrr@10"] = 0.0 if out["first_relevant_rank"] is None else 1.0 / out["first_relevant_rank"]

    gains = [1 if sid in correct else 0 for sid in ranked_ids[:10]]
    ideal = [1] * min(len(correct), 10)
    ideal_dcg = dcg(ideal)
    out["ndcg@10"] = dcg(gains) / ideal_dcg if ideal_dcg > 0 else 0.0
    return out


def average_metrics(per_query: OrderedDict[str, dict[str, Any]]) -> dict[str, float]:
    if not per_query:
        return {}
    return {key: sum(float(row[key]) for row in per_query.values()) / len(per_query) for key in METRIC_KEYS}


def build_breakdown(per_query: OrderedDict[str, dict[str, Any]]) -> dict[str, Any]:
    buckets = {
        "rank_1": 0,
        "rank_2_3": 0,
        "rank_4_5": 0,
        "rank_6_10": 0,
        "missing@10": 0,
    }
    missing_at_5 = []
    missing_at_10 = []
    per_type: dict[str, dict[str, list[float]]] = defaultdict(lambda: defaultdict(list))

    for query_id, row in per_query.items():
        rank = row["first_relevant_rank"]
        if rank is None:
            buckets["missing@10"] += 1
        elif rank == 1:
            buckets["rank_1"] += 1
        elif rank <= 3:
            buckets["rank_2_3"] += 1
        elif rank <= 5:
            buckets["rank_4_5"] += 1
        else:
            buckets["rank_6_10"] += 1

        if row["recall_any@5"] < 1:
            missing_at_5.append(query_id)
        if row["recall_any@10"] < 1:
            missing_at_10.append(query_id)

        qtype = row["question_type"]
        for key in METRIC_KEYS:
            per_type[qtype][key].append(float(row[key]))

    per_type_summary = {
        qtype: {
            "count": len(next(iter(vals.values()))) if vals else 0,
            **{key: sum(values) / len(values) for key, values in vals.items() if values},
        }
        for qtype, vals in sorted(per_type.items())
    }

    return {
        "first_relevant_rank_buckets": buckets,
        "queries_missing_at_5": missing_at_5,
        "queries_missing_at_10": missing_at_10,
        "per_type": per_type_summary,
    }


def md_cell(value: Any, limit: int = 120) -> str:
    text = str(value).replace("\n", " ").replace("|", "\\|").strip()
    if len(text) > limit:
        text = text[: limit - 3].rstrip() + "..."
    return text


def write_results(payload: dict[str, Any], results_dir: Path, max_queries: int | None) -> tuple[Path, Path]:
    results_dir.mkdir(parents=True, exist_ok=True)
    config = payload["config"]
    mode = "no-rerank" if config["no_rerank"] else "rerank"
    source = "vector-only" if config["vector_only"] else "fts-only" if config["fts_only"] else config["fusion"]
    if source == "vector-bm25":
        source = f"vector-bm25-w{weight_label(config['bm25_weight'])}-c{config['rerank_candidates']}"
    label = f"-{config['run_label']}" if config.get("run_label") else ""
    smoke = f"-smoke-{max_queries}" if max_queries else ""
    suffix = f"longmemeval-{config['doc_mode']}{label}-{source}-{mode}{smoke}"
    stem = f"{datetime.now(timezone.utc).strftime('%Y%m%dT%H%M%SZ')}-{suffix}"
    json_path = results_dir / f"{stem}.json"
    md_path = results_dir / f"{stem}.md"

    with json_path.open("w", encoding="utf-8") as f:
        json.dump(payload, f, indent=2, ensure_ascii=False)
        f.write("\n")

    metrics = payload["metrics"]
    breakdown = payload["breakdown"]
    with md_path.open("w", encoding="utf-8") as f:
        f.write("# Mnemosyne LongMemEval Results\n\n")
        f.write(f"- Generated: {payload['generated_at']}\n")
        f.write(f"- Dataset: `longmemeval_s_cleaned`\n")
        f.write(f"- Questions: {payload['query_count']}\n")
        f.write(f"- Document mode: `{config['doc_mode']}`\n")
        f.write(f"- Search limit: {config['limit']}\n")
        f.write(f"- Rerank candidates: {config['rerank_candidates']}\n")
        f.write(f"- Fusion: `{config['fusion']}`\n")
        f.write(f"- BM25 weight: `{config['bm25_weight']}`\n")
        f.write(f"- Rerank enabled: {not config['no_rerank']}\n")
        if config.get("max_sessions"):
            f.write(f"- Max sessions per question: {config['max_sessions']} plus answer sessions\n")
        if config.get("run_label"):
            f.write(f"- Run label: `{config['run_label']}`\n")
        if config.get("mnemosyne_config"):
            f.write(f"- Mnemosyne config: `{config['mnemosyne_config']}`\n")
        f.write("\n")
        f.write("| Metric | Score |\n")
        f.write("| --- | ---: |\n")
        for key in METRIC_KEYS:
            f.write(f"| `{key}` | {metrics[key]:.4f} |\n")

        f.write("\n## First Relevant Rank\n\n")
        f.write("| Bucket | Queries |\n")
        f.write("| --- | ---: |\n")
        for key, value in breakdown["first_relevant_rank_buckets"].items():
            f.write(f"| `{key}` | {value} |\n")

        f.write("\n## Per Type\n\n")
        f.write("| Question Type | Count | R@5 | R@10 | nDCG@10 |\n")
        f.write("| --- | ---: | ---: | ---: | ---: |\n")
        for qtype, row in breakdown["per_type"].items():
            f.write(
                f"| `{qtype}` | {row['count']} | {row['recall_any@5']:.4f} | "
                f"{row['recall_any@10']:.4f} | {row['ndcg@10']:.4f} |\n"
            )

        if breakdown["queries_missing_at_5"]:
            f.write("\n## Missing At 5\n\n")
            f.write("| Query ID | Type | Question | Answer Sessions | Top 10 Sessions |\n")
            f.write("| --- | --- | --- | --- | --- |\n")
            for query_id in breakdown["queries_missing_at_5"][:30]:
                row = payload["per_query"][query_id]
                f.write(
                    f"| {md_cell(query_id)} | {md_cell(row['question_type'])} | "
                    f"{md_cell(row['question'])} | {md_cell(', '.join(row['answer_session_ids']))} | "
                    f"{md_cell(', '.join(row['top_10_session_ids']))} |\n"
                )

    return json_path, md_path


def main() -> int:
    args = parse_args()
    repo_root = Path(__file__).resolve().parents[2]
    binary = args.binary.resolve()
    config_path = args.config.resolve() if args.config else None
    run_label = sanitize_label(args.run_label or config_path.stem) if (args.run_label or config_path) else None

    if args.fts_only and args.vector_only:
        print("cannot use both --fts-only and --vector-only", file=sys.stderr)
        return 2
    fusion = args.fusion
    if fusion is None:
        fusion = "rrf" if (args.vector_only or args.fts_only) else "vector-bm25"

    if fusion == "vector-bm25" and args.fts_only:
        print("--fusion vector-bm25 cannot be used with --fts-only", file=sys.stderr)
        return 2
    if args.bm25_weight < 0 or args.bm25_weight > 1:
        print("--bm25-weight must be between 0 and 1", file=sys.stderr)
        return 2
    if not binary.exists():
        print(f"mnemosyne binary not found at {binary}; run `task build` first", file=sys.stderr)
        return 2
    if config_path is not None and not config_path.exists():
        print(f"mnemosyne config not found at {config_path}", file=sys.stderr)
        return 2

    data_file = args.data_file.resolve()
    download_dataset(data_file, args.skip_download)
    data = load_dataset(data_file)
    if args.max_queries is not None:
        data = data[: args.max_queries]
    if not data:
        raise ValueError("No LongMemEval questions selected")

    work_dir = args.work_dir.resolve() / args.doc_mode
    if run_label:
        work_dir = work_dir / run_label
    work_dir.mkdir(parents=True, exist_ok=True)

    started = time.monotonic()
    per_query: OrderedDict[str, dict[str, Any]] = OrderedDict()
    total_docs = 0

    for index, original_entry in enumerate(data, start=1):
        entry = subset_entry(original_entry, args.max_sessions)
        query_id = str(entry.get("question_id") or f"query-{index}")
        question_dir = work_dir / query_id
        db_path = question_dir / "mnemosyne.db"
        import_file = question_dir / "import.jsonl"

        if not args.reuse_db:
            if question_dir.exists():
                shutil.rmtree(question_dir)
            question_dir.mkdir(parents=True, exist_ok=True)
            docs = write_import_file(import_file, args.collection, entry, args.doc_mode)
            remove_db_files(db_path)
            import_question_db(binary, import_file, args.collection, db_path, config_path, repo_root)
        else:
            if not db_path.exists():
                raise FileNotFoundError(f"--reuse-db requested but {db_path} does not exist")
            docs = documents_for_entry(entry, args.doc_mode)

        total_docs += len(docs)
        results = query_mnemosyne(
            binary=binary,
            collection=args.collection,
            query=str(entry["question"]),
            limit=args.limit,
            rerank_candidates=args.rerank_candidates,
            no_rerank=args.no_rerank,
            fusion=fusion,
            bm25_weight=args.bm25_weight,
            vector_only=args.vector_only,
            fts_only=args.fts_only,
            db_path=db_path,
            config_path=config_path,
            repo_root=repo_root,
        )
        ranked_ids = ranked_session_ids(results)
        row = metrics_for_query(ranked_ids, [str(sid) for sid in entry["answer_session_ids"]])
        row.update(
            {
                "question": entry["question"],
                "question_type": entry.get("question_type", ""),
                "question_date": entry.get("question_date", ""),
                "doc_count": len(docs),
            }
        )
        per_query[query_id] = row

        if index == 1 or index % 10 == 0 or index == len(data):
            print(
                f"Evaluated {index}/{len(data)} questions "
                f"R@5={row['recall_any@5']:.0f} R@10={row['recall_any@10']:.0f}",
                flush=True,
            )

    elapsed = time.monotonic() - started
    metrics = average_metrics(per_query)
    breakdown = build_breakdown(per_query)
    payload = {
        "generated_at": datetime.now(timezone.utc).replace(microsecond=0).isoformat().replace("+00:00", "Z"),
        "dataset": "longmemeval_s_cleaned",
        "source_url": DATASET_URL,
        "query_count": len(per_query),
        "total_imported_docs": total_docs,
        "elapsed_seconds": round(elapsed, 3),
        "config": {
            "binary": str(binary),
            "work_dir": str(work_dir),
            "mnemosyne_config": str(config_path) if config_path else None,
            "run_label": run_label,
            "doc_mode": args.doc_mode,
            "collection": args.collection,
            "limit": args.limit,
            "rerank_candidates": args.rerank_candidates,
            "no_rerank": args.no_rerank,
            "fusion": fusion,
            "bm25_weight": args.bm25_weight,
            "vector_only": args.vector_only,
            "fts_only": args.fts_only,
            "max_queries": args.max_queries,
            "max_sessions": args.max_sessions,
        },
        "metrics": metrics,
        "breakdown": breakdown,
        "per_query": per_query,
    }

    json_path, md_path = write_results(payload, args.results_dir.resolve(), args.max_queries)

    print("\nResults", flush=True)
    for key in METRIC_KEYS:
        print(f"  {key}: {metrics[key]:.4f}", flush=True)
    print(f"\nWrote {json_path}", flush=True)
    print(f"Wrote {md_path}", flush=True)
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
