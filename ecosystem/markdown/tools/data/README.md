# HTML named-entity data

The canonical [entities.json](../../../html/data/entities.json) now belongs to
ecosystem::html. It contains all 2,125 semicolon-terminated HTML5 named references
from CPython 3.12.3, with terminal semicolons removed. The PSF license is retained
in [HTML's license](../../../html/LICENSE.entities.txt) and Markdown's original copy.

SHA-256: `ea49a2b75ac0cd3028224804e1925509528ddadf3271131545b546f740c71192`.

The native GoML tool validates the count, names, values and checksum, sorts the
entries and groups lookup functions by first ASCII byte. Its native test verifies
both HTML's generated lookup and Markdown's generated compatibility wrapper.
No reference data is obtained from the implementation being tested.
