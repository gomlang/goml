package validate

import (
	"testing"

	"golang.org/x/tools/go/packages"
	"goml.run/tools/goml-go-meta/internal/protocol"
)

func TestPackageDiagnosticAttribution(t *testing.T) {
	request := protocol.Request{
		Bindings:    []protocol.Binding{{ID: "first"}, {ID: "duplicate"}, {ID: "unrelated"}},
		TypeQueries: []protocol.TypeQuery{{ID: "shared-import"}, {ID: "unrelated"}},
	}
	witnesses := &witnessSet{
		targets: map[string]witnessTargets{
			"binding.go": {bindings: []int{0, 1}},
			"query.go":   {queries: []int{0}},
		},
		origins: witnessOrigins{path: "bundle.go", lines: map[int][]string{3: {"binding.go", "query.go"}}},
	}
	for _, test := range []struct {
		name         string
		position     string
		dependency   bool
		wantBindings []bool
		wantQueries  []bool
		wantCode     string
	}{
		{"binding", "binding.go:1:1", false, []bool{true, true, false}, []bool{false, false}, "ffi-signature-mismatch"},
		{"shared bundle import", "bundle.go:3:2", false, []bool{true, true, false}, []bool{true, false}, "ffi-signature-mismatch"},
		{"dependency", "query.go:1:1", true, []bool{false, false, false}, []bool{true, false}, "ffi-package-load"},
		{"unmapped", "caller.go:1:1", false, []bool{true, true, true}, []bool{true, true}, "ffi-signature-mismatch"},
		{"unpositioned", "", true, []bool{true, true, true}, []bool{true, true}, "ffi-package-load"},
	} {
		t.Run(test.name, func(t *testing.T) {
			root := &packages.Package{}
			loaded := root
			if test.dependency {
				loaded = &packages.Package{}
			}
			problem := packages.Error{Pos: test.position, Msg: "invalid type", Kind: packages.TypeError}
			loaded.Errors = []packages.Error{problem, problem}
			results := newCheckResults(request)
			results.packageErrors(root, []*packages.Package{loaded}, witnesses)
			check := func(id, status string, diagnostics []protocol.Diagnostic, wantFailed bool) {
				t.Helper()
				if !wantFailed {
					if status != "verified" || len(diagnostics) != 0 {
						t.Errorf("unrelated result %s: %s %+v", id, status, diagnostics)
					}
					return
				}
				if status != "failed" || len(diagnostics) != 1 || diagnostics[0].Code != test.wantCode || diagnostics[0].Message != problem.Msg {
					t.Errorf("result %s: expected one %s diagnostic, got %s %+v", id, test.wantCode, status, diagnostics)
				}
			}
			for index, result := range results.Bindings {
				check(result.ID, result.Status, result.Diagnostics, test.wantBindings[index])
			}
			for index, result := range results.TypeResults {
				check(result.ID, result.Status, result.Diagnostics, test.wantQueries[index])
			}
		})
	}
}
