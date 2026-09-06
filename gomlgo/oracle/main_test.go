package main

import (
	"encoding/json"
	"runtime"
	"testing"
)

func encoded(t *testing.T, response protocolResponse) []byte {
	t.Helper()
	result, err := json.Marshal(response)
	if err != nil {
		t.Fatal(err)
	}
	return result
}

func encodedValue(t *testing.T, response any) []byte {
	t.Helper()
	result, err := json.Marshal(response)
	if err != nil {
		t.Fatal(err)
	}
	return result
}

func typeCheckInput(files string) []byte {
	return []byte(`{"package_path":"example/p","go_version":"go1.26","goos":"linux","goarch":"amd64","files":` + files + `,"config":{"enable_alias":true}}`)
}

func TestOutputIsDeterministic(t *testing.T) {
	source := []byte("package main\n\nfunc main() { println(1) }\n")
	first := encoded(t, execute("parse-file", source))
	second := encoded(t, execute("parse-file", source))
	if string(first) != string(second) {
		t.Fatalf("output changed between runs\nfirst: %s\nsecond: %s", first, second)
	}
}

func TestScanUsesByteOffsetsAndMarksSyntheticSemicolons(t *testing.T) {
	response := execute("scan", []byte("package π\n"))
	if !response.OK {
		t.Fatalf("scan failed: %#v", response.Diagnostics)
	}
	if len(response.Tokens) != 4 {
		t.Fatalf("got %d tokens", len(response.Tokens))
	}
	identifier := response.Tokens[1]
	if identifier.Kind != "IDENT" || identifier.Literal != "π" || identifier.Start != 8 || identifier.End != 10 {
		t.Fatalf("unexpected identifier: %#v", identifier)
	}
	semicolon := response.Tokens[2]
	if semicolon.Kind != ";" || semicolon.Literal != "\n" || semicolon.Start != 10 || semicolon.End != 10 || !semicolon.Synthetic {
		t.Fatalf("unexpected semicolon: %#v", semicolon)
	}
}

func TestScanOrdinaryFile(t *testing.T) {
	response := execute("scan", []byte("package main\nvar answer = 42\n"))
	if !response.OK || len(response.Diagnostics) != 0 {
		t.Fatalf("scan failed: %#v", response.Diagnostics)
	}
	want := []string{"package", "IDENT", ";", "var", "IDENT", "=", "INT", ";", "EOF"}
	if len(response.Tokens) != len(want) {
		t.Fatalf("got %d tokens, want %d", len(response.Tokens), len(want))
	}
	for index, kind := range want {
		if response.Tokens[index].Kind != kind {
			t.Fatalf("token %d: got %q, want %q", index, response.Tokens[index].Kind, kind)
		}
	}
}

func TestParseOrdinaryFile(t *testing.T) {
	response := execute("parse-file", []byte("package main\nfunc main() {}\n"))
	if !response.OK || len(response.Diagnostics) != 0 {
		t.Fatalf("parse failed: %#v", response.Diagnostics)
	}
	node, ok := response.AST.(normalizedNode)
	if !ok || node.Kind != "File" || node.Start != 0 || node.End != 27 {
		t.Fatalf("unexpected AST: %#v", response.AST)
	}
}

func TestParseExpression(t *testing.T) {
	response := execute("parse-expr", []byte("f[A, B](x + 1)"))
	if !response.OK || len(response.Diagnostics) != 0 {
		t.Fatalf("parse failed: %#v", response.Diagnostics)
	}
	node, ok := response.AST.(normalizedNode)
	if !ok || node.Kind != "CallExpr" || node.Start != 0 || node.End != 14 {
		t.Fatalf("unexpected AST: %#v", response.AST)
	}
}

func TestInvalidFileReturnsPartialASTAndDiagnostics(t *testing.T) {
	response := execute("parse-file", []byte("package main\nfunc main( {\n"))
	if response.OK || len(response.Diagnostics) == 0 {
		t.Fatalf("invalid input was accepted: %#v", response)
	}
	if _, ok := response.AST.(normalizedNode); !ok {
		t.Fatalf("missing partial AST: %#v", response.AST)
	}
}

func TestVersionReportsRunningToolchain(t *testing.T) {
	response := execute("version", nil)
	if response.GoVersion != runtime.Version() {
		t.Fatalf("got %q", response.GoVersion)
	}
	if response.GOROOT != runtime.GOROOT() {
		t.Fatalf("got %q", response.GOROOT)
	}
}

func TestTypeCheckPackageOutputIsDeterministic(t *testing.T) {
	input := typeCheckInput(`[{"name":"a.go","source":"package p\nvar X int\nfunc F(x int) int { return x + X }\n"}]`)
	first := encodedValue(t, typeCheckPackage(input))
	second := encodedValue(t, typeCheckPackage(input))
	if string(first) != string(second) {
		t.Fatalf("output changed between runs\nfirst: %s\nsecond: %s", first, second)
	}
}

