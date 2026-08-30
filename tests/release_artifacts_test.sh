#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$ROOT_DIR"

fail() {
  echo "release_artifacts_test.sh: $*" >&2
  exit 1
}

assert_contains() { grep -F -- "$2" "$1" >/dev/null || fail "missing text '$2' in $1"; }
assert_not_contains() { ! grep -F -- "$2" "$1" >/dev/null || fail "unexpected text '$2' in $1"; }

assert_contains .goreleaser.yml "project_name: mnemoteca"
assert_contains .goreleaser.yml "binary: mnemoteca"
assert_contains .goreleaser.yml "github.com/gandazgul/mnemoteca/cmd.Version"
assert_not_contains .goreleaser.yml "github.com/gandazgul/mnemosyne"
assert_not_contains .goreleaser.yml "project_name: mnemosyne"

assert_contains .github/workflows/ci.yml "mnemoteca-linux-amd64"
assert_contains .github/workflows/ci.yml "mnemoteca-windows-amd64.exe"
assert_contains .github/workflows/release.yml "mnemoteca_\${VERSION}_linux_"
assert_contains .github/workflows/release.yml "mnemoteca_\${VERSION}_windows_amd64.zip"
assert_not_contains .github/workflows/ci.yml "mnemosyne-linux"
assert_not_contains .github/workflows/release.yml "github.com/gandazgul/mnemosyne"
assert_not_contains install.sh "MNEMOSYNE_REPO"

work="$(mktemp -d)"
trap 'rm -rf "$work"' EXIT
mkdir -p "$work/bin"

make_fake_binary() {
  local path="$1"
  local expected="$2"
  cat >"$path" <<SH
#!/usr/bin/env sh
case "\${1:-}" in
  version) printf '%s\\n' '$expected' ;;
  *) exit 64 ;;
esac
SH
  chmod +x "$path"
}

check_tar_archive() {
  local os="$1"
  local arch="$2"
  local expected="mnemoteca 1.2.3 ($os/$arch)"
  local asset="$work/mnemoteca_1.2.3_${os}_${arch}.tar.gz"
  rm -f "$work/bin/mnemoteca"
  make_fake_binary "$work/bin/mnemoteca" "$expected"
  tar -czf "$asset" -C "$work/bin" mnemoteca
  member_count="$(tar -tzf "$asset" | sed 's#^.*/##' | grep -Ec '^mnemoteca$')"
  [ "$member_count" = "1" ] || fail "$os/$arch archive mnemoteca member count = $member_count"
  if tar -tzf "$asset" | sed 's#^.*/##' | grep -Eq '^mnemosyne$'; then
    fail "$os/$arch archive contained mnemosyne member"
  fi
  extract="$work/extract-$os-$arch"
  mkdir -p "$extract"
  tar -xzf "$asset" -C "$extract"
  actual="$($extract/mnemoteca version | sed -n '1p')"
  [ "$actual" = "$expected" ] || fail "$os/$arch version output = $actual"
}

check_tar_archive darwin amd64
check_tar_archive darwin arm64
check_tar_archive linux amd64
check_tar_archive linux arm64

make_fake_binary "$work/bin/mnemoteca.exe" "mnemoteca 1.2.3 (windows/amd64)"
( cd "$work/bin" && zip -q "$work/mnemoteca_1.2.3_windows_amd64.zip" mnemoteca.exe )
zip_count="$(unzip -Z1 "$work/mnemoteca_1.2.3_windows_amd64.zip" | sed 's#^.*/##' | grep -Ec '^mnemoteca.exe$')"
[ "$zip_count" = "1" ] || fail "windows archive mnemoteca.exe member count = $zip_count"
if unzip -Z1 "$work/mnemoteca_1.2.3_windows_amd64.zip" | sed 's#^.*/##' | grep -Eq '^mnemosyne.exe$'; then
  fail "windows archive contained mnemosyne.exe member"
fi

windows_extract="$work/extract-windows-amd64"
mkdir -p "$windows_extract"
unzip -q "$work/mnemoteca_1.2.3_windows_amd64.zip" -d "$windows_extract"
actual="$($windows_extract/mnemoteca.exe version | sed -n '1p')"
[ "$actual" = "mnemoteca 1.2.3 (windows/amd64)" ] || fail "windows version output = $actual"

echo "release_artifacts_test.sh: release contracts passed"
