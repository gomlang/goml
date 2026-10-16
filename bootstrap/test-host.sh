set -euo pipefail

repository_root="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
temporary="$(mktemp -d)"
trap 'rm -rf "$temporary"' EXIT

mkdir -p "$temporary/bin" "$temporary/repo/bootstrap" "$temporary/repo/tools/lib"
cp "$repository_root/.justfile" "$temporary/repo/.justfile"
cp "$repository_root/bootstrap/bootstrap.sh" "$repository_root/bootstrap/check-host.sh" \
    "$repository_root/bootstrap/common.sh" "$repository_root/bootstrap/stage0.env" \
    "$temporary/repo/bootstrap/"
cp "$repository_root/tools/lib/common.sh" "$temporary/repo/tools/lib/"
cat > "$temporary/bin/uname" <<'UNAME'
#!/bin/sh
case "$1" in
    -s) printf '%s\n' "$GOML_TEST_HOST_OS" ;;
    -m) printf '%s\n' "$GOML_TEST_HOST_ARCH" ;;
    *) exit 1 ;;
esac
UNAME
chmod +x "$temporary/bin/uname"

cd "$temporary/repo"
for directory in stage0 stage1 stage2 stage3 \
    gomlc/_bootstrap/stage1 gomlc/_bootstrap/stage2 gomlc/_bootstrap/stage3 gomlc/_bootstrap/stage3-fixed \
    goml/_bootstrap/stage1 goml/_bootstrap/stage2 goml/_bootstrap/stage3 goml/_bootstrap/stage3-fixed; do
    mkdir -p "$directory"
    printf '%s\n' retained > "$directory/retained"
done

export PATH="$temporary/bin:$PATH"
for platform in Darwin/x86_64 Linux/aarch64 FreeBSD/x86_64; do
    export GOML_TEST_HOST_OS="${platform%/*}"
    export GOML_TEST_HOST_ARCH="${platform#*/}"
    for recipe in make bootstrap ci _bootstrap-stage3; do
        if just "$recipe" > "$temporary/stdout" 2> "$temporary/stderr"; then
            printf 'unexpected bootstrap success on %s: %s\n' "$platform" "$recipe" >&2
            exit 1
        fi
        grep -Fq "source bootstrap requires Linux amd64 or macOS arm64; detected $platform" "$temporary/stderr"
        test "$(find . -name retained -type f | wc -l | tr -d ' ')" = 12
        test ! -e _bootstrap
    done
done

export GOML_TEST_HOST_OS=Darwin GOML_TEST_HOST_ARCH=arm64
if just ci > "$temporary/stdout" 2> "$temporary/stderr"; then
    exit 1
fi
grep -Fq 'this command requires linux-amd64; detected darwin-arm64' "$temporary/stderr"
test "$(find . -name retained -type f | wc -l | tr -d ' ')" = 12

source bootstrap/common.sh
source bootstrap/stage0.env
for platform in Linux/x86_64 Darwin/arm64; do
    export GOML_TEST_HOST_OS="${platform%/*}"
    export GOML_TEST_HOST_ARCH="${platform#*/}"
    bash bootstrap/check-host.sh
    bootstrap_stage0 "$(bootstrap_platform)"
    if test "$platform" = Linux/x86_64; then
        expected_platform=linux-amd64
        test "$GOML_STAGE0_SHA256" = "$GOML_STAGE0_LINUX_AMD64_SHA256"
    else
        expected_platform=darwin-arm64
        test "$GOML_STAGE0_SHA256" = "$GOML_STAGE0_DARWIN_ARM64_SHA256"
    fi
    test "$GOML_STAGE0_ASSET_NAME" = "goml-${GOML_STAGE0_RELEASE_TAG#v}-$expected_platform.tar.gz"
    test "$GOML_STAGE0_URL" = "https://github.com/gomlang/goml/releases/download/$GOML_STAGE0_RELEASE_TAG/$GOML_STAGE0_ASSET_NAME"
    printf '%s\n' corrupt > "$temporary/corrupt.tar.gz"
    if GOML_STAGE0_ARCHIVE="$temporary/corrupt.tar.gz" \
        bash bootstrap/bootstrap.sh bootstrap/stage0.env stage0 \
        > "$temporary/stdout" 2> "$temporary/stderr"; then
        exit 1
    fi
    test ! -e stage0/.stage0.sha256
    test "$(find . -name retained -type f | wc -l | tr -d ' ')" = 12
done

unset GOML_STAGE0_DARWIN_ARM64_SHA256
if bootstrap_stage0 darwin-arm64 > "$temporary/stdout" 2> "$temporary/stderr"; then
    exit 1
fi
grep -Fq 'missing or invalid stage0 SHA-256 for darwin-arm64' "$temporary/stderr"
