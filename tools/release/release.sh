set -euo pipefail

parse_version() {
    release_value="${1#v}"
    if [[ ! "$release_value" =~ ^(0|[1-9][0-9]*)\.(0|[1-9][0-9]*)\.(0|[1-9][0-9]*)$ ]]; then
        return 1
    fi
    release_major="${BASH_REMATCH[1]}"
    release_minor="${BASH_REMATCH[2]}"
    release_patch="${BASH_REMATCH[3]}"
}

version_greater() {
    test "$1" -gt "$4" ||
        { test "$1" -eq "$4" && test "$2" -gt "$5"; } ||
        { test "$1" -eq "$4" && test "$2" -eq "$5" && test "$3" -gt "$6"; }
}

check_version() {
    parse_version "$1"
    expected="$release_major.$release_minor.$release_patch"
    goml_version="$(sed -n -E 's/^[[:space:]]*"([0-9]+\.[0-9]+\.[0-9]+)"[[:space:]]*$/\1/p' goml/version/version.gom)"
    gomlc_version="$(sed -n -E 's/^[[:space:]]*"([0-9]+\.[0-9]+\.[0-9]+)"[[:space:]]*$/\1/p' gomlc/version/version.gom)"
    vscode_version="$(sed -n -E 's/^[[:space:]]*"version": "([^"]+)",?$/\1/p' editors/vscode/package.json | head -n 1)"
    lock_versions="$(sed -n -E 's/^[[:space:]]*"version": "([^"]+)",?$/\1/p' editors/vscode/package-lock.json | head -n 2)"
    test "$(cat VERSION)" = "$expected"
    test "$goml_version" = "$expected"
    test "$gomlc_version" = "$expected"
    test "$vscode_version" = "$expected"
    test "$lock_versions" = "$expected
$expected"
}

check_stage0() {
    parse_version "$1"
    expected="v$release_major.$release_minor.$release_patch"
    source bootstrap/stage0.env
    test "$GOML_STAGE0_RELEASE_TAG" = "$expected"
}

check_next() {
    parse_version "$1"
    previous_major="$release_major"
    previous_minor="$release_minor"
    previous_patch="$release_patch"
    parse_version "$2"
    { test "$release_major" -eq "$previous_major" &&
        test "$release_minor" -eq "$previous_minor" &&
        test "$release_patch" -eq "$((previous_patch + 1))"; } ||
        { test "$release_major" -eq "$previous_major" &&
            test "$release_minor" -eq "$((previous_minor + 1))" &&
            test "$release_patch" -eq 0; } ||
        { test "$release_major" -eq "$((previous_major + 1))" &&
            test "$release_minor" -eq 0 &&
            test "$release_patch" -eq 0; }
}

latest() {
    latest_found=false
    latest_major=0
    latest_minor=0
    latest_patch=0
    for tag in "$@"; do
        if [[ "$tag" != v* ]] || ! parse_version "$tag"; then
            continue
        fi
        if ! $latest_found || version_greater "$release_major" "$release_minor" "$release_patch" "$latest_major" "$latest_minor" "$latest_patch"; then
            latest_found=true
            latest_major="$release_major"
            latest_minor="$release_minor"
            latest_patch="$release_patch"
        fi
    done
    $latest_found
    printf 'v%s.%s.%s\n' "$latest_major" "$latest_minor" "$latest_patch"
}

set_version() {
    parse_version "$1"
    value="$release_major.$release_minor.$release_patch"
    printf '%s\n' "$value" > VERSION
    sed -i -E "s/^([[:space:]]*)\"[0-9]+\\.[0-9]+\\.[0-9]+\"[[:space:]]*$/\\1\"$value\"/" goml/version/version.gom
    sed -i -E "s/^([[:space:]]*)\"[0-9]+\\.[0-9]+\\.[0-9]+\"[[:space:]]*$/\\1\"$value\"/" gomlc/version/version.gom
}

set_stage0() {
    parse_version "$1"
    value="$release_major.$release_minor.$release_patch"
    checksum="$2"
    [[ "$checksum" =~ ^[0-9a-f]{64}$ ]]
    asset="goml-$value-linux-amd64.tar.gz"
    printf '%s\n' \
        "GOML_STAGE0_RELEASE_TAG=v$value" \
        "GOML_STAGE0_ASSET_NAME=$asset" \
        "GOML_STAGE0_URL=https://github.com/gomlang/goml/releases/download/v$value/$asset" \
        "GOML_STAGE0_SHA256=$checksum" \
        > bootstrap/stage0.env
}

command="${1:-}"
shift || true
case "$command" in
    check-version) test "$#" = 1; check_version "$1" ;;
    check-stage0) test "$#" = 1; check_stage0 "$1" ;;
    check-next) test "$#" = 2; check_next "$1" "$2" ;;
    latest) test "$#" -gt 0; latest "$@" ;;
    set-version) test "$#" = 1; set_version "$1" ;;
    set-stage0) test "$#" = 2; set_stage0 "$1" "$2" ;;
    *) exit 2 ;;
esac
