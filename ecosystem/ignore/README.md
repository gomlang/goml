# ignore

`ecosystem::ignore` is a GoML implementation of byte-oriented glob matching, hierarchical Gitignore rules, and bounded filesystem traversal. The matcher and traversal policy are implemented in GoML. Filesystem access uses the standard library's Linux amd64 descriptor APIs. No external matching library or Git subprocess is needed at runtime.

Add `"ecosystem::ignore" = "0.1.0"` under your module's `[dependencies]`. `ecosystem::tempfile` is used by the black-box filesystem tests.

```gom
use ecosystem::ignore;
use std::context;

fn sources(root: string) -> Result[Vec[string], ignore::Error] {
    let wanted = ignore::Glob::new("**/*.gom")?;
    let options = ignore::WalkOptions {
        exclude: Option::Some(|relative, _| relative == "_artifact"),
        ..ignore::WalkOptions::standard()
    };
    let walker = ignore::Walker::new(root, options)?;
    defer { walker.close(); };
    let result: Vec[string] = Vec::new();
    while let Some(entry) = walker.next(context::Context::background()) {
        let entry = entry?;
        if entry.file_type().is_file() && wanted.is_match(entry.relative_path())? {
            result.push(entry.path());
        }
    }
    Result::Ok(result)
}
```

## Matching

`Glob::new(pattern)` compiles a full-path matcher. `Glob::with_options(pattern, limits, ascii_case_insensitive)` adds bounds and optional ASCII case folding. `is_match(path)` returns a checked result. `GlobSet::new(patterns)` compiles several patterns; `matches(path)` returns all matching indices in input order, including duplicates. Inputs and returned vectors cannot mutate a compiled set.

Patterns support literals, backslash escapes, `?`, `*`, bracket ranges, `!`/`^` class negation, literal initial `]`, and the C-locale POSIX classes `alnum`, `alpha`, `blank`, `cntrl`, `digit`, `graph`, `lower`, `print`, `punct`, `space`, `upper`, and `xdigit`. Wildcards and classes operate on UTF-8 **bytes**; `??` matches `é`, while `?` does not. Literal Unicode names work. `/` is the only separator; backslashes escape pattern bytes and are ordinary bytes in paths. Dotfiles receive no special treatment. There is no brace expansion, extglob, shell expansion, Unicode normalization, or locale-dependent folding.

`*` and `?` do not cross separators. A run of two or more stars occupying a complete path component is recursive: `**/x` matches `x` and any descendant `x`; `a/**/b` permits zero or more intervening directories; `a/**` matches contents below `a`. Other star runs behave like `*`. Compilation rejects invalid standalone globs. Matching uses bounded dynamic programming rather than recursive backtracking; memory is linear in path length.

`Gitignore::new()` creates an immutable ruleset. `with_source(base, source_name, text)` returns a new ruleset with a `.gitignore` source rooted at a normalized relative directory. `with_excludes(source_name, text)` adds explicitly provided low-priority global or repository excludes. Deeper `.gitignore` sources override shallower ones regardless of insertion order; later rules at the same depth win. Multiple excludes sources follow their insertion order. `with_options(limits, ascii_case_insensitive)` creates a configured ruleset.

The parser handles CRLF, an initial UTF-8 BOM, comments, blank lines, escaped initial `#`/`!`, escaped trailing spaces, negation, leading `/` anchors, directory-only trailing `/`, and patterns scoped relative to each ignore file. Unanchored patterns without `/` match basenames at any depth. Malformed Gitignore glob patterns are retained as nonmatching rules, as with Git; size-limit failures remain errors.

`matched(relative_path, is_directory)` returns a `Match` with `is_ignored()`, `is_whitelisted()`, `rule()`, and `inherited_from()`. A `Rule` exposes source name, 1-based line, original pattern, base, negation, and directory-only status. Every ancestor is checked before the target: a negation cannot reinclude a file beneath an ignored directory. The explicitly supplied root `""` or `"."` is always included. Absolute paths, NUL, and `..` components are rejected; redundant separators and `.` components are normalized.

An explicit trailing slash is preserved for Git's query semantics. For example, with `build/*`, `matched("build", true)` permits entering that directory, while `matched("build/", true)` matches the empty suffix after `/`, as `git check-ignore build/` does. Filesystem walking uses directory names without trailing slashes. Case-insensitive matching folds pattern bytes only; source bases continue to identify exact directory paths.

## Walking

`Walker::new(root, options)` accepts an explicit file or directory root. `next(context)` yields `Option[Result[Entry, Error]]` in deterministic, sorted, depth-first preorder. Metadata and filesystem errors appear as entries. The root is emitted before children or ignore files are read. `Entry` exposes absolute path, root-relative path, basename, effective file type, depth, and whether the original node is a symlink.

