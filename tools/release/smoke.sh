#!/usr/bin/env bash

set -euo pipefail

test "$#" = 1 || test "$#" = 3

repository_root="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
release_version="$1"
source "$repository_root/tools/release/common.sh"
target_goos="${2:-$(go env GOHOSTOS)}"
target_goarch="${3:-$(go env GOHOSTARCH)}"
platform="$(release_platform "$target_goos" "$target_goarch")"
test "$(go env GOHOSTOS)" = "$target_goos"
test "$(go env GOHOSTARCH)" = "$target_goarch"
export GOOS="$target_goos" GOARCH="$target_goarch"
package="goml-$release_version-$platform"
smoke_root="$(mktemp -d)"
smoke_root="$(cd "$smoke_root" && pwd -P)"
trap 'rm -rf "$smoke_root"' EXIT
export GOML_HOME="$smoke_root/home"

(
    cd "$repository_root/dist"
    release_sha256 -c "$package.tar.gz.sha256"
)

tar -xzf "$repository_root/dist/$package.tar.gz" -C "$smoke_root"
test "$("$smoke_root/$package/bin/goml" version)" = "goml $release_version"
test "$("$smoke_root/$package/bin/gomlc" version --format json | jq -r .version)" = "$release_version"
test "$("$smoke_root/$package/bin/gomlfmt" --version)" = "gomlfmt $release_version"
test "$("$smoke_root/$package/bin/gomldoc" --version)" = "gomldoc $release_version"
test -f "$smoke_root/$package/lib/builtin/contract.goml"
test -f "$smoke_root/$package/lib/builtin/goml.toml"
test -f "$smoke_root/$package/lib/builtin/runtime.goml"
test -f "$smoke_root/$package/lib/builtin/impls.goml"
test -f "$smoke_root/$package/lib/builtin/language.goml"
test -f "$smoke_root/$package/lib/builtin/numeric.goml"
test -f "$smoke_root/$package/lib/builtin/derive.goml"
test -f "$smoke_root/$package/lib/prelude/prelude.goml"
test -f "$smoke_root/$package/lib/prelude/goml.toml"
test -f "$smoke_root/$package/lib/std/goml.toml"
test "$("$smoke_root/$package/bin/goml-c-bind" --runtime-dir)" = "$smoke_root/$package/lib/cabi"
test ! -e "$smoke_root/$package/lib/compiler/compiler-world-v2.gaf"
"$smoke_root/$package/bin/goml" __toolchain-finalize --prefix "$smoke_root/$package"
test -f "$smoke_root/$package/lib/compiler/compiler-world-v2.gaf"
cp -R "$repository_root/tools/release/testdata/smoke" "$smoke_root/project"
cd "$smoke_root/project"
"$smoke_root/$package/bin/goml" check
"$smoke_root/$package/bin/goml" build
test "$("$smoke_root/$package/bin/goml" run)" = "std/works"
"$smoke_root/$package/bin/goml" test
"$smoke_root/$package/bin/gomlfmt" -w ./*.goml
"$smoke_root/$package/bin/gomlfmt" --check ./*.goml
"$smoke_root/$package/bin/goml" fmt --check
"$smoke_root/$package/bin/goml" doc --format json
jq -e '.schema_version == 1 and (.packages | length) > 0' _artifact/doc/module.json >/dev/null
"$smoke_root/$package/bin/gomldoc" --output _artifact/doc-html
test -f _artifact/doc-html/index.html
test -f _artifact/doc-html/style.css
bash "$repository_root/tools/release/lsp_smoke.sh" "$smoke_root/$package/bin/gomllsp" "$release_version"
mkdir -p "$smoke_root/ffi/gen"
cat > "$smoke_root/ffi/go.mod" <<'GOMOD'
module example.com/ffi-smoke

go 1.26.0
GOMOD
jq -n \
    --arg go_executable "$(command -v go)" \
    --arg toolchain "$(GOTOOLCHAIN=local go env GOVERSION)" \
    --arg module_dir "$smoke_root/ffi" \
    --arg caller_dir "$smoke_root/ffi/gen" \
    --arg goos "$target_goos" \
    --arg goarch "$target_goarch" \
    '{protocol_version: 1,
      build_context: {go_executable: $go_executable, toolchain: $toolchain,
        module_dir: $module_dir, goos: $goos, goarch: $goarch, cgo_enabled: "0",
        goflags: "", build_tags: [], go111module: "on", gowork: "off",
        dependencies: "readonly", network: "off"},
      caller_context: {load_mode: "package", import_path: "example.com/ffi-smoke/gen", directory: $caller_dir, package: "gen"},
      bindings: [{binding_id: "smoke::abs", import_path: "math", symbol: "Abs",
        bridge_parameter_types: [{tag: "float64"}], bridge_result_types: [{tag: "float64"}],
        call_mode: "ordinary"}]}' > "$smoke_root/ffi/request.json"
"$smoke_root/$package/bin/goml-go-meta" < "$smoke_root/ffi/request.json" > "$smoke_root/ffi/response.json"
jq -e '.protocol_version == 1 and (.bindings | length) == 1 and .bindings[0].status == "verified"' "$smoke_root/ffi/response.json" >/dev/null
test ! -e "$smoke_root/ffi/go.sum"
test ! -e "$smoke_root/ffi/gen/goml_ffi_witness_0.go"
cp -R "$repository_root/tools/release/testdata/go-export" "$smoke_root/go-export"
cd "$smoke_root/go-export"
export GOWORK=off GOTOOLCHAIN=local GOPROXY=off GOSUMDB=off
"$smoke_root/$package/bin/goml" check --compiler "$smoke_root/$package/bin/gomlc" --ffi-check required
"$smoke_root/$package/bin/goml" export-go . --compiler "$smoke_root/$package/bin/gomlc" --import-path example.com/host/gen/calclib --out ./gen/calclib
release_sha256 gen/calclib/goml_generated.go gen/calclib/goml_exports.json > export-digests
"$smoke_root/$package/bin/goml" export-go . --compiler "$smoke_root/$package/bin/gomlc" --import-path example.com/host/gen/calclib --out ./gen/calclib
release_sha256 -c export-digests
rm calc.goml goml.toml
(
    export PATH="$(dirname "$(command -v go)"):/usr/bin:/bin"
    go test ./...
    test "$(go run .)" = "42 4 7 8"
)
test ! -e go.sum
mkdir -p "$smoke_root/file-read"
for source in goml.toml go.mod main.goml data.txt; do
    cp "$repository_root/goml/testdata/ffi/file-read/$source" "$smoke_root/file-read/$source"
done
cp -R "$repository_root/goml/testdata/ffi/file-read/shim" "$smoke_root/file-read/shim"
cd "$smoke_root/file-read"
"$smoke_root/$package/bin/goml" fmt --check
"$smoke_root/$package/bin/goml" check --ffi-check required
expected_file_read="$(cat <<'OUTPUT'
read: 3
partial data: abc
Go round-trip preserved partial read: true
read error: unexpected EOF
explicit close succeeded: true
original handle is closed: true
OUTPUT
)"
for attempt in 1 2; do
    test "$("$smoke_root/$package/bin/goml" run --ffi-check required)" = "$expected_file_read"
    if test "$attempt" = 1; then
        release_sha256 _artifact/build/pkg/ffi_file_read/goml_generated.go > file-read-digests
    else
        release_sha256 -c file-read-digests
    fi
done
test ! -e go.sum
mkdir -p "$smoke_root/callbacks/shim"
for source in goml.toml go.mod main.goml shim/shim.go; do
    cp "$repository_root/goml/testdata/ffi/callbacks/$source" "$smoke_root/callbacks/$source"
done
cd "$smoke_root/callbacks"
"$smoke_root/$package/bin/goml" fmt --check
"$smoke_root/$package/bin/goml" check --ffi-check required
expected_callbacks="$(cat <<'OUTPUT'
deferred until release: true
concurrent retained callbacks: true
unregistered callbacks stopped: true
recursive reentry: true
callback can unregister itself: true
registration stress balanced: true
OUTPUT
)"
for attempt in 1 2; do
    test "$("$smoke_root/$package/bin/goml" run --ffi-check required)" = "$expected_callbacks"
    if test "$attempt" = 1; then
        release_sha256 _artifact/build/pkg/ffi_callbacks/goml_generated.go > callback-digests
    else
        release_sha256 -c callback-digests
    fi
done
test ! -e go.sum
mkdir -p "$smoke_root/bind-go"
for source in goml.toml go.mod main.goml bindings.json; do
    cp "$repository_root/goml/testdata/ffi/bind-go/$source" "$smoke_root/bind-go/$source"
done
cd "$smoke_root/bind-go"
"$smoke_root/$package/bin/goml" bind-go bindings.json --dry-run
test ! -e bindings
test ! -e native
"$smoke_root/$package/bin/goml" bind-go bindings.json
release_sha256 bindings/generated.goml native/generated.go bindings.json.goml-bind.json > binding-digests
"$smoke_root/$package/bin/goml" bind-go bindings.json
release_sha256 -c binding-digests
"$smoke_root/$package/bin/goml" fmt --check
"$smoke_root/$package/bin/goml" check
expected_bindings="$(cat <<'OUTPUT'
3
42
true
OUTPUT
)"
test "$("$smoke_root/$package/bin/goml" run)" = "$expected_bindings"
printf '\nfn handwritten() -> () {}\n' >> bindings/generated.goml
if "$smoke_root/$package/bin/goml" bind-go bindings.json > binding-rejected.log 2>&1; then
    exit 1
fi
test ! -e go.sum
test ! -e .goml-bind-go-lock
"$smoke_root/$package/bin/goml" bind-c --help > "$smoke_root/bind-c-help.txt"
"$smoke_root/$package/bin/goml-c-bind" --help > "$smoke_root/bind-c-helper-help.txt"
mkdir -p "$smoke_root/bind-c"
for source in goml.toml go.mod main.goml bindings_test.goml bindings.json sample.h; do
    cp "$repository_root/goml/testdata/ffi/bind-c/$source" "$smoke_root/bind-c/$source"
done
cd "$smoke_root/bind-c"
export CGO_ENABLED=1
"$smoke_root/$package/bin/goml" bind-c bindings.json
"$smoke_root/$package/bin/goml" check --ffi-check required
"$smoke_root/$package/bin/goml" test
test "$("$smoke_root/$package/bin/goml" run)" = "$(printf '42\nGoML calls C\n255')"

if test "$platform" = darwin-arm64; then
    "$smoke_root/$package/bin/gomlc" run-single "$repository_root/tools/release/testdata/unsupported-linux.goml"
    jq '.backend = "dynamic" | .libraries = ["libc.so.6"] |
        .package = "dynamic_bindings" | .output = "dynamic_bindings/generated.goml" |
        .go_package = "dynamic_native" | .go_output = "dynamic_native/generated.go"' bindings.json > dynamic.json
    if CGO_ENABLED=0 "$smoke_root/$package/bin/goml" bind-c dynamic.json > dynamic.log 2>&1; then
        exit 1
    fi
    grep -q 'dynamic C bindings require a Linux amd64 host' dynamic.log
fi
