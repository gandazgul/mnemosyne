#!/bin/bash
set -e

mkdir -p lib_linux_amd64 lib_linux_arm64 lib_darwin_amd64 lib_darwin_arm64

ORT_VER="1.23.1"
TOK_VER="1.25.0"

# linux amd64
curl -fsSL "https://github.com/microsoft/onnxruntime/releases/download/v${ORT_VER}/onnxruntime-linux-x64-${ORT_VER}.tgz" | tar xz -C lib_linux_amd64 --strip-components=1 --include="*/lib/libonnxruntime*"
curl -fsSL "https://github.com/daulet/tokenizers/releases/download/v${TOK_VER}/libtokenizers.linux-x86_64.tar.gz" | tar xz -C lib_linux_amd64

# linux arm64
curl -fsSL "https://github.com/microsoft/onnxruntime/releases/download/v${ORT_VER}/onnxruntime-linux-aarch64-${ORT_VER}.tgz" | tar xz -C lib_linux_arm64 --strip-components=1 --include="*/lib/libonnxruntime*"
curl -fsSL "https://github.com/daulet/tokenizers/releases/download/v${TOK_VER}/libtokenizers.linux-aarch64.tar.gz" | tar xz -C lib_linux_arm64

# darwin amd64
curl -fsSL "https://github.com/microsoft/onnxruntime/releases/download/v${ORT_VER}/onnxruntime-osx-x86_64-${ORT_VER}.tgz" | tar xz -C lib_darwin_amd64 --strip-components=1 --include="*/lib/libonnxruntime*"
curl -fsSL "https://github.com/daulet/tokenizers/releases/download/v${TOK_VER}/libtokenizers.darwin-x86_64.tar.gz" | tar xz -C lib_darwin_amd64

# darwin arm64
curl -fsSL "https://github.com/microsoft/onnxruntime/releases/download/v${ORT_VER}/onnxruntime-osx-arm64-${ORT_VER}.tgz" | tar xz -C lib_darwin_arm64 --strip-components=1 --include="*/lib/libonnxruntime*"
curl -fsSL "https://github.com/daulet/tokenizers/releases/download/v${TOK_VER}/libtokenizers.darwin-aarch64.tar.gz" | tar xz -C lib_darwin_arm64
