package protocol

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"reflect"
	"strings"
	"testing"
)

func validRequest() Request {
	return Request{
		ProtocolVersion: Version,
		BuildContext: BuildContext{
			GoExecutable: `C:\Program Files\Go\bin\go.exe`, Toolchain: "go1.25.8",
			ModuleDir: `C:\my project`, GOOS: "windows", GOARCH: "amd64", CGOEnabled: "0",
			GO111MODULE: "on", GOWORK: "off", Dependencies: "readonly", Network: "off",
		},
		CallerContext: CallerContext{LoadMode: "package", ImportPath: "example.com/host/gen", Directory: `C:\my project\gen`, Package: "gen"},
		Bindings:      []Binding{{ID: "demo::calc::abs", ImportPath: "math", Symbol: "Abs", Parameters: []Type{{Tag: "float64"}}, Results: []Type{{Tag: "float64"}}, CallMode: "ordinary"}},
	}
}

func TestRoundTrip(t *testing.T) {
	want := validRequest()
	data, err := json.Marshal(want)
	if err != nil {
		t.Fatal(err)
	}
	got, err := DecodeRequest(context.Background(), bytes.NewReader(data))
	if err != nil || !reflect.DeepEqual(got, want) {
		t.Fatalf("round trip: %v, %#v", err, got)
	}
}

func TestMalformedRequests(t *testing.T) {
	valid, err := json.Marshal(validRequest())
	if err != nil {
		t.Fatal(err)
	}
	for name, data := range map[string][]byte{
		"truncated":     valid[:len(valid)-1],
		"trailing":      append(append([]byte{}, valid...), []byte(" {}")...),
		"unknown field": []byte(`{"protocol_version":1,"surprise":true}`),
		"oversized":     bytes.Repeat([]byte(" "), MaxMessageBytes+1),
	} {
		t.Run(name, func(t *testing.T) {
			if _, err := DecodeRequest(context.Background(), bytes.NewReader(data)); err == nil {
				t.Fatal("accepted malformed request")
			}
		})
	}
}

func TestInvalidRequests(t *testing.T) {
	cases := map[string]func(*Request){
		"version":           func(r *Request) { r.ProtocolVersion++ },
		"duplicate id":      func(r *Request) { r.Bindings = append(r.Bindings, r.Bindings[0]) },
		"unknown tag":       func(r *Request) { r.Bindings[0].Parameters[0].Tag = "injected()" },
		"missing element":   func(r *Request) { r.Bindings[0].Parameters[0].Tag = "slice" },
		"extra element":     func(r *Request) { r.Bindings[0].Parameters[0].Element = &Type{Tag: "bool"} },
		"channel direction": func(r *Request) { r.Bindings[0].Parameters[0] = Type{Tag: "channel", Element: &Type{Tag: "int"}} },
		"nested marker":     func(r *Request) { r.Bindings[0].Parameters[0] = Type{Tag: "slice", Element: &Type{Tag: "any"}} },
		"negative array": func(r *Request) {
			n := int64(-1)
			r.Bindings[0].Parameters[0] = Type{Tag: "array", Length: &n, Element: &Type{Tag: "int"}}
		},
		"workspace": func(r *Request) { r.BuildContext.GOWORK = "auto" },
		"network":   func(r *Request) { r.BuildContext.Network = "on" },
		"mode":      func(r *Request) { r.Bindings[0].CallMode = "spread" },
		"depth": func(r *Request) {
			ty := Type{Tag: "int"}
			for i := 0; i < MaxTypeDepth+2; i++ {
				inner := ty
				ty = Type{Tag: "slice", Element: &inner}
			}
			r.Bindings[0].Parameters[0] = ty
		},
	}
	for name, mutate := range cases {
		t.Run(name, func(t *testing.T) {
			r := validRequest()
			mutate(&r)
			if err := r.Validate(); err == nil {
				t.Fatal("accepted invalid request")
			}
		})
	}
}

func TestCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if _, err := DecodeRequest(ctx, strings.NewReader("")); !errors.Is(err, context.Canceled) {
		t.Fatalf("expected cancellation, got %v", err)
	}
}

func TestStableResponse(t *testing.T) {
	r := Response{ProtocolVersion: Version, Bindings: []BindingResult{{ID: "abs", Status: "verified", ActualSignature: "func(float64) float64"}}}
	var a, b bytes.Buffer
	if err := EncodeResponse(&a, r); err != nil {
		t.Fatal(err)
	}
	if err := EncodeResponse(&b, r); err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(a.Bytes(), b.Bytes()) {
		t.Fatal("unstable encoding")
	}
	r.Bindings[0].ActualSignature = strings.Repeat("x", MaxMessageBytes)
	b.Reset()
	if err := EncodeResponse(&b, r); err == nil || b.Len() != 0 {
		t.Fatal("oversized response was written")
	}
}

