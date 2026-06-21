#!/usr/bin/env python3
"""Run Mnemosyne on a BEIR-format retrieval benchmark.

The default target is SciFact, a small BEIR dataset useful for quick,
reproducible retrieval evaluation. The harness intentionally drives the
compiled mnemosyne CLI and consumes `search -f json` output so it exercises the
same surface area that external benchmark scripts will use.
"""

from __future__ import annotations

import argparse
import json
import math
import os
import re
import subprocess
import sys
import time
import urllib.request
import zipfile
from collections import OrderedDict
from datetime import datetime, timezone
from pathlib import Path
from typing import Any


BEIR_DATASET_URL = "https://public.ukp.informatik.tu-darmstadt.de/thakur/BEIR/datasets/{dataset}.zip"
DEFAULT_DATASET = "scifact"
FORMAT_VERSION = 1
METRIC_KEYS = ["ndcg@10", "mrr@10", "recall@10", "recall@100", "map@100"]


def parse_args() -> argparse.Namespace:
    repo_root = Path(__file__).resolve().parents[2]
    parser = argparse.ArgumentParser(
        description="Run Mnemosyne against a BEIR-format retrieval dataset.",
    )
    parser.add_argument(
        "--dataset",
        default=DEFAULT_DATASET,
        help="BEIR dataset name to download and run. Defaults to scifact.",
    )
    parser.add_argument(
        "--binary",
        type=Path,
        default=repo_root / "mnemosyne",
        help="Path to the mnemosyne binary. Defaults to ./mnemosyne.",
    )
    parser.add_argument(
        "--data-dir",
        type=Path,
        default=repo_root / "benchmarks" / "data" / "beir",
        help="Directory for downloaded/extracted BEIR datasets.",
    )
    parser.add_argument(
        "--work-dir",
        type=Path,
        default=repo_root / "benchmarks" / "work" / "beir",
        help="Directory for generated imports and benchmark SQLite DBs.",
    )
    parser.add_argument(
        "--results-dir",
        type=Path,
        default=repo_root / "benchmarks" / "results",
        help="Directory where JSON and Markdown result summaries are written.",
    )
    parser.add_argument(
        "--collection",
        default=None,
        help="Mnemosyne collection name. Defaults to bench_beir_<dataset>.",
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
        help="Label for this model/config run. Used in result filenames and DB work directory. Defaults to the config filename stem.",
    )
    parser.add_argument(
        "--limit",
        type=int,
        default=100,
        help="Search result limit. Defaults to 100 for Recall@100/MAP@100.",
    )
    parser.add_argument(
        "--rerank-candidates",
        type=int,
        default=100,
        help="Candidates passed to the reranker. Defaults to 100.",
    )
    parser.add_argument(
        "--max-queries",
        type=int,
        default=None,
        help="Run only the first N judged queries. Useful for smoke tests.",
    )
    parser.add_argument(
        "--max-docs",
        type=int,
        default=None,
        help="Import only a corpus subset. Relevant docs for selected queries are always kept. Useful for smoke tests.",
    )
    parser.add_argument(
        "--no-rerank",
        action="store_true",
        help="Disable cross-encoder reranking during search.",
    )
    parser.add_argument(
        "--reuse-db",
        action="store_true",
        help="Reuse an existing benchmark DB instead of rebuilding/importing.",
    )
    parser.add_argument(
        "--skip-download",
        action="store_true",
        help="Require the dataset to already exist under --data-dir.",
    )
    return parser.parse_args()


def load_jsonl(path: Path) -> OrderedDict[str, dict[str, Any]]:
    rows: OrderedDict[str, dict[str, Any]] = OrderedDict()
    with path.open("r", encoding="utf-8") as f:
        for line in f:
            if not line.strip():
                continue
            row = json.loads(line)
            rows[str(row["_id"])] = row
    return rows


