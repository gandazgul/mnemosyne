#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
INSTALL_SH="$ROOT_DIR/install.sh"
TEST_COUNT=0
CURRENT_WORK=""

on_error() {
  local status="$?"
  echo "install_test.sh: failed near line $1 with status $status" >&2
  if [ -n "$CURRENT_WORK" ] && [ -f "$CURRENT_WORK/output" ]; then
    echo "install_test.sh: captured installer output:" >&2
    sed -n '1,200p' "$CURRENT_WORK/output" >&2
  fi
  exit "$status"
}
trap 'on_error $LINENO' ERR

fail() {
  echo "install_test.sh: $*" >&2
  exit 1
}

assert_file() { [ -f "$1" ] || fail "missing file: $1"; }
assert_absent() { [ ! -e "$1" ] || fail "unexpected path exists: $1"; }
assert_contains() { grep -F -- "$2" "$1" >/dev/null || fail "missing text '$2' in $1"; }
assert_not_contains() { ! grep -F -- "$2" "$1" >/dev/null || fail "unexpected text '$2' in $1"; }

make_release() {
  local work="$1"
  local bin_mode="${2:-normal}"
  mkdir -p "$work/release/bin"
  cat >"$work/release/bin/mnemoteca" <<'SH'
#!/usr/bin/env sh
set -eu
state="${TEST_STATE:?TEST_STATE required}"
db="${TEST_DB_PATH:-$state/mnemoteca.db}"
cmd="${1:-}"
shift || true
case "$cmd" in
  version)
    printf 'mnemoteca 1.2.3 (linux/amd64)\n'
    ;;
  stats)
    collections=0
    documents=0
    if [ -f "$state/imported" ]; then
      collections="$(cat "$state/collections")"
      documents="$(cat "$state/documents")"
    fi
    printf '{"database_path":"%s","collection_count":%s,"document_count":%s}\n' "$db" "$collections" "$documents"
    ;;
  import)
    printf 'mnemoteca import %s\n' "$*" >>"$state/mnemoteca.log"
    if [ "${MNEMOTECA_IMPORT_FAIL:-0}" = "1" ]; then exit 42; fi
    if [ "${MNEMOTECA_IMPORT_MISMATCH:-0}" = "1" ]; then
      printf '99' >"$state/collections"
      printf '99' >"$state/documents"
    fi
    touch "$state/imported"
    ;;
  search)
    printf 'mnemoteca search %s\n' "$*" >>"$state/mnemoteca.log"
    printf 'expected migrated memory\n'
    ;;
  *)
    printf 'unexpected mnemoteca command: %s %s\n' "$cmd" "$*" >&2
    exit 64
    ;;
esac
SH
  chmod +x "$work/release/bin/mnemoteca"
  if [ "$bin_mode" = "legacy-member" ]; then
    printf '#!/usr/bin/env sh\n' >"$work/release/bin/mnemosyne"
    chmod +x "$work/release/bin/mnemosyne"
  fi
  tar -czf "$work/mnemoteca_1.2.3_linux_amd64.tar.gz" -C "$work/release/bin" .
}

make_fakes() {
  local work="$1"
  mkdir -p "$work/fakebin"
  cat >"$work/fakebin/uname" <<'SH'
#!/usr/bin/env sh
case "${1:-}" in
  -s) printf 'Linux\n' ;;
  -m) printf 'x86_64\n' ;;
  *) printf 'Linux\n' ;;
esac
SH
  chmod +x "$work/fakebin/uname"
  cat >"$work/fakebin/curl" <<'SH'
#!/usr/bin/env sh
set -eu
archive="${TEST_ARCHIVE:?TEST_ARCHIVE required}"
if [ "$#" -ge 4 ] && [ "$1" = "-fsSL" ] && [ "$3" = "-o" ]; then
  cp "$archive" "$4"
  exit 0
fi
printf '{"assets":[{"browser_download_url":"https://example.invalid/releases/download/v1.2.3/mnemoteca_1.2.3_linux_amd64.tar.gz"}]}'
SH
  chmod +x "$work/fakebin/curl"
}

