set -euo pipefail

repo_root="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$repo_root"
source bootstrap/common.sh
source tools/lib/common.sh
temporary="$(mktemp -d)"
trap 'rm -rf "$temporary"' EXIT
prefix="$temporary/stage0"

mkdir -p "$prefix/bin" "$prefix/lib/builtin" "$prefix/lib/prelude" "$prefix/lib/std/ascii" \
    "$prefix/lib/compiler" "$prefix/lib/obsolete" "$prefix/custom"
for file in builtin/obsolete.gom builtin/obsolete.goml prelude/obsolete.goml \
    std/ascii/obsolete.goml compiler/compiler-world.gaf; do
    printf '%s\n' invalid > "$prefix/lib/$file"
done
printf '%s\n' obsolete > "$prefix/bin/obsolete"
printf '%s\n' obsolete > "$prefix/lib/obsolete/old.goml"
printf '%s\n' retained > "$prefix/custom/retained"

bash bootstrap/bootstrap.sh bootstrap/stage0.env "$prefix"
for file in builtin/obsolete.gom builtin/obsolete.goml prelude/obsolete.goml \
    std/ascii/obsolete.goml; do
    test ! -e "$prefix/lib/$file"
done
test ! -e "$prefix/bin/obsolete"
test ! -e "$prefix/lib/obsolete"
test -f "$prefix/custom/retained"
test -f "$prefix/lib/compiler/compiler-world.gaf"
source bootstrap/stage0.env
bootstrap_stage0 "$(bootstrap_platform)"
printf '%s\n' obsolete > "$prefix/bin/obsolete"
mkdir -p "$prefix/lib/obsolete"
printf '%s\n' obsolete > "$prefix/lib/obsolete/old.goml"
legacy_tree="$(toolchain_hash_files "$prefix/bin" "$prefix/lib" | toolchain_sha256 | cut -d ' ' -f 1)"
printf '%s\n%s\n' "$GOML_STAGE0_SHA256" "$legacy_tree" > "$prefix/.stage0.sha256"
bash bootstrap/bootstrap.sh bootstrap/stage0.env "$prefix"
test ! -e "$prefix/bin/obsolete"
test ! -e "$prefix/lib/obsolete"
test -f "$prefix/custom/retained"
first_world_mtime="$(toolchain_file_stamp "$prefix/lib/compiler/compiler-world.gaf")"
bash bootstrap/bootstrap.sh bootstrap/stage0.env "$prefix"
second_world_mtime="$(toolchain_file_stamp "$prefix/lib/compiler/compiler-world.gaf")"
test "$first_world_mtime" = "$second_world_mtime"

package="$temporary/package"
mkdir -p "$package/bin" "$package/lib/builtin" "$package/lib/prelude" "$package/lib/std"
cat > "$package/bin/goml" <<'DRIVER'
#!/bin/sh
if test "$1" = version; then
    printf '%s\n' "$GOML_BOOTSTRAP_TEST_VERSION"
    exit 0
fi
exit 1
DRIVER
chmod +x "$package/bin/goml"
for tool in gomlc gomlfmt gomldoc gomllsp; do
    cp "$package/bin/goml" "$package/bin/$tool"
done
for resource in builtin/goml.toml builtin/contract.goml builtin/runtime.goml builtin/impls.goml \
    builtin/language.goml builtin/numeric.goml builtin/derive.goml prelude/goml.toml std/goml.toml; do
    printf '%s\n' placeholder > "$package/lib/$resource"
done
before_tree="$(toolchain_hash_files "$prefix/bin" "$prefix/lib" | toolchain_sha256)"
before_stamp="$(cat "$prefix/.stage0.sha256")"
for failure in missing-resource finalize; do
    if test "$failure" = finalize; then
        printf '%s\n' placeholder > "$package/lib/prelude/prelude.goml"
    fi
    tar -czf "$temporary/invalid.tar.gz" -C "$temporary" package
    cp bootstrap/stage0.env "$temporary/invalid.env"
    invalid_checksum="$(toolchain_sha256 "$temporary/invalid.tar.gz" | cut -d ' ' -f 1)"
    printf '\nGOML_STAGE0_LINUX_AMD64_SHA256=%s\nGOML_STAGE0_DARWIN_ARM64_SHA256=%s\n' \
        "$invalid_checksum" "$invalid_checksum" >> "$temporary/invalid.env"
    if GOML_STAGE0_ARCHIVE="$temporary/invalid.tar.gz" \
        GOML_BOOTSTRAP_TEST_VERSION="goml ${GOML_STAGE0_RELEASE_TAG#v}" \
        bash bootstrap/bootstrap.sh "$temporary/invalid.env" "$prefix" \
        > "$temporary/$failure.stdout" 2> "$temporary/$failure.stderr"; then
        exit 1
    fi
    after_tree="$(toolchain_hash_files "$prefix/bin" "$prefix/lib" | toolchain_sha256)"
    test "$before_tree" = "$after_tree"
    test "$before_stamp" = "$(cat "$prefix/.stage0.sha256")"
    test -f "$prefix/custom/retained"
done