def download_and_extract(dataset: str, data_dir: Path, skip_download: bool) -> Path:
    dataset_dir = data_dir / dataset
    if dataset_dir.exists():
        return dataset_dir

    if skip_download:
        raise FileNotFoundError(f"{dataset_dir} does not exist and --skip-download was provided")

    data_dir.mkdir(parents=True, exist_ok=True)
    zip_path = data_dir / f"{dataset}.zip"
    url = BEIR_DATASET_URL.format(dataset=dataset)

    if not zip_path.exists():
        print(f"Downloading {dataset} from {url}", flush=True)
        urllib.request.urlretrieve(url, zip_path)  # noqa: S310 - benchmark CLI downloads a public dataset URL.

    print(f"Extracting {zip_path}", flush=True)
    with zipfile.ZipFile(zip_path) as zf:
        zf.extractall(data_dir)

    if not dataset_dir.exists():
        raise FileNotFoundError(f"Expected extracted dataset at {dataset_dir}")
    return dataset_dir


def read_qrels(path: Path) -> OrderedDict[str, dict[str, int]]:
    qrels: OrderedDict[str, dict[str, int]] = OrderedDict()
    with path.open("r", encoding="utf-8") as f:
        for i, line in enumerate(f):
            parts = line.rstrip("\n").split("\t")
            if i == 0 and parts and parts[0] in {"query-id", "query_id", "qid"}:
                continue
            if len(parts) < 3:
                continue
            query_id, doc_id, score = parts[0], parts[1], int(parts[2])
            if score <= 0:
                continue
            qrels.setdefault(query_id, {})[doc_id] = score
    return qrels


def corpus_content(row: dict[str, Any]) -> str:
    title = str(row.get("title") or "").strip()
    text = str(row.get("text") or "").strip()
    if title and text:
        return f"{title}\n\n{text}"
    return title or text


def write_import_file(dataset: str, collection: str, corpus: OrderedDict[str, dict[str, Any]], path: Path) -> None:
    path.parent.mkdir(parents=True, exist_ok=True)
    with path.open("w", encoding="utf-8") as f:
        header = {
            "version": FORMAT_VERSION,
            "exported_at": datetime.now(timezone.utc).replace(microsecond=0).isoformat().replace("+00:00", "Z"),
            "collection": collection,
            "doc_count": len(corpus),
        }
        f.write(json.dumps(header, ensure_ascii=False) + "\n")

        for doc_id, row in corpus.items():
            metadata = {
                "beir_dataset": dataset,
                "beir_doc_id": doc_id,
            }
            title = str(row.get("title") or "").strip()
            if title:
                metadata["title"] = title

            record = {
                "content": corpus_content(row),
                "metadata": json.dumps(metadata, ensure_ascii=False),
            }
            f.write(json.dumps(record, ensure_ascii=False) + "\n")


def select_corpus_subset(
    corpus: OrderedDict[str, dict[str, Any]],
    qrels: OrderedDict[str, dict[str, int]],
    judged_queries: OrderedDict[str, dict[str, Any]],
    max_docs: int | None,
) -> OrderedDict[str, dict[str, Any]]:
    if max_docs is None or max_docs >= len(corpus):
        return corpus

    selected: OrderedDict[str, dict[str, Any]] = OrderedDict()

    for query_id in judged_queries:
        for doc_id in qrels[query_id]:
            if doc_id in corpus:
                selected[doc_id] = corpus[doc_id]

    for doc_id, row in corpus.items():
        if len(selected) >= max_docs:
            break
        selected.setdefault(doc_id, row)

    return selected


def remove_db_files(db_path: Path) -> None:
    for candidate in [db_path, Path(f"{db_path}-wal"), Path(f"{db_path}-shm")]:
        if candidate.exists():
            candidate.unlink()


