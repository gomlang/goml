make:
    bash bootstrap/bootstrap.sh bootstrap/stage0.env stage0
    mkdir -p stage2/bin
    bash tools/goml-go-meta/build.sh stage2
    bash bootstrap/build-stage.sh stage2 stage0/bin/goml stage0/bin/gomlc
    cp gomlc/_bootstrap/stage2/bin/cmd/gomlc/gomlc stage2/bin/gomlc
    cp gomlc/_bootstrap/stage2/bin/cmd/gomlfmt/gomlfmt stage2/bin/gomlfmt
    cp gomlc/_bootstrap/stage2/bin/cmd/gomllsp/gomllsp stage2/bin/gomllsp
    bash tools/lib/install.sh stage2
    cp goml/_bootstrap/stage2/bin/cmd/goml/goml stage2/bin/goml
    bash tools/lib/finalize-toolchain.sh stage2 stage2/bin/goml stage2/bin/gomlc

make-tools: make

test: make-tools
    cd tools/goml-go-meta && go test ./...
    bash tools/lib/install.sh _artifact/gomlc-test/test
    cd gomlc && GOML_TEST_GOML=../stage2/bin/goml GOML_TEST_GOMLC=../stage2/bin/gomlc GOML_TEST_COMPILER_WORLD=../stage2/lib/compiler/compiler-world-v2.gaf ../stage2/bin/goml test --target-dir ../_artifact/gomlc-test --compiler ../stage2/bin/gomlc --jobs 16 --timeout 10m
    cd goml && GOML_TEST_GOML=../stage2/bin/goml GOML_TEST_GOMLC=../stage2/bin/gomlc ../stage2/bin/goml test --compiler ../stage2/bin/gomlc --jobs 16 --timeout 10m

gomlgo-oracle:
    mkdir -p _artifact/gomlgo
    cd gomlgo/oracle && "${GOMLGO_GO:-/usr/lib/go-1.26/bin/go}" test ./...
    cd gomlgo/oracle && "${GOMLGO_GO:-/usr/lib/go-1.26/bin/go}" build -o ../../_artifact/gomlgo/oracle .

gomlgo-test: make gomlgo-oracle
    bash tools/lib/install.sh _artifact/gomlgo-test/test
    cd gomlgo && GOMLGO_ORACLE=../_artifact/gomlgo/oracle ../stage2/bin/goml test --target-dir ../_artifact/gomlgo-test --compiler ../stage2/bin/gomlc --jobs 16 --timeout 10m

gomlgo-type-test: gomlgo-test

gomlgo-build: make gomlgo-oracle
    cd gomlgo && ../stage2/bin/goml build --target-dir ../_artifact/gomlgo-build --compiler ../stage2/bin/gomlc

gomlgo-diff: gomlgo-build
    GOMLGO_ORACLE="$(pwd)/_artifact/gomlgo/oracle" _artifact/gomlgo-build/bin/cmd/gomlgo_scan_diff/gomlgo_scan_diff gomlgo/testdata
    GOMLGO_ORACLE="$(pwd)/_artifact/gomlgo/oracle" _artifact/gomlgo-build/bin/cmd/gomlgo_parse_diff/gomlgo_parse_diff ast-comments gomlgo/testdata/valid
    GOMLGO_ORACLE="$(pwd)/_artifact/gomlgo/oracle" _artifact/gomlgo-build/bin/cmd/gomlgo_parse_diff/gomlgo_parse_diff parse-acceptance gomlgo/testdata/invalid
    GOMLGO_ORACLE="$(pwd)/_artifact/gomlgo/oracle" _artifact/gomlgo-build/bin/cmd/gomlgo_parse_diff/gomlgo_parse_diff parse-acceptance gomlgo/testdata/regressions

gomlgo-type-diff: gomlgo-build
    GOMLGO_ORACLE="$(pwd)/_artifact/gomlgo/oracle" _artifact/gomlgo-build/bin/cmd/gomlgo_type_diff/gomlgo_type_diff gomlgo/testdata/typecheck

gomlgo-type-diff-official: gomlgo-build
    systemd-run --user --scope -p MemoryHigh=8G -p MemoryMax=10G -p MemorySwapMax=2G -p TasksMax=256 -p CPUQuota=300% sh -c 'find /usr/lib/go-1.26/src/internal/types/testdata/check /usr/lib/go-1.26/src/internal/types/testdata/spec /usr/lib/go-1.26/src/internal/types/testdata/examples /usr/lib/go-1.26/src/internal/types/testdata/fixedbugs /usr/lib/go-1.26/src/go/types/testdata/local -type f -name "*.go" -print0 | sort -z | xargs -0 -n1 env GOMLGO_ORACLE="$(pwd)/_artifact/gomlgo/oracle" _artifact/gomlgo-build/bin/cmd/gomlgo_type_diff/gomlgo_type_diff info'

