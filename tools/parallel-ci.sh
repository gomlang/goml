set -euo pipefail

repository_root="$(cd "$(dirname "$0")/.." && pwd)"
cd "$repository_root"

if test "${GOML_CI_SEQUENTIAL:-0}" = 1; then
    for recipe in _ci-gomlc-test _ci-scripts _bootstrap-stage3 _ci-goml-test _ci-vscode _ci-release-smoke; do
        just "$recipe"
    done
    exit 0
fi

ci_pids=()
just _ci-gomlc-test &
ci_pids+=("$!")

GOML_BUILD_JOBS=2 nice -n 10 just _ci-scripts &
ci_pids+=("$!")

GOML_BUILD_JOBS=2 just _bootstrap-stage3 &
ci_pids+=("$!")

for recipe in _ci-goml-test _ci-vscode _ci-release-smoke; do
    nice -n 10 just "$recipe" &
    ci_pids+=("$!")
done

ci_status=0
set +e
for ci_pid in "${ci_pids[@]}"; do
    wait "$ci_pid"
    process_status="$?"
    if test "$process_status" != 0 && test "$ci_status" = 0; then
        ci_status="$process_status"
    fi
done
set -e

exit "$ci_status"
