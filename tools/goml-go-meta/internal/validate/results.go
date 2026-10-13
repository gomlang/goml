package validate

import (
	"golang.org/x/tools/go/packages"
	"goml.dev/tools/goml-go-meta/internal/protocol"
)

type checkResults struct {
	protocol.Response
}

func newCheckResults(request protocol.Request) *checkResults {
	results := &checkResults{Response: protocol.Response{ProtocolVersion: protocol.Version}}
	for _, binding := range request.Bindings {
		results.Bindings = append(results.Bindings, protocol.BindingResult{ID: binding.ID, Status: "verified"})
	}
	for _, query := range request.TypeQueries {
		results.TypeResults = append(results.TypeResults, protocol.TypeResult{ID: query.ID, Status: "verified"})
	}
	return results
}

func appendDiagnostic(diagnostics []protocol.Diagnostic, code, message string) []protocol.Diagnostic {
	for _, diagnostic := range diagnostics {
		if diagnostic.Code == code && diagnostic.Message == message {
			return diagnostics
		}
	}
	return append(diagnostics, protocol.Diagnostic{Code: code, Message: message})
}

func (r *checkResults) failBinding(index int, code, message string) {
	result := &r.Bindings[index]
	result.Status = "failed"
	result.Diagnostics = appendDiagnostic(result.Diagnostics, code, message)
}

func (r *checkResults) failQuery(index int, code, message string) {
	result := &r.TypeResults[index]
	result.Status = "failed"
	result.Diagnostics = appendDiagnostic(result.Diagnostics, code, message)
}

func (r *checkResults) failAll(code, message string) {
	for index := range r.Bindings {
		r.failBinding(index, code, message)
	}
	for index := range r.TypeResults {
		r.failQuery(index, code, message)
	}
}

func (r *checkResults) packageErrors(root *packages.Package, loaded []*packages.Package, witnesses *witnessSet) {
	for _, p := range loaded {
		for _, problem := range p.Errors {
			code := "ffi-package-load"
			if p == root && problem.Kind == packages.TypeError {
				code = "ffi-signature-mismatch"
			}
			matched := false
			for _, file := range witnesses.origins.diagnostic(problem.Pos) {
				targets := witnesses.targets[file]
				for _, index := range targets.bindings {
					r.failBinding(index, code, problem.Msg)
					matched = true
				}
				for _, index := range targets.queries {
					r.failQuery(index, code, problem.Msg)
					matched = true
				}
			}
			if !matched {
				r.failAll(code, problem.Msg)
			}
		}
	}
}
