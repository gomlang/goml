bootstrap_platform() {
    local host_os host_arch
    host_os="$(uname -s)"
    host_arch="$(uname -m)"
    case "$host_os/$host_arch" in
        Linux/x86_64) printf '%s\n' linux-amd64 ;;
        Darwin/arm64) printf '%s\n' darwin-arm64 ;;
        *)
            printf 'source bootstrap requires Linux amd64 or macOS arm64; detected %s/%s\n' \
                "$host_os" "$host_arch" >&2
            return 1
            ;;
    esac
}

bootstrap_stage0() {
    local platform="$1"
    case "$platform" in
        linux-amd64) GOML_STAGE0_SHA256="${GOML_STAGE0_LINUX_AMD64_SHA256:-}" ;;
        darwin-arm64) GOML_STAGE0_SHA256="${GOML_STAGE0_DARWIN_ARM64_SHA256:-}" ;;
        *) printf 'unsupported stage0 platform: %s\n' "$platform" >&2; return 1 ;;
    esac
    if [[ ! "$GOML_STAGE0_SHA256" =~ ^[0-9a-f]{64}$ ]]; then
        printf 'missing or invalid stage0 SHA-256 for %s\n' "$platform" >&2
        return 1
    fi
    GOML_STAGE0_ASSET_NAME="goml-${GOML_STAGE0_RELEASE_TAG#v}-$platform.tar.gz"
    GOML_STAGE0_URL="https://github.com/gomlang/goml/releases/download/$GOML_STAGE0_RELEASE_TAG/$GOML_STAGE0_ASSET_NAME"
}
