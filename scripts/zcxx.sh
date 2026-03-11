#!/bin/bash
export ZIG_LOCAL_CACHE_DIR="/tmp/zig-local-cache"
export ZIG_GLOBAL_CACHE_DIR="/tmp/zig-global-cache"
exec zig c++ "$@"
