import calendar
import datetime as dt
import hashlib
import random
import subprocess
from pathlib import Path
from zoneinfo import ZoneInfo


ROOT = Path(__file__).resolve().parent
FIXTURES = ROOT / "fixtures"
BINARY = ROOT.parent / "consumers/datetime/_artifact/bin/datetime"
UTC = dt.timezone.utc
EPOCH = dt.datetime(1970, 1, 1)


def timestamp(value):
    difference = value.replace(tzinfo=None) - EPOCH
    return difference.days * 86400 + difference.seconds


def iso(value, nanos=None):
    if nanos is None:
        nanos = value.microsecond * 1000
    result = f"{value.year:04d}-{value.month:02d}-{value.day:02d}T{value.hour:02d}:{value.minute:02d}:{value.second:02d}"
    return result + ("." + f"{nanos:09d}".rstrip("0") if nanos else "")


def add_months(date, months, policy):
    index = (date.year - 1) * 12 + date.month - 1 + months
    if not 0 <= index < 9999 * 12:
        return "error"
    year, month = divmod(index, 12)
    year += 1
    month += 1
    last = calendar.monthrange(year, month)[1]
    if date.day > last and policy == "reject":
        return "error"
    try:
        value = dt.date(year, month, min(date.day, last))
        if policy == "carry":
            value = dt.date(year, month, 1) + dt.timedelta(days=date.day - 1)
        return value.isoformat()
    except (ValueError, OverflowError):
        return "error"


