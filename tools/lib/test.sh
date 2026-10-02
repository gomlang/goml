#!/usr/bin/env bash

set -euo pipefail

test "$#" = 1

repo_root="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
prefix="$(cd "$1" && pwd)"
temporary="$(mktemp -d)"
trap 'rm -rf "$temporary"' EXIT

test -f "$prefix/lib/builtin/contract.goml"
test -f "$prefix/lib/builtin/goml.toml"
test -f "$prefix/lib/builtin/runtime.goml"
test -f "$prefix/lib/builtin/impls.goml"
test -f "$prefix/lib/builtin/language.goml"
test -f "$prefix/lib/builtin/numeric.goml"
test -f "$prefix/lib/builtin/derive.goml"
test -f "$prefix/lib/prelude/prelude.goml"
test -f "$prefix/lib/prelude/goml.toml"
test -f "$prefix/lib/std/goml.toml"
test -f "$prefix/lib/cabi/go.mod"
test -f "$prefix/lib/cabi/runtime_linux_amd64.s"
test ! -e "$prefix/lib/builtin_contract.gom"
test ! -e "$prefix/lib/builtin_prelude.gom"
test ! -e "$prefix/lib/builtin_numeric.gom"
test ! -e "$prefix/lib/builtin_derive.gom"

mkdir -p "$temporary/install/lib/std/obsolete" "$temporary/install/lib/std/fs" \
    "$temporary/install/lib/compiler" "$temporary/install/lib/custom"
printf '%s\n' stale > "$temporary/install/lib/std/obsolete/old.goml"
printf '%s\n' stale > "$temporary/install/lib/std/fs/removed.goml"
printf '%s\n' retained > "$temporary/install/lib/compiler/retained"
printf '%s\n' retained > "$temporary/install/lib/custom/retained"
printf '%s\n' retained > "$temporary/install/lib/custom/retained.gom"
bash "$repo_root/tools/lib/install.sh" "$temporary/install"
test ! -e "$temporary/install/lib/std/obsolete"
test ! -e "$temporary/install/lib/std/fs/removed.goml"
test -f "$temporary/install/lib/std/fs/fs.goml"
test -f "$temporary/install/lib/compiler/retained"
test -f "$temporary/install/lib/custom/retained"
test -f "$temporary/install/lib/custom/retained.gom"
if bash "$repo_root/tools/lib/install.sh" "$repo_root" > "$temporary/install-stdout" 2> "$temporary/install-stderr"; then
    exit 1
fi
grep -F 'cannot install toolchain resources over their source directory' "$temporary/install-stderr" >/dev/null

mkdir -p "$temporary/toolchain/bin"
cp "$prefix/bin/goml" "$temporary/toolchain/bin/goml"
cp "$prefix/bin/gomlc" "$temporary/toolchain/bin/gomlc"
mkdir -p "$temporary/toolchain/lib"
cp -R "$prefix/lib/builtin" "$temporary/toolchain/lib/builtin"
cp -R "$prefix/lib/prelude" "$temporary/toolchain/lib/prelude"
cp -R "$prefix/lib/std" "$temporary/toolchain/lib/std"
test ! -e "$temporary/toolchain/lib/compiler/compiler-world-v2.gaf"
"$temporary/toolchain/bin/goml" __toolchain-finalize --prefix "$temporary/toolchain"
test -f "$temporary/toolchain/lib/compiler/compiler-world-v2.gaf"
test ! -e "$temporary/toolchain/lib/compiler/compiler-world-v2.gaf.tmp"
first_world_hash="$(sha256sum "$temporary/toolchain/lib/compiler/compiler-world-v2.gaf")"
"$temporary/toolchain/bin/goml" __toolchain-finalize --prefix "$temporary/toolchain"
second_world_hash="$(sha256sum "$temporary/toolchain/lib/compiler/compiler-world-v2.gaf")"
test "$first_world_hash" = "$second_world_hash"
rm -f "$temporary/toolchain/lib/compiler/finalize-input.sha256"
bash "$repo_root/tools/lib/finalize-toolchain.sh" \
    "$temporary/toolchain" \
    "$temporary/toolchain/bin/goml" \
    "$temporary/toolchain/bin/gomlc"
