import re
import subprocess
import sys
import unicodedata
from pathlib import Path
from urllib.parse import unquote, urlsplit

ROOT = Path(__file__).resolve().parent.parent


def markdown_files():
    result = subprocess.run(
        ['git', 'ls-files', '--cached', '--others', '--exclude-standard', '-z', '--', '*.md'],
        cwd=ROOT, check=True, capture_output=True, text=True,
    )
    return sorted({ROOT / name for name in result.stdout.split('\0') if name and (ROOT / name).is_file()})


def sections(text):
    prose = []
    blocks = []
    block = []
    fence = None
    for line in text.splitlines():
        marker = re.match(r'^\s{0,3}(`{3,}|~{3,})(.*)$', line)
        if fence is None:
            if marker:
                fence = marker.group(1)
                block = [line]
            else:
                prose.append(line)
        else:
            block.append(line)
            if marker and marker.group(1)[0] == fence[0] and len(marker.group(1)) >= len(fence) and not marker.group(2).strip():
                blocks.append('\n'.join(block))
                fence = None
    if fence is not None:
        raise ValueError('unclosed code fence')
    return '\n'.join(prose), blocks


def slug(heading):
    heading = re.sub(r'<[^>]+>', '', heading).lower()
    heading = ''.join(c for c in heading if c in '-_ ' or unicodedata.category(c)[0] in 'LNM')
    return heading.replace(' ', '-')


def headings(prose):
    return re.findall(r'^ {0,3}(#{1,6})\s+(.+?)(?:\s+#+)?\s*$', prose, re.MULTILINE)


def anchors(prose):
    result = set(re.findall(r'<a\s+(?:id|name)=["\']([^"\']+)["\']\s*>', prose))
    used = set()
    for _, title in headings(prose):
        base = slug(title)
        name = base
        index = 0
        while name in used:
            index += 1
            name = f'{base}-{index}'
        used.add(name)
        result.add(name)
    return result


def destinations(prose):
    prose = re.sub(r'(`+).*?\1', '', prose)
    inline = re.findall(r'\]\(<?([^\s)>]+)>?(?:\s+["\'][^\n]*["\'])?\)', prose)
    references = re.findall(r'^\s{0,3}\[[^\]]+\]:\s*<?([^\s>]+)>?', prose, re.MULTILINE)
    return inline + references


def main():
    errors = []
    documents = {}
    files = markdown_files()
    for path in files:
        try:
            documents[path] = sections(path.read_text())
        except (OSError, ValueError) as error:
            errors.append(f'{path.relative_to(ROOT)}: {error}')
    for path, (prose, blocks) in documents.items():
        relative = path.relative_to(ROOT)
        chinese = path.name.endswith('.zh-CN.md')
        english = path.with_name(path.name.replace('.zh-CN.md', '.md')) if chinese else path
        translation = english.with_name(english.stem + '.zh-CN.md')
        counterpart = english if chinese else translation
        if counterpart not in documents:
            errors.append(f'{relative}: missing counterpart {counterpart.name}')
        switch = f'[English]({english.name}) | [简体中文]({translation.name})'
        if switch not in '\n'.join(prose.splitlines()[:10]):
            errors.append(f'{relative}: missing language switch near the top')
        if not chinese and translation in documents:
            translated_prose, translated_blocks = documents[translation]
            if [level for level, _ in headings(prose)] != [level for level, _ in headings(translated_prose)]:
                errors.append(f'{relative}: translated heading structure differs')
            for anchor in sorted(anchors(prose) - anchors(translated_prose)):
                errors.append(f'{relative}: translation does not preserve heading anchor #{anchor}')
            if blocks != translated_blocks:
                errors.append(f'{relative}: translated code examples differ')
        for destination in destinations(prose):
            url = urlsplit(destination)
            if url.scheme or url.netloc:
                continue
            target = (path.parent / unquote(url.path)).resolve() if url.path else path
            if not target.exists():
                errors.append(f'{relative}: missing link target {destination}')
                continue
            if url.fragment and target.suffix == '.md':
                target_prose = documents[target][0] if target in documents else sections(target.read_text())[0]
                if unquote(url.fragment) not in anchors(target_prose):
                    errors.append(f'{relative}: missing anchor {destination}')
            if chinese and target in documents and not target.name.endswith('.zh-CN.md') and target != english:
                errors.append(f'{relative}: link should use Chinese counterpart: {destination}')
    if errors:
        print('\n'.join(errors), file=sys.stderr)
        return 1
    print(f'Checked {len(documents) // 2} English/Chinese documentation pairs: coverage, language switches, heading structure, code examples, and local links/anchors.')
    return 0


if __name__ == '__main__':
    sys.exit(main())
