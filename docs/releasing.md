# Releasing goml

Releases use strict `vX.Y.Z` tags and publish Linux amd64 and macOS arm64 binaries.

The Go compatibility baseline is Go 1.26. Build and validate releases with Go 1.26.x; users need Go 1.26 or newer to compile generated programs, including their explicit foreign calls. The archives do not bundle a Go toolchain.

The root [VERSION](../VERSION) file is authoritative. `goml`, `gomlc`, `gomlfmt`, `gomldoc`, `gomllsp`, and the VS Code extension must use the same version. Official archives contain these five executables and the builtin, prelude and standard-library resources.

## Version policy

Each release must be one continuous SemVer step from the latest published release:

- patch: `0.1.0` to `0.1.1`
- minor: `0.1.0` to `0.2.0`
- major: `0.1.0` to `1.0.0`

Skipped steps such as `0.1.0` to `0.1.2` or `0.3.0` are rejected.

During the early bootstrap period, releases are limited to continuous `0.1.x` patch versions.

## Publish

Run from the repository root with the prerequisites in the [development guide](development.md), plus an authenticated GitHub CLI (`gh`) for release operations. Replace `X.Y.Z` with the next continuous version after the latest published release:

```sh
release_version=X.Y.Z
just set-version "$release_version"
just ci
```

`just set-version` synchronizes `VERSION`, both GoML version modules, and the VS Code package and lockfile versions. Commit and push the change, then wait for main CI on that exact commit before creating the tag. Keep `release_version` set in the same shell:

```sh
(
    set -eu
    release_sha="$(git rev-parse HEAD)"
    git push origin main
    ci_run_id="$(gh run list --repo gomlang/goml --workflow CI --branch main --commit "$release_sha" --event push --limit 1 --json databaseId --jq '.[0].databaseId // empty')"
    test -n "$ci_run_id"
    gh run watch "$ci_run_id" --repo gomlang/goml --exit-status
    git tag -a "v$release_version" -m "goml v$release_version"
    git push origin "v$release_version"
)
```

The block stops before tagging if the push fails, the CI run is not registered yet, or CI fails. If the run is not yet visible, wait for it to appear and rerun the block.

