#!/usr/bin/env bash
set -euo pipefail

if [[ ${GOMLGO_GO+x} ]]; then
    gomlgo_go=${GOMLGO_GO:-go}
elif [[ -x /usr/lib/go-1.26/bin/go ]]; then
    gomlgo_go=/usr/lib/go-1.26/bin/go
else
    gomlgo_go=go
fi

export GOTOOLCHAIN=local
gomlgo_version=$("$gomlgo_go" env GOVERSION)
case "$gomlgo_version" in
    go1.26|go1.26.*) ;;
    *) printf 'gomlgo requires Go 1.26.x, found %s\n' "$gomlgo_version" >&2; exit 1 ;;
esac

if [[ ${1:-} == --source-root ]]; then
    gomlgo_root=$("$gomlgo_go" env GOROOT)
    printf '%s/src\n' "$gomlgo_root"
else
    exec "$gomlgo_go" "$@"
fi
