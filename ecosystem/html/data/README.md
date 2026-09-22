# HTML named-character references

`entities.json` retains 2,125 semicolon-terminated HTML5 names from CPython 3.12.3
`html.entities.html5`, with terminal semicolons removed. Names without semicolons
are not a separate dataset here; this file alone does not define HTML's legacy
semicolon-optional matching rules. The PSF license is in
[LICENSE.entities.txt](../LICENSE.entities.txt).

SHA-256: `ea49a2b75ac0cd3028224804e1925509528ddadf3271131545b546f740c71192`.

From `ecosystem/markdown/tools`, run `../../../stage2/bin/goml run -- generate ..`
to generate the HTML lookup and Markdown compatibility wrapper; use `check` to
verify both. The generator validates the checksum, entry count, names and values.
The HTML library tests also compare every public lookup with this independent data.

`legacy.json` contains the 106 semicolon-optional names transcribed from Go 1.26.8
`src/html/entity.go`. SHA-256:
`a6a154ef2aceab185d73f76c7afc646cc3cca813de4fc0755dfb91146fed3912`.
The same native generator verifies this input and generates `legacy.gom`.
Its BSD license is retained in [LICENSE.go.txt](../LICENSE.go.txt). The numeric C1
replacement values were cross-checked against that toolchain's `html/escape.go`.