def main():
    for line in (FIXTURES / "SHA256SUMS").read_text().splitlines():
        digest, name = line.split("  ")
        if hashlib.sha256((FIXTURES / name).read_bytes()).hexdigest() != digest:
            raise AssertionError(f"fixture checksum mismatch: {name}")
    names = ["America/New_York", "Australia/Lord_Howe", "Pacific/Apia", "Europe/Dublin", "Asia/Kathmandu", "Etc/UTC", "Synthetic/Shifted", "Synthetic/Extreme", "Synthetic/Julian", "Synthetic/NegativeDst", "Synthetic/AllYear"]
    zones = {name: ZoneInfo.from_file((FIXTURES / name).open("rb"), key=name) for name in names}
    rng = random.Random(221937)
    cases = []
    expected = []

    def add(command, output):
        cases.append(command)
        expected.append(output)

    dates = [dt.date(1, 1, 1), dt.date(9999, 12, 31), dt.date(1900, 2, 28), dt.date(2000, 2, 29), dt.date(2021, 1, 1)]
    dates += [dt.date.fromordinal(rng.randint(1, dt.date.max.toordinal())) for _ in range(1800)]
    for date in dates:
        days = rng.randint(-4000, 4000)
        months = rng.randint(-120000, 120000) if rng.random() < 0.1 else rng.randint(-100, 100)
        try:
            shifted = (date + dt.timedelta(days=days)).isoformat()
        except (ValueError, OverflowError):
            shifted = "error"
        week = date.isocalendar()
        output = [str((date - dt.date(1970, 1, 1)).days), str(date.timetuple().tm_yday), str(date.isoweekday()), str(week.year), str(week.week), str(week.weekday), shifted]
        output += [add_months(date, months, policy) for policy in ("reject", "clamp", "carry")]
        add(f"date\t{date.isoformat()}\t{days}\t{months}", "\t".join(output))

    seconds = [-5364662400, -2208988800, -2147483648, -1, 0, 1, 2147483647, 2147483648, 253402000000]
    seconds += [rng.randint(-5000000000, 50000000000) for _ in range(350)]
    for name, zone in zones.items():
        if name == "Synthetic/AllYear":
            zone = dt.timezone(dt.timedelta(hours=-4), "EDT")
        selected_seconds = [value for value in seconds if not name.startswith("Synthetic/") or value >= 0]
        if name.startswith("Synthetic/"):
            for year in (2000, 2024, 2500):
                start = timestamp(dt.datetime(year, 1, 1))
                selected_seconds += [start + delta for delta in (-691200, -86400, -1, 0, 1, 86400, 691200)]
        for second in selected_seconds:
            nanos = rng.randrange(1000000000)
            utc = (EPOCH + dt.timedelta(seconds=second)).replace(tzinfo=UTC)
            local = utc.astimezone(zone)
            output = f"{iso(utc, nanos)}Z\t{iso(local, nanos)}\t{int(local.utcoffset().total_seconds())}\t{local.tzname()}\t{str(name == "Synthetic/AllYear" or bool(local.dst())).lower()}"
            add(f"instant\t{second}\t{nanos}\t{name}", output)

    centers = {
        "America/New_York": [dt.datetime(2024, 3, 10, 2), dt.datetime(2024, 11, 3, 1), dt.datetime(2500, 3, 14, 2), dt.datetime(1883, 11, 18, 12)],
        "Australia/Lord_Howe": [dt.datetime(2024, 4, 7, 1, 30), dt.datetime(2024, 10, 6, 2)],
        "Pacific/Apia": [dt.datetime(2011, 12, 30, 12)],
        "Europe/Dublin": [dt.datetime(1916, 10, 1, 2, 30), dt.datetime(2024, 10, 27, 1)],
        "Asia/Kathmandu": [dt.datetime(1986, 1, 1, 0)],
        "Etc/UTC": [dt.datetime(1969, 12, 31, 23, 59)],
    }
    for name, zone in zones.items():
        if name == "Synthetic/AllYear":
            zone = dt.timezone(dt.timedelta(hours=-4), "EDT")
        locals = [center + dt.timedelta(minutes=minute) for center in centers.get(name, [dt.datetime(2024, 1, 1), dt.datetime(2024, 3, 1), dt.datetime(2024, 12, 31)]) for minute in range(-125, 126, 5)]
        locals += [EPOCH + dt.timedelta(seconds=rng.randint(0 if name.startswith("Synthetic/") else -5000000000, 50000000000)) for _ in range(80)]
        for local in locals:
            nanos = rng.randrange(1000000000)
            valid = set()
            for fold in (0, 1):
                candidate = local.replace(tzinfo=zone, fold=fold).astimezone(UTC)
                if candidate.astimezone(zone).replace(tzinfo=None) == local:
                    valid.add(timestamp(candidate))
            values = sorted(valid)
            if len(values) == 0:
                output = "gap"
            elif len(values) == 1:
                output = f"unique\t{values[0]}\t{nanos}"
            else:
                output = f"fold\t{values[0]}\t{values[1]}\t{nanos}"
            add(f"local\t{iso(local, nanos)}\t{name}", output)

    for _ in range(250):
        second = rng.randint(-5000000000, 50000000000)
        offset = rng.randint(-1439, 1439)
        nanos = rng.randrange(1000000000)
        utc = (EPOCH + dt.timedelta(seconds=second)).replace(tzinfo=UTC)
        value = utc.astimezone(dt.timezone(dt.timedelta(minutes=offset)))
        absolute = abs(offset)
        suffix = "Z" if offset == 0 else f"{'-' if offset < 0 else '+'}{absolute // 60:02d}:{absolute % 60:02d}"
        canonical = iso(value, nanos) + suffix
        add(f"rfc\t{canonical}", f"{second}\t{nanos}\t{canonical}")

    synthetic = [index for index, case in enumerate(cases) if "Synthetic/" in case and "Synthetic/AllYear" not in case]
    reference = ROOT / "_artifact" / "go-time-oracle"
    reference.parent.mkdir(parents=True, exist_ok=True)
    subprocess.run(["go", "build", "-o", str(reference), str(ROOT / "reference/main.go")], check=True, timeout=60)
    oracle = subprocess.run([str(reference), str(FIXTURES)], input="\n".join(cases[index] for index in synthetic) + "\n", text=True, capture_output=True, check=True, timeout=60).stdout.splitlines()
    if len(oracle) != len(synthetic):
        raise AssertionError("Go time oracle output count mismatch")
    for index, value in zip(synthetic, oracle):
        expected[index] = value

    result = subprocess.run([str(BINARY), "--oracle", str(FIXTURES)], input="\n".join(cases) + "\n", capture_output=True, text=True, timeout=60, check=True)
    actual = result.stdout.splitlines()
    if len(actual) != len(expected):
        raise AssertionError(f"expected {len(expected)} lines, got {len(actual)}: {result.stderr}")
    for index, (actual, expected) in enumerate(zip(actual, expected)):
        if actual != expected:
            raise AssertionError(f"case {index}: {cases[index]}\nGoML: {actual!r}\nReference: {expected!r}")
    version = (FIXTURES / "VERSION").read_text().strip()
    print(f"datetime: {len(cases)} Python datetime/zoneinfo, Go time and RFC9636 all-year DST comparisons passed against checksum-pinned tzdata {version} and synthetic fixtures")


if __name__ == "__main__":
    main()
