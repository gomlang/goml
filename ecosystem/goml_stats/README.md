# goml_stats

A GoML project statistics tool using the `ecosystem::ignore` library for
hierarchical Git ignore rules and filesystem traversal. It reports source files,
physical lines, code lines, comment-only lines,
blank lines, UTF-8 bytes, test files, discovered modules and package directories.
Text tables and JSON reports support both interactive use and automation.

## Build and run

From the repository root:

```sh
python3 ecosystem/goml_stats/verify.py
ecosystem/goml_stats/_artifact/bin/cmd/goml_stats/goml_stats .
ecosystem/goml_stats/_artifact/bin/cmd/goml_stats/goml_stats gomlc --modules --packages
ecosystem/goml_stats/_artifact/bin/cmd/goml_stats/goml_stats ecosystem --files
ecosystem/goml_stats/_artifact/bin/cmd/goml_stats/goml_stats . --exclude gomlc/testdata --exclude gomlgo/testdata
ecosystem/goml_stats/_artifact/bin/cmd/goml_stats/goml_stats . --json --files > ecosystem/goml_stats/_artifact/source-stats.json
```

The verification command resolves versioned dependencies from an isolated local
registry snapshot and builds the tool. With the dependency available in a
configured registry, `goml build` inside this directory also works. The resulting
`goml_stats` executable can be copied onto `PATH`. Filesystem traversal currently
targets Linux amd64, matching the standard filesystem APIs.

## Options

```text
goml_stats [OPTIONS] [PATH]
```

`PATH` defaults to the current directory. It can be a directory or a single `.gom`
file. Options can appear before or after the path.

| Option | Behavior |
| --- | --- |
| `--files` | Add per-file rows, or the `files` array in JSON |
| `--modules` | Add module-directory totals to the text report |
| `--packages` | Add package-directory totals to the text report |
| `--json` | Emit JSON with totals, module and package summaries |
| `--exclude NAME_OR_PATH` | Exclude a basename anywhere or one root-relative path; repeatable |
| `--exclude=NAME_OR_PATH` | Alternative spelling for an exclusion |
| `--no-default-excludes` | Disable the built-in exclusions; explicit exclusions and Git ignore rules still apply |
| `--no-ignore` | Disable `.gitignore` and root `.git/info/exclude` loading |
| `-h`, `--help` | Show usage |
| `--` | End options, for example `goml_stats -- -example.gom` |

Default exclusions are `.git`, `.hg`, `.svn`, `.goml`, `_artifact`, `_bootstrap`,
`stage0`, `stage1`, `stage2`, `stage3`, `node_modules` and `__pycache__`. Matching
directories are pruned before reading their contents. Basenames match at any
depth; paths containing `/` match relative to the scan root. Exclusions can also
match files. Absolute paths, empty exclusions and paths escaping the root are
rejected. Explicit exclusions remain literal names or paths.

By default, the scanner reads `.gitignore` in the scan root and visited child
directories, plus the root `.git/info/exclude` at lower priority. These rules
support Git glob patterns, escaping, anchoring and negation. A negated file rule
cannot restore a file inside a pruned parent directory. Ancestor `.gitignore`
files and user-global Git configuration are not loaded. `--no-ignore` disables
these files; it does not disable the built-in or explicit exclusions. Explicit
exclusions take precedence over rule negation.

For a linked Git worktree, the root `.git` file and its `commondir` metadata
locate the shared `info/exclude`. This explicit Git metadata reference can point
outside the scanned directory; ancestor `.gitignore` files are still not loaded.

An explicitly supplied scan root is always visited, even when its name is on the
exclusion list or matches an ignore rule. A single-file scan bypasses ignore
rules. Symbolic links inside the tree are skipped, including dangling
links and directory loops. A symbolic link supplied as the root is rejected.
Only regular files with the case-sensitive `.gom` extension are counted; generated
Go files, compiler snapshots and other file types are ignored. Distinct hard-link
paths count as distinct files.

## Counting rules

