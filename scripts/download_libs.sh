#!/bin/bash
set -e

mkdir -p lib_linux_amd64 lib_linux_arm64 lib_darwin_amd64 lib_darwin_arm64 tmp_ort

ORT_VER="1.23.1"
TOK_VER="1.25.0"

extract_ort() {
  local url=$1
  local dest=$2
  echo "Downloading $url"
  curl -fsSL "$url" | tar xz -C tmp_ort
  cp -P tmp_ort/*/lib/libonnxruntime* "$dest/"
  rm -rf tmp_ort/*
}

# linux amd64
extract_ort "https://github.com/microsoft/onnxruntime/releases/download/v${ORT_VER}/onnxruntime-linux-x64-${ORT_VER}.tgz" lib_linux_amd64
curl -fsSL "https://github.com/daulet/tokenizers/releases/download/v${TOK_VER}/libtokenizers.linux-x86_64.tar.gz" | tar xz -C lib_linux_amd64

# linux arm64
extract_ort "https://github.com/microsoft/onnxruntime/releases/download/v${ORT_VER}/onnxruntime-linux-aarch64-${ORT_VER}.tgz" lib_linux_arm64
curl -fsSL "https://github.com/daulet/tokenizers/releases/download/v${TOK_VER}/libtokenizers.linux-aarch64.tar.gz" | tar xz -C lib_linux_arm64

# darwin amd64
extract_ort "https://github.com/microsoft/onnxruntime/releases/download/v${ORT_VER}/onnxruntime-osx-x86_64-${ORT_VER}.tgz" lib_darwin_amd64
curl -fsSL "https://github.com/daulet/tokenizers/releases/download/v${TOK_VER}/libtokenizers.darwin-x86_64.tar.gz" | tar xz -C lib_darwin_amd64

# darwin arm64
extract_ort "https://github.com/microsoft/onnxruntime/releases/download/v${ORT_VER}/onnxruntime-osx-arm64-${ORT_VER}.tgz" lib_darwin_arm64
curl -fsSL "https://github.com/daulet/tokenizers/releases/download/v${TOK_VER}/libtokenizers.darwin-aarch64.tar.gz" | tar xz -C lib_darwin_arm64

rm -rf tmp_ort
