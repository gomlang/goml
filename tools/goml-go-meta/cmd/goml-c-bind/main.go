package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"time"

	"goml.dev/tools/goml-go-meta/internal/cbind"
)

func main() {
	executable, err := os.Executable()
	if err == nil {
		executable, err = filepath.EvalSymlinks(executable)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, "bind-c:", err)
		os.Exit(1)
	}
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	ctx, timeout := context.WithTimeout(ctx, 2*time.Minute)
	defer timeout()
	if err := cbind.Execute(ctx, os.Args[1:], filepath.Join(filepath.Dir(executable), "gomlfmt"), os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, "bind-c:", err)
		os.Exit(1)
	}
}