- `lines = code + comments + blank`. A final unterminated line counts; a trailing
  line terminator does not add an extra line. LF, CRLF and lone CR are accepted.
  An empty file has zero lines and still contributes one file.
- A line containing code and an inline comment counts once as code. A line whose
  only non-whitespace content is a `//` comment counts as a comment. GoML does
  not support block comments.
- Strings, escaped quotes, characters, byte strings, raw string hash delimiters,
  interpolated expressions and `\\` multiline-string markers are recognized.
  Comment-looking text and empty physical lines inside a raw string count as
  code. Incomplete raw strings continue through the end of the file. The counter
  works on unfinished source and does not perform compiler syntax validation.
- `bytes` includes all source bytes, including whitespace, comments and line
  terminators. Invalid UTF-8 and filesystem errors fail the scan. No partial
  success report is emitted.
- `test_files` counts files ending in `_test.gom` or located beneath a directory
  named `tests`. This is a filename/path convention, not a count of `#[test]`
  declarations. It also applies when scanning a test file directly.
- Modules are directories containing a discovered regular `goml.toml` file,
  including modules with zero source files. Manifest contents are not parsed.
  Each source belongs to the closest discovered ancestor module; nested modules
  are not counted again in their parent. Sources outside these modules appear in
  `unassigned`. Ancestors outside the requested tree are not searched; a
  single-file scan therefore discovers no modules.
- Packages are distinct directories with at least one counted `.gom` file.
  Standalone fixtures and sources without a valid package declaration are still
  included. This measures source layout rather than compiler-resolved packages.

Paths in summaries are relative to the scanned directory, or the parent of a
single-file root. `.` identifies that base directory. The report root is absolute.
Files and groups are sorted lexicographically for reproducible output.

## JSON and library API

JSON schema version `1` includes `root`, `module_count`, `package_count`, `totals`,
`unassigned`, `modules` and `packages`. Every counts object contains `files`,
`test_files`, `lines`, `code`, `comments`, `blank` and `bytes`. Each group contains
`path` and `counts`. With `--files`, file entries also contain `path`, `package`,
`module` (a directory path or `null`) and `counts`.

The module path is `ecosystem::goml_stats`; the reusable root package exposes:

```text
count_source(source: string) -> Counts
scan(root: string, options: Options) -> Result[Report, string]
render_json(report: Report, include_files: bool) -> string
render_text(report: Report, include_modules: bool, include_packages: bool,
            include_files: bool) -> string
Options::new() -> Options
default_excludes() -> Vec[string]
normalize_exclude(value: string) -> Result[string, string]
```

`Options` has public `excludes: Vec[string]`, `use_default_excludes: bool`
and `use_gitignore: bool` fields. `Report`, `Group`, `FileStats` and `Counts` expose
their data as public fields. The scanner reads one source file at a time and retains counts and paths;
it does not retain all project source text.

Traversal inherits the ignore library's standard resource limits, including
100,000 entries per directory, 1,000,000 visited entries, 32,768 bytes per relative
path and 1 MiB per ignore file. Exceeding a limit fails the scan instead of
returning partial totals. The scanner disables the walker's depth cutoff so deep
files are not silently omitted.

Exit status is `0` for success, `1` for a scan/read failure and `2` for invalid
arguments. Errors go to stderr; JSON goes only to stdout.

## Verification

```sh
python3 ecosystem/goml_stats/verify.py
```

The verifier checks formatting, runs nine GoML tests, builds the executable and
runs CLI checks on temporary projects. Cases cover nested and empty modules,
unassigned sources, per-file and aggregate consistency, test-file detection,
literal exclusions, hierarchical Git ignore rules and negation, symbolic links,
special files, invalid UTF-8, argument errors, JSON
escaping, deterministic output and 48 generated source files with known counts.
Temporary fixtures, command logs and a verification report stay under
`_artifact/verification/`. Verification needs Python 3 and the GoML toolchain, with
no manual registry setup or external Python packages.
