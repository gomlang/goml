set -euo pipefail

test "$#" = 1
repository_root="$(cd "$(dirname "$0")/../.." && pwd)"
output_prefix="$(realpath -m "$1")"
mkdir -p "$output_prefix/bin"
host_goos="$(go env GOHOSTOS)"
host_goarch="$(go env GOHOSTARCH)"
(
    cd "$repository_root/tools/goml-go-meta"
    GOWORK=off GO111MODULE=on GOFLAGS= CGO_ENABLED=0 GOOS="$host_goos" GOARCH="$host_goarch" go build -mod=readonly -trimpath -o "$output_prefix/bin/goml-go-meta" .
    GOWORK=off GO111MODULE=on GOFLAGS= CGO_ENABLED=0 GOOS="$host_goos" GOARCH="$host_goarch" go build -mod=readonly -trimpath -o "$output_prefix/bin/goml-c-bind" ./cmd/goml-c-bind
)
