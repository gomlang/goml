source "$(dirname "${BASH_SOURCE[0]}")/../lib/common.sh"

release_platform() {
    case "$1-$2" in
        linux-amd64|darwin-arm64) printf '%s-%s\n' "$1" "$2" ;;
        *) echo "unsupported release target: $1/$2" >&2; return 1 ;;
    esac
}

release_sha256() {
    toolchain_sha256 "$@"
}