def sanitize_label(label: str) -> str:
    sanitized = re.sub(r"[^A-Za-z0-9_.-]+", "-", label.strip()).strip("-")
    if not sanitized:
        raise ValueError("--run-label cannot be empty after sanitization")
    return sanitized


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
        joined = " ".join(cmd)
        raise RuntimeError(
            f"command failed with exit code {proc.returncode}: {joined}\n"
            f"stdout:\n{proc.stdout}\n"
            f"stderr:\n{proc.stderr}"
        )
    return proc


def import_corpus(
    binary: Path,
    import_file: Path,
    collection: str,
    db_path: Path,
    config_path: Path | None,
    repo_root: Path,
) -> None:
    env = mnemosyne_env(db_path, config_path)
    print(f"Importing corpus into {db_path}", flush=True)
    proc = run([str(binary), "import", str(import_file), "--name", collection], env, repo_root)
    if proc.stdout.strip():
        print(proc.stdout.strip(), flush=True)


def query_mnemosyne(
    binary: Path,
    collection: str,
    query: str,
    limit: int,
    rerank_candidates: int,
    no_rerank: bool,
    db_path: Path,
    config_path: Path | None,
    repo_root: Path,
) -> list[str]:
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
        "--no-threshold",
    ]
    if no_rerank:
        cmd.append("--no-rerank")
    cmd.append(query)

    env = mnemosyne_env(db_path, config_path)
    proc = run(cmd, env, repo_root)
    payload = json.loads(proc.stdout)

    doc_ids: list[str] = []
    for result in payload.get("results", []):
        metadata = result.get("metadata") or {}
        doc_id = metadata.get("beir_doc_id")
        if doc_id is None:
            raise ValueError(f"search result is missing metadata.beir_doc_id: {result}")
        doc_ids.append(str(doc_id))
    return doc_ids


def dcg(relevances: list[int]) -> float:
    return sum(((2**rel) - 1) / math.log2(rank + 1) for rank, rel in enumerate(relevances, start=1))


def metrics_for_query(ranked_doc_ids: list[str], relevant: dict[str, int]) -> dict[str, Any]:
    relevant_ids = set(relevant)
    rel_count = len(relevant_ids)
    if rel_count == 0:
        return {
            "ndcg@10": 0.0,
            "mrr@10": 0.0,
            "recall@10": 0.0,
            "recall@100": 0.0,
            "map@100": 0.0,
            "relevant_doc_ids": [],
            "found_relevant_ranks": {},
            "missing_relevant_doc_ids": [],
            "first_relevant_rank": None,
            "top_10_doc_ids": ranked_doc_ids[:10],
        }

    top10 = ranked_doc_ids[:10]
    top100 = ranked_doc_ids[:100]

    gains = [relevant.get(doc_id, 0) for doc_id in top10]
    ideal_gains = sorted(relevant.values(), reverse=True)[:10]
    ideal_dcg = dcg(ideal_gains)
    ndcg10 = dcg(gains) / ideal_dcg if ideal_dcg > 0 else 0.0

    mrr10 = 0.0
    for rank, doc_id in enumerate(top10, start=1):
        if doc_id in relevant_ids:
            mrr10 = 1.0 / rank
            break

    recall10 = len(set(top10) & relevant_ids) / rel_count
    recall100 = len(set(top100) & relevant_ids) / rel_count

    hits = 0
    precision_sum = 0.0
    for rank, doc_id in enumerate(top100, start=1):
        if doc_id in relevant_ids:
            hits += 1
            precision_sum += hits / rank
    map100 = precision_sum / rel_count
    found_relevant_ranks = {
        doc_id: rank
        for rank, doc_id in enumerate(ranked_doc_ids[:100], start=1)
        if doc_id in relevant_ids
    }
    first_relevant_rank = min(found_relevant_ranks.values()) if found_relevant_ranks else None
    missing_relevant_doc_ids = [doc_id for doc_id in relevant if doc_id not in found_relevant_ranks]

    return {
        "ndcg@10": ndcg10,
        "mrr@10": mrr10,
        "recall@10": recall10,
        "recall@100": recall100,
        "map@100": map100,
        "relevant_doc_ids": list(relevant.keys()),
        "found_relevant_ranks": found_relevant_ranks,
        "missing_relevant_doc_ids": missing_relevant_doc_ids,
        "first_relevant_rank": first_relevant_rank,
        "top_10_doc_ids": ranked_doc_ids[:10],
    }


