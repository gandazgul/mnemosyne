#!/usr/bin/env sh
set -eu

REPO="${MNEMOTECA_REPO:-gandazgul/mnemoteca}"
BIN_NAME="mnemoteca"
LEGACY_BIN_NAME="mnemosyne"
INSTALL_DIR="${INSTALL_DIR:-"$HOME/.local/bin"}"
VERSION="${VERSION:-latest}"

err() {
  printf 'error: %s\n' "$*" >&2
  exit 1
}

warn() {
  printf 'warning: %s\n' "$*" >&2
}

need() {
  command -v "$1" >/dev/null 2>&1 || err "missing required command: $1"
}

download() {
  url="$1"
  dest="$2"

  if command -v curl >/dev/null 2>&1; then
    curl -fsSL "$url" -o "$dest"
  elif command -v wget >/dev/null 2>&1; then
    wget -qO "$dest" "$url"
  else
    err "missing required command: curl or wget"
  fi
}

fetch() {
  url="$1"

  if command -v curl >/dev/null 2>&1; then
    curl -fsSL "$url"
  elif command -v wget >/dev/null 2>&1; then
    wget -qO- "$url"
  else
    err "missing required command: curl or wget"
  fi
}

have_tty() {
  { : </dev/tty; } 2>/dev/null && { : >/dev/tty; } 2>/dev/null
}

say_tty() {
  if have_tty; then
    printf '%s\n' "$*" >/dev/tty
  else
    printf '%s\n' "$*"
  fi
}

ask_yes_no() {
  prompt="$1"
  if ! have_tty; then
    return 1
  fi
  printf '%s [y/N] ' "$prompt" >/dev/tty
  if ! IFS= read -r answer </dev/tty; then
    return 1
  fi
  case "$answer" in
    y|Y|yes|YES|Yes) return 0 ;;
    *) return 1 ;;
  esac
}

read_tty_line() {
  prompt="$1"
  if ! have_tty; then
    return 1
  fi
  printf '%s' "$prompt" >/dev/tty
  IFS= read -r answer </dev/tty || return 1
  printf '%s\n' "$answer"
}

command_path() {
  command -v "$1" 2>/dev/null || true
}

is_absolute_executable_file() {
  case "$1" in
    /*) [ -f "$1" ] && [ -x "$1" ] ;;
    *) return 1 ;;
  esac
}

is_mnemoteca_link() {
  path="$1"
  [ -L "$path" ] || return 1
  target="$(readlink "$path" 2>/dev/null || true)"
  case "$target" in
    *mnemoteca*) return 0 ;;
    *) return 1 ;;
  esac
}

legacy_supports_export() {
  help="$($1 export --help 2>/dev/null || true)"
  printf '%s\n' "$help" | grep -- '--all' >/dev/null 2>&1 || return 1
  printf '%s\n' "$help" | grep -- '--output' >/dev/null 2>&1 || return 1
  printf '%s\n' "$help" | grep -- '--yes' >/dev/null 2>&1 || return 1
}

qualify_legacy_export_source() {
  path="$1"
  [ -n "$path" ] || return 1
  is_absolute_executable_file "$path" || return 1
  is_mnemoteca_link "$path" && return 1
  first_line="$($path version 2>/dev/null | sed -n '1p' || true)"
  printf '%s\n' "$first_line" | grep '^mnemosyne [^ ][^ ]* ([^)]*)$' >/dev/null 2>&1 || return 1
  legacy_supports_export "$path" || return 1
}

canonical_legacy_paths_present() {
  [ -e "$HOME/.config/mnemosyne/config.yaml" ] && return 0
  [ -e "$HOME/.local/share/mnemosyne/mnemosyne.db" ] && return 0
  [ -e "$HOME/.local/share/mnemosyne/models" ] && return 0
  [ -e "$HOME/.local/share/mnemosyne/lib" ] && return 0
  return 1
}

json_field_string() {
  key="$1"
  sed -n 's/.*"'"$key"'"[[:space:]]*:[[:space:]]*"\([^"]*\)".*/\1/p' | head -n 1
}

json_field_int() {
  key="$1"
  sed -n 's/.*"'"$key"'"[[:space:]]*:[[:space:]]*\([0-9][0-9]*\).*/\1/p' | head -n 1
}

mnemoteca_stats_json() {
  "$INSTALL_DIR/$BIN_NAME" stats --format json 2>/dev/null
}

