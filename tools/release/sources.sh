set -euo pipefail

test "$#" = 1
repository_root="$(cd "$(dirname "$0")/../.." && pwd)"
cd "$repository_root"
mkdir -p "$(dirname "$1")"
tar -czf "$1" \
    goml/_bootstrap/stage2/build/pkg/gomlang/bootstrap_goml/cmd/goml/goml_generated.go \
    gomlc/_bootstrap/stage2/build/pkg/gomlc/cmd/gomlc/goml_generated.go \
    gomlc/_bootstrap/stage2/build/pkg/gomlc/cmd/gomlfmt/goml_generated.go \
    gomlc/_bootstrap/stage2/build/pkg/gomlc/cmd/gomldoc/goml_generated.go \
    gomlc/_bootstrap/stage2/build/pkg/gomlc/cmd/gomllsp/goml_generated.go
