#!/usr/bin/env sh
set -eu

REPO="${MNEMOSYNE_REPO:-gandazgul/mnemosyne}"
BIN_NAME="mnemosyne"
INSTALL_DIR="${INSTALL_DIR:-"$HOME/.local/bin"}"
VERSION="${VERSION:-latest}"

err() {
  printf 'error: %s\n' "$*" >&2
  exit 1
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

need grep
need head
need find
need mktemp
need mv
need tar
need chmod
need mkdir

if [ "$VERSION" = "latest" ]; then
  release_url="https://api.github.com/repos/$REPO/releases/latest"
else
  release_url="https://api.github.com/repos/$REPO/releases/tags/$VERSION"
fi

asset_suffix="${os}_${arch}.tar.gz"
asset_url="$(
  fetch "$release_url" |
    grep -Eo 'https://[^"]+' |
    grep '/releases/download/' |
    grep "$asset_suffix" |
    grep '\.tar\.gz$' |
    head -n 1
)"

if [ -z "$asset_url" ]; then
  err "no release archive found for $os/$arch at $release_url"
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