make_legacy() {
  local work="$1"
  local mode="${2:-good}"
  cat >"$work/fakebin/mnemosyne" <<'SH'
#!/usr/bin/env sh
set -eu
state="${TEST_STATE:?TEST_STATE required}"
legacy_db="${TEST_LEGACY_DB:?TEST_LEGACY_DB required}"
mode="${TEST_LEGACY_MODE:-good}"
cmd="${1:-}"
shift || true
case "$cmd" in
  version)
    if [ "$mode" = "bad-version" ]; then printf 'mnemoteca 1.2.3 (linux/amd64)\n'; else printf 'mnemosyne 0.9.9 (linux/amd64)\n'; fi
    ;;
  export)
    if [ "${1:-}" = "--help" ]; then
      if [ "$mode" = "bad-help" ]; then printf 'export help\n'; else printf 'export --all --output --yes\n'; fi
      exit 0
    fi
    printf 'mnemosyne export %s\n' "$*" >>"$state/legacy.log"
    if [ "${MNEMOSYNE_EXPORT_FAIL:-0}" = "1" ]; then exit 41; fi
    out=""
    while [ "$#" -gt 0 ]; do
      case "$1" in
        --output) shift; out="$1" ;;
      esac
      shift || true
    done
    mkdir -p "$out"
    if [ "${MNEMOSYNE_MALFORMED_EXPORT:-0}" = "1" ]; then
      printf 'not a valid header\n' >"$out/bad.jsonl"
      exit 0
    fi
    printf '{"collection":"alpha"}\n{"text":"one"}\n{"text":"two"}\n' >"$out/alpha.jsonl"
    printf '{"collection":"empty"}\n' >"$out/empty.jsonl"
    printf '2' >"$state/collections"
    printf '2' >"$state/documents"
    ;;
  stats)
    if [ "${MNEMOSYNE_AMBIGUOUS_STATS:-0}" = "1" ]; then
      printf 'Database Path: %s\nDatabase Path: other\n' "$legacy_db"
    else
      printf 'Mnemosyne Statistics\nDatabase Path:   %s\n' "$legacy_db"
    fi
    ;;
  *)
    printf 'unexpected mnemosyne command: %s %s\n' "$cmd" "$*" >&2
    exit 64
    ;;
esac
SH
  chmod +x "$work/fakebin/mnemosyne"
  TEST_LEGACY_MODE="$mode"
  export TEST_LEGACY_MODE
}

run_nonterminal() {
  local work="$1"
  CURRENT_WORK="$work"
  shift
  (cd "$ROOT_DIR" && python3 - "$INSTALL_SH" "$work" "$@" <<'PY') >"$work/output" 2>&1
import os, subprocess, sys
install_sh, work, *env_pairs = sys.argv[1:]
env = os.environ.copy()
env.update({
    "PATH": f"{work}/fakebin:/usr/bin:/bin",
    "HOME": f"{work}/home",
    "INSTALL_DIR": f"{work}/install",
    "TEST_ARCHIVE": f"{work}/mnemoteca_1.2.3_linux_amd64.tar.gz",
    "TEST_STATE": f"{work}/state",
    "TEST_DB_PATH": f"{work}/home/.local/share/mnemoteca/mnemoteca.db",
    "TEST_LEGACY_DB": f"{work}/home/.local/share/mnemosyne/mnemosyne.db",
})
for pair in env_pairs:
    key, value = pair.split("=", 1)
    env[key] = value
with open(os.devnull, "rb") as stdin:
    proc = subprocess.run(["/bin/sh", install_sh], cwd=os.path.dirname(install_sh), env=env, stdin=stdin, text=True, capture_output=True, preexec_fn=os.setsid)
sys.stdout.write(proc.stdout)
sys.stderr.write(proc.stderr)
sys.exit(proc.returncode)
PY
}

