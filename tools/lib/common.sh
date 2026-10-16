toolchain_sha256() {
    if command -v sha256sum >/dev/null 2>&1; then
        sha256sum "$@"
    else
        shasum -a 256 "$@"
    fi
}

toolchain_hash_files() {
    local hash_command=(shasum -a 256)
    if command -v sha256sum >/dev/null 2>&1; then
        hash_command=(sha256sum)
    fi
    find "$@" -type f -print0 \
        | LC_ALL=C sort -z \
        | xargs -0 "${hash_command[@]}"
}

toolchain_file_stamp() {
    if test "$(uname -s)" = Darwin; then
        stat -f '%m:%i' "$1"
    else
        stat -c '%y:%i' "$1"
    fi
}
