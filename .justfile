make:
    bash bootstrap/bootstrap.sh bootstrap/stage0.env stage0
    mkdir -p stage2/bin
    bash tools/goml-go-meta/build.sh stage2
    bash bootstrap/build-stage.sh stage2 stage0/bin/goml stage0/bin/gomlc
    cp gomlc/_bootstrap/stage2/bin/cmd/gomlc/gomlc stage2/bin/gomlc
    cp gomlc/_bootstrap/stage2/bin/cmd/gomlfmt/gomlfmt stage2/bin/gomlfmt
    cp gomlc/_bootstrap/stage2/bin/cmd/gomldoc/gomldoc stage2/bin/gomldoc
    cp gomlc/_bootstrap/stage2/bin/cmd/gomllsp/gomllsp stage2/bin/gomllsp
    bash tools/lib/install.sh stage2
    cp goml/_bootstrap/stage2/bin/cmd/goml/goml stage2/bin/goml
    bash tools/lib/finalize-toolchain.sh stage2 stage2/bin/goml stage2/bin/gomlc

make-tools: make

[positional-arguments]
test: make-tools
    cd toolchain/manifest && ../../stage2/bin/goml test --compiler ../../stage2/bin/gomlc --jobs 4
    cd tools/goml-go-meta && go test ./...
    cd lib/cabi && CGO_ENABLED=0 go test -gcflags=all=-d=checkptr=2 ./...
    bash tools/lib/install.sh _artifact/gomlc-test/test
    cd gomlc && GOML_TEST_GOML=../stage2/bin/goml GOML_TEST_GOMLC=../stage2/bin/gomlc GOML_TEST_COMPILER_WORLD=../stage2/lib/compiler/compiler-world-v2.gaf ../stage2/bin/goml test --target-dir ../_artifact/gomlc-test --compiler ../stage2/bin/gomlc --jobs 16 --timeout 10m
    cd goml && GOML_TEST_GOML=../stage2/bin/goml GOML_TEST_GOMLC=../stage2/bin/gomlc ../stage2/bin/goml test --compiler ../stage2/bin/gomlc --jobs 16 --timeout 10m

all: test

cloc:
    find ./goml -type f -name '*.goml' -exec cat {} + | wc -l | awk '{ print "goml:", $1 }'
    find ./gomlc -type f -name '*.goml' -exec cat {} + | wc -l | awk '{ print "gomlc:", $1 }'
    find ./lib -type f -name '*.goml' -exec cat {} + | wc -l | awk '{ print "lib:", $1 }'

clean:
    rm -rf _artifact _bootstrap
    rm -rf gomlc/_artifact gomlc/_bootstrap
    rm -rf goml/_artifact goml/_bootstrap
    rm -rf toolchain/manifest/_artifact
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
    diff -ru --exclude='*.goml-*-fingerprint' --exclude='*.goml-output-epoch' gomlc/_bootstrap/stage3/build/deps gomlc/_bootstrap/stage3-fixed/build/deps
    diff -ru --exclude='*.goml-*-fingerprint' --exclude='*.goml-output-epoch' goml/_bootstrap/stage3/build/pkg goml/_bootstrap/stage3-fixed/build/pkg
    diff -ru --exclude='*.goml-*-fingerprint' --exclude='*.goml-output-epoch' goml/_bootstrap/stage3/build/deps goml/_bootstrap/stage3-fixed/build/deps

bootstrap:
    rm -rf gomlc/_bootstrap/stage1 gomlc/_bootstrap/stage2 gomlc/_bootstrap/stage3 gomlc/_bootstrap/stage3-fixed
    rm -rf goml/_bootstrap/stage1 goml/_bootstrap/stage2 goml/_bootstrap/stage3 goml/_bootstrap/stage3-fixed
    rm -rf stage0 stage1 stage2 stage3
    just make
    just _bootstrap-stage3

_ci-scripts:
    cd toolchain/manifest && ../../stage2/bin/goml test --compiler ../../stage2/bin/gomlc --jobs 4
    go run tools/generate_unicode_tables.go --check
    python3 tools/generate_unicode_casefold.py --check
    python3 tools/syscall/generate.py --check
    python3 tools/syscall/layouts.py --check --verify-headers
    bash -n tools/release/release.sh tools/release/test.sh tools/release/common.sh tools/release/sources.sh tools/release/checksums.sh tools/release/package.sh tools/release/smoke.sh tools/release/lsp_smoke.sh
    bash -n tools/lib/install.sh tools/lib/test.sh tools/lib/finalize-toolchain.sh tools/goml-go-meta/build.sh
    cd tools/goml-go-meta && go test -race ./...
    cd lib/cabi && CGO_ENABLED=0 go test -gcflags=all=-d=checkptr=2 ./...
    bash tools/release/test.sh
    bash tools/lib/test.sh stage2
    bash tools/release/release.sh check-version "$(cat VERSION)"
    bash -n bootstrap/bootstrap.sh bootstrap/test.sh
    bash bootstrap/test.sh

_ci-gomlc-test: _ci-gomlc-test-build _ci-gomlc-test-run

