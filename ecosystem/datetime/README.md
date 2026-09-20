# datetime

A GoML calendar, nanosecond timestamp, and IANA timezone library. Calendar algorithms, checked arithmetic, text handling, Serde implementations, TZif decoding, POSIX rule evaluation, and local-time resolution are implemented in GoML. No Go time adapter is used by the library. The standard library provides the clock and bounded Linux file I/O.

## Types and arithmetic

- `Date`: validated proleptic Gregorian dates, years 1 through 9999; epoch-day and ordinal conversions; Monday-based weekday numbers; ISO week dates and inverse construction; checked day, month, and year arithmetic.
- `Time`: hour/minute/second/nanosecond validation with nanosecond precision. Leap seconds and `24:00:00` are rejected.
- `LocalDateTime`: a date and wall-clock time with no inferred offset. `assume_utc` and `assume_offset` make interpretation explicit.
- `UtcInstant`: POSIX epoch seconds plus nonnegative fractional nanoseconds, covering years 1 through 9999 UTC. Negative epoch milliseconds and nanoseconds use floor normalization: `-1 ns` is second `-1` plus `999999999 ns`. Conversion back to signed 64-bit nanoseconds is checked; seconds and milliseconds cover the full calendar range. `now` uses `std::time::SystemTime`.
- `Duration`: signed elapsed time represented by floor seconds plus nonnegative fractional nanoseconds. Normalized seconds are bounded to `-315537897600..315537897600`; addition and negation are checked. `whole_seconds` truncates toward zero, while `floor_seconds` exposes the normalized component. `checked_nanoseconds` rejects values outside `i64`.
- `Offset`: second-precision offsets in `-25:59:59..25:59:59`, including historical non-minute offsets. TZif decoding additionally enforces RFC 9636's narrower negative bound.
- `OffsetDateTime`: a checked instant and offset whose local date also fits the supported range. Equality includes both fields; compare `instant()` values when only chronology matters.
- `TimeZone` and `ZonedDateTime`: immutable parsed transition tables, future rules, zone designation/DST metadata, conversions, and calendar arithmetic that returns a local-time resolution result.

Month/year arithmetic requires `InvalidDay::Reject`, `Clamp`, or `Carry`. For example, January 31 plus one month in 2023 errors, becomes February 28, or becomes March 3 respectively. Carry starts at the target month's first day and adds the original day minus one; this rule also applies to negative month deltas.

`ZonedDateTime::checked_add(Duration::days(1)?)` adds 86400 elapsed seconds. `checked_add_days(1)` retains local clock time, then resolves the resulting date. A spring transition can therefore make the latter span 23 elapsed hours.

```goml
use ecosystem::datetime as dt;

fn appointment() -> Result[dt::ZonedDateTime, dt::Error] {
    let zone = dt::TimeZone::load("America/New_York")?;
    let local = dt::LocalDateTime::parse("2024-11-03T01:30:00.123456789")?;
    zone.resolve_local(local)?.resolve(dt::FoldPolicy::Later)
}
```

`resolve_local` returns `Unique`, `Ambiguous(earlier, later)`, or `Nonexistent`. Candidates retain the original nanoseconds. `FoldPolicy` selects an earlier/later instant or rejects ambiguity; gaps always remain errors when resolving. A fabricated timezone with more than two candidates returns `Unsupported`. Out-of-calendar-range conversions return `OutOfRange`, not a DST gap.

## Text and serialization

`Date::parse`, `Time::parse`, and `LocalDateTime::parse` accept their exact ISO-shaped representations, requiring seconds and at most nine fractional digits. `OffsetDateTime::parse_rfc3339` additionally requires a known numeric offset or `Z`; RFC-permitted lowercase `t`/`z` are accepted. It rejects leap seconds, extra precision, trailing input, offset hours above 23, second-precision offsets, and unknown local offset `-00:00`. `to_rfc3339` rejects offsets the format cannot represent. `to_string` on an offset datetime remains lossless for historical second offsets.

Local and offset datetimes support formatting directives `%Y`, `%m`, `%d`, `%H`, `%M`, `%S`, `%f` (nine digits), `%j`, `%u`, `%G`, `%V`, `%F`, `%T`, `%z`, `%:z`, and `%%`. Literal Unicode is preserved. Offset directives on a local datetime and unknown directives return errors. Arbitrary-pattern parsing, localized month/day names, and locale databases are not implemented.

Serde validates values on decoding:

| Type | Representation in all formats |
| --- | --- |
| Date, Time, LocalDateTime | Canonical string |
| UtcInstant | Canonical UTC RFC3339 string |
| Offset | Signed string, including seconds when needed |
| OffsetDateTime | Tuple of instant string and offset string |
| Duration | Tuple of normalized floor seconds and nonnegative fractional nanoseconds |

JSON and Bincode tests cover the independent registry boundary. `TimeZone` and `ZonedDateTime` do not implement implicit Serde lookup: persist an instant and zone name with the application's selected timezone-data version, then reload explicitly.

## Timezone data

