set -euo pipefail

test "$#" = 1 || test "$#" = 2
repository_root="$(cd "$(dirname "$0")/../.." && pwd)"
cd "$repository_root"
bash tools/release/release.sh check-version "$1"
source tools/release/common.sh
cd "${2:-dist}"

archives=("goml-$1-linux-amd64.tar.gz" "goml-$1-darwin-arm64.tar.gz")
for archive in "${archives[@]}"; do
    release_sha256 -c "$archive.sha256"
done
release_sha256 "${archives[@]}" > SHA256SUMS.tmp
mv SHA256SUMS.tmp SHA256SUMS
