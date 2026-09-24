#!/usr/bin/env bash

set -euo pipefail

test "$#" = 1

repository_root="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
release_version="$1"
package="goml-$release_version-linux-amd64"
smoke_root="$(mktemp -d)"
trap 'rm -rf "$smoke_root"' EXIT

tar -xzf "$repository_root/dist/$package.tar.gz" -C "$smoke_root"
test -f "$smoke_root/$package/lib/builtin/contract.gom"
test -f "$smoke_root/$package/lib/builtin/goml.toml"
test -f "$smoke_root/$package/lib/builtin/runtime.gom"
test -f "$smoke_root/$package/lib/builtin/impls.gom"
test -f "$smoke_root/$package/lib/builtin/language.gom"
test -f "$smoke_root/$package/lib/builtin/numeric.gom"
test -f "$smoke_root/$package/lib/builtin/derive.gom"
test -f "$smoke_root/$package/lib/prelude/prelude.gom"
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
"$smoke_root/$package/bin/gomlfmt" -w ./*.gom
"$smoke_root/$package/bin/gomlfmt" --check ./*.gom
"$smoke_root/$package/bin/goml" fmt --check
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
    '{protocol_version: 1,
      build_context: {go_executable: $go_executable, toolchain: $toolchain,
        module_dir: $module_dir, goos: "linux", goarch: "amd64", cgo_enabled: "0",
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
export GOML_HOME="$smoke_root/export-home"
export GOWORK=off GOTOOLCHAIN=local GOPROXY=off GOSUMDB=off
"$smoke_root/$package/bin/goml" check --compiler "$smoke_root/$package/bin/gomlc" --ffi-check required
"$smoke_root/$package/bin/goml" export-go . --compiler "$smoke_root/$package/bin/gomlc" --import-path example.com/host/gen/calclib --out ./gen/calclib
sha256sum gen/calclib/goml_generated.go gen/calclib/goml_exports.json > export-digests
"$smoke_root/$package/bin/goml" export-go . --compiler "$smoke_root/$package/bin/gomlc" --import-path example.com/host/gen/calclib --out ./gen/calclib
sha256sum -c export-digests
rm calc.gom goml.toml
export PATH="$(dirname "$(command -v go)"):/usr/bin:/bin"
go test ./...
test "$(go run .)" = "42 4 7 8"
test ! -e go.sum
mkdir -p "$smoke_root/file-read"
for source in goml.toml go.mod main.gom data.txt; do
    cp "$repository_root/examples/ffi-file-read/$source" "$smoke_root/file-read/$source"
done
cp -R "$repository_root/examples/ffi-file-read/shim" "$smoke_root/file-read/shim"
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
        sha256sum _artifact/build/pkg/ffi_file_read/goml_generated.go > file-read-digests
    else
        sha256sum -c file-read-digests
    fi
done
test ! -e go.sum
mkdir -p "$smoke_root/callbacks/shim"
for source in goml.toml go.mod main.gom shim/shim.go; do
    cp "$repository_root/examples/ffi-callbacks/$source" "$smoke_root/callbacks/$source"
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
        sha256sum _artifact/build/pkg/ffi_callbacks/goml_generated.go > callback-digests
    else
        sha256sum -c callback-digests
    fi
done
test ! -e go.sum
mkdir -p "$smoke_root/bind-go"
for source in goml.toml go.mod main.gom bindings.json; do
    cp "$repository_root/examples/ffi-bind-go/$source" "$smoke_root/bind-go/$source"
done
cd "$smoke_root/bind-go"
"$smoke_root/$package/bin/goml" bind-go bindings.json --dry-run
test ! -e bindings
test ! -e native
"$smoke_root/$package/bin/goml" bind-go bindings.json
sha256sum bindings/generated.gom native/generated.go bindings.json.goml-bind.json > binding-digests
"$smoke_root/$package/bin/goml" bind-go bindings.json
sha256sum -c binding-digests
"$smoke_root/$package/bin/goml" fmt --check
"$smoke_root/$package/bin/goml" check
expected_bindings="$(cat <<'OUTPUT'
3
42
true
OUTPUT
)"
test "$("$smoke_root/$package/bin/goml" run)" = "$expected_bindings"
printf '\nfn handwritten() -> () {}\n' >> bindings/generated.gom
if "$smoke_root/$package/bin/goml" bind-go bindings.json > binding-rejected.log 2>&1; then
    exit 1
fi
test ! -e go.sum
test ! -e .goml-bind-go-lock
"$smoke_root/$package/bin/goml" bind-c --help > "$smoke_root/bind-c-help.txt"
"$smoke_root/$package/bin/goml-c-bind" --help > "$smoke_root/bind-c-helper-help.txt"
