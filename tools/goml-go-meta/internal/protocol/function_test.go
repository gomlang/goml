package protocol

import (
	"bytes"
	"context"
	"encoding/json"
	"reflect"
	"testing"
)

func functionType(parameters, results []Type) Type {
	return Type{Tag: "function", Signature: &FunctionSignature{Parameters: parameters, Results: results}}
}

func TestFunctionSignatureProtocol(t *testing.T) {
	empty := functionType([]Type{}, []Type{})
	nested := functionType([]Type{empty, {Tag: "string"}}, []Type{{Tag: "int64"}, {Tag: "named", Name: "error"}})
	r := validRequest()
	r.Bindings[0].Parameters = []Type{nested}
	r.Bindings[0].Results = []Type{empty}
	data, err := json.Marshal(r)
	if err != nil {
		t.Fatal(err)
	}
	decoded, err := DecodeRequest(context.Background(), bytes.NewReader(data))
	if err != nil || !reflect.DeepEqual(decoded, r) {
		t.Fatalf("round trip: %v, %#v", err, decoded)
	}
	for name, ty := range map[string]Type{
		"missing signature":           {Tag: "function"},
		"extra signature":             {Tag: "int64", Signature: empty.Signature},
		"missing parameters":          functionType(nil, []Type{}),
		"missing results":             functionType([]Type{}, nil),
		"parameter marker":            functionType([]Type{{Tag: "any"}}, []Type{}),
		"result marker":               functionType([]Type{}, []Type{{Tag: "any"}}),
		"marker in generic parameter": functionType([]Type{{Tag: "named", ImportPath: "example.com/host", Name: "Box", Arguments: []Type{{Tag: "any"}}}}, []Type{}),
		"bad child":                   functionType([]Type{{Tag: "not_a_type"}}, []Type{}),
		"container metadata":          {Tag: "function", Signature: empty.Signature, Element: &empty},
	} {
		t.Run(name, func(t *testing.T) {
			r.Bindings[0].Parameters = []Type{ty}
			if r.Validate() == nil {
				t.Fatal("accepted malformed function signature")
			}
		})
	}
	deep := empty
	for i := 0; i <= MaxTypeDepth; i++ {
		deep = functionType([]Type{deep}, []Type{})
	}
	r.Bindings[0].Parameters = []Type{deep}
	if r.Validate() == nil {
		t.Fatal("accepted excessive signature depth")
	}
	r.Bindings[0].Parameters = []Type{functionType(make([]Type, MaxTypeNodes+1), []Type{})}
	for i := range r.Bindings[0].Parameters[0].Signature.Parameters {
		r.Bindings[0].Parameters[0].Signature.Parameters[i] = Type{Tag: "int64"}
	}
	if r.Validate() == nil {
		t.Fatal("accepted excessive signature nodes")
	}
}