def average_metrics(per_query: OrderedDict[str, dict[str, Any]]) -> dict[str, float]:
    if not per_query:
        return {}
    return {key: sum(float(row[key]) for row in per_query.values()) / len(per_query) for key in METRIC_KEYS}


def md_cell(value: Any, limit: int = 90) -> str:
    text = str(value).replace("\n", " ").replace("|", "\\|").strip()
    if len(text) > limit:
        text = text[: limit - 3].rstrip() + "..."
    return text


def first_relevant_rank_bucket(rank: int | None) -> str:
    if rank is None:
        return "missing@100"
    if rank == 1:
        return "rank_1"
    if rank <= 3:
        return "rank_2_3"
    if rank <= 10:
        return "rank_4_10"
    return "rank_11_100"


def build_breakdown(
    per_query: OrderedDict[str, dict[str, Any]],
    queries: OrderedDict[str, dict[str, Any]],
) -> dict[str, Any]:
    buckets = {
        "rank_1": 0,
        "rank_2_3": 0,
        "rank_4_10": 0,
        "rank_11_100": 0,
        "missing@100": 0,
    }

    for row in per_query.values():
        buckets[first_relevant_rank_bucket(row["first_relevant_rank"])] += 1

    missing_at_100 = [
        {
            "query_id": query_id,
            "query": queries[query_id]["text"],
            "relevant_doc_ids": row["relevant_doc_ids"],
            "missing_relevant_doc_ids": row["missing_relevant_doc_ids"],
            "top_10_doc_ids": row["top_10_doc_ids"],
        }
        for query_id, row in per_query.items()
        if row["recall@100"] < 1.0
    ]

    lowest_mrr = [
        {
            "query_id": query_id,
            "query": queries[query_id]["text"],
            "mrr@10": row["mrr@10"],
            "recall@100": row["recall@100"],
            "first_relevant_rank": row["first_relevant_rank"],
            "relevant_doc_ids": row["relevant_doc_ids"],
            "top_10_doc_ids": row["top_10_doc_ids"],
        }
        for query_id, row in sorted(per_query.items(), key=lambda item: (item[1]["mrr@10"], item[1]["recall@100"], item[0]))
    ]

    return {
        "summary": {
            "queries": len(per_query),
            "queries_with_rank_1_hit": buckets["rank_1"],
            "queries_with_top_10_hit": buckets["rank_1"] + buckets["rank_2_3"] + buckets["rank_4_10"],
            "queries_with_top_100_hit": len(per_query) - buckets["missing@100"],
            "queries_missing_at_100": buckets["missing@100"],
        },
        "first_relevant_rank_buckets": buckets,
        "missing_at_100": missing_at_100,
        "lowest_mrr@10": lowest_mrr[:20],
    }