`TimeZone::from_tzif(name, bytes)` reads TZif versions 1–4. It validates header counts, reserved bytes, sorted transition times, type indexes, flags, designation termination/UTF-8, footer framing, and final-transition/footer agreement. Version 2 rejects the signed/extended transition-time syntax reserved by TZif for version 3 or later. Transition lookup uses binary search; timestamps before the first transition use type zero as specified by RFC 9636.

Future rules support `Jn`, zero-based `n`, `Mm.w.d`, quoted designations, explicit/default daylight offsets, negative DST, southern-hemisphere seasons, and signed transition times through `±167:59:59`. Annual daylight intervals may cross year boundaries or cover the whole year. `from_posix` provides the same rule engine without a file. Clock suffix extensions such as `/2u` and implementation-dependent DST rules with omitted start/end dates are rejected.

Files containing leap records are rejected: this library uses POSIX seconds, so `right/` leap-second datasets cannot be treated as ordinary zoneinfo. An empty future footer on a zone with transitions allows lookups only through the final recorded whole second; later timestamps return `Unsupported`. A zone with no transitions and no footer is fixed at type zero. The designation `-00` produces `Unsupported` when selected because TZif marks that interval's local time as unspecified. Future conversions therefore do not silently extend an unknown last offset.

`load` reads `/usr/share/zoneinfo`; `load_from` accepts an explicit trusted data directory. Names reject absolute paths, empty/dot/dot-dot components, NULs, backslashes, and characters outside ASCII letters/digits/`_+-/`. Normal distribution symlinks are followed, so the chosen directory is a trust boundary rather than a filesystem sandbox. File loading is bounded before parsing and closes its descriptor on success or failure. Both loading and decoding have a 16 MiB limit; transition counts are limited to 1,000,000, types to 256, designation storage to 65,536 bytes, and POSIX footers to 4,096 bytes. Existing zone handles retain their parsed snapshot when source files change. There is no process-global timezone mutation or mutable lookup cache.

System lookup depends on the installed timezone data and does not invent a version identifier: TZif has no database-version field. Applications needing reproducibility should ship a selected dataset and call `load_from` or `from_tzif`.

`fixtures/VERSION` records `2026c`, as reported by the source system's `tzdata.zi`. The six bundled zone files are compiled distribution data, with individual hashes in `fixtures/SHA256SUMS`; the Dublin file uses the distribution's positive-DST compatibility encoding. These are test fixtures, not a bundled global timezone database. The IANA database is [public-domain data](https://data.iana.org/time-zones/tz-link.html). `fixture_builder.py` deterministically regenerates the separately authored `Synthetic/` files and refreshes the fixture checksums without replacing the six IANA files.

## Verification and reference differences

Run from the repository root:

```sh
python3 ecosystem/verify.py datetime
```

The verification runs 18 library tests, the independent versioned consumer, cached-build checks, 8,140 calendar/timezone/reference cases, and all 18 tests under Go's race detector. The concurrent test shares one immutable zone across 12 workers performing local resolution and reverse conversion.

Python `datetime` validates Gregorian/ISO-week arithmetic and month policies. Python `zoneinfo` reads the exact bundled zone bytes for historical second offsets, negative epochs, DST gaps/folds, skipped dates, non-hour transitions, and future timestamps beyond explicit records. Go `time.LoadLocationFromTZData` is an independent oracle for synthetic POSIX cases at nonnegative epochs. RFC 9636 supplies fixed expected values for the all-year DST fixture. The Go reference executable is test tooling only.

Oracle selection follows the specifications, because the reference implementations also have edge cases:

- POSIX defines bare `n` as zero-based with leap days included, while `Jn` starts at one and excludes leap days. For `STD0DST,59,300` in leap year 2660, the end is October 27 at 02:00 daylight time (01:00 UTC). Installed Python 3.12.3 `zoneinfo` instead ends on October 26; `2660-10-26T14:40:17Z` must still have offset `+01:00`. Go's reference produces that specified result. See the [POSIX TZ rule definitions](https://pubs.opengroup.org/onlinepubs/9799919799/basedefs/V1_chap08.html).
- RFC 9636 explicitly defines `XXX3EDT4,0/0,J365/23` as permanently `EDT`, offset `-04:00`. Installed Go 1.25 selects unused `XXX` at `2024-01-01T00:00:00Z`; the corresponding fixture therefore uses the RFC's fixed result, including calendar boundaries. See [RFC 9636 §3.3.1](https://www.rfc-editor.org/rfc/rfc9636.html#section-3.3.1).
- Go 1.25's negative-epoch POSIX-tail calculation also differs at some year boundaries: `STD0DST,M1.1.0/-167,M12.5.0/167` at `1900-01-01T00:00:01Z` reports standard time despite lying inside the specified daylight interval. Synthetic Go comparisons use nonnegative epochs; historical negative-epoch comparisons remain covered by the actual IANA files and Python.

The library implements calendar and timezone policies, not leap-second/TAI arithmetic, alternative calendars, localized rendering, recurrence scheduling, or automatic timezone-data updates. Bounds failures, unsupported data, and parse failures are recoverable `Error` values.