_ci-gomlc-test-build:
    bash tools/lib/install.sh _artifact/gomlc-test/test
    cd gomlc && ../stage2/bin/goml test --no-run --target-dir ../_artifact/gomlc-test --compiler ../stage2/bin/gomlc --jobs "${GOML_TEST_JOBS:-16}"

_ci-gomlc-test-run:
    cd gomlc && GOML_TEST_GOML=../stage2/bin/goml GOML_TEST_GOMLC=../stage2/bin/gomlc GOML_TEST_COMPILER_WORLD=../stage2/lib/compiler/compiler-world-v2.gaf ../stage2/bin/goml test --no-build --target-dir ../_artifact/gomlc-test --compiler ../stage2/bin/gomlc --jobs "${GOML_TEST_JOBS:-16}" --timeout 10m

_ci-goml-test: _ci-goml-test-build _ci-goml-test-run

_ci-goml-test-build:
    cd goml && ../stage2/bin/goml test --no-run --compiler ../stage2/bin/gomlc --jobs "${GOML_TEST_JOBS:-4}"

_ci-goml-test-run:
    cd goml && GOML_TEST_GOML=../stage2/bin/goml GOML_TEST_GOMLC=../stage2/bin/gomlc ../stage2/bin/goml test --no-build --compiler ../stage2/bin/gomlc --jobs "${GOML_TEST_JOBS:-4}" --timeout 10m

_ci-vscode:
    mkdir -p editors/vscode/bin
    find editors/vscode/bin -mindepth 1 -type f -delete
    find editors/vscode/bin -mindepth 1 -depth -type d -empty -delete
    cp stage2/bin/gomllsp editors/vscode/bin/gomllsp
    cp stage2/bin/goml-go-meta editors/vscode/bin/goml-go-meta
    bash tools/lib/install.sh editors/vscode
    cd editors/vscode && npm install
    cd editors/vscode && npm run package

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
    cp stage2/bin/goml-go-meta editors/vscode/bin/goml-go-meta
    bash tools/lib/install.sh editors/vscode
    cd editors/vscode && npm install
    cd editors/vscode && npm run compile

package-vscode-ext: vscode-ext
    cd editors/vscode && npm run package

install: make-tools
    mkdir -p "${GOML_HOME:-$HOME/.goml}/bin"
    cp stage2/bin/gomlc "${GOML_HOME:-$HOME/.goml}/bin/gomlc"
    cp stage2/bin/goml "${GOML_HOME:-$HOME/.goml}/bin/goml"
    cp stage2/bin/gomlfmt "${GOML_HOME:-$HOME/.goml}/bin/gomlfmt"
    cp stage2/bin/gomldoc "${GOML_HOME:-$HOME/.goml}/bin/gomldoc"
    cp stage2/bin/gomllsp "${GOML_HOME:-$HOME/.goml}/bin/gomllsp"
    cp stage2/bin/goml-go-meta "${GOML_HOME:-$HOME/.goml}/bin/goml-go-meta"
    cp stage2/bin/goml-c-bind "${GOML_HOME:-$HOME/.goml}/bin/goml-c-bind"
    bash tools/lib/install.sh "${GOML_HOME:-$HOME/.goml}"
    bash tools/lib/finalize-toolchain.sh "${GOML_HOME:-$HOME/.goml}" "${GOML_HOME:-$HOME/.goml}/bin/goml" "${GOML_HOME:-$HOME/.goml}/bin/gomlc"

update-golden: make
    bash tools/lib/install.sh _artifact/gomlc-test/test
    cd gomlc && UPDATE_EXPECT=1 ../stage2/bin/goml test formatter --target-dir ../_artifact/gomlc-test --compiler ../stage2/bin/gomlc --jobs "${GOML_TEST_JOBS:-16}" --timeout 10m
    cd gomlc && UPDATE_EXPECT=1 ../stage2/bin/goml test pipeline_test --target-dir ../_artifact/gomlc-test --compiler ../stage2/bin/gomlc --jobs "${GOML_TEST_JOBS:-16}" --timeout 10m
    cd gomlc && UPDATE_EXPECT=1 GOML_TEST_GOML=../stage2/bin/goml GOML_TEST_GOMLC=../stage2/bin/gomlc ../stage2/bin/goml test compiler_test --target-dir ../_artifact/gomlc-test --compiler ../stage2/bin/gomlc --jobs "${GOML_TEST_JOBS:-16}" --timeout 10m

verify-golden: make
    bash tools/lib/install.sh _artifact/gomlc-test/test
    cd gomlc && ../stage2/bin/goml test formatter --target-dir ../_artifact/gomlc-test --compiler ../stage2/bin/gomlc --jobs "${GOML_TEST_JOBS:-16}" --timeout 10m
    cd gomlc && ../stage2/bin/goml test pipeline_test --target-dir ../_artifact/gomlc-test --compiler ../stage2/bin/gomlc --jobs "${GOML_TEST_JOBS:-16}" --timeout 10m
    cd gomlc && GOML_TEST_GOML=../stage2/bin/goml GOML_TEST_GOMLC=../stage2/bin/gomlc ../stage2/bin/goml test compiler_test --target-dir ../_artifact/gomlc-test --compiler ../stage2/bin/gomlc --jobs "${GOML_TEST_JOBS:-16}" --timeout 10m
