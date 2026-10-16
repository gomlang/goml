set -euo pipefail

source "$(dirname "${BASH_SOURCE[0]}")/common.sh"
platform="$(bootstrap_platform)"
if test "$#" = 1 && test "$platform" != "$1"; then
    printf 'this command requires %s; detected %s\n' "$1" "$platform" >&2
    exit 1
fi
