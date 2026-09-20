import argparse
import array
import hashlib
import json
from concurrent.futures import ThreadPoolExecutor
from pathlib import Path
import urllib.request

ROOT = Path(__file__).resolve().parent
DATA = ROOT / 'data'
MANIFEST = json.loads((DATA / 'manifest.json').read_text())
SOURCES = {}
MAX_SOURCE_BYTES = 16777216


def fetch_source(item):
    name, metadata = item
    request = urllib.request.Request(metadata['url'], headers={'Cache-Control': 'no-cache', 'Pragma': 'no-cache'})
    with urllib.request.urlopen(request, timeout=30) as response:
        raw = response.read(MAX_SOURCE_BYTES + 1)
    if len(raw) > MAX_SOURCE_BYTES:
        raise RuntimeError('Unicode source exceeds download limit: ' + name)
    if hashlib.sha256(raw).hexdigest() != metadata['sha256']:
        raise RuntimeError('Downloaded Unicode checksum mismatch: ' + name)
    return name, raw


def download_sources():
    SOURCES.clear()
    selected = [(name, item) for name, item in MANIFEST['sources'].items() if name != 'LICENSE.txt']
    with ThreadPoolExecutor(max_workers=4) as executor:
        downloaded = dict(executor.map(fetch_source, selected))
    SOURCES.update(downloaded)


def source(name):
    if name not in SOURCES:
        raise RuntimeError('Unicode source has not been downloaded in this run: ' + name)
    return SOURCES[name].decode('utf-8')


def rows(name):
    for line in source(name).splitlines():
        body = line.split('#', 1)[0].strip()
        if body:
            fields = [field.strip() for field in body.split(';')]
            interval = fields[0].split('..')
            yield int(interval[0], 16), int(interval[-1], 16), fields[1:]


def verify_sources():
    for name, item in MANIFEST['sources'].items():
        if name == 'LICENSE.txt':
            if hashlib.sha256((DATA / name).read_bytes()).hexdigest() != item['sha256']:
                raise RuntimeError('Unicode license checksum mismatch')
        else:
            source(name)


def generate():
    download_sources()
    verify_sources()
    properties = array.array('I', [512 << 17]) * 0x110000
    output = ['package unicode_text;', '', 'pub const UNICODE_VERSION: string = "16.0.0";', '']
    def assign(lo, hi, value, mask):
        for cp in range(lo, hi + 1):
            properties[cp] = (properties[cp] & ~mask) | value
    def flag(lo, hi, bit):
        for cp in range(lo, hi + 1):
            properties[cp] |= bit << 17
    for name, prefix, shift, bits, default in [
        ('GraphemeBreakProperty.txt', 'G', 0, 4, 'Other'),
        ('WordBreakProperty.txt', 'W', 4, 5, 'Other'),
        ('LineBreak.txt', 'L', 9, 6, 'XX'),
    ]:
        entries = list(rows(name))
        names = [default] + sorted({fields[0] for _, _, fields in entries} - {default})
        values = {name: number for number, name in enumerate(names)}
        for name, number in values.items():
            output.append(f'const {prefix}_{name.upper()}: isize = {number};')
        output.append('')
        for lo, hi, fields in entries:
            assign(lo, hi, values[fields[0]] << shift, ((1 << bits) - 1) << shift)
    indic = {'None': 0, 'Consonant': 1, 'Extend': 2, 'Linker': 3}
    for name, number in indic.items():
        output.append(f'const I_{name.upper()}: isize = {number};')
    output.append('')
    for lo, hi, fields in rows('DerivedCoreProperties.txt'):
        if fields[0] == 'InCB':
            assign(lo, hi, indic[fields[1]] << 15, 3 << 15)
        if fields[0] == 'Default_Ignorable_Code_Point':
            flag(lo, hi, 1)
    flags = {'ZERO': 1, 'PICTOGRAPHIC': 2, 'EMOJI': 4, 'EMOJI_PRESENTATION': 8,
             'WIDE': 16, 'AMBIGUOUS': 32, 'EAST_ASIAN': 64, 'INITIAL_QUOTE': 128,
             'FINAL_QUOTE': 256, 'UNASSIGNED': 512, 'MARK': 1024, 'WORD': 2048}
    for name, number in flags.items():
        output.append(f'const F_{name}: isize = {number};')
    output.append('')
    start = None
    for line in source('UnicodeData.txt').splitlines():
        fields = line.split(';'); cp = int(fields[0], 16); category = fields[2]
        if fields[1].endswith(', First>'):
            start = cp
            continue
        lo = start if fields[1].endswith(', Last>') else cp
        start = None
        assign(lo, cp, 0, 512 << 17)
        if category in {'Mn', 'Mc', 'Me', 'Cf', 'Cc', 'Zl', 'Zp'}:
            flag(lo, cp, 1)
        if category.startswith(('L', 'N')):
            flag(lo, cp, 2048)
        if category in {'Mn', 'Mc'}:
            flag(lo, cp, 1024)
        if category == 'Pi':
            flag(lo, cp, 128)
        if category == 'Pf':
            flag(lo, cp, 256)
    for lo, hi, fields in rows('EastAsianWidth.txt'):
        if fields[0] in {'W', 'F'}:
            flag(lo, hi, 16)
        if fields[0] == 'A':
            flag(lo, hi, 32)
        if fields[0] in {'W', 'F', 'H'}:
            flag(lo, hi, 64)
    for lo, hi, fields in rows('emoji-data.txt'):
        bit = {'Extended_Pictographic': 2, 'Emoji': 4, 'Emoji_Presentation': 8}.get(fields[0], 0)
        if bit:
            flag(lo, hi, bit)
    packed = []
    lo = 0
    while lo < len(properties):
        hi = lo
        while hi + 1 < len(properties) and properties[hi + 1] == properties[lo]:
            hi += 1
        packed.append(f'{lo:06x}{hi:06x}{properties[lo]:08x}')
        lo = hi + 1
    output.append('const PROPERTY_DATA: string = r"' + '\n'.join(packed) + '";')
    return '\n'.join(line + ('\n' if line.startswith('const ') and not line.startswith('const PROPERTY_DATA') else '') for line in output).replace('\n\n\n', '\n\n') + '\n'


def main():
    parser = argparse.ArgumentParser()
    parser.add_argument('--check', action='store_true')
    args = parser.parse_args()
    generated = generate()
    target = ROOT / 'properties.gom'
    if args.check:
        if target.read_text() != generated:
            raise RuntimeError('Generated Unicode tables differ; run generate.py')
        print('Unicode 16.0.0 sources freshly downloaded; hashes and generated tables verified')
    else:
        target.write_text(generated)


if __name__ == '__main__':
    main()
