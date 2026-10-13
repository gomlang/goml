package cbind

import (
	"context"
	"testing"
	"time"
)

func TestRunBoundsInheritedOutputWait(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()
	started := time.Now()
	_, err := run(ctx, t.TempDir(), "", "sh", "-c", "sleep 3 & wait")
	if err == nil {
		t.Fatal("canceled command succeeded")
	}
	if elapsed := time.Since(started); elapsed >= 2500*time.Millisecond {
		t.Fatalf("inherited output pipes delayed cancellation: %s", elapsed)
	}
}
