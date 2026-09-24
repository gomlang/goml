# Unicode data

All public tables are pinned to Unicode 15.0.0. The seven tables_*.gom files are
generated from the Go 1.26 unicode package by tools/generate_unicode_tables.go.
The generator rejects any other Unicode version, sorts named tables and emits
canonical GoML using stage2/bin/gomlfmt (overridable with GOMLFMT).

From the repository root with Go 1.26 on PATH:

```sh
go run tools/generate_unicode_tables.go --check
go run tools/generate_unicode_tables.go --write
```

Check mode compares generated, formatted output without modifying sources and
runs in the script CI gate. Write mode regenerates the seven derived files.
The existing full casefold.gom uses tools/generate_unicode_casefold.py instead;
it requires Python Unicode data 15.0.0 and supports multi-scalar folds. Both
generators support non-mutating `--check`, canonical formatting and the GOMLFMT
override; both checks run in the script CI gate. Regenerate full folds with
`python3 tools/generate_unicode_casefold.py --write`, or verify them with `--check`.
Invoking that Python script without arguments retains its historical write mode.

Range records are lowercase hexadecimal: six digits each for inclusive low and
high bounds, then eight digits for a nonzero u32 stride (20 characters total).
Case records contain a code point and its upper/lower/title mappings, each six
digits (24 total). Simple-fold records contain a code point and its successor
(12 total). Records are sorted; lookup uses binary search without native calls.
The encoding is private, not a supported serialization format.

The Unicode package has no standard-package dependencies. Its private output
helpers assemble validated string fragments or hexadecimal digits into builtin
byte vectors, using the existing UTF-8 string boundary primitive to finish.
This keeps text -> unicode acyclic without moving public Unicode types or
duplicating the versioned tables. The full-casefold generator uses these same
helpers; its mapping data is unchanged.

The external standard-algorithms consumer compares classifications and simple
case mappings over every valid scalar to Go, and checks all named range tables.
Go calls occur only in generation and tests. See ../THIRD_PARTY_NOTICES for the
upstream Go copyright and license notice.
