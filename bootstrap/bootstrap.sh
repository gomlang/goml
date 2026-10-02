set -euo pipefail

test "$#" = 2
test "$(uname -s)" = Linux
test "$(uname -m)" = x86_64

stage0_file="$1"
stage0_output="$2"
source "$stage0_file"

stage0_cache="_bootstrap/cache"
stage0_archive="${GOML_STAGE0_ARCHIVE:-$stage0_cache/$GOML_STAGE0_ASSET_NAME}"
stage0_stamp="$stage0_output/.stage0.sha256"
stage0_tree_fingerprint() {
    find "$stage0_output/bin" "$stage0_output/lib" -type f -print0 \
        | sort -z \
        | xargs -0 -r sha256sum \
        | sha256sum \
        | cut -d ' ' -f 1
}
if test -f "$stage0_stamp" \
    && test "$(sed -n '1p' "$stage0_stamp")" = "$GOML_STAGE0_SHA256" \
    && test -x "$stage0_output/bin/goml" \
    && test -x "$stage0_output/bin/gomlc" \
    && test -x "$stage0_output/bin/gomllsp" \
    && test -f "$stage0_output/lib/compiler/compiler-world-v2.gaf" \
    && test "$(sed -n '2p' "$stage0_stamp")" = "$(stage0_tree_fingerprint)"; then
    exit 0
fi
mkdir -p "$stage0_cache" "$stage0_output/bin" "$stage0_output/lib"

if test -z "${GOML_STAGE0_ARCHIVE:-}"; then
    if ! test -f "$stage0_archive" || ! printf '%s  %s\n' "$GOML_STAGE0_SHA256" "$stage0_archive" | sha256sum --check --status; then
        stage0_temporary="$(mktemp "$stage0_cache/$GOML_STAGE0_ASSET_NAME.XXXXXX")"
        trap 'rm -f "$stage0_temporary"' EXIT
        curl --fail --location --retry 3 --output "$stage0_temporary" "$GOML_STAGE0_URL"
        printf '%s  %s\n' "$GOML_STAGE0_SHA256" "$stage0_temporary" | sha256sum --check --status
        mv "$stage0_temporary" "$stage0_archive"
        trap - EXIT
    fi
fi

printf '%s  %s\n' "$GOML_STAGE0_SHA256" "$stage0_archive" | sha256sum --check --status

stage0_extracted="$(mktemp -d)"
trap 'rm -rf "$stage0_extracted"' EXIT
tar -xzf "$stage0_archive" --strip-components=1 -C "$stage0_extracted"

rm -rf -- "$stage0_output/lib/builtin" "$stage0_output/lib/prelude" \
    "$stage0_output/lib/std" "$stage0_output/lib/cabi" "$stage0_output/lib/compiler"
cp -R "$stage0_extracted/bin/." "$stage0_output/bin/"
cp -R "$stage0_extracted/lib/." "$stage0_output/lib/"

trap - EXIT
rm -rf "$stage0_extracted"
test -x "$stage0_output/bin/goml"
test -x "$stage0_output/bin/gomlc"
test -x "$stage0_output/bin/gomllsp"
test -f "$stage0_output/lib/builtin/contract.goml"
test -f "$stage0_output/lib/builtin/numeric.goml"
test -f "$stage0_output/lib/builtin/derive.goml"
test -f "$stage0_output/lib/prelude/prelude.goml"
test -f "$stage0_output/lib/std/goml.toml"

stage0_version="${GOML_STAGE0_RELEASE_TAG#v}"
test "$("$stage0_output/bin/goml" version)" = "goml $stage0_version"
bash tools/lib/finalize-toolchain.sh \
    "$stage0_output" \
    "$stage0_output/bin/goml" \
    "$stage0_output/bin/gomlc"
test -f "$stage0_output/lib/compiler/compiler-world-v2.gaf"
test ! -e "$stage0_output/lib/compiler/compiler-world-v2.gaf.tmp"
printf '%s\n%s\n' "$GOML_STAGE0_SHA256" "$(stage0_tree_fingerprint)" > "$stage0_stamp"
