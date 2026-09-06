package nativebridge

import (
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"testing"
)

func TestImageBuildLockWorker(t *testing.T) {
	root := os.Getenv("GOMLGO_TEST_IMAGE_LOCK_ROOT")
	if root == "" {
		return
	}
	for range 20 {
		lock, message := LockImageBuild(filepath.Join(root, "image.lock"))
		if message != "" {
			t.Fatal(message)
		}
		func() {
			defer UnlockImageBuild(lock)
			file := filepath.Join(root, "counter")
			data, err := os.ReadFile(file)
			if err != nil {
				t.Fatal(err)
			}
			count, err := strconv.Atoi(string(data))
			if err != nil {
				t.Fatal(err)
			}
			if err := os.WriteFile(file, []byte(strconv.Itoa(count+1)), 0o600); err != nil {
				t.Fatal(err)
			}
		}()
	}
}

func TestImageBuildLockSerializesProcesses(t *testing.T) {
	root := t.TempDir()
	file := filepath.Join(root, "counter")
	if err := os.WriteFile(file, []byte("0"), 0o600); err != nil {
		t.Fatal(err)
	}
	workers := make([]*exec.Cmd, 4)
	for index := range workers {
		worker := exec.Command(os.Args[0], "-test.run=^TestImageBuildLockWorker$")
		worker.Env = append(os.Environ(), "GOMLGO_TEST_IMAGE_LOCK_ROOT="+root)
		worker.Stdout = os.Stdout
		worker.Stderr = os.Stderr
		if err := worker.Start(); err != nil {
			t.Fatal(err)
		}
		workers[index] = worker
	}
	for _, worker := range workers {
		if err := worker.Wait(); err != nil {
			t.Fatal(err)
		}
	}
	data, err := os.ReadFile(file)
	if err != nil {
		t.Fatal(err)
	}
	if string(data) != "80" {
		t.Fatalf("counter = %s, want 80", data)
	}
}

func TestImageBuildLockReportsOpenFailure(t *testing.T) {
	lock, message := LockImageBuild(filepath.Join(t.TempDir(), "missing", "image.lock"))
	if lock >= 0 || message == "" {
		t.Fatalf("lock = %v, error = %q", lock, message)
	}
}
