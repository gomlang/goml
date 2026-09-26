set -euo pipefail

test "$#" = 3 || test "$#" = 4

stage_name="$1"
driver_path="$(realpath "$2")"
compiler_path="$(realpath "$3")"
repository_root="$(cd "$(dirname "$0")/.." && pwd)"
build_mode="${4:-full}"
case "$build_mode" in
    full|compiler|artifacts) ;;
    *) exit 1 ;;
esac
build_jobs="${GOML_BUILD_JOBS:-$(nproc)}"
case "$build_jobs" in
    ''|*[!0-9]*) exit 1 ;;
esac
test "$build_jobs" -gt 0
if test "$build_jobs" -gt 22; then
    build_jobs=22
fi
go_compile_jobs="$build_jobs"
if test "$go_compile_jobs" -gt 12; then
    go_compile_jobs=12
fi
bootstrap_go_flags="${GOFLAGS:-}"
if test -n "$bootstrap_go_flags"; then
    bootstrap_go_flags="$bootstrap_go_flags "
fi
bootstrap_go_flags="${bootstrap_go_flags}-gcflags=-c=$go_compile_jobs -ldflags=-s"

(
    cd "$repository_root/gomlc"
    if test "$build_mode" = compiler; then
        GOFLAGS="$bootstrap_go_flags" "$driver_path" run cmd/gomlc --jobs "$build_jobs" --target-dir "_bootstrap/$stage_name" --compiler "$compiler_path" -- version >/dev/null
        "$driver_path" __build-artifacts cmd/gomldoc --jobs "$build_jobs" --target-dir "_bootstrap/$stage_name" --compiler "$compiler_path"
    elif test "$build_mode" = artifacts; then
        "$driver_path" __build-artifacts cmd/gomlc --jobs "$build_jobs" --target-dir "_bootstrap/$stage_name" --compiler "$compiler_path"
        "$driver_path" __build-artifacts cmd/gomldoc --jobs "$build_jobs" --target-dir "_bootstrap/$stage_name" --compiler "$compiler_path"
    else
        GOFLAGS="$bootstrap_go_flags" "$driver_path" build --jobs "$build_jobs" --target-dir "_bootstrap/$stage_name" --compiler "$compiler_path"
    fi
) &
compiler_build_pid="$!"

(
    cd "$repository_root/goml"
    if test "$build_mode" = artifacts; then
        "$driver_path" __build-artifacts cmd/goml --jobs "$build_jobs" --target-dir "_bootstrap/$stage_name" --compiler "$compiler_path"
    else
        GOFLAGS="$bootstrap_go_flags" "$driver_path" build --jobs "$build_jobs" --target-dir "_bootstrap/$stage_name" --compiler "$compiler_path"
    fi
) &
driver_build_pid="$!"

set +e
wait "$compiler_build_pid"
compiler_build_status="$?"
wait "$driver_build_pid"
driver_build_status="$?"
set -e

test "$compiler_build_status" = 0
test "$driver_build_status" = 0
