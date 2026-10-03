release_platform() {
    case "$1-$2" in
        linux-amd64|darwin-arm64) printf '%s-%s\n' "$1" "$2" ;;
        *) echo "unsupported release target: $1/$2" >&2; return 1 ;;
    esac
}

release_sha256() {
    if command -v sha256sum >/dev/null 2>&1; then
        sha256sum "$@"
    else
        shasum -a 256 "$@"
    fi
}
