#!/usr/bin/env bash

set -euo pipefail

test "$#" = 1

repo_root="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
prefix="$1"

if [[ "$prefix/lib/std" -ef "$repo_root/lib/std" ]]; then
    echo 'cannot install toolchain resources over their source directory' >&2
    exit 1
fi

mkdir -p "$prefix/lib"
rm -f \
    "$prefix/lib/builtin_contract.gom" \
    "$prefix/lib/builtin_prelude.gom" \
    "$prefix/lib/builtin_numeric.gom" \
    "$prefix/lib/builtin_derive.gom"
rm -rf -- "$prefix/lib/std" "$prefix/lib/cabi"
cp -R "$repo_root/lib/." "$prefix/lib/"
