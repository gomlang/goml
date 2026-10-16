set -euo pipefail

test "$#" = 2
source "$(dirname "${BASH_SOURCE[0]}")/common.sh"
source "$(dirname "${BASH_SOURCE[0]}")/../tools/lib/common.sh"
stage0_platform="$(bootstrap_platform)"

stage0_file="$1"
stage0_output="$2"
source "$stage0_file"
bootstrap_stage0 "$stage0_platform"

stage0_cache="_bootstrap/cache"
stage0_archive="${GOML_STAGE0_ARCHIVE:-$stage0_cache/$GOML_STAGE0_ASSET_NAME}"
stage0_stamp="$stage0_output/.stage0.sha256"
stage0_tree_fingerprint() {
    toolchain_hash_files "$stage0_output/bin" "$stage0_output/lib" \
        | toolchain_sha256 \
        | cut -d ' ' -f 1
}
if test -f "$stage0_stamp" \
    && test "$(sed -n '1p' "$stage0_stamp")" = "$GOML_STAGE0_SHA256" \
    && test "$(sed -n '3p' "$stage0_stamp")" = managed-bin-lib-v1 \
    && test -x "$stage0_output/bin/goml" \
    && test -x "$stage0_output/bin/gomlc" \
    && test -x "$stage0_output/bin/gomllsp" \
    && test -f "$stage0_output/lib/compiler/compiler-world-v2.gaf" \
    && test "$(sed -n '2p' "$stage0_stamp")" = "$(stage0_tree_fingerprint)"; then
    exit 0
fi
mkdir -p "$stage0_cache" "$stage0_output/bin" "$stage0_output/lib"

if test -z "${GOML_STAGE0_ARCHIVE:-}"; then
    if ! test -f "$stage0_archive" || ! printf '%s  %s\n' "$GOML_STAGE0_SHA256" "$stage0_archive" | toolchain_sha256 -c - >/dev/null; then
        stage0_temporary="$(mktemp "$stage0_cache/$GOML_STAGE0_ASSET_NAME.XXXXXX")"
        trap 'rm -f "$stage0_temporary"' EXIT
        curl --fail --location --retry 3 --output "$stage0_temporary" "$GOML_STAGE0_URL"
        printf '%s  %s\n' "$GOML_STAGE0_SHA256" "$stage0_temporary" | toolchain_sha256 -c - >/dev/null
        mv "$stage0_temporary" "$stage0_archive"
        trap - EXIT
    fi
fi

printf '%s  %s\n' "$GOML_STAGE0_SHA256" "$stage0_archive" | toolchain_sha256 -c - >/dev/null

stage0_extracted="$(mktemp -d "$stage0_output/.stage0-install.XXXXXX")"
cleanup_stage0_install() {
    stage0_status="$?"
    set +e
    stage0_restore_failed=0
    if test "$stage0_status" != 0; then
        for directory in bin lib; do
            if test -e "$stage0_extracted/previous-$directory" \
                || test -L "$stage0_extracted/previous-$directory"; then
                if ! rm -rf -- "$stage0_output/$directory" \
                    || ! mv "$stage0_extracted/previous-$directory" "$stage0_output/$directory"; then
                    stage0_restore_failed=1
                fi
            fi
        done
    fi
    if test "$stage0_restore_failed" = 0; then
        rm -rf -- "$stage0_extracted"
    else
        printf 'stage0 restoration failed; backup retained at %s\n' "$stage0_extracted" >&2
    fi
    exit "$stage0_status"
}
trap cleanup_stage0_install EXIT
trap 'exit 129' HUP
trap 'exit 130' INT
trap 'exit 143' TERM
tar -xzf "$stage0_archive" --strip-components=1 -C "$stage0_extracted"

for tool in goml gomlc gomlfmt gomldoc gomllsp; do
    test -x "$stage0_extracted/bin/$tool"
done
for resource in builtin/goml.toml builtin/contract.goml builtin/runtime.goml builtin/impls.goml \
    builtin/language.goml builtin/numeric.goml builtin/derive.goml \
    prelude/goml.toml prelude/prelude.goml std/goml.toml; do
    test -f "$stage0_extracted/lib/$resource"
done

stage0_version="${GOML_STAGE0_RELEASE_TAG#v}"
test "$("$stage0_extracted/bin/goml" version)" = "goml $stage0_version"
for directory in bin lib; do
    mv "$stage0_output/$directory" "$stage0_extracted/previous-$directory"
    mv "$stage0_extracted/$directory" "$stage0_output/$directory"
done
bash tools/lib/finalize-toolchain.sh \
    "$stage0_output" \
    "$stage0_output/bin/goml" \
    "$stage0_output/bin/gomlc"
test -f "$stage0_output/lib/compiler/compiler-world-v2.gaf"
test ! -e "$stage0_output/lib/compiler/compiler-world-v2.gaf.tmp"
stage0_fingerprint="$(stage0_tree_fingerprint)"
printf '%s\n%s\n%s\n' "$GOML_STAGE0_SHA256" "$stage0_fingerprint" managed-bin-lib-v1 > "$stage0_extracted/stamp"
mv "$stage0_extracted/stamp" "$stage0_stamp"
