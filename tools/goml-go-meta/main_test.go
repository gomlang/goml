package main

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"strings"
	"testing"
	"time"

	"goml.dev/tools/goml-go-meta/internal/protocol"
)

func TestMalformedRequestResponse(t *testing.T) {
	var output bytes.Buffer
	if serve(context.Background(), io.NopCloser(strings.NewReader(`{"protocol_version":99}`)), &output) != 1 {
		t.Fatal("malformed request succeeded")
	}
	var response protocol.Response
	if err := json.Unmarshal(output.Bytes(), &response); err != nil {
		t.Fatal(err)
	}
	if response.ProtocolVersion != 1 || len(response.Diagnostics) != 1 || response.Diagnostics[0].Code != "ffi-protocol" {
		t.Fatalf("unexpected response: %+v", response)
	}
}

func TestBlockedInputCancellation(t *testing.T) {
	reader, writer := io.Pipe()
	defer writer.Close()
	ctx, cancel := context.WithCancel(context.Background())
	var output bytes.Buffer
	done := make(chan int, 1)
	go func() { done <- serve(ctx, reader, &output) }()
	cancel()
	select {
	case status := <-done:
		if status != 1 {
			t.Fatal("canceled request succeeded")
		}
	case <-time.After(5 * time.Second):
		t.Fatal("blocked input was not canceled")
	}
}
