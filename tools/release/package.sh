#!/usr/bin/env bash

set -euo pipefail

test "$#" = 1 || test "$#" = 3

repository_root="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
cd "$repository_root"

release_version="$1"
source tools/release/common.sh
target_goos="${2:-$(go env GOHOSTOS)}"
target_goarch="${3:-$(go env GOHOSTARCH)}"
platform="$(release_platform "$target_goos" "$target_goarch")"
package="goml-$release_version-$platform"

bash tools/release/release.sh check-version "$release_version"
rm -rf "dist/$package"
rm -f "dist/$package.tar.gz" "dist/$package.tar.gz.sha256"
mkdir -p "dist/$package/bin"

export CGO_ENABLED=0 GOOS="$target_goos" GOARCH="$target_goarch"
export GOWORK=off GO111MODULE=off GOFLAGS=
go build -trimpath -o "dist/$package/bin/goml" goml/_bootstrap/stage2/build/pkg/gomlang/bootstrap_goml/cmd/goml/goml_generated.go
for tool in gomlc gomlfmt gomldoc gomllsp; do
    go build -trimpath -o "dist/$package/bin/$tool" "gomlc/_bootstrap/stage2/build/pkg/gomlc/cmd/$tool/goml_generated.go"
done

bash tools/goml-go-meta/build.sh "dist/$package" "$target_goos" "$target_goarch"
bash tools/lib/install.sh "dist/$package"
test -f "dist/$package/lib/builtin/contract.goml"
test -f "dist/$package/lib/builtin/goml.toml"
test -f "dist/$package/lib/builtin/runtime.goml"
test -f "dist/$package/lib/builtin/impls.goml"
test -f "dist/$package/lib/builtin/language.goml"
test -f "dist/$package/lib/builtin/numeric.goml"
test -f "dist/$package/lib/builtin/derive.goml"
test -f "dist/$package/lib/prelude/prelude.goml"
test -f "dist/$package/lib/prelude/goml.toml"
test -f "dist/$package/lib/std/goml.toml"
test -f "dist/$package/lib/cabi/go.mod"
test -f "dist/$package/lib/cabi/runtime_linux_amd64.s"
test ! -e "dist/$package/lib/compiler/compiler-world-v2.gaf"
for tool in goml gomlc gomlfmt gomldoc gomllsp goml-go-meta goml-c-bind; do
    metadata="$(go version -m "dist/$package/bin/$tool")"
    printf '%s\n' "$metadata" | grep -q "build[[:space:]]GOOS=$target_goos$"
    printf '%s\n' "$metadata" | grep -q "build[[:space:]]GOARCH=$target_goarch$"
    printf '%s\n' "$metadata" | grep -q 'build[[:space:]]CGO_ENABLED=0$'
done

tar -C dist -czf "dist/$package.tar.gz" "$package"
(
    cd dist
    release_sha256 "$package.tar.gz" > "$package.tar.gz.sha256"
)
