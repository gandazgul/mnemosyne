#!/bin/bash
set -e

mkdir -p tmp_ort include lib

# Download SQLite headers for cross-compilation
if [ ! -f "include/sqlite3.h" ] || [ ! -f "include/sqlite3ext.h" ]; then
  echo "Downloading SQLite headers..."
  curl -fsSL https://sqlite.org/2024/sqlite-amalgamation-3450300.zip -o sqlite.zip
  unzip -q -o sqlite.zip
  cp sqlite-amalgamation-3450300/sqlite3.h sqlite-amalgamation-3450300/sqlite3ext.h include/
  rm -rf sqlite-amalgamation-3450300 sqlite.zip
else
  echo "SQLite headers already exist in include/, skipping download."
fi

ORT_VER="1.24.2"
TOK_VER="1.26.0"

download_arch() {
  local os_arch=$1
  local ort_url=$2
  local tok_url=$3
  local dest="lib/${os_arch}"

  # Check if libraries already exist
  if [ -d "$dest" ] && ls "$dest"/libonnxruntime* >/dev/null 2>&1 && ls "$dest"/libtokenizers.a >/dev/null 2>&1; then
    echo "Libraries for $os_arch already exist in $dest, skipping download."
    return
  fi

  echo "Downloading ONNX Runtime for $os_arch..."
  curl -fsSL "$ort_url" | tar xz -C tmp_ort
  mkdir -p "$dest"
  cp -rP tmp_ort/*/lib/libonnxruntime* "$dest/"
  rm -rf tmp_ort/*

  echo "Downloading tokenizers for $os_arch..."
  curl -fsSL "$tok_url" | tar xz -C "$dest"
}

# linux amd64
download_arch "linux_amd64" \
  "https://github.com/microsoft/onnxruntime/releases/download/v${ORT_VER}/onnxruntime-linux-x64-${ORT_VER}.tgz" \
  "https://github.com/daulet/tokenizers/releases/download/v${TOK_VER}/libtokenizers.linux-x86_64.tar.gz"

# linux arm64
download_arch "linux_arm64" \
  "https://github.com/microsoft/onnxruntime/releases/download/v${ORT_VER}/onnxruntime-linux-aarch64-${ORT_VER}.tgz" \
  "https://github.com/daulet/tokenizers/releases/download/v${TOK_VER}/libtokenizers.linux-aarch64.tar.gz"

# darwin arm64
download_arch "darwin_arm64" \
  "https://github.com/microsoft/onnxruntime/releases/download/v${ORT_VER}/onnxruntime-osx-arm64-${ORT_VER}.tgz" \
  "https://github.com/daulet/tokenizers/releases/download/v${TOK_VER}/libtokenizers.darwin-aarch64.tar.gz"

rm -rf tmp_ort