legacy_database_path() {
  stats="$($1 stats 2>/dev/null || true)"
  count="$(printf '%s\n' "$stats" | grep -c '^Database Path:[[:space:]]*' || true)"
  [ "$count" = "1" ] || return 1
  printf '%s\n' "$stats" | sed -n 's/^Database Path:[[:space:]]*//p' | head -n 1
}

count_export() {
  dir="$1"
  collections=0
  documents=0

  [ -d "$dir" ] || return 1
  entries="$(find "$dir" -mindepth 1 -maxdepth 1 -print)"
  [ -n "$entries" ] || return 1

  for entry in $entries; do
    case "$entry" in
      *.jsonl) ;;
      *) return 1 ;;
    esac
    [ -f "$entry" ] || return 1
    [ -s "$entry" ] || return 1
    header="$(sed -n '1p' "$entry")"
    printf '%s\n' "$header" | grep '"collection"' >/dev/null 2>&1 || return 1
    lines="$(wc -l <"$entry" | tr -d ' ')"
    [ "$lines" -ge 1 ] || return 1
    collections=$((collections + 1))
    documents=$((documents + lines - 1))
  done

  printf '%s %s\n' "$collections" "$documents"
}

same_file_or_path() {
  left="$1"
  right="$2"
  [ "$left" = "$right" ] && return 0
  if [ -e "$left" ] && [ -e "$right" ]; then
    [ "$left" -ef "$right" ] 2>/dev/null && return 0
  fi
  return 1
}

safe_remove_file() {
  path="$1"
  label="$2"
  [ -e "$path" ] || return 0
  if [ -L "$path" ]; then
    warn "not removing symlinked $label: $path"
    return 0
  fi
  [ -f "$path" ] || { warn "not removing non-file $label: $path"; return 0; }
  say_tty "Cleanup candidate: $path"
  if ask_yes_no "Remove this $label?"; then
    rm -f "$path"
    say_tty "Removed $path"
  fi
}

safe_remove_dir() {
  path="$1"
  label="$2"
  [ -e "$path" ] || return 0
  case "$path" in
    ""|/|"$HOME") warn "not removing unsafe $label path: $path"; return 0 ;;
    "$HOME/.local/share/mnemosyne/models"|"$HOME/.local/share/mnemosyne/lib") ;;
    *) warn "not removing non-canonical $label path: $path"; return 0 ;;
  esac
  if [ -L "$path" ]; then
    warn "not removing symlinked $label: $path"
    return 0
  fi
  [ -d "$path" ] || { warn "not removing non-directory $label: $path"; return 0; }
  say_tty "Cleanup candidate: $path"
  if ask_yes_no "Remove this $label?"; then
    rm -rf "$path"
    say_tty "Removed $path"
  fi
}

maybe_cleanup() {
  legacy_cmd="$1"
  legacy_db="$2"
  [ "$migration_verified" = "true" ] || return 0
  ask_yes_no "Review optional legacy Mnemosyne cleanup now?" || return 0

  if [ "$legacy_cmd" = "$INSTALL_DIR/$LEGACY_BIN_NAME" ]; then
    safe_remove_file "$legacy_cmd" "legacy Mnemosyne binary"
  else
    say_tty "Legacy command is not $INSTALL_DIR/$LEGACY_BIN_NAME. Remove it manually if needed: $legacy_cmd"
  fi

  safe_remove_file "$HOME/.config/mnemosyne/config.yaml" "legacy config file"
  case "$legacy_db" in
    "$HOME/.local/share/mnemosyne/mnemosyne.db") safe_remove_file "$legacy_db" "legacy database" ;;
    "") ;;
    *) say_tty "Legacy database is not canonical. Remove it manually if needed: $legacy_db" ;;
  esac
  safe_remove_dir "$HOME/.local/share/mnemosyne/models" "legacy models directory"
  safe_remove_dir "$HOME/.local/share/mnemosyne/lib" "legacy runtime library directory"
}

maybe_link() {
  [ "$migration_verified" = "true" ] || return 0
  ask_yes_no "Create optional mnemosyne compatibility symlink to mnemoteca?" || return 0

  link_path="$INSTALL_DIR/$LEGACY_BIN_NAME"
  target_path="$INSTALL_DIR/$BIN_NAME"
  if [ -e "$link_path" ]; then
    say_tty "Cannot create compatibility link because $link_path exists. Remove the identified legacy binary first, then rerun the installer."
    return 0
  fi
  ln -s "$target_path" "$link_path"
  say_tty "Created $link_path -> $target_path"
  resolved="$(command_path "$LEGACY_BIN_NAME")"
  if [ "$resolved" = "$link_path" ]; then
    say_tty "mnemosyne resolves to the compatibility link: $resolved"
  else
    warn "mnemosyne resolves to another PATH entry before the compatibility link: ${resolved:-not found}"
  fi
}

