# HTML named-entity data

`entities.json` contains all 2,125 semicolon-terminated HTML5 named character references from CPython 3.12.3 `html.entities.html5`, stripped of their terminal semicolons. This is the same independent source used to create the existing `entities.gom`, retained as data so regeneration no longer depends on Python. The Python Software Foundation source license is in [LICENSE.entities.txt](../../LICENSE.entities.txt).

SHA-256: `ea49a2b75ac0cd3028224804e1925509528ddadf3271131545b546f740c71192`.

The native GoML tool validates the case count, names, values and checksum, sorts the entries and groups their lookup functions by the first ASCII byte. Its native test checks exact generated output against `entities.gom`. No data is obtained from the implementation being tested.