func TestTypeCheckPackageSupportsMultipleFiles(t *testing.T) {
	response := typeCheckPackage(typeCheckInput(`[
		{"name":"a.go","source":"package p\nvar X int\n"},
		{"name":"b.go","source":"package p\nfunc F() int { return X }\n"}
	]`))
	if !response.OK || len(response.Diagnostics) != 0 {
		t.Fatalf("type check failed: %#v", response.Diagnostics)
	}
	if response.Package.Path != "example/p" || response.Package.Name != "p" || !response.Package.Complete {
		t.Fatalf("unexpected package: %#v", response.Package)
	}
	if len(response.PackageScope) != 2 || response.PackageScope[0].Name != "F" || response.PackageScope[1].Name != "X" {
		t.Fatalf("unexpected package scope: %#v", response.PackageScope)
	}
	if len(response.UniverseScope) != 44 || len(response.UnsafeScope) != 9 {
		t.Fatalf("unexpected predeclared scopes: universe=%d unsafe=%d", len(response.UniverseScope), len(response.UnsafeScope))
	}
	if len(response.Uses) == 0 || len(response.Types) == 0 || len(response.Scopes) == 0 {
		t.Fatalf("missing type information: %#v", response)
	}
}

func TestTypeCheckPackageSupportsStandardLibraryImports(t *testing.T) {
	response := typeCheckPackage(typeCheckInput(`[{"name":"a.go","source":"package p\nimport \"fmt\"\nvar X = fmt.Sprint(1)\n"}]`))
	if !response.OK || len(response.Diagnostics) != 0 {
		t.Fatalf("type check failed: %#v", response.Diagnostics)
	}
	if len(response.PackageScope) != 1 || response.PackageScope[0].Name != "X" || response.PackageScope[0].Type.Text != "string" {
		t.Fatalf("unexpected package scope: %#v", response.PackageScope)
	}
	importFound := false
	for _, fact := range response.Implicits {
		if fact.Object != nil && fact.Object.Kind == "PackageName" && fact.Object.ImportedPath == "fmt" {
			importFound = true
		}
	}
	if !importFound {
		t.Fatalf("missing fmt import: %#v", response.Implicits)
	}
}

func TestTypeCheckPackageReturnsInvalidFactsAndDiagnostics(t *testing.T) {
	response := typeCheckPackage(typeCheckInput(`[{"name":"a.go","source":"package p\nvar X int = \"bad\"\n"}]`))
	if response.OK || len(response.Diagnostics) == 0 {
		t.Fatalf("invalid package was accepted: %#v", response)
	}
	if response.Diagnostics[0].File != "a.go" || response.Diagnostics[0].Start != 22 {
		t.Fatalf("unexpected diagnostic: %#v", response.Diagnostics[0])
	}
	if len(response.PackageScope) != 1 || response.PackageScope[0].Name != "X" {
		t.Fatalf("missing partial package scope: %#v", response.PackageScope)
	}
}

func TestTypeCheckPackageCanSkipInfo(t *testing.T) {
	response := typeCheckPackage([]byte(`{"package_path":"example/p","go_version":"go1.26","goos":"linux","goarch":"amd64","files":[{"name":"a.go","source":"package p\nvar X int\n"}],"config":{"enable_alias":true,"collect_info":false}}`))
	if !response.OK || len(response.Diagnostics) != 0 {
		t.Fatalf("type check failed: %#v", response.Diagnostics)
	}
	if len(response.PackageScope) != 0 || len(response.Types) != 0 || len(response.Defs) != 0 || len(response.Uses) != 0 {
		t.Fatalf("unexpected type information: %#v", response)
	}
}

func TestConstantEvalPreservesArbitraryPrecision(t *testing.T) {
	response := constantEval([]byte(`{"expression":"(1 << 100) + 3","go_version":"go1.26"}`))
	if !response.OK || len(response.Diagnostics) != 0 {
		t.Fatalf("constant evaluation failed: %#v", response.Diagnostics)
	}
	if response.Type == nil || response.Type.Text != "untyped int" {
		t.Fatalf("unexpected type: %#v", response.Type)
	}
	if response.Value == nil || response.Value.Kind != "Int" || response.Value.Exact != "1267650600228229401496703205379" {
		t.Fatalf("unexpected value: %#v", response.Value)
	}
}

func TestConstantEvalReportsInvalidExpression(t *testing.T) {
	response := constantEval([]byte(`{"expression":"1 / 0","go_version":"go1.26"}`))
	if response.OK || len(response.Diagnostics) == 0 || response.Value != nil {
		t.Fatalf("invalid constant was accepted: %#v", response)
	}
}