def write_results(
    payload: dict[str, Any],
    results_dir: Path,
    dataset: str,
    max_queries: int | None,
) -> tuple[Path, Path]:
    results_dir.mkdir(parents=True, exist_ok=True)
    mode = "no-rerank" if payload["config"]["no_rerank"] else "rerank"
    label = payload["config"].get("run_label")
    label_part = f"-{label}" if label else ""
    suffix = f"{dataset}{label_part}-{mode}-smoke-{max_queries}" if max_queries else f"{dataset}{label_part}-{mode}"
    stem = f"{datetime.now(timezone.utc).strftime('%Y%m%dT%H%M%SZ')}-{suffix}"
    json_path = results_dir / f"{stem}.json"
    md_path = results_dir / f"{stem}.md"

    with json_path.open("w", encoding="utf-8") as f:
        json.dump(payload, f, indent=2, ensure_ascii=False)
        f.write("\n")

    metrics = payload["metrics"]
    breakdown = payload["breakdown"]
    with md_path.open("w", encoding="utf-8") as f:
        f.write(f"# Mnemosyne BEIR {dataset} Results\n\n")
        f.write(f"- Generated: {payload['generated_at']}\n")
        f.write(f"- Dataset: `{dataset}`\n")
        f.write(f"- Queries: {payload['query_count']}\n")
        f.write(f"- Corpus documents: {payload['corpus_doc_count']}\n")
        if payload["config"].get("max_docs"):
            f.write(f"- Corpus subset limit: {payload['config']['max_docs']}\n")
        f.write(f"- Search limit: {payload['config']['limit']}\n")
        f.write(f"- Rerank candidates: {payload['config']['rerank_candidates']}\n")
        f.write(f"- Rerank enabled: {not payload['config']['no_rerank']}\n")
        if payload["config"].get("run_label"):
            f.write(f"- Run label: `{payload['config']['run_label']}`\n")
        if payload["config"].get("mnemosyne_config"):
            f.write(f"- Mnemosyne config: `{payload['config']['mnemosyne_config']}`\n")
        f.write("\n")
        f.write("| Metric | Score |\n")
        f.write("| --- | ---: |\n")
        for key in ["ndcg@10", "mrr@10", "recall@10", "recall@100", "map@100"]:
            f.write(f"| `{key}` | {metrics[key]:.4f} |\n")

        f.write("\n## Breakdown\n\n")
        f.write("| Detail | Count |\n")
        f.write("| --- | ---: |\n")
        for key, value in breakdown["summary"].items():
            f.write(f"| `{key}` | {value} |\n")

        f.write("\n### First Relevant Rank\n\n")
        f.write("| Bucket | Queries |\n")
        f.write("| --- | ---: |\n")
        for key, value in breakdown["first_relevant_rank_buckets"].items():
            f.write(f"| `{key}` | {value} |\n")

        if breakdown["missing_at_100"]:
            f.write("\n### Missing At 100\n\n")
            f.write("| Query ID | Missing Doc IDs | Query | Top 10 Doc IDs |\n")
            f.write("| --- | --- | --- | --- |\n")
            for row in breakdown["missing_at_100"]:
                f.write(
                    "| "
                    + md_cell(row["query_id"])
                    + " | "
                    + md_cell(", ".join(row["missing_relevant_doc_ids"]))
                    + " | "
                    + md_cell(row["query"], limit=120)
                    + " | "
                    + md_cell(", ".join(row["top_10_doc_ids"]), limit=120)
                    + " |\n"
                )

        f.write("\n### Lowest MRR@10\n\n")
        f.write("| Query ID | MRR@10 | Recall@100 | First Relevant Rank | Query |\n")
        f.write("| --- | ---: | ---: | ---: | --- |\n")
        for row in breakdown["lowest_mrr@10"][:10]:
            first_rank = row["first_relevant_rank"] if row["first_relevant_rank"] is not None else "missing"
            f.write(
                "| "
                + md_cell(row["query_id"])
                + f" | {row['mrr@10']:.4f}"
                + f" | {row['recall@100']:.4f}"
                + " | "
                + md_cell(first_rank)
                + " | "
                + md_cell(row["query"], limit=120)
                + " |\n"
            )

    return json_path, md_path