Default options read the root and descendant `.gitignore` files, root `.git/info/exclude`, and prune `.git` entries. Linked worktrees and separate Git directories are supported: a bounded `gitdir:` marker is resolved relative to the scan root, and an optional `commondir` is resolved relative to that Git directory before loading `info/exclude`. Absolute metadata paths also work. Malformed pointers and inaccessible metadata produce errors. No ancestor ignore files, user global Git configuration, Git index, or tracked-file state are read. Call `with_excludes` explicitly for global patterns.

Set `gitignore: false` to disable automatic Git metadata and ignore-source loading. Explicit `initial` rules still apply. `skip_git: false` independently allows visiting `.git` entries. `exclude: Some(callback)` receives each non-root relative path and effective directory flag; returning `true` skips that node and prunes its descendants. It takes precedence over ignore rules. Excluded directory ignore files are never opened. `skip_current_dir()` prunes the directory most recently yielded; `close()` immediately ends iteration and releases queued work.

The default `Symlinks::Skip` emits links without following them. `Symlinks::Follow` uses target types and follows directory links; canonical ancestor paths detect symbolic-link cycles and return `ErrorKind::Cycle`. Separate links to the same non-ancestor directory are each traversed. Following links can leave the scan root. This is ordinary filesystem traversal, not a filesystem sandbox or snapshot: concurrent renames and mount changes can alter results. Bind-mount cycles are bounded by depth/work limits rather than inode identity. Symlink ignore files and symlink `.git` markers are not loaded.

Ignore files are read in bounded chunks with `O_NOFOLLOW` and `O_NONBLOCK`; nonregular sources are rejected. Directory entries are fetched in bounded kernel batches, then accumulated up to the configured bound and sorted. Descriptors are closed before every `next` return, including errors and cancellation. No descriptor survives abandonment of the iterator. Filenames and source text must be valid UTF-8; undecodable names report an error. An I/O or source error skips the affected entry or subtree and permits later siblings; callers requiring complete results should propagate the first error. Cancellation emits one cancellation error and permanently ends iteration. Cancellation is checked between filesystem calls and cannot interrupt an already blocked kernel call.

`Walker` copies share one cursor and must be used sequentially, including its callbacks. Immutable `Glob`, `GlobSet`, `Gitignore`, `Rule`, and `Match` values can be shared across tasks. Use separate walkers for independent scans.

`walk_parallel(root, options, context, workers, callback)` streams a single traversal through a queue of `workers` entries to exactly that many callback workers (1–1024). It bounds callback concurrency and joins every callback before returning. Callbacks receive an entry and a child cancellation context. The successful count includes all callback-processed entries, including the root. Traversal and ignore decisions remain serial and ordered; callback execution and completion order are unspecified. A callback or traversal error cancels further admission, waits for in-flight callbacks, and returns an error, preferring an observed callback error over derivative cancellation. Callbacks must cooperate with cancellation; a callback that never returns prevents completion. There is no panic recovery.

## Bounds and validation

`Limits::standard()` allows 4,096 bytes per pattern, 32,768 per relative path, 100,000 rules, 1 MiB per source, and 16,777,216 matching work units per query. Matching work includes ancestor rule scans and character-class range checks. Work limits are shared across a complete Gitignore or GlobSet query. `WalkOptions` defaults to depth 256, 100,000 entries per directory, and 1,000,000 visited or queued entries combined. `max_depth` is an intentional cutoff; other exhausted limits produce recoverable errors. Walker limits also constrain initial rules and automatically loaded sources. Rulesets use immutable flat rule arrays; extending a ruleset copies that array, and matching scans applicable rules rather than compiling a multi-pattern automaton.

From the repository root:

```sh
python3 ecosystem/verify.py ignore
```

The suite covers matching, rule explanations and precedence, path validation, aliases, worktree metadata, symlink loops, malformed or oversized sources, deterministic traversal, cancellation, resource bounds, immutable concurrent matching, and bounded parallel callbacks. The independent versioned consumer exercises the exported API. `interop.py` compares 200 generated rule hierarchies and 9,400 path queries with isolated real Git repositories, including the distinction between explicit trailing-slash queries and directory traversal. `race.py` runs all library tests under the Go race detector. CI integration is intentionally left to the parent project.

Git semantics are based on the [Gitignore documentation](https://git-scm.com/docs/gitignore), [repository layout](https://git-scm.com/docs/gitrepository-layout), and [git check-ignore](https://git-scm.com/docs/git-check-ignore). The JSON consumer's directory requests append a slash to reproduce the exact paths used by the Git oracle.