first_world_mtime="$(stat -c %y "$temporary/toolchain/lib/compiler/compiler-world-v2.gaf")"
bash "$repo_root/tools/lib/finalize-toolchain.sh" \
    "$temporary/toolchain" \
    "$temporary/toolchain/bin/goml" \
    "$temporary/toolchain/bin/gomlc"
second_world_mtime="$(stat -c %y "$temporary/toolchain/lib/compiler/compiler-world-v2.gaf")"
test "$first_world_mtime" = "$second_world_mtime"
expected_world_hash="$(sha256sum "$temporary/toolchain/lib/compiler/compiler-world-v2.gaf")"
printf '%s\n' invalid > "$temporary/toolchain/lib/compiler/compiler-world-v2.gaf"
bash "$repo_root/tools/lib/finalize-toolchain.sh" \
    "$temporary/toolchain" \
    "$temporary/toolchain/bin/goml" \
    "$temporary/toolchain/bin/gomlc"
repaired_world_hash="$(sha256sum "$temporary/toolchain/lib/compiler/compiler-world-v2.gaf")"
test "$expected_world_hash" = "$repaired_world_hash"
cp -R "$repo_root/tools/release/testdata/smoke" "$temporary/project"
(
    cd "$temporary/project"
    "$temporary/toolchain/bin/goml" check --dry-run > "$temporary/project-plan"
)
grep -F -- "--world $temporary/toolchain/lib/compiler/compiler-world-v2.gaf" "$temporary/project-plan" >/dev/null
"$temporary/toolchain/bin/gomlc" build \
    --package tests::toml \
    --input "$repo_root/gomlc/testdata/module/project055_toml/main.goml" \
    --output "$temporary/toolchain/smoke/main" \
    --world "$temporary/toolchain/lib/compiler/compiler-world-v2.gaf"
test -f "$temporary/toolchain/smoke/main.interface"
test -f "$temporary/toolchain/smoke/main.core"

mkdir -p "$temporary/bin"
cp "$prefix/bin/gomlc" "$temporary/bin/gomlc"

if "$temporary/bin/gomlc" __builtin-interface > "$temporary/stdout" 2> "$temporary/stderr"; then
    exit 1
fi

grep -F "could not read builtin resource $temporary/lib/builtin/contract.goml" "$temporary/stderr" >/dev/null

cd "$temporary"
cp -R "$prefix/lib" "$temporary/lib"
"$temporary/bin/gomlc" __builtin-interface >/dev/null
"$temporary/bin/gomlc" __prelude-interface >/dev/null

mv "$temporary/lib/prelude/prelude.goml" "$temporary/lib/prelude/prelude.goml.missing"
if "$temporary/bin/gomlc" __prelude-interface > "$temporary/prelude-stdout" 2> "$temporary/prelude-stderr"; then
    exit 1
fi
grep -F "could not read prelude resource $temporary/lib/prelude/prelude.goml" "$temporary/prelude-stderr" >/dev/null
mv "$temporary/lib/prelude/prelude.goml.missing" "$temporary/lib/prelude/prelude.goml"
mv "$temporary/lib/builtin/contract.goml" "$temporary/lib/builtin/contract.gom"
if "$temporary/bin/gomlc" __builtin-interface > "$temporary/legacy-stdout" 2> "$temporary/legacy-stderr"; then
    exit 1
fi
grep -F "could not read builtin resource $temporary/lib/builtin/contract.goml" "$temporary/legacy-stderr" >/dev/null
mv "$temporary/lib/builtin/contract.gom" "$temporary/lib/builtin/contract.goml"
mv "$temporary/lib/prelude/prelude.goml" "$temporary/lib/prelude/prelude.gom"
if "$temporary/bin/gomlc" __prelude-interface > "$temporary/legacy-stdout" 2> "$temporary/legacy-stderr"; then
    exit 1
fi
grep -F "could not read prelude resource $temporary/lib/prelude/prelude.goml" "$temporary/legacy-stderr" >/dev/null
