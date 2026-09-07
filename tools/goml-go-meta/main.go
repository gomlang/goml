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

	"goml.dev/tools/goml-go-meta/internal/protocol"
	"goml.dev/tools/goml-go-meta/internal/validate"
)

func serve(ctx context.Context, input io.ReadCloser, output io.Writer) int {
	stop := context.AfterFunc(ctx, func() { input.Close() })
	defer stop()
	request, err := protocol.DecodeRequest(ctx, input)
	var response protocol.Response
	if err != nil {
		response = protocol.Response{ProtocolVersion: protocol.Version, Diagnostics: []protocol.Diagnostic{{Code: "ffi-protocol", Message: err.Error()}}}
	} else {
		cleanup, resolveErr := selectGoCommand(request.BuildContext.GoExecutable)
		if resolveErr != nil {
			response = protocol.Response{ProtocolVersion: protocol.Version, Diagnostics: []protocol.Diagnostic{{Code: "ffi-tool-unavailable", Message: resolveErr.Error()}}}
		} else {
			response = validate.Check(ctx, request)
			cleanup()
		}
	}
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

func selectGoCommand(value string) (func(), error) {
	selected, err := exec.LookPath(value)
	if err != nil {
		return nil, err
	}
	selected, err = filepath.Abs(selected)
	if err != nil {
		return nil, err
	}
	previousPath := os.Getenv("PATH")
	directory, err := os.MkdirTemp("", "goml-go-command-")
	if err != nil {
		return nil, err
	}
	cleanup := func() { os.Setenv("PATH", previousPath); os.RemoveAll(directory) }
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
