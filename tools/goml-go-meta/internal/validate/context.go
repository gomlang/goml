package validate

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"goml.dev/tools/goml-go-meta/internal/protocol"
)

func environment(c protocol.BuildContext) []string {
	values := map[string]string{}
	for _, value := range os.Environ() {
		if key, value, ok := strings.Cut(value, "="); ok {
			values[key] = value
		}
	}
	for key, value := range map[string]string{
		"GOOS": c.GOOS, "GOARCH": c.GOARCH, "CGO_ENABLED": c.CGOEnabled,
		"GOFLAGS": c.GOFLAGS, "GO111MODULE": c.GO111MODULE, "GOWORK": c.GOWORK,
		"GOTOOLCHAIN": "local", "GOPROXY": "off", "GOSUMDB": "off", "GOVCS": "*:off",
		"GOPACKAGESDRIVER": "off", "GONOPROXY": "none", "GONOSUMDB": "*",
	} {
		values[key] = value
	}
	keys := make([]string, 0, len(values))
	for key := range values {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	result := make([]string, 0, len(keys))
	for _, key := range keys {
		result = append(result, key+"="+values[key])
	}
	return result
}

func verifyToolchain(ctx context.Context, c protocol.BuildContext, env []string) error {
	selected, err := exec.LookPath(c.GoExecutable)
	if err != nil {
		return err
	}
	actual, err := exec.LookPath("go")
	if err != nil {
		return err
	}
	selected, err = filepath.EvalSymlinks(selected)
	if err != nil {
		return err
	}
	actual, err = filepath.EvalSymlinks(actual)
	if err != nil {
		return err
	}
	if selected != actual {
		return fmt.Errorf("selected Go executable must be the helper's resolved go command")
	}
	command := exec.CommandContext(ctx, selected, "env", "GOVERSION")
	command.WaitDelay = time.Second
	command.Env = env
	command.Dir = c.ModuleDir
	output, err := command.Output()
	if err != nil {
		return err
	}
	if strings.TrimSpace(string(output)) != c.Toolchain {
		return fmt.Errorf("Go toolchain changed: expected %s, found %s", c.Toolchain, strings.TrimSpace(string(output)))
	}
	return nil
}

func splitGoFlags(value string) ([]string, error) {
	var result []string
	for {
		value = strings.TrimLeft(value, " \t\r\n")
		if value == "" {
			return result, nil
		}
		if value[0] == '\'' || value[0] == '"' {
			quote := value[0]
			end := strings.IndexByte(value[1:], quote)
			if end < 0 {
				return nil, fmt.Errorf("unterminated quote in GOFLAGS")
			}
			result = append(result, value[1:1+end])
			value = value[2+end:]
		} else {
			end := strings.IndexAny(value, " \t\r\n")
			if end < 0 {
				return append(result, value), nil
			}
			result = append(result, value[:end])
			value = value[end:]
		}
	}
}