run_pty() {
  local work="$1"
  CURRENT_WORK="$work"
  local answers="$2"
  shift 2
  (cd "$ROOT_DIR" && python3 - "$INSTALL_SH" "$work" "$answers" "$@" <<'PY') >"$work/output" 2>&1
import os, pty, sys, time
install_sh, work, answers, *env_pairs = sys.argv[1:]
env = os.environ.copy()
env.update({
    "PATH": f"{work}/fakebin:/usr/bin:/bin",
    "HOME": f"{work}/home",
    "INSTALL_DIR": f"{work}/install",
    "TEST_ARCHIVE": f"{work}/mnemoteca_1.2.3_linux_amd64.tar.gz",
    "TEST_STATE": f"{work}/state",
    "TEST_DB_PATH": f"{work}/home/.local/share/mnemoteca/mnemoteca.db",
    "TEST_LEGACY_DB": f"{work}/home/.local/share/mnemosyne/mnemosyne.db",
})
for pair in env_pairs:
    key, value = pair.split("=", 1)
    env[key] = value
pid, fd = pty.fork()
if pid == 0:
    os.chdir(os.path.dirname(install_sh))
    os.execve("/bin/sh", ["sh", install_sh], env)
for line in answers.split("\\n"):
    if line:
        time.sleep(0.15)
        try:
            os.write(fd, (line + "\n").encode())
        except OSError:
            break
chunks = []
status = 0
while True:
    try:
        data = os.read(fd, 4096)
        if not data:
            break
        chunks.append(data)
    except OSError:
        break
_, status = os.waitpid(pid, 0)
sys.stdout.buffer.write(b"".join(chunks))
if os.WIFEXITED(status):
    sys.exit(os.WEXITSTATUS(status))
sys.exit(1)
PY
}

setup_case() {
  local work
  work="$(mktemp -d)"
  mkdir -p "$work/home" "$work/install" "$work/state"
  make_release "$work" "${1:-normal}"
  make_fakes "$work"
  printf '%s\n' "$work"
}

fresh_nonterminal_install() {
  local work
  work="$(setup_case)"
  run_nonterminal "$work"
  assert_file "$work/install/mnemoteca"
  assert_contains "$work/output" "Installed mnemoteca"
  assert_contains "$work/output" "No terminal is available"
  assert_absent "$work/home/.mnemoteca-migration-exports"
  TEST_COUNT=$((TEST_COUNT + 1))
}

archive_without_mnemoteca_fails() {
  local work
  work="$(mktemp -d)"
  mkdir -p "$work/home" "$work/install" "$work/state" "$work/fakebin" "$work/empty"
  tar -czf "$work/mnemoteca_1.2.3_linux_amd64.tar.gz" -C "$work/empty" .
  make_fakes "$work"
  if run_nonterminal "$work"; then
    fail "installer accepted archive without mnemoteca"
  fi
  assert_contains "$work/output" "archive did not contain mnemoteca"
  TEST_COUNT=$((TEST_COUNT + 1))
}

archive_with_legacy_member_fails() {
  local work
  work="$(setup_case legacy-member)"
  if run_nonterminal "$work"; then
    fail "installer accepted archive with mnemosyne member"
  fi
  assert_contains "$work/output" "archive contained legacy mnemosyne executable"
  TEST_COUNT=$((TEST_COUNT + 1))
}

uncertain_legacy_does_not_migrate() {
  local work
  work="$(setup_case)"
  make_legacy "$work" bad-version
  mkdir -p "$work/home/.config/mnemosyne"
  touch "$work/home/.config/mnemosyne/config.yaml"
  run_pty "$work" ""
  assert_file "$work/install/mnemoteca"
  assert_contains "$work/output" "not recognized as a final Mnemosyne export source"
  [ ! -f "$work/state/legacy.log" ] || fail "uncertain legacy was exported"
  assert_not_contains "$work/output" "Remove this legacy"
  TEST_COUNT=$((TEST_COUNT + 1))
}

