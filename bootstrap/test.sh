set -euo pipefail

repo_root="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$repo_root"
temporary="$(mktemp -d)"
trap 'rm -rf "$temporary"' EXIT
prefix="$temporary/stage0"

mkdir -p "$prefix/lib/builtin" "$prefix/lib/prelude" "$prefix/lib/std/ascii" \
    "$prefix/lib/cabi" "$prefix/lib/compiler" "$prefix/lib/custom"
for file in builtin/obsolete.gom builtin/obsolete.goml prelude/obsolete.goml \
    std/ascii/obsolete.goml cabi/obsolete.go compiler/compiler-world-v2.gaf; do
    printf '%s\n' invalid > "$prefix/lib/$file"
done
printf '%s\n' retained > "$prefix/lib/custom/retained"

bash bootstrap/bootstrap.sh bootstrap/stage0.env "$prefix"
for file in builtin/obsolete.gom builtin/obsolete.goml prelude/obsolete.goml \
    std/ascii/obsolete.goml cabi/obsolete.go; do
    test ! -e "$prefix/lib/$file"
done
test -f "$prefix/lib/custom/retained"
test -f "$prefix/lib/compiler/compiler-world-v2.gaf"
first_world_mtime="$(stat -c %y "$prefix/lib/compiler/compiler-world-v2.gaf")"
bash bootstrap/bootstrap.sh bootstrap/stage0.env "$prefix"
second_world_mtime="$(stat -c %y "$prefix/lib/compiler/compiler-world-v2.gaf")"
test "$first_world_mtime" = "$second_world_mtime"
