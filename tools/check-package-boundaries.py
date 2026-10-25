import graphlib
import re
import subprocess
import sys
from pathlib import Path

ROOT = Path(__file__).resolve().parent.parent
IMPORT = re.compile(r'^(?:pub\s+)?use\s+module::([A-Za-z_0-9:]+)', re.MULTILINE)
BACKEND = {'core', 'mono', 'lift', 'anf', 'repr', 'go_ir', 'go_types', 'go_runtime', 'go_backend'}
FRONTEND_CLIENTS = {'checker', 'compiler', 'query', 'lsp', 'cmd'}
ALLOWED = {
    'gomlc/go_ir': {'ast', 'pretty'},
    'gomlc/diagnostic': {'ast'},
    'gomlc/names': set(),
    'gomlc/builtin_source': {'ast', 'hir', 'names', 'parser', 'resources'},
    'gomlc/go_runtime': {'ast', 'go_ir', 'go_types', 'repr', 'resources'},
    'goml/build_model': set(),
}
FORBIDDEN = {
    'gomlc/tast': FRONTEND_CLIENTS | {'go_backend', 'parser', 'resources'},
    'goml/project': {'build', 'build_plan', 'go_export'},
    'goml/build': {'build_plan', 'go_export'},
    'goml/build_plan': {'build', 'go_export'},
}


def source_files():
    result = subprocess.run(
        ['git', 'ls-files', '--cached', '--others', '--exclude-standard', '-z', '--',
         'gomlc/*.goml', 'goml/*.goml'],
        cwd=ROOT, check=True, capture_output=True, text=True,
    )
    return sorted({
        Path(name) for name in result.stdout.split('\0')
        if name and (ROOT / name).is_file()
        and not Path(name).name.endswith('_test.goml')
        and not {'testdata', 'tests', 'examples'}.intersection(Path(name).parts)
    })


def check(files):
    packages = {str(path.parent) for path in files}
    graph = {package: set() for package in packages}
    errors = []
    for path in files:
        package = str(path.parent)
        module = path.parts[0]
        for match in IMPORT.finditer((ROOT / path).read_text()):
            parts = [part for part in match.group(1).split('::') if part]
            dependency = module + '/' + '/'.join(parts)
            while dependency not in packages and '/' in dependency:
                dependency = dependency.rsplit('/', 1)[0]
            if dependency not in packages:
                errors.append(f'{path}: unknown local import {match.group(0)}')
                continue
            graph[package].add(dependency)
            owner = dependency.split('/')[1]
            forbidden = set(FORBIDDEN.get(package, set()))
            if module == 'gomlc' and path.parts[1] in BACKEND:
                forbidden |= FRONTEND_CLIENTS
            if owner == 'cmd' or owner in forbidden:
                errors.append(f'{path}: forbidden dependency on {dependency}')
            if package in ALLOWED and owner not in ALLOWED[package]:
                errors.append(f'{path}: dependency outside its layer: {dependency}')
    try:
        tuple(graphlib.TopologicalSorter(graph).static_order())
    except graphlib.CycleError as error:
        errors.append('Package cycle: ' + ' -> '.join(error.args[1]))
    return errors, len(graph)


def main():
    errors, count = check(source_files())
    if errors:
        print('\n'.join(errors), file=sys.stderr)
        return 1
    print(f'Checked {count} compiler/driver packages: local imports, dependency layers, and cycles.')
    return 0


if __name__ == '__main__':
    sys.exit(main())
