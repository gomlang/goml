import * as fs from 'fs';
import * as path from 'path';

function isExecutable(file: string): boolean {
  try {
    if (!fs.statSync(file).isFile()) {
      return false;
    }
    fs.accessSync(file, process.platform === 'win32' ? fs.constants.F_OK : fs.constants.X_OK);
    return true;
  } catch {
    return false;
  }
}

export function findExecutable(
  configured: string | undefined,
  directories: readonly string[],
  names: readonly string[]
): string | undefined {
  if (configured) {
    return isExecutable(configured) ? configured : undefined;
  }
  for (const directory of directories) {
    for (const name of names) {
      const candidate = path.join(directory, name);
      if (isExecutable(candidate)) {
        return candidate;
      }
    }
  }
  return undefined;
}