run_migration() {
  legacy_cmd="$1"
  migration_verified=false

  if ! have_tty; then
    say_tty "Installed Mnemoteca. No terminal is available for prompts, so migration, cleanup, and compatibility link creation were skipped. Rerun install.sh from a terminal to migrate legacy Mnemosyne data."
    return 0
  fi

  if [ -z "$legacy_cmd" ]; then
    if canonical_legacy_paths_present; then
      say_tty "Legacy Mnemosyne data was found, but no qualifying mnemosyne command was found. Install the final Mnemosyne release and rerun this installer, or export manually with mnemosyne export --all --yes --output DIR and import with mnemoteca import --dir DIR."
    fi
    return 0
  fi

  say_tty "Mnemoteca replaces Mnemosyne. A qualifying legacy Mnemosyne command was found at: $legacy_cmd"
  say_tty "This installer can export legacy JSONL data and import it into an empty Mnemoteca store. Stop all agent hosts before you continue."
  ask_yes_no "Have you stopped agent hosts and do you want to migrate now?" || { say_tty "Migration skipped. Legacy data was not changed."; return 0; }

  export_root="$HOME/.mnemoteca-migration-exports"
  mkdir -p "$export_root"
  chmod 700 "$export_root"
  timestamp="$(date -u '+%Y%m%dT%H%M%SZ')"
  export_dir="$export_root/mnemosyne-export-$timestamp-$$"
  [ ! -e "$export_dir" ] || { say_tty "Export directory already exists. Migration blocked: $export_dir"; return 0; }
  old_umask="$(umask)"
  umask 077
  mkdir "$export_dir"
  umask "$old_umask"

  say_tty "Exporting legacy Mnemosyne data to persistent directory: $export_dir"
  if ! "$legacy_cmd" export --all --yes --output "$export_dir"; then
    say_tty "Legacy export failed. The retained export directory may be incomplete: $export_dir"
    return 0
  fi

  counts="$(count_export "$export_dir" || true)"
  if [ -z "$counts" ]; then
    say_tty "Export verification failed. The export is retained at: $export_dir"
    return 0
  fi
  expected_collections="$(printf '%s\n' "$counts" | awk '{print $1}')"
  expected_documents="$(printf '%s\n' "$counts" | awk '{print $2}')"

  legacy_db="$(legacy_database_path "$legacy_cmd" || true)"
  if [ -z "$legacy_db" ]; then
    say_tty "Migration blocked because legacy stats did not provide exactly one Database Path label. Export retained at: $export_dir"
    return 0
  fi

  stats="$(mnemoteca_stats_json || true)"
  dest_db="$(printf '%s\n' "$stats" | json_field_string database_path)"
  dest_collections="$(printf '%s\n' "$stats" | json_field_int collection_count)"
  dest_documents="$(printf '%s\n' "$stats" | json_field_int document_count)"
  if [ -z "$dest_db" ] || [ -z "$dest_collections" ] || [ -z "$dest_documents" ]; then
    say_tty "Migration blocked because mnemoteca stats --format json failed. Export retained at: $export_dir"
    return 0
  fi
  if [ "$dest_collections" != "0" ] || [ "$dest_documents" != "0" ]; then
    say_tty "Migration blocked because the Mnemoteca destination is not empty: $dest_collections collections, $dest_documents documents. Export retained at: $export_dir"
    return 0
  fi
  if same_file_or_path "$legacy_db" "$dest_db"; then
    say_tty "Migration blocked because legacy and Mnemoteca database paths are the same filesystem object: $dest_db"
    return 0
  fi

  say_tty "Warning: stopped agents are required. The empty check and import are not transactionally locked together."
  if ! "$INSTALL_DIR/$BIN_NAME" import --dir "$export_dir"; then
    say_tty "Import failed or was interrupted. Legacy data and export remain intact. Do not retry until the partial Mnemoteca destination is removed or resolved: $dest_db"
    return 0
  fi

  stats_after="$(mnemoteca_stats_json || true)"
  actual_collections="$(printf '%s\n' "$stats_after" | json_field_int collection_count)"
  actual_documents="$(printf '%s\n' "$stats_after" | json_field_int document_count)"
  if [ "$actual_collections" != "$expected_collections" ] || [ "$actual_documents" != "$expected_documents" ]; then
    say_tty "Import count mismatch. Expected $expected_collections collections and $expected_documents documents, got ${actual_collections:-unknown} and ${actual_documents:-unknown}. Legacy data and export remain intact. Cleanup and link creation are blocked."
    return 0
  fi

  query="$(read_tty_line 'Enter a representative memory search query to verify migration: ' || true)"
  [ -n "$query" ] || { say_tty "Search verification declined. Cleanup and link creation are blocked. Export retained at: $export_dir"; return 0; }
  if ! "$INSTALL_DIR/$BIN_NAME" search --fts-only --no-rerank --limit 5 "$query"; then
    say_tty "Search verification failed. Cleanup and link creation are blocked. Export retained at: $export_dir"
    return 0
  fi
  ask_yes_no "Did the search return the expected memory?" || { say_tty "Search verification was not confirmed. Cleanup and link creation are blocked. Export retained at: $export_dir"; return 0; }

  migration_verified=true
  say_tty "Migration verified. Legacy store and export remain available. Export retained at: $export_dir"
  maybe_cleanup "$legacy_cmd" "$legacy_db"
  maybe_link
}