Main CI builds stage2 directly from the pinned Linux amd64 stage0. Its Linux checks cover compiler, driver, standard-library and script tests, extension compilation and VSIX packaging, archive smoke tests, and the stage3 fixed point. Fixed-point verification compares the compiler and driver artifacts built by stage2 with a rebuild using stage3. A macOS 15 arm64 job builds and tests its release archive using the generated Go sources exported by the Linux build, then independently bootstraps from the pinned macOS arm64 stage0, verifies its native fixed point, and runs toolchain script tests. It receives no Linux executables or compiler world. Both platforms must pass the final `test` job. The independent [gomlgo project](https://github.com/gomlang/gomlgo) owns its separate test suite.

The Release workflow requires successful main CI for the exact tagged commit and verifies the version, previous release, and stage0. It rebuilds stage2 on Linux, exports the five generated Go entry points, then builds all five tools natively in parallel on Ubuntu amd64 and macOS arm64. Each job verifies binary target metadata, packages resources, and runs archive relocation, finalization, project commands, formatting, documentation, LSP, direct Go FFI and Go export smoke tests. macOS also checks rejection of Linux syscalls. Release CI does not repeat the complete Linux suite or fixed-point verification.

Each platform uploads its archive and an archive-specific `.sha256` file. The publish job waits for both platforms, checks both digests, generates one `SHA256SUMS` containing exactly the two archives, and uploads them to a draft release before publishing. Only this job has release write permission. Published releases are not overwritten.

To reproduce packaging after `just make-tools` on Linux:

```sh
bash tools/release/package.sh "$(cat VERSION)" linux amd64
bash tools/release/smoke.sh "$(cat VERSION)" linux amd64
bash tools/release/sources.sh _artifact/release-sources.tar.gz
```

Transfer the source archive to a checkout of the same commit on macOS arm64, with Go 1.26.x, Bash, tar, jq and a C compiler available, then run:

```sh
tar -xzf release-sources.tar.gz
bash tools/release/package.sh "$(cat VERSION)" darwin arm64
bash tools/release/smoke.sh "$(cat VERSION)" darwin arm64
```

The package script also supports cross compilation with an explicit target; smoke tests require the target host. The scripts use `shasum -a 256` when `sha256sum` is unavailable. Once both archives and their `.sha256` files are in `dist`, run `bash tools/release/checksums.sh "$(cat VERSION)"` to generate `dist/SHA256SUMS`.

## Compiler world loading

Toolchain finalization writes and fully validates `lib/compiler/compiler-world-v2.gaf`. The world contains builtin initialization metadata and indexed package interfaces and Core. The compiler treats the world at its own executable-relative installation path as an immutable toolchain resource: normal loading checks its structure, format, compiler identity and dependency fingerprints without recomputing file checksums. The driver uses the installed world's size and modification time for cache freshness, alongside the compiler fingerprint; it does not rehash the world on every command. Worlds supplied from other paths and finalization validation retain complete checksum verification. Release archives and stage0 downloads retain SHA-256 verification.

Replace toolchain resources through installation or finalization rather than editing an installed world in place. Normal compilation does not guarantee detection of corruption in the installed world. Changes that preserve its size and modification time can retain a stale driver cache. Package checking loads interfaces, types and compile-time metadata; runtime Core is loaded for linking.

## Installation

Install Go 1.26 or newer and make it available on `PATH`. The standalone POSIX shell installer is maintained in [gomlang/install](https://github.com/gomlang/install) and requires curl, tar, and sha256sum or shasum:

```sh
curl -fsSL https://raw.githubusercontent.com/gomlang/install/main/install.sh -o install-goml.sh && sh install-goml.sh
export PATH="${GOML_HOME:-$HOME/.goml}/bin:$PATH"
goml version
```

The installer detects Linux amd64 or macOS arm64, downloads the latest published release and its `SHA256SUMS`, verifies the archive, and initializes the compiler world. It installs under `${GOML_HOME:-$HOME/.goml}` and prints a PATH command to add to your shell configuration. It does not modify shell configuration files or require `just` or a source checkout.

To select a release or installation directory:

```sh
sh install-goml.sh --version X.Y.Z --prefix "$HOME/tools/goml"
```

Run the installer again to upgrade. It replaces the five GoML executables and the `builtin`, `prelude`, `std`, and `compiler` resource directories, preserving dependency caches and other files. If replacement or initialization fails, it restores the previous toolchain. Concurrent installations to the same prefix are rejected; after a forced termination, inspect the reported `.goml-install` directory and any backups before retrying. The installer rejects unsupported hosts, Go versions older than 1.26, and installation over a GoML source checkout.

For a manual installation, follow the steps below.

Download the archive for your host and `SHA256SUMS` from the same release:

| Host | Archive | Native release validation |
| --- | --- | --- |
| Linux amd64 | `goml-X.Y.Z-linux-amd64.tar.gz` | Ubuntu 24.04 |
| macOS arm64 (Apple Silicon) | `goml-X.Y.Z-darwin-arm64.tar.gz` | macOS 15 |

For example, on macOS, replace `X.Y.Z` with the release version and verify the downloaded archive before extracting it:

```sh
release_version=X.Y.Z
archive="goml-$release_version-darwin-arm64.tar.gz"
awk -v archive="$archive" '$2 == archive' SHA256SUMS > archive.sha256
test "$(wc -l < archive.sha256 | tr -d ' ')" = 1
shasum -a 256 -c archive.sha256
tar -xzf "$archive"
prefix="$PWD/goml-$release_version-darwin-arm64"
"$prefix/bin/goml" __toolchain-finalize --prefix "$prefix"
export PATH="$prefix/bin:$PATH"
goml version
```

Linux users select `linux-amd64` and may use `sha256sum -c archive.sha256`. Keep `bin` and `lib` together when relocating an installation. Go 1.26 or newer must be on `PATH` to build programs. User-supplied cgo adapters additionally require a C compiler.

macOS supports the compiler, driver, formatter, documentation generator, LSP, Go FFI, user-supplied cgo adapters, and native source bootstrap through `just bootstrap`. The bootstrap uses the published macOS arm64 archive and its pinned checksum in `bootstrap/stage0.env`. `std::os::linux`, the current Linux syscall-backed socket operations in `std::net` and dependent networking APIs remain Linux amd64-specific. SIMD uses its scalar implementation on arm64.

Release archives use a complete toolchain prefix:

```text
goml-X.Y.Z-<os>-<arch>/
├── bin/
│   ├── goml
│   ├── gomlc
│   ├── gomlfmt
│   ├── gomldoc
│   └── gomllsp
└── lib/
    ├── builtin/
    │   ├── goml.toml
    │   └── ...
    ├── prelude/
    │   ├── goml.toml
    │   └── prelude.goml
    └── std/
        ├── goml.toml
        └── ...
```

The compiler resolves `lib` relative to its executable. The archive must preserve this layout exactly. `builtin`, `prelude`, and `std` are separate GoML projects; flat `builtin_*.goml` files must not be packaged.

Release archives contain the toolchain project sources and manifests, but do not contain `lib/compiler/compiler-world-v2.gaf`. The installation command above finalizes it once with the binaries from that same archive.

Finalization builds and validates the compiler world in a temporary path, then atomically installs it at `lib/compiler/compiler-world-v2.gaf`. Normal `goml check`, `goml build`, and `goml test` commands only read that executable-relative world. A missing world is an incomplete installation and must not trigger an implicit source build.

## Advance stage0

The release is built by the previous Linux amd64 release as stage0; native macOS bootstrap uses the macOS arm64 archive of that same release. After publishing, keep `release_version` set to the published version and take both platform checksums from `SHA256SUMS`, then advance stage0:

```sh
linux_sha256=LINUX_AMD64_SHA256_FROM_SHA256SUMS
darwin_sha256=DARWIN_ARM64_SHA256_FROM_SHA256SUMS
just set-bootstrap-stage0 "$release_version" "$linux_sha256" "$darwin_sha256"
just bootstrap
```

Commit the updated `bootstrap/stage0.env` before using language features that the previous stage0 cannot compile. The next release workflow requires stage0 to match the latest published release.

For an offline bootstrap, download the pinned archive for the host and run, for example on macOS arm64:

```sh
GOML_STAGE0_ARCHIVE=/path/to/goml-X.Y.Z-darwin-arm64.tar.gz just bootstrap
```