successful_migration_retains_export_and_links_after_consent() {
  local work
  work="$(setup_case)"
  make_legacy "$work" good
  mkdir -p "$work/home/.local/share/mnemosyne"
  touch "$work/home/.local/share/mnemosyne/mnemosyne.db"
  run_pty "$work" "y\nrepresentative memory\ny\nn\ny"
  assert_file "$work/install/mnemoteca"
  assert_file "$work/install/mnemosyne"
  [ -L "$work/install/mnemosyne" ] || fail "compatibility link was not a symlink"
  assert_contains "$work/state/legacy.log" "mnemosyne export --all --yes --output"
  assert_contains "$work/state/mnemoteca.log" "mnemoteca import --dir"
  assert_contains "$work/state/mnemoteca.log" "mnemoteca search --fts-only --no-rerank --limit 5 representative memory"
  exports="$(find "$work/home/.mnemoteca-migration-exports" -type d -name 'mnemosyne-export-*' | wc -l | tr -d ' ')"
  [ "$exports" = "1" ] || fail "expected one retained export, got $exports"
  TEST_COUNT=$((TEST_COUNT + 1))
}

malformed_export_blocks_import_and_link() {
  local work
  work="$(setup_case)"
  make_legacy "$work" good
  mkdir -p "$work/home/.local/share/mnemosyne"
  touch "$work/home/.local/share/mnemosyne/mnemosyne.db"
  run_pty "$work" "y" MNEMOSYNE_MALFORMED_EXPORT=1
  assert_contains "$work/output" "Export verification failed"
  [ ! -f "$work/state/mnemoteca.log" ] || fail "malformed export reached mnemoteca"
  assert_absent "$work/install/mnemosyne"
  TEST_COUNT=$((TEST_COUNT + 1))
}

same_file_destination_blocks_import() {
  local work
  work="$(setup_case)"
  make_legacy "$work" good
  mkdir -p "$work/home/.local/share/mnemosyne"
  touch "$work/home/.local/share/mnemosyne/mnemosyne.db"
  run_pty "$work" "y" TEST_DB_PATH="$work/home/.local/share/mnemosyne/mnemosyne.db"
  assert_contains "$work/output" "same filesystem object"
  [ ! -f "$work/state/mnemoteca.log" ] || fail "same-file destination reached mnemoteca"
  TEST_COUNT=$((TEST_COUNT + 1))
}

import_failure_blocks_cleanup_and_link() {
  local work
  work="$(setup_case)"
  make_legacy "$work" good
  mkdir -p "$work/home/.local/share/mnemosyne"
  touch "$work/home/.local/share/mnemosyne/mnemosyne.db"
  run_pty "$work" "y" MNEMOTECA_IMPORT_FAIL=1
  assert_contains "$work/output" "Import failed or was interrupted"
  assert_absent "$work/install/mnemosyne"
  TEST_COUNT=$((TEST_COUNT + 1))
}

harness_integrity_check() {
  local backup work
  backup="$(mktemp)"
  cp "$INSTALL_SH" "$backup"
  trap 'cp "$backup" "$INSTALL_SH"; rm -f "$backup"' RETURN
  cat >"$INSTALL_SH" <<'SH'
#!/usr/bin/env sh
echo known failing installer fixture >&2
exit 77
SH
  chmod +x "$INSTALL_SH"
  work="$(setup_case)"
  if run_nonterminal "$work"; then
    fail "harness did not fail when production installer failed"
  fi
  assert_contains "$work/output" "known failing installer fixture"
  cp "$backup" "$INSTALL_SH"
  rm -f "$backup"
  trap - RETURN
  TEST_COUNT=$((TEST_COUNT + 1))
}

fresh_nonterminal_install
archive_without_mnemoteca_fails
archive_with_legacy_member_fails
uncertain_legacy_does_not_migrate
successful_migration_retains_export_and_links_after_consent
malformed_export_blocks_import_and_link
same_file_destination_blocks_import
import_failure_blocks_cleanup_and_link
harness_integrity_check

echo "install_test.sh: $TEST_COUNT tests passed"