gomlgo-type-diagnostics-official: gomlgo-build
    systemd-run --user --scope -p MemoryHigh=8G -p MemoryMax=10G -p MemorySwapMax=2G -p TasksMax=256 -p CPUQuota=300% sh -c 'find /usr/lib/go-1.26/src/internal/types/testdata/check /usr/lib/go-1.26/src/internal/types/testdata/spec /usr/lib/go-1.26/src/internal/types/testdata/examples /usr/lib/go-1.26/src/internal/types/testdata/fixedbugs /usr/lib/go-1.26/src/go/types/testdata/local -type f -name "*.go" -print0 | sort -z | xargs -0 -n1 env GOMLGO_ORACLE="$(pwd)/_artifact/gomlgo/oracle" _artifact/gomlgo-build/bin/cmd/gomlgo_type_diff/gomlgo_type_diff diagnostics'

gomlgo-type-diff-goroot limit="-1" scope="std" mode="acceptance" goos="linux" goarch="amd64" cgo="0" tests="none": gomlgo-build
    systemd-run --user --scope -p MemoryHigh=6G -p MemoryMax=8G -p MemorySwapMax=1G -p TasksMax=512 -p CPUQuota=300% env GOMLGO_ORACLE="$(pwd)/_artifact/gomlgo/oracle" _artifact/gomlgo-build/bin/cmd/gomlgo_type_goroot_diff/gomlgo_type_goroot_diff "{{mode}}" "{{scope}}" "{{limit}}" /usr/lib/go-1.26/src "" "{{goos}}" "{{goarch}}" "{{cgo}}" "{{tests}}"

gomlgo-type-diff-goroot-matrix limit="16": gomlgo-build
    systemd-run --user --scope -p MemoryHigh=6G -p MemoryMax=8G -p MemorySwapMax=1G -p TasksMax=512 -p CPUQuota=300% sh -c 'for target in "linux amd64" "linux arm64" "windows amd64"; do set -- $target; GOMLGO_ORACLE="$(pwd)/_artifact/gomlgo/oracle" _artifact/gomlgo-build/bin/cmd/gomlgo_type_goroot_diff/gomlgo_type_goroot_diff acceptance std "{{limit}}" /usr/lib/go-1.26/src "" "$1" "$2" 0 none || exit $?; done'

gomlgo-type-diff-goroot-tests limit="16" tests="external": gomlgo-build
    systemd-run --user --scope -p MemoryHigh=6G -p MemoryMax=8G -p MemorySwapMax=1G -p TasksMax=512 -p CPUQuota=300% env GOMLGO_ORACLE="$(pwd)/_artifact/gomlgo/oracle" _artifact/gomlgo-build/bin/cmd/gomlgo_type_goroot_diff/gomlgo_type_goroot_diff acceptance std "{{limit}}" /usr/lib/go-1.26/src "" linux amd64 0 "{{tests}}"

gomlgo-type-mutate-diff limit="25": gomlgo-build
    GOMLGO_ORACLE="$(pwd)/_artifact/gomlgo/oracle" _artifact/gomlgo-build/bin/cmd/gomlgo_type_mutate_diff/gomlgo_type_mutate_diff gomlgo/testdata/typecheck/valid "{{limit}}" "$(pwd)/_artifact/gomlgo-type-diff/failures"

gomlgo-diff-goroot: gomlgo-build
    GOMLGO_ORACLE="$(pwd)/_artifact/gomlgo/oracle" _artifact/gomlgo-build/bin/cmd/gomlgo_scan_diff/gomlgo_scan_diff /usr/lib/go-1.26/src
    GOMLGO_ORACLE="$(pwd)/_artifact/gomlgo/oracle" _artifact/gomlgo-build/bin/cmd/gomlgo_parse_diff/gomlgo_parse_diff ast-comments /usr/lib/go-1.26/src

gomlgo-parse-diff: gomlgo-build
    GOMLGO_ORACLE="$(pwd)/_artifact/gomlgo/oracle" _artifact/gomlgo-build/bin/cmd/gomlgo_parse_diff/gomlgo_parse_diff parse-acceptance gomlgo/testdata/valid

gomlgo-parse-diff-goroot: gomlgo-build
    GOMLGO_ORACLE="$(pwd)/_artifact/gomlgo/oracle" _artifact/gomlgo-build/bin/cmd/gomlgo_parse_diff/gomlgo_parse_diff parse-acceptance /usr/lib/go-1.26/src

gomlgo-ast-shape-goroot: gomlgo-build
    GOMLGO_ORACLE="$(pwd)/_artifact/gomlgo/oracle" _artifact/gomlgo-build/bin/cmd/gomlgo_parse_diff/gomlgo_parse_diff ast-shape /usr/lib/go-1.26/src

