#!/usr/bin/env bash

set -euo pipefail

test "$#" = 1

repository_root="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
cd "$repository_root"

release_version="$1"
package="goml-$release_version-linux-amd64"

bash tools/release/release.sh check-version "$release_version"
rm -rf "dist/$package"
rm -f "dist/$package.tar.gz" dist/SHA256SUMS
mkdir -p "dist/$package/bin"

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -o "dist/$package/bin/goml" goml/_bootstrap/stage2/build/pkg/gomlang/bootstrap_goml/cmd/goml/goml_generated.go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -o "dist/$package/bin/gomlc" gomlc/_bootstrap/stage2/build/pkg/gomlc/cmd/gomlc/goml_generated.go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -o "dist/$package/bin/gomlfmt" gomlc/_bootstrap/stage2/build/pkg/gomlc/cmd/gomlfmt/goml_generated.go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -o "dist/$package/bin/gomllsp" gomlc/_bootstrap/stage2/build/pkg/gomlc/cmd/gomllsp/goml_generated.go

bash tools/goml-go-meta/build.sh "dist/$package"
bash tools/lib/install.sh "dist/$package"
test -f "dist/$package/lib/builtin/contract.gom"
test -f "dist/$package/lib/builtin/goml.toml"
test -f "dist/$package/lib/builtin/runtime.gom"
test -f "dist/$package/lib/builtin/impls.gom"
test -f "dist/$package/lib/builtin/language.gom"
test -f "dist/$package/lib/builtin/numeric.gom"
test -f "dist/$package/lib/builtin/derive.gom"
test -f "dist/$package/lib/prelude/prelude.gom"
test -f "dist/$package/lib/prelude/goml.toml"
test -f "dist/$package/lib/std/goml.toml"
test -f "dist/$package/lib/cabi/go.mod"
test -f "dist/$package/lib/cabi/runtime_linux_amd64.s"
test ! -e "dist/$package/lib/compiler/compiler-world-v2.gaf"
test "$(dist/$package/bin/goml version)" = "goml $release_version"
test "$(dist/$package/bin/gomlc version --format json | jq -r .version)" = "$release_version"
test "$(dist/$package/bin/gomlfmt --version)" = "gomlfmt $release_version"

tar -C dist -czf "dist/$package.tar.gz" "$package"
(
    cd dist
    sha256sum "$package.tar.gz" > SHA256SUMS
)
