package protocol

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"go/token"
	"io"
	"path/filepath"
	"strings"
)

const Version = 1
const MaxMessageBytes = 8 << 20
const MaxTypeNodes = 65536
const MaxTypeDepth = 64
const MaxBindings = 4096

type BuildContext struct {
	GoExecutable string   `json:"go_executable"`
	Toolchain    string   `json:"toolchain"`
	ModuleDir    string   `json:"module_dir"`
	GOOS         string   `json:"goos"`
	GOARCH       string   `json:"goarch"`
	CGOEnabled   string   `json:"cgo_enabled"`
	GOFLAGS      string   `json:"goflags"`
	BuildTags    []string `json:"build_tags"`
	GO111MODULE  string   `json:"go111module"`
	GOWORK       string   `json:"gowork"`
	Dependencies string   `json:"dependencies"`
	Network      string   `json:"network"`
}

type CallerContext struct {
	GeneratedSource string `json:"generated_source,omitempty"`
	LoadMode        string `json:"load_mode"`
	ImportPath      string `json:"import_path"`
	Directory       string `json:"directory"`
	Package         string `json:"package"`
}

type FunctionSignature struct {
	Parameters []Type `json:"parameters"`
	Results    []Type `json:"results"`
}

type Type struct {
	Signature  *FunctionSignature `json:"signature,omitempty"`
	Key        *Type              `json:"key,omitempty"`
	ImportPath string             `json:"import_path,omitempty"`
	Name       string             `json:"name,omitempty"`
	Arguments  []Type             `json:"type_arguments,omitempty"`
	Tag        string             `json:"tag"`
	Element    *Type              `json:"element,omitempty"`
	Length     *int64             `json:"length,omitempty"`
	Direction  string             `json:"direction,omitempty"`
}

type Binding struct {
	ID         string `json:"binding_id"`
	ImportPath string `json:"import_path"`
	Symbol     string `json:"symbol"`
	Parameters []Type `json:"bridge_parameter_types"`
	Results    []Type `json:"bridge_result_types"`
	CallMode   string `json:"call_mode"`
}

type TypeQuery struct {
	Mode       string `json:"query_mode,omitempty"`
	ID         string `json:"type_id"`
	ImportPath string `json:"import_path"`
	Name       string `json:"name"`
	Arguments  []Type `json:"type_arguments,omitempty"`
}

type TypeResult struct {
	ID           string       `json:"type_id"`
	Status       string       `json:"status"`
	Kind         string       `json:"declaration_kind,omitempty"`
	DeclaredType string       `json:"declared_type,omitempty"`
	Type         string       `json:"type,omitempty"`
	Diagnostics  []Diagnostic `json:"diagnostics,omitempty"`
}

type SourceOverlay struct {
	Path   string `json:"path"`
	Source string `json:"source"`
}

type Request struct {
	SourceOverlay   *SourceOverlay `json:"source_overlay,omitempty"`
	TypeQueries     []TypeQuery    `json:"type_queries,omitempty"`
	ProtocolVersion int            `json:"protocol_version"`
	BuildContext    BuildContext   `json:"build_context"`
	CallerContext   CallerContext  `json:"caller_context"`
	Bindings        []Binding      `json:"bindings"`
}

type Diagnostic struct {
	Code      string `json:"code"`
	Message   string `json:"message"`
	Direction string `json:"direction,omitempty"`
	Index     *int   `json:"index,omitempty"`
}

type BindingResult struct {
	ID              string       `json:"binding_id"`
	Status          string       `json:"status"`
	ActualSignature string       `json:"actual_signature"`
	SignatureType   string       `json:"signature_type,omitempty"`
	Diagnostics     []Diagnostic `json:"diagnostics"`
}

type TypeField struct {
	Name       string `json:"name"`
	ImportPath string `json:"import_path,omitempty"`
	Type       string `json:"type"`
	Tag        string `json:"tag,omitempty"`
	Embedded   bool   `json:"embedded,omitempty"`
}

type TypeTerm struct {
	Type  string `json:"type"`
	Tilde bool   `json:"tilde,omitempty"`
}

type ReferencedType struct {
	RuntimeInterface *bool        `json:"runtime_interface,omitempty"`
	InterfaceMethods *[]TypeField `json:"interface_methods,omitempty"`
	ID               string       `json:"id"`
	Tag              string       `json:"tag"`
	ImportPath       string       `json:"import_path,omitempty"`
	Name             string       `json:"name,omitempty"`
	Element          string       `json:"element,omitempty"`
	Key              string       `json:"key,omitempty"`
	Underlying       string       `json:"underlying,omitempty"`
	Origin           string       `json:"origin,omitempty"`
	Length           *int64       `json:"length,omitempty"`
	Direction        string       `json:"direction,omitempty"`
	Parameters       []string     `json:"parameters,omitempty"`
	Results          []string     `json:"results,omitempty"`
	TypeArguments    []string     `json:"type_arguments,omitempty"`
	TypeParameters   []string     `json:"type_parameters,omitempty"`
	Receiver         string       `json:"receiver,omitempty"`
	Variadic         bool         `json:"variadic,omitempty"`
	Fields           []TypeField  `json:"fields,omitempty"`
	Methods          []TypeField  `json:"methods,omitempty"`
	Embedded         []string     `json:"embedded,omitempty"`
	Constraint       string       `json:"constraint,omitempty"`
	Terms            []TypeTerm   `json:"terms,omitempty"`
}