func TestBridgeShapes(t *testing.T) {
	for _, tag := range []string{"bool", "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64", "float32", "float64", "string", "any"} {
		t.Run(tag, func(t *testing.T) {
			r := validRequest()
			r.Bindings[0].Parameters = []Type{{Tag: tag}}
			r.Bindings[0].Results = []Type{{Tag: tag}}
			if err := r.Validate(); err != nil {
				t.Fatal(err)
			}
		})
	}
	length := int64(2)
	for _, ty := range []Type{
		{Tag: "array", Length: &length, Element: &Type{Tag: "int32"}},
		{Tag: "slice", Element: &Type{Tag: "uint8"}},
		{Tag: "channel", Direction: "both", Element: &Type{Tag: "string"}},
		{Tag: "channel", Direction: "send", Element: &Type{Tag: "string"}},
		{Tag: "channel", Direction: "receive", Element: &Type{Tag: "string"}},
	} {
		r := validRequest()
		r.Bindings[0].Parameters = []Type{ty}
		r.Bindings[0].Results = nil
		if err := r.Validate(); err != nil {
			t.Fatal(err)
		}
	}
	r := validRequest()
	r.Bindings[0].Results = []Type{{Tag: "string"}, {Tag: "bool"}}
	if err := r.Validate(); err != nil {
		t.Fatal(err)
	}
	r.Bindings[0].Results[0] = Type{Tag: "any"}
	if err := r.Validate(); err == nil {
		t.Fatal("accepted marker in tuple")
	}
}

func TestTypeNodeLimit(t *testing.T) {
	r := validRequest()
	r.Bindings[0].Parameters = make([]Type, MaxTypeNodes+1)
	for i := range r.Bindings[0].Parameters {
		r.Bindings[0].Parameters[i] = Type{Tag: "int"}
	}
	if err := r.Validate(); err == nil {
		t.Fatal("accepted excessive type nodes")
	}
}

func TestTypeQueryProtocol(t *testing.T) {
	r := validRequest()
	r.TypeQueries = []TypeQuery{{ID: "type", ImportPath: "time", Name: "Duration"}}
	data, err := json.Marshal(r)
	if err != nil {
		t.Fatal(err)
	}
	decoded, err := DecodeRequest(context.Background(), bytes.NewReader(data))
	if err != nil || !reflect.DeepEqual(r, decoded) {
		t.Fatalf("type request roundtrip: %v", err)
	}
	r.TypeQueries = append(r.TypeQueries, r.TypeQueries[0])
	if r.Validate() == nil {
		t.Fatal("duplicate type id")
	}
	r.TypeQueries = r.TypeQueries[:1]
	r.TypeQueries[0].Arguments = []Type{{Tag: "unknown"}}
	if r.Validate() == nil {
		t.Fatal("invalid type argument")
	}
	r.TypeQueries = make([]TypeQuery, MaxBindings)
	if r.Validate() == nil {
		t.Fatal("combined request bound")
	}
}

func TestNamedTypeValidation(t *testing.T) {
	r := validRequest()
	r.TypeQueries = []TypeQuery{{ID: "named", ImportPath: "host", Name: "Box", Arguments: []Type{{Tag: "named", ImportPath: "time", Name: "Duration"}}}}
	if err := r.Validate(); err != nil {
		t.Fatal(err)
	}
	for _, argument := range []Type{
		{Tag: "named", Name: "Duration"},
		{Tag: "named", ImportPath: "time", Name: "Duration;panic"},
		{Tag: "named", ImportPath: "time", Name: "Duration", Element: &Type{Tag: "int64"}},
		{Tag: "int64", ImportPath: "time", Name: "Duration"},
	} {
		r.TypeQueries[0].Arguments = []Type{argument}
		if r.Validate() == nil {
			t.Fatalf("invalid argument: %+v", argument)
		}
	}
}

func TestTypeQueryModes(t *testing.T) {
	r := validRequest()
	r.TypeQueries = []TypeQuery{{ID: "type", ImportPath: "host", Name: "Box", Mode: "declaration"}}
	if err := r.Validate(); err != nil {
		t.Fatal(err)
	}
	r.TypeQueries[0].Arguments = []Type{{Tag: "int64"}}
	if r.Validate() == nil {
		t.Fatal("declaration accepted instantiation arguments")
	}
	r.TypeQueries[0].Arguments = nil
	r.TypeQueries[0].Mode = "guess"
	if r.Validate() == nil {
		t.Fatal("unknown mode")
	}
}