os="$(uname -s | tr '[:upper:]' '[:lower:]')"
case "$os" in
  darwin|linux) ;;
  *) err "unsupported OS: $os" ;;
esac

arch="$(uname -m)"
case "$arch" in
  x86_64|amd64) arch="amd64" ;;
  arm64|aarch64) arch="arm64" ;;
  *) err "unsupported architecture: $arch" ;;
esac

legacy_candidate="$(command_path "$LEGACY_BIN_NAME")"
legacy_export_source=""
if qualify_legacy_export_source "$legacy_candidate"; then
  legacy_export_source="$legacy_candidate"
elif [ -n "$legacy_candidate" ] && is_absolute_executable_file "$legacy_candidate"; then
  legacy_export_source=""
fi

need grep
need head
need find
need mktemp
need mv
need tar
need chmod
need mkdir
need sed
need awk
need wc
need date
need readlink

if [ "$VERSION" = "latest" ]; then
  release_url="https://api.github.com/repos/$REPO/releases/latest"
else
  release_url="https://api.github.com/repos/$REPO/releases/tags/$VERSION"
fi

asset_suffix="${os}_${arch}.tar.gz"
asset_url="$(
  fetch "$release_url" |
    grep -Eo 'https://[^" ]+' |
    grep '/releases/download/' |
    grep "mnemoteca_" |
    grep "$asset_suffix" |
    grep '\.tar\.gz$' |
    head -n 1
)"

if [ -z "$asset_url" ]; then
  err "no Mnemoteca release archive found for $os/$arch at $release_url"
fi

tmp_dir="$(mktemp -d)"
trap 'rm -rf "$tmp_dir"' EXIT INT TERM

archive="$tmp_dir/$BIN_NAME.tar.gz"
download "$asset_url" "$archive"
tar -xzf "$archive" -C "$tmp_dir"

bin_path="$(find "$tmp_dir" -type f -name "$BIN_NAME" | head -n 1)"
if [ -z "$bin_path" ]; then
  err "archive did not contain $BIN_NAME"
fi
if find "$tmp_dir" -type f -name "$LEGACY_BIN_NAME" | grep . >/dev/null 2>&1; then
  err "archive contained legacy $LEGACY_BIN_NAME executable"
fi

mkdir -p "$INSTALL_DIR"
chmod +x "$bin_path"
mv "$bin_path" "$INSTALL_DIR/$BIN_NAME"

printf 'Installed %s to %s\n' "$BIN_NAME" "$INSTALL_DIR/$BIN_NAME"

case ":$PATH:" in
  *":$INSTALL_DIR:"*) ;;
  *)
    printf '\n%s is not on your PATH yet.\n' "$INSTALL_DIR"
    printf 'Add this to your shell profile:\n'
    printf '  export PATH="$HOME/.local/bin:$PATH"\n'
    ;;
esac

if [ -n "$legacy_candidate" ] && [ -z "$legacy_export_source" ]; then
  if have_tty; then
    say_tty "A mnemosyne command was found but was not recognized as a final Mnemosyne export source: $legacy_candidate"
    say_tty "No migration, cleanup, or compatibility link action will run for this uncertain command."
  fi
fi

run_migration "$legacy_export_source"