type Response struct {
	TypeResults     []TypeResult     `json:"type_results,omitempty"`
	ProtocolVersion int              `json:"protocol_version"`
	GoWorldIdentity string           `json:"go_world_identity"`
	ReferencedTypes []ReferencedType `json:"referenced_types"`
	Bindings        []BindingResult  `json:"bindings"`
	Diagnostics     []Diagnostic     `json:"diagnostics"`
}

func DecodeRequest(ctx context.Context, input io.Reader) (Request, error) {
	var request Request
	if err := ctx.Err(); err != nil {
		return request, err
	}
	data, err := io.ReadAll(io.LimitReader(input, MaxMessageBytes+1))
	if err != nil {
		return request, fmt.Errorf("ffi-protocol: read request: %w", err)
	}
	if err := ctx.Err(); err != nil {
		return request, err
	}
	if len(data) > MaxMessageBytes {
		return request, fmt.Errorf("ffi-protocol: request exceeds %d bytes", MaxMessageBytes)
	}
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&request); err != nil {
		return request, fmt.Errorf("ffi-protocol: decode request: %w", err)
	}
	var trailing any
	if err := decoder.Decode(&trailing); err != io.EOF {
		return request, fmt.Errorf("ffi-protocol: expected exactly one JSON request")
	}
	return request, request.Validate()
}

func (r Request) Validate() error {
	if r.ProtocolVersion != Version {
		return fmt.Errorf("ffi-protocol: unsupported protocol version %d", r.ProtocolVersion)
	}
	c := r.BuildContext
	if c.GoExecutable == "" || c.Toolchain == "" || c.ModuleDir == "" || c.GOOS == "" || c.GOARCH == "" {
		return fmt.Errorf("ffi-protocol: incomplete Go build context")
	}
	if value := r.SourceOverlay; value != nil {
		relative, err := filepath.Rel(c.ModuleDir, value.Path)
		name := filepath.Base(value.Path)
		if !filepath.IsAbs(value.Path) || filepath.Clean(value.Path) != value.Path || err != nil || relative == ".." || strings.HasPrefix(relative, ".."+string(filepath.Separator)) || filepath.Ext(name) != ".go" || strings.HasPrefix(name, "goml_ffi_witness_") || strings.HasPrefix(name, "goml_type_query_") || value.Source == "" || len(value.Source) > MaxMessageBytes {
			return fmt.Errorf("ffi-protocol: invalid module source overlay")
		}
	}
	if c.GOWORK != "off" || (c.GO111MODULE != "on" && c.GO111MODULE != "off") || (c.CGOEnabled != "0" && c.CGOEnabled != "1") {
		return fmt.Errorf("ffi-protocol: invalid Go build mode")
	}
	if c.Dependencies != "readonly" || c.Network != "off" {
		return fmt.Errorf("ffi-protocol: v1 requires readonly dependencies and network off")
	}
	if r.CallerContext.ImportPath == "" || r.CallerContext.Directory == "" || r.CallerContext.Package == "" {
		return fmt.Errorf("ffi-protocol: incomplete caller context")
	}
	if r.CallerContext.GeneratedSource != "" && r.CallerContext.LoadMode != "package" {
		return fmt.Errorf("ffi-protocol: generated source requires package mode")
	}
	if r.CallerContext.LoadMode != "package" && r.CallerContext.LoadMode != "files" {
		return fmt.Errorf("ffi-protocol: caller load mode must be package or files")
	}
	if r.CallerContext.LoadMode == "files" && r.CallerContext.ImportPath != "command-line-arguments" {
		return fmt.Errorf("ffi-protocol: file callers must use command-line-arguments identity")
	}
	if len(r.Bindings)+len(r.TypeQueries) > MaxBindings {
		return fmt.Errorf("ffi-protocol: too many bindings")
	}
	ids := make(map[string]bool)
	nodes := 0
	for _, b := range r.Bindings {
		if b.ID == "" || ids[b.ID] {
			return fmt.Errorf("ffi-protocol: empty or duplicate binding id %q", b.ID)
		}
		ids[b.ID] = true
		if b.Symbol == "" || (b.CallMode != "ordinary" && b.CallMode != "method") ||
			(b.CallMode == "ordinary" && b.ImportPath == "") ||
			(b.CallMode == "method" && (b.ImportPath != "" || len(b.Parameters) == 0)) {
			return fmt.Errorf("ffi-protocol: binding %q has invalid target or call mode", b.ID)
		}
		if len(b.Results) > 1 {
			for _, result := range b.Results {
				if result.Tag == "any" {
					return fmt.Errorf("ffi-protocol: binding %q: marker data is not supported in tuple results", b.ID)
				}
			}
		}
		for _, group := range [][]Type{b.Parameters, b.Results} {
			for _, ty := range group {
				if err := validateType(ty, 0, &nodes); err != nil {
					return fmt.Errorf("ffi-protocol: binding %q: %w", b.ID, err)
				}
			}
		}
	}
	queryIDs := map[string]bool{}
	for _, q := range r.TypeQueries {
		if q.ID == "" || queryIDs[q.ID] || q.ImportPath == "" || q.Name == "" {
			return fmt.Errorf("ffi-protocol: invalid or duplicate type query %q", q.ID)
		}
		if q.Mode != "" && q.Mode != "instance" && q.Mode != "declaration" && q.Mode != "function" {
			return fmt.Errorf("ffi-protocol: invalid type query mode")
		}
		if q.Mode == "declaration" && len(q.Arguments) != 0 {
			return fmt.Errorf("ffi-protocol: declaration queries cannot have type arguments")
		}
		queryIDs[q.ID] = true
		for _, argument := range q.Arguments {
			if err := validateType(argument, 0, &nodes); err != nil {
				return fmt.Errorf("ffi-protocol: type query %q: %w", q.ID, err)
			}
		}
	}
	return nil
}

