set -euo pipefail

test "$#" = 1 || test "$#" = 3
repository_root="$(cd "$(dirname "$0")/../.." && pwd)"
mkdir -p "$1/bin"
output_prefix="$(cd "$1" && pwd -P)"
target_goos="${2:-$(go env GOHOSTOS)}"
target_goarch="${3:-$(go env GOHOSTARCH)}"
(
    cd "$repository_root/tools/goml-go-meta"
    GOWORK=off GO111MODULE=on GOFLAGS= CGO_ENABLED=0 GOOS="$target_goos" GOARCH="$target_goarch" go build -mod=readonly -trimpath -o "$output_prefix/bin/goml-go-meta" .
    GOWORK=off GO111MODULE=on GOFLAGS= CGO_ENABLED=0 GOOS="$target_goos" GOARCH="$target_goarch" go build -mod=readonly -trimpath -o "$output_prefix/bin/goml-c-bind" ./cmd/goml-c-bind
)