func TestPointerTypeValidation(t *testing.T) {
	for _, ty := range []Type{
		{Tag: "pointer"},
		{Tag: "pointer", Element: &Type{Tag: "any"}},
		{Tag: "pointer", Element: &Type{Tag: "int64"}, Direction: "both"},
	} {
		r := validRequest()
		r.Bindings[0].Parameters = []Type{ty}
		data, err := json.Marshal(r)
		if err != nil {
			t.Fatal(err)
		}
		if _, err := DecodeRequest(context.Background(), bytes.NewReader(data)); err == nil {
			t.Fatalf("accepted invalid pointer: %+v", ty)
		}
	}
	r := validRequest()
	r.Bindings[0].Parameters = []Type{{Tag: "pointer", Element: &Type{Tag: "pointer", Element: &Type{Tag: "int64"}}}}
	data, err := json.Marshal(r)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := DecodeRequest(context.Background(), bytes.NewReader(data)); err != nil {
		t.Fatal(err)
	}
}

func TestMethodBindingRequiresReceiverAndNoPackageTarget(t *testing.T) {
	r := validRequest()
	r.Bindings[0].CallMode = "method"
	r.Bindings[0].ImportPath = ""
	r.Bindings[0].Parameters = []Type{{Tag: "named", ImportPath: "time", Name: "Duration"}}
	if err := r.Validate(); err != nil {
		t.Fatal(err)
	}
	r.Bindings[0].ImportPath = "time"
	if r.Validate() == nil {
		t.Fatal("accepted a package target for a receiver method")
	}
	r.Bindings[0].ImportPath = ""
	r.Bindings[0].Parameters = nil
	if r.Validate() == nil {
		t.Fatal("accepted a method without a receiver")
	}
}

func TestBuiltinErrorNamedIdentity(t *testing.T) {
	r := validRequest()
	r.Bindings[0].Parameters = []Type{{Tag: "named", Name: "error"}}
	if err := r.Validate(); err != nil {
		t.Fatal(err)
	}
	for _, invalid := range []Type{
		{Tag: "named", Name: "any"},
		{Tag: "named", Name: "Error"},
		{Tag: "named", Name: "error", Arguments: []Type{{Tag: "int64"}}},
	} {
		r.Bindings[0].Parameters = []Type{invalid}
		if r.Validate() == nil {
			t.Fatalf("accepted invalid builtin identity: %+v", invalid)
		}
	}
}

func TestMapTypeProtocol(t *testing.T) {
	key := Type{Tag: "string"}
	value := Type{Tag: "slice", Element: &Type{Tag: "uint8"}}
	r := validRequest()
	r.Bindings[0].Parameters = []Type{{Tag: "map", Key: &key, Element: &value}}
	data, err := json.Marshal(r)
	if err != nil {
		t.Fatal(err)
	}
	got, err := DecodeRequest(context.Background(), bytes.NewReader(data))
	if err != nil || !reflect.DeepEqual(got, r) {
		t.Fatalf("map round trip: %v, %#v", err, got)
	}
	for name, ty := range map[string]Type{
		"missing key":       {Tag: "map", Element: &value},
		"missing value":     {Tag: "map", Key: &key},
		"extra key":         {Tag: "slice", Key: &key, Element: &value},
		"marker key":        {Tag: "map", Key: &Type{Tag: "any"}, Element: &value},
		"marker value":      {Tag: "map", Key: &key, Element: &Type{Tag: "any"}},
		"nested marker key": {Tag: "map", Key: &Type{Tag: "pointer", Element: &Type{Tag: "any"}}, Element: &value},
		"unknown key":       {Tag: "map", Key: &Type{Tag: "unknown"}, Element: &value},
	} {
		t.Run(name, func(t *testing.T) {
			invalid := validRequest()
			invalid.Bindings[0].Parameters = []Type{ty}
			if err := invalid.Validate(); err == nil {
				t.Fatal("accepted invalid map type")
			}
		})
	}
	deep := key
	for i := 0; i < MaxTypeDepth+1; i++ {
		inner := deep
		deep = Type{Tag: "map", Key: &inner, Element: &key}
	}
	r.Bindings[0].Parameters = []Type{deep}
	if err := r.Validate(); err == nil {
		t.Fatal("accepted deep map keys")
	}
}