func validateType(t Type, depth int, nodes *int) error {
	*nodes += 1
	if depth > MaxTypeDepth || *nodes > MaxTypeNodes {
		return fmt.Errorf("type complexity limit exceeded")
	}
	if t.Tag == "any" && depth > 0 {
		return fmt.Errorf("marker data is only supported as a direct value")
	}
	container := false
	switch t.Tag {
	case "bool", "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64", "float32", "float64", "string", "any":
	case "named":
		if (t.ImportPath == "" && (t.Name != "error" || len(t.Arguments) != 0)) || !token.IsIdentifier(t.Name) || t.Name == "_" {
			return fmt.Errorf("invalid named type")
		}
	case "function":
	case "pointer", "array", "slice", "channel", "map":
		container = true
	default:
		return fmt.Errorf("unknown type tag %q", t.Tag)
	}
	if (t.Tag == "function") != (t.Signature != nil) {
		return fmt.Errorf("invalid signature for %s", t.Tag)
	}
	if t.Signature != nil {
		if t.Signature.Parameters == nil || t.Signature.Results == nil {
			return fmt.Errorf("function signature requires parameter and result arrays")
		}
		for _, group := range [][]Type{t.Signature.Parameters, t.Signature.Results} {
			for _, child := range group {
				if child.Tag == "any" {
					return fmt.Errorf("marker data is not supported in function signatures")
				}
				if err := validateType(child, depth+1, nodes); err != nil {
					return err
				}
			}
		}
	}
	if t.Tag != "named" && (t.ImportPath != "" || t.Name != "" || len(t.Arguments) != 0) {
		return fmt.Errorf("named metadata on %s", t.Tag)
	}
	for _, argument := range t.Arguments {
		if err := validateType(argument, depth+1, nodes); err != nil {
			return err
		}
	}
	if (t.Tag == "map") != (t.Key != nil) {
		return fmt.Errorf("invalid key for %s", t.Tag)
	}
	if t.Key != nil {
		if t.Key.Tag == "any" {
			return fmt.Errorf("marker data is only supported as a direct value")
		}
		if err := validateType(*t.Key, depth+1, nodes); err != nil {
			return err
		}
	}
	if container != (t.Element != nil) {
		return fmt.Errorf("invalid element for %s", t.Tag)
	}
	if (t.Tag == "array") != (t.Length != nil) || (t.Length != nil && *t.Length < 0) {
		return fmt.Errorf("invalid array length")
	}
	if t.Tag == "channel" {
		if t.Direction != "both" && t.Direction != "send" && t.Direction != "receive" {
			return fmt.Errorf("invalid channel direction %q", t.Direction)
		}
	} else if t.Direction != "" {
		return fmt.Errorf("direction on non-channel type")
	}
	if t.Element != nil {
		if t.Element.Tag == "any" {
			return fmt.Errorf("marker data is only supported as a direct value")
		}
		return validateType(*t.Element, depth+1, nodes)
	}
	return nil
}

func EncodeResponse(output io.Writer, response Response) error {
	if response.ProtocolVersion != Version {
		return fmt.Errorf("ffi-protocol: unsupported response version")
	}
	if len(response.ReferencedTypes) > MaxTypeNodes || len(response.Bindings)+len(response.TypeResults) > MaxBindings {
		return fmt.Errorf("ffi-protocol: response complexity limit exceeded")
	}
	data, err := json.Marshal(response)
	if err != nil {
		return err
	}
	if len(data)+1 > MaxMessageBytes {
		return fmt.Errorf("ffi-protocol: response size limit exceeded")
	}
	_, err = output.Write(append(data, '\n'))
	return err
}