gomlgo-ast-position-goroot: gomlgo-build
    GOMLGO_ORACLE="$(pwd)/_artifact/gomlgo/oracle" _artifact/gomlgo-build/bin/cmd/gomlgo_parse_diff/gomlgo_parse_diff ast-position /usr/lib/go-1.26/src

gomlgo-ast-comments-goroot: gomlgo-build
    GOMLGO_ORACLE="$(pwd)/_artifact/gomlgo/oracle" _artifact/gomlgo-build/bin/cmd/gomlgo_parse_diff/gomlgo_parse_diff ast-comments /usr/lib/go-1.26/src

gomlgo-mutate-diff limit="25": gomlgo-build
    GOMLGO_ORACLE="$(pwd)/_artifact/gomlgo/oracle" _artifact/gomlgo-build/bin/cmd/gomlgo_mutate_diff/gomlgo_mutate_diff /usr/lib/go-1.26/src "{{limit}}" "$(pwd)/_artifact/gomlgo-diff/failures"

all: test

cloc:
    find ./goml -type f -name '*.gom' -exec cat {} + | wc -l | awk '{ print "goml:", $1 }'
    find ./gomlc -type f -name '*.gom' -exec cat {} + | wc -l | awk '{ print "gomlc:", $1 }'
    find ./lib -type f -name '*.gom' -exec cat {} + | wc -l | awk '{ print "lib:", $1 }'

clean:
    rm -rf _artifact _bootstrap
    rm -rf gomlc/_artifact gomlc/_bootstrap
    rm -rf goml/_artifact goml/_bootstrap
    rm -rf stage1 stage2 stage3
    rm -rf editors/vscode/bin editors/vscode/lib

_bootstrap-stage3:
    rm -rf gomlc/_bootstrap/stage3 gomlc/_bootstrap/stage3-fixed
    rm -rf goml/_bootstrap/stage3 goml/_bootstrap/stage3-fixed stage3
    mkdir -p stage3/bin
    bash tools/goml-go-meta/build.sh stage3
    bash bootstrap/build-stage.sh stage3 stage2/bin/goml stage2/bin/gomlc compiler
    cp gomlc/_bootstrap/stage3/bin/cmd/gomlc/gomlc stage3/bin/gomlc
    bash tools/lib/install.sh stage3
    cp goml/_bootstrap/stage3/bin/cmd/goml/goml stage3/bin/goml
    bash tools/lib/finalize-toolchain.sh stage3 stage3/bin/goml stage3/bin/gomlc
    bash bootstrap/build-stage.sh stage3-fixed stage3/bin/goml stage3/bin/gomlc artifacts
    diff -ru --exclude='*.goml-*-fingerprint' --exclude='*.goml-output-epoch' gomlc/_bootstrap/stage3/build/pkg gomlc/_bootstrap/stage3-fixed/build/pkg
    diff -ru --exclude='*.goml-*-fingerprint' --exclude='*.goml-output-epoch' goml/_bootstrap/stage3/build/pkg goml/_bootstrap/stage3-fixed/build/pkg

bootstrap:
    rm -rf gomlc/_bootstrap/stage1 gomlc/_bootstrap/stage2 gomlc/_bootstrap/stage3 gomlc/_bootstrap/stage3-fixed
    rm -rf goml/_bootstrap/stage1 goml/_bootstrap/stage2 goml/_bootstrap/stage3 goml/_bootstrap/stage3-fixed
    rm -rf stage0 stage1 stage2 stage3
    just make
    just _bootstrap-stage3

_ci-scripts:
    bash -n tools/release/release.sh tools/release/test.sh tools/release/package.sh tools/release/smoke.sh tools/release/lsp_smoke.sh
    bash -n tools/lib/install.sh tools/lib/test.sh tools/lib/finalize-toolchain.sh tools/goml-go-meta/build.sh
    cd tools/goml-go-meta && go test -race ./...
    bash tools/release/test.sh
    bash tools/lib/test.sh stage2
    bash tools/release/release.sh check-version "$(cat VERSION)"
    bash -n bootstrap/bootstrap.sh

_ci-gomlc-test:
    bash tools/lib/install.sh _artifact/gomlc-test/test
    cd gomlc && GOML_TEST_GOML=../stage2/bin/goml GOML_TEST_GOMLC=../stage2/bin/gomlc GOML_TEST_COMPILER_WORLD=../stage2/lib/compiler/compiler-world-v2.gaf ../stage2/bin/goml test --target-dir ../_artifact/gomlc-test --compiler ../stage2/bin/gomlc --jobs "${GOML_TEST_JOBS:-16}" --timeout 10m

_ci-goml-test:
    cd goml && GOML_TEST_GOML=../stage2/bin/goml GOML_TEST_GOMLC=../stage2/bin/gomlc ../stage2/bin/goml test --compiler ../stage2/bin/gomlc --jobs "${GOML_TEST_JOBS:-4}" --timeout 10m

