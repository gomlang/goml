package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"time"

	"goml.run/tools/goml-go-meta/internal/protocol"
	"goml.run/tools/goml-go-meta/internal/validate"
)

func serve(ctx context.Context, input io.ReadCloser, output io.Writer) int {
	stop := context.AfterFunc(ctx, func() { input.Close() })
	defer stop()
	response := checkRequest(ctx, input)
	if err := protocol.EncodeResponse(output, response); err != nil {
		fmt.Fprintln(os.Stderr, "ffi-protocol:", err)
		return 1
	}
	if len(response.Diagnostics) > 0 {
		return 1
	}
	for _, binding := range response.Bindings {
		if binding.Status != "verified" {
			return 1
		}
	}
	for _, result := range response.TypeResults {
		if result.Status != "verified" {
			return 1
		}
	}
	return 0
}

func checkRequest(ctx context.Context, input io.Reader) protocol.Response {
	request, err := protocol.DecodeRequest(ctx, input)
	if err != nil {
		return failureResponse("ffi-protocol", err)
	}
	cleanup, err := selectGoCommand(request.BuildContext.GoExecutable)
	if err != nil {
		return failureResponse("ffi-tool-unavailable", err)
	}
	defer cleanup()
	return validate.Check(ctx, request)
}

func failureResponse(code string, err error) protocol.Response {
	return protocol.Response{ProtocolVersion: protocol.Version, Diagnostics: []protocol.Diagnostic{{Code: code, Message: err.Error()}}}
}

func selectGoCommand(value string) (func(), error) {
	selected, err := exec.LookPath(value)
	if err != nil {
		return nil, err
	}
	selected, err = filepath.Abs(selected)
	if err != nil {
		return nil, err
	}
	previousPath, hadPath := os.LookupEnv("PATH")
	directory, err := os.MkdirTemp("", "goml-go-command-")
	if err != nil {
		return nil, err
	}
	cleanup := func() {
		if hadPath {
			os.Setenv("PATH", previousPath)
		} else {
			os.Unsetenv("PATH")
		}
		os.RemoveAll(directory)
	}
	if err := os.Symlink(selected, filepath.Join(directory, "go")); err != nil {
		cleanup()
		return nil, err
	}
	if err := os.Setenv("PATH", directory+string(os.PathListSeparator)+previousPath); err != nil {
		cleanup()
		return nil, err
	}
	return cleanup, nil
}

func main() {
	timeout := flag.Duration("timeout", 2*time.Minute, "maximum request duration")
	flag.Parse()
	if *timeout <= 0 || *timeout > 10*time.Minute || flag.NArg() != 0 {
		fmt.Fprintln(os.Stderr, "ffi-protocol: timeout must be positive and at most 10m; positional arguments are unsupported")
		os.Exit(2)
	}
	ctx, cancelSignal := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancelSignal()
	ctx, cancelTimeout := context.WithTimeout(ctx, *timeout)
	defer cancelTimeout()
	os.Exit(serve(ctx, os.Stdin, os.Stdout))
}