def main() -> int:
    args = parse_args()
    repo_root = Path(__file__).resolve().parents[2]
    dataset = args.dataset.lower()
    collection = args.collection or f"bench_beir_{dataset}"
    binary = args.binary.resolve()
    config_path = args.config.resolve() if args.config else None
    run_label = sanitize_label(args.run_label or config_path.stem) if (args.run_label or config_path) else None

    if not binary.exists():
        print(f"mnemosyne binary not found at {binary}; run `task build` first", file=sys.stderr)
        return 2
    if config_path is not None and not config_path.exists():
        print(f"mnemosyne config not found at {config_path}", file=sys.stderr)
        return 2

    dataset_dir = download_and_extract(dataset, args.data_dir.resolve(), args.skip_download)
    corpus = load_jsonl(dataset_dir / "corpus.jsonl")
    queries = load_jsonl(dataset_dir / "queries.jsonl")
    qrels = read_qrels(dataset_dir / "qrels" / "test.tsv")

    judged_queries = OrderedDict((qid, queries[qid]) for qid in queries if qid in qrels)
    if args.max_queries is not None:
        judged_queries = OrderedDict(list(judged_queries.items())[: args.max_queries])
    if not judged_queries:
        raise ValueError("No judged queries selected")

    corpus_for_import = select_corpus_subset(corpus, qrels, judged_queries, args.max_docs)

    work_dir = args.work_dir.resolve() / dataset
    if run_label:
        work_dir = work_dir / run_label
    work_dir.mkdir(parents=True, exist_ok=True)
    db_path = work_dir / "mnemosyne.db"
    import_file = work_dir / "import.jsonl"

    if not args.reuse_db:
        remove_db_files(db_path)
        write_import_file(dataset, collection, corpus_for_import, import_file)
        import_corpus(binary, import_file, collection, db_path, config_path, repo_root)
    elif not db_path.exists():
        raise FileNotFoundError(f"--reuse-db requested but {db_path} does not exist")

    started = time.monotonic()
    per_query: OrderedDict[str, dict[str, Any]] = OrderedDict()
    for index, (query_id, row) in enumerate(judged_queries.items(), start=1):
        query = str(row["text"])
        ranked = query_mnemosyne(
            binary=binary,
            collection=collection,
            query=query,
            limit=args.limit,
            rerank_candidates=args.rerank_candidates,
            no_rerank=args.no_rerank,
            db_path=db_path,
            config_path=config_path,
            repo_root=repo_root,
        )
        per_query[query_id] = metrics_for_query(ranked, qrels[query_id])
        if index == 1 or index % 10 == 0 or index == len(judged_queries):
            print(f"Evaluated {index}/{len(judged_queries)} queries", flush=True)

    elapsed = time.monotonic() - started
    metrics = average_metrics(per_query)
    breakdown = build_breakdown(per_query, judged_queries)
    payload = {
        "generated_at": datetime.now(timezone.utc).replace(microsecond=0).isoformat().replace("+00:00", "Z"),
        "dataset": dataset,
        "source_url": BEIR_DATASET_URL.format(dataset=dataset),
        "collection": collection,
        "corpus_doc_count": len(corpus_for_import),
        "source_corpus_doc_count": len(corpus),
        "query_count": len(judged_queries),
        "elapsed_seconds": round(elapsed, 3),
        "config": {
            "binary": str(binary),
            "db_path": str(db_path),
            "mnemosyne_config": str(config_path) if config_path else None,
            "run_label": run_label,
            "limit": args.limit,
            "rerank_candidates": args.rerank_candidates,
            "no_rerank": args.no_rerank,
            "max_queries": args.max_queries,
            "max_docs": args.max_docs,
        },
        "metrics": metrics,
        "breakdown": breakdown,
        "per_query": per_query,
    }

    json_path, md_path = write_results(payload, args.results_dir.resolve(), dataset, args.max_queries)

    print("\nResults", flush=True)
    for key in ["ndcg@10", "mrr@10", "recall@10", "recall@100", "map@100"]:
        print(f"  {key}: {metrics[key]:.4f}", flush=True)
    print(f"\nWrote {json_path}", flush=True)
    print(f"Wrote {md_path}", flush=True)
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
