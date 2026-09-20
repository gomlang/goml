# Releasing goml

Releases use strict `vX.Y.Z` tags and currently publish Linux amd64 binaries.

The Go compatibility baseline is Go 1.26. Build and validate releases with Go 1.26.x; users need Go 1.26 or newer to compile generated programs and validate Go FFI. The archives do not bundle a Go toolchain.

The root [VERSION](../VERSION) file is authoritative. `goml`, `gomlc`, `gomlfmt`, `gomllsp`, and the VS Code extension must use the same version. The metadata helper `goml-go-meta` is packaged alongside them.

## Version policy

Each release must be one continuous SemVer step from the latest published release:

- patch: `0.1.0` to `0.1.1`
- minor: `0.1.0` to `0.2.0`
- major: `0.1.0` to `1.0.0`

Skipped steps such as `0.1.0` to `0.1.2` or `0.3.0` are rejected.

During the early bootstrap period, releases are limited to continuous `0.1.x` patch versions.

## Publish

Run from the repository root with the development prerequisites in the [README](../README.md#development), plus an authenticated GitHub CLI (`gh`) for release operations. Replace `X.Y.Z` with the next continuous version after the latest published release:

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
    ci_run_id="$(gh run list --repo lijunchen/goml --workflow CI --branch main --commit "$release_sha" --event push --limit 1 --json databaseId --jq '.[0].databaseId // empty')"
    test -n "$ci_run_id"
    gh run watch "$ci_run_id" --repo lijunchen/goml --exit-status
    git tag -a "v$release_version" -m "goml v$release_version"
    git push origin "v$release_version"
)
```

The block stops before tagging if the push fails, the CI run is not registered yet, or CI fails. If the run is not yet visible, wait for it to appear and rerun the block.

Main CI builds stage2 directly from the pinned stage0. Its checks cover compiler and driver tests, Go metadata and scripts, extension compilation and VSIX packaging, archive smoke tests, and the stage3 fixed point. Fixed-point verification compares the compiler and driver artifacts built by stage2 with a rebuild using stage3. The independent gomlgo suite runs separately. The Release workflow requires successful main CI for the tagged commit, verifies the version, previous release, and stage0, rebuilds stage2, and runs archive and LSP smoke tests before publishing.

Release archives use a complete toolchain prefix:

```text
goml-X.Y.Z-linux-amd64/
├── bin/
│   ├── goml
│   ├── gomlc
│   ├── gomlfmt
│   ├── goml-go-meta
│   └── gomllsp
└── lib/
    ├── builtin/
    │   ├── goml.toml
    │   ├── contract.gom
    │   ├── runtime.gom
    │   ├── impls.gom
    │   ├── intrinsics.gom
    │   ├── language.gom
    │   ├── derive.gom
    │   ├── ordering.gom
    │   └── numeric.gom
    ├── prelude/
    │   ├── goml.toml
    │   └── prelude.gom
    └── std/
        ├── goml.toml
        └── ...
```

The compiler resolves `lib` relative to its executable. The archive must preserve this layout exactly. `builtin`, `prelude`, and `std` are separate GoML projects; flat `builtin_*.gom` files must not be packaged.

Release archives contain the toolchain project sources and manifests, but do not contain `lib/compiler/compiler-world-v2.gaf`. After extracting an archive, installation must finalize it once with the binaries from that same archive:

```sh
prefix=/path/to/goml-X.Y.Z-linux-amd64
"$prefix/bin/goml" __toolchain-finalize --prefix "$prefix"
```

Finalization builds and validates the compiler world in a temporary path, then atomically installs it at `lib/compiler/compiler-world-v2.gaf`. Normal `goml check`, `goml build`, and `goml test` commands only read that executable-relative world. A missing world is an incomplete installation and must not trigger an implicit source build.

## Advance stage0

The release is built by the previous release as stage0. After publishing, keep `release_version` set to the published version and replace `SHA256_FROM_SHA256SUMS` with its archive checksum, then advance stage0:

```sh
archive_sha256=SHA256_FROM_SHA256SUMS
just set-bootstrap-stage0 "$release_version" "$archive_sha256"
just bootstrap
```

Commit the updated `bootstrap/stage0.env` before using language features that the previous stage0 cannot compile. The next release workflow requires stage0 to match the latest published release.

For an offline bootstrap, download the pinned archive and run:

```sh
GOML_STAGE0_ARCHIVE=/path/to/goml-X.Y.Z-linux-amd64.tar.gz just bootstrap
```
