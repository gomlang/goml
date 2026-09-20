# GNU interoperability inputs

`gnu-cases.json` preserves the original eight edge cases and 40 deterministic input pairs (seed 20260920) from the former Python driver. It contains only input text; every test invocation obtains fresh independent results from installed GNU `diff` and `patch`.

The native GoML test produces unified patches using both Myers and Hirschberg algorithms and asks GNU `patch` to apply each. It also parses and applies fresh GNU `diff -u` output and compares the bytes to the target text. Empty documents, missing final newlines, CRLF, tabs, CJK, emoji, blank lines and separated hunks are covered.
