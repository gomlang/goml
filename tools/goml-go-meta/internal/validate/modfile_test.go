package validate

import (
	"context"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"testing"

	"goml.dev/tools/goml-go-meta/internal/protocol"
)

func TestModuleFileFlagWithSpaces(t *testing.T) {
	request := fixture(t)
	file := filepath.Join(request.BuildContext.ModuleDir, "alternate go.mod")
	if err := os.WriteFile(file, []byte("module example.com/host\n\ngo 1.26.0\n"), 0644); err != nil {
		t.Fatal(err)
	}
	original := []byte("module example.com/unselected\n\ngo 1.26.0\n")
	if err := os.WriteFile(filepath.Join(request.BuildContext.ModuleDir, "go.mod"), original, 0644); err != nil {
		t.Fatal(err)
	}
	request.BuildContext.GOFLAGS = "-trimpath " + strconv.Quote("-modfile="+file)
	i64 := []protocol.Type{{Tag: "int64"}}
	request.Bindings = []protocol.Binding{binding("Scalar", i64, i64)}
	result := Check(context.Background(), request)
	if len(result.Bindings) != 1 || result.Bindings[0].Status != "verified" {
		t.Fatalf("alternate module failed: %+v", result)
	}
	unchanged, err := os.ReadFile(filepath.Join(request.BuildContext.ModuleDir, "go.mod"))
	if err != nil || string(unchanged) != string(original) {
		t.Fatalf("root module changed: %s, %v", unchanged, err)
	}
}

func TestGoFlagQuoting(t *testing.T) {
	for _, test := range []struct {
		input    string
		expected []string
	}{
		{"", nil},
		{" \t-trimpath\r\n-mod=readonly ", []string{"-trimpath", "-mod=readonly"}},
		{`"-modfile=a b'c" '-tags=x y'`, []string{"-modfile=a b'c", "-tags=x y"}},
		{`-ldflags=-X=a'b`, []string{"-ldflags=-X=a'b"}},
		{`"-tags=a"-trimpath`, []string{"-tags=a", "-trimpath"}},
	} {
		actual, err := splitGoFlags(test.input)
		if err != nil || !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("%q: got %v, %v; want %v", test.input, actual, err, test.expected)
		}
	}
	if _, err := splitGoFlags(`"-modfile=unfinished`); err == nil {
		t.Fatal("accepted unterminated flags")
	}
}