_ci-vscode:
    mkdir -p editors/vscode/bin
    find editors/vscode/bin -mindepth 1 -type f -delete
    find editors/vscode/bin -mindepth 1 -depth -type d -empty -delete
    cp stage2/bin/gomllsp editors/vscode/bin/gomllsp
    bash tools/lib/install.sh editors/vscode
    cd editors/vscode && npm install
    cd editors/vscode && npm run compile

_ci-release-smoke:
    bash tools/release/package.sh "$(cat VERSION)"
    bash tools/release/smoke.sh "$(cat VERSION)"

ci:
    rm -rf gomlc/_bootstrap/stage1 gomlc/_bootstrap/stage2 gomlc/_bootstrap/stage3 gomlc/_bootstrap/stage3-fixed
    rm -rf goml/_bootstrap/stage1 goml/_bootstrap/stage2 goml/_bootstrap/stage3 goml/_bootstrap/stage3-fixed
    rm -rf stage1 stage2 stage3
    just make
    bash tools/parallel-ci.sh

set-version version:
    bash tools/release/release.sh set-version "{{version}}"
    cd editors/vscode && npm version --no-git-tag-version "{{version}}"
    bash tools/release/release.sh check-version "$(cat VERSION)"

set-bootstrap-stage0 version sha256:
    bash tools/release/release.sh set-stage0 "{{version}}" "{{sha256}}"
    bash bootstrap/bootstrap.sh bootstrap/stage0.env stage0

vscode-ext: make-tools
    mkdir -p editors/vscode/bin
    find editors/vscode/bin -mindepth 1 -type f -delete
    find editors/vscode/bin -mindepth 1 -depth -type d -empty -delete
    cp stage2/bin/gomllsp editors/vscode/bin/gomllsp
    bash tools/lib/install.sh editors/vscode
    cd editors/vscode && npm install
    cd editors/vscode && npm run compile

package-vscode-ext: vscode-ext
    cd editors/vscode && npx @vscode/vsce package --allow-missing-repository --skip-license

install: make-tools
    mkdir -p "${GOML_HOME:-$HOME/.goml}/bin"
    cp stage2/bin/gomlc "${GOML_HOME:-$HOME/.goml}/bin/gomlc"
    cp stage2/bin/goml "${GOML_HOME:-$HOME/.goml}/bin/goml"
    cp stage2/bin/gomlfmt "${GOML_HOME:-$HOME/.goml}/bin/gomlfmt"
    cp stage2/bin/gomllsp "${GOML_HOME:-$HOME/.goml}/bin/gomllsp"
    cp stage2/bin/goml-go-meta "${GOML_HOME:-$HOME/.goml}/bin/goml-go-meta"
    bash tools/lib/install.sh "${GOML_HOME:-$HOME/.goml}"
    bash tools/lib/finalize-toolchain.sh "${GOML_HOME:-$HOME/.goml}" "${GOML_HOME:-$HOME/.goml}/bin/goml" "${GOML_HOME:-$HOME/.goml}/bin/gomlc"

update-golden: make
    bash tools/lib/install.sh _artifact/gomlc-test/test
    cd gomlc && UPDATE_EXPECT=1 ../stage2/bin/goml test formatter --target-dir ../_artifact/gomlc-test --compiler ../stage2/bin/gomlc --jobs 16 --timeout 10m
    cd gomlc && UPDATE_EXPECT=1 ../stage2/bin/goml test pipeline_test --target-dir ../_artifact/gomlc-test --compiler ../stage2/bin/gomlc --jobs 16 --timeout 10m
    cd gomlc && UPDATE_EXPECT=1 GOML_TEST_GOML=../stage2/bin/goml GOML_TEST_GOMLC=../stage2/bin/gomlc ../stage2/bin/goml test compiler_test --target-dir ../_artifact/gomlc-test --compiler ../stage2/bin/gomlc --jobs 16 --timeout 10m

verify-golden: make
    bash tools/lib/install.sh _artifact/gomlc-test/test
    cd gomlc && ../stage2/bin/goml test formatter --target-dir ../_artifact/gomlc-test --compiler ../stage2/bin/gomlc --jobs 16 --timeout 10m
    cd gomlc && ../stage2/bin/goml test pipeline_test --target-dir ../_artifact/gomlc-test --compiler ../stage2/bin/gomlc --jobs 16 --timeout 10m
    cd gomlc && GOML_TEST_GOML=../stage2/bin/goml GOML_TEST_GOMLC=../stage2/bin/gomlc ../stage2/bin/goml test compiler_test --target-dir ../_artifact/gomlc-test --compiler ../stage2/bin/gomlc --jobs 16 --timeout 10m
