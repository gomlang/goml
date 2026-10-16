#!/usr/bin/env bash

set -euo pipefail

test "$#" = 3 || test "$#" = 4
source "$(dirname "${BASH_SOURCE[0]}")/common.sh"

prefix="$(realpath "$1")"
driver="$(realpath "$2")"
compiler="$(realpath "$3")"
compiler_dir="$prefix/lib/compiler"
world="$compiler_dir/${4:-compiler-world.gaf}"
fingerprint="$compiler_dir/finalize-input.sha256"
temporary="$fingerprint.tmp"

mkdir -p "$compiler_dir"
input_fingerprint="$({
    toolchain_sha256 "$driver" "$compiler"
    toolchain_hash_files "$prefix/lib/builtin" "$prefix/lib/prelude" "$prefix/lib/std"
} | toolchain_sha256 | cut -d ' ' -f 1)"

if test -f "$world" && test -f "$fingerprint"; then
    stored_input_fingerprint="$(sed -n '1p' "$fingerprint")"
    stored_world_fingerprint="$(sed -n '2p' "$fingerprint")"
    world_fingerprint="$(toolchain_sha256 "$world" | cut -d ' ' -f 1)"
    if test "$stored_input_fingerprint" = "$input_fingerprint" \
        && test "$stored_world_fingerprint" = "$world_fingerprint"; then
        exit 0
    fi
fi

"$driver" __toolchain-finalize --prefix "$prefix" --compiler "$compiler"
world_fingerprint="$(toolchain_sha256 "$world" | cut -d ' ' -f 1)"
printf '%s\n%s\n' "$input_fingerprint" "$world_fingerprint" > "$temporary"
mv "$temporary" "$fingerprint"
