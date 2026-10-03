set -euo pipefail

repository_root="$(cd "$(dirname "$0")/../.." && pwd)"
cd "$repository_root"
release_script="$repository_root/tools/release/release.sh"

bash "$release_script" check-next v0.1.0 v0.1.1
bash "$release_script" check-next v0.1.0 v0.2.0
bash "$release_script" check-next v0.1.0 v1.0.0
! bash "$release_script" check-next v0.1.0 v0.1.2
! bash "$release_script" check-next v0.1.0 v0.3.0
test "$(bash "$release_script" latest v0.1.9 v0.2.0 v0.1.10)" = v0.2.0

source tools/release/common.sh
test "$(release_platform linux amd64)" = linux-amd64
test "$(release_platform darwin arm64)" = darwin-arm64
if release_platform darwin amd64; then
    exit 1
fi

temporary="$(mktemp -d)"
trap 'rm -rf "$temporary"' EXIT
release_version="$(cat VERSION)"
mkdir -p "$temporary/archives with spaces" "$temporary/hash-bin"
archive_dir="$temporary/archives with spaces"
for platform in linux-amd64 darwin-arm64; do
    archive="goml-$release_version-$platform.tar.gz"
    printf '%s\n' "$platform" > "$archive_dir/$archive"
    (
        cd "$archive_dir"
        release_sha256 "$archive" > "$archive.sha256"
    )
done
bash tools/release/checksums.sh "$release_version" "$archive_dir"
test "$(wc -l < "$archive_dir/SHA256SUMS" | tr -d ' ')" = 2
(
    cd "$archive_dir"
    release_sha256 -c SHA256SUMS
    ln -s "$(command -v shasum)" "$temporary/hash-bin/shasum"
    PATH="$temporary/hash-bin" release_sha256 -c SHA256SUMS
    PATH="$temporary/hash-bin" release_sha256 "goml-$release_version-linux-amd64.tar.gz" > fallback.sha256
    cmp fallback.sha256 "goml-$release_version-linux-amd64.tar.gz.sha256"
)
rm "$archive_dir/SHA256SUMS"
printf 'corrupt\n' >> "$archive_dir/goml-$release_version-darwin-arm64.tar.gz"
if bash tools/release/checksums.sh "$release_version" "$archive_dir"; then
    exit 1
fi
test ! -e "$archive_dir/SHA256SUMS"
rm "$archive_dir/goml-$release_version-darwin-arm64.tar.gz"
if bash tools/release/checksums.sh "$release_version" "$archive_dir"; then
    exit 1
fi
test ! -e "$archive_dir/SHA256SUMS"
