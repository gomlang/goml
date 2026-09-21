package cbind

import (
	"fmt"
	"go/format"
	"path/filepath"
	"strconv"
	"strings"
)

type representation struct {
	gom       string
	goType    string
	cgo       string
	handle    string
	canonical string
}

func cgoType(value string) (string, error) {
	words := strings.Fields(strings.ReplaceAll(value, "*", " * "))
	filtered := []string{}
	for _, word := range words {
		if word != "const" && word != "volatile" && word != "restrict" {
			filtered = append(filtered, word)
		}
	}
	pointers := 0
	for len(filtered) > 0 && filtered[len(filtered)-1] == "*" {
		pointers++
		filtered = filtered[:len(filtered)-1]
	}
	base := strings.Join(filtered, " ")
	names := map[string]string{
		"signed char": "schar", "unsigned char": "uchar", "short int": "short",
		"signed short": "short", "signed short int": "short", "unsigned short": "ushort",
		"unsigned short int": "ushort", "signed": "int", "signed int": "int",
		"unsigned": "uint", "unsigned int": "uint", "long int": "long",
		"signed long": "long", "signed long int": "long", "unsigned long": "ulong",
		"unsigned long int": "ulong", "long long": "longlong", "long long int": "longlong",
		"signed long long": "longlong", "signed long long int": "longlong",
		"unsigned long long": "ulonglong", "unsigned long long int": "ulonglong", "_Bool": "bool",
	}
	if base == "void" && pointers == 1 {
		return "unsafe.Pointer", nil
	}
	if name := names[base]; name != "" {
		base = name
	} else if strings.HasPrefix(base, "struct ") || strings.HasPrefix(base, "enum ") {
		base = strings.ReplaceAll(base, " ", "_")
	}
	if !identifier.MatchString(base) {
		return "", fmt.Errorf("unsupported C type %q", value)
	}
	return strings.Repeat("*", pointers) + "C." + base, nil
}

func (w *World) representation(value, requested string) (representation, error) {
	canonical := w.normalize(value)
	cgo, err := cgoType(value)
	if err != nil {
		return representation{}, err
	}
	handle := ""
	for _, ty := range w.Project.Config.Types {
		if w.Handles[ty.Name] == canonical && (requested == "" || requested == ty.Name) {
			if handle != "" {
				return representation{}, fmt.Errorf("ambiguous opaque type %s; select type explicitly", value)
			}
			handle = ty.Name
		}
	}
	if requested != "" && handle == "" {
		return representation{}, fmt.Errorf("C type %s does not match opaque type %s", value, requested)
	}
	if handle != "" {
		return representation{gom: handle, goType: handle, cgo: cgo, handle: handle, canonical: canonical}, nil
	}
	if canonical == "void" {
		return representation{gom: "()", canonical: canonical}, nil
	}
	if strings.HasPrefix(canonical, "enum ") {
		return representation{gom: "i64", goType: "int64", cgo: cgo, canonical: canonical}, nil
	}
	if canonical == "_Bool" || canonical == "bool" {
		return representation{gom: "bool", goType: "bool", cgo: cgo, canonical: canonical}, nil
	}
	if canonical == "float" {
		return representation{gom: "f32", goType: "float32", cgo: cgo, canonical: canonical}, nil
	}
	if canonical == "double" {
		return representation{gom: "f64", goType: "float64", cgo: cgo, canonical: canonical}, nil
	}
	unsigned := canonical == "unsigned" || strings.HasPrefix(canonical, "unsigned ")
	base := strings.TrimPrefix(strings.TrimPrefix(canonical, "unsigned "), "signed ")
	base = strings.TrimSuffix(base, " int")
	if base == "" || base == "unsigned" || base == "signed" {
		base = "int"
	}
	if base == "char" && canonical == "char" {
		unsigned = w.Sizes["char_signed"] == 2
	}
	size := w.Sizes[strings.ReplaceAll(base, " ", "_")]
	if size != 1 && size != 2 && size != 4 && size != 8 {
		return representation{}, fmt.Errorf("unsupported C type %q; pointers require an explicit mapping", value)
	}
	prefix, goPrefix := "i", "int"
	if unsigned {
		prefix, goPrefix = "u", "uint"
	}
	return representation{gom: fmt.Sprintf("%s%d", prefix, size*8), goType: fmt.Sprintf("%s%d", goPrefix, size*8), cgo: cgo, canonical: canonical}, nil
}

type generated struct {
	Goml []byte
	Go   []byte
}

type generator struct {
	w          *World
	gom        strings.Builder
	native     strings.Builder
	assertions strings.Builder
}

func (g *generator) abi(ty string, rep representation) {
	if rep.handle != "" || rep.gom == "()" {
		return
	}
	condition := ""
	if strings.HasPrefix(rep.canonical, "enum ") {
		condition = fmt.Sprintf("sizeof(%s) <= 4", ty)
	} else if rep.gom == "bool" {
		condition = fmt.Sprintf("sizeof(%s) == sizeof(_Bool)", ty)
	} else {
		bits := strings.TrimLeft(rep.gom, "iuf")
		width, _ := strconv.Atoi(bits)
		condition = fmt.Sprintf("sizeof(%s) == %d", ty, width/8)
		if rep.gom[0] == 'i' || rep.gom[0] == 'u' {
			sign := "1"
			if rep.gom[0] == 'u' {
				sign = "0"
			}
			condition += fmt.Sprintf(" && (((%s)-1 < (%s)0) == %s)", ty, ty, sign)
		}
	}
	fmt.Fprintf(&g.assertions, "_Static_assert(%s, \"C ABI changed; regenerate GoML bindings\");\n", condition)
}

func rawGom(rep representation) string {
	if rep.handle != "" {
		return "GomlC" + rep.handle
	}
	return rep.gom
}

func goZero(ty string) string {
	switch {
	case ty == "string":
		return "\"\""
	case ty == "bool":
		return "false"
	case ty == "[]byte":
		return "nil"
	case strings.HasPrefix(ty, "int"), strings.HasPrefix(ty, "uint"), strings.HasPrefix(ty, "float"):
		return "0"
	default:
		return ty + "{}"
	}
}

type output struct {
	goTypes  []string
	rawTypes []string
	gomType  string
	goValues []string
	gomValue string
}

func (g *generator) outputValue(rep representation, expression string, index int) output {
	gom := fmt.Sprintf("r%d", index)
	value := rep.goType + "(" + expression + ")"
	if rep.handle != "" {
		value = rep.goType + "{raw: " + expression + "}"
		gom = rep.handle + " { raw: " + gom + " }"
	}
	return output{goTypes: []string{rep.goType}, rawTypes: []string{rawGom(rep)}, gomType: rep.gom, goValues: []string{value}, gomValue: gom}
}

func (g *generator) stringOutput(expression string, index int) output {
	return output{
		goTypes: []string{"string", "bool"}, rawTypes: []string{"ffi::String", "bool"},
		gomType: "Option[c::CString]", goValues: []string{expression + "Text", expression + " != nil"},
		gomValue: fmt.Sprintf("if r%d { Option::Some(c::CString::from_raw(r%d)?) } else { Option::None }", index+1, index),
	}
}

func (g *generator) release(symbol, pointer string) (string, error) {
	if symbol == "" {
		return "", nil
	}
	node, exists := g.w.Functions[symbol]
	params := parameters(node)
	if !exists || node.Variadic || len(params) != 1 || g.w.normalize(resultType(node)) != "void" {
		return "", fmt.Errorf("%s must be a void release function with one pointer argument", symbol)
	}
	canonical := g.w.normalize(params[0].Type.Qual)
	switch canonical {
	case "char *":
		return "defer C." + symbol + "(" + pointer + ")\n", nil
	case "void *":
		return "defer C." + symbol + "(unsafe.Pointer(" + pointer + "))\n", nil
	default:
		return "", fmt.Errorf("unsupported release parameter %s", canonical)
	}
}

func (g *generator) function(fn Function) error {
	node := g.w.Functions[fn.Symbol]
	params := parameters(node)
	nativeParameters, publicParameters, rawParameters := []string{}, []string{}, []string{}
	cArguments, goArguments := []string{}, []string{}
	outputs := []output{}
	before, after := strings.Builder{}, strings.Builder{}
	rawCount := 0
	returnType := resultType(node)
	retCanonical := g.w.normalize(returnType)
	callPrefix := ""
	if fn.Return.Kind == "cstring" {
		if retCanonical != "char *" {
			return fmt.Errorf("%s: cstring return requires char pointer", fn.Name)
		}
		callPrefix = "cResult := "
		outputs = append(outputs, g.stringOutput("cResult", rawCount))
		rawCount += 2
	} else if retCanonical != "void" {
		rep, err := g.w.representation(returnType, fn.Return.Type)
		if err != nil {
			return fmt.Errorf("%s return: %w", fn.Name, err)
		}
		g.abi(returnType, rep)
		callPrefix = "cResult := "
		outputs = append(outputs, g.outputValue(rep, "cResult", rawCount))
		rawCount++
	}
	type stringCopy struct {
		expression, release string
		limit               int
	}
	copies := []stringCopy{}
	if fn.Return.Kind == "cstring" {
		copies = append(copies, stringCopy{"cResult", fn.Return.Release, fn.Return.MaxBytes})
	}
	for i, mapping := range fn.Parameters {
		ty := params[i].Type.Qual
		canonical := g.w.normalize(ty)
		name := fmt.Sprintf("a%d", i)
		cName := fmt.Sprintf("cArg%d", i)
		kind := mapping.Kind
		if kind == "" {
			kind = "value"
		}
		switch kind {
		case "value":
			rep, err := g.w.representation(ty, mapping.Type)
			if err != nil {
				return fmt.Errorf("%s parameter %s: %w", fn.Name, mapping.Name, err)
			}
			g.abi(ty, rep)
			nativeParameters = append(nativeParameters, name+" "+rep.goType)
			publicParameters = append(publicParameters, mapping.Name+": "+rep.gom)
			rawParameters = append(rawParameters, name+": "+rawGom(rep))
			if rep.handle != "" {
				cArguments = append(cArguments, name+".raw")
				goArguments = append(goArguments, mapping.Name+".raw")
			} else {
				cArguments = append(cArguments, rep.cgo+"("+name+")")
				goArguments = append(goArguments, mapping.Name)
				if strings.HasPrefix(rep.canonical, "enum ") {
					fmt.Fprintf(&before, "if int64(%s(%s)) != %s { ERROR(%q) }\n", rep.cgo, name, name, "C enum argument is out of range")
				}
			}
		case "cstring":
			if canonical != "char *" {
				return fmt.Errorf("%s: cstring requires char pointer", fn.Name)
			}
			nativeParameters = append(nativeParameters, name+" string")
			publicParameters = append(publicParameters, mapping.Name+": c::CString")
			rawParameters = append(rawParameters, name+": ffi::String")
			goArguments = append(goArguments, mapping.Name+".as_raw()")
			fmt.Fprintf(&before, "if strings.IndexByte(%s, 0) >= 0 || len(%s) > 67108864 { ERROR(%q) }\n", name, name, "C string contains NUL or exceeds 64 MiB")
			fmt.Fprintf(&before, "%s := C.CString(%s)\ndefer C.free(unsafe.Pointer(%s))\n", cName, name, cName)
			cArguments = append(cArguments, cName)
		case "bytes", "inout_bytes":
			if canonical != "void *" && canonical != "char *" && canonical != "unsigned char *" && canonical != "signed char *" {
				return fmt.Errorf("%s: bytes requires void or byte pointer", fn.Name)
			}
			cty, err := cgoType(ty)
			if err != nil {
				return err
			}
			nativeParameters = append(nativeParameters, name+" []byte")
			publicParameters = append(publicParameters, mapping.Name+": Bytes")
			rawParameters = append(rawParameters, name+": ffi::RawSlice[u8]")
			goArguments = append(goArguments, "ffi::slice_from_vec_copy("+mapping.Name+".to_vec())")
			fmt.Fprintf(&before, "if len(%s) > 67108864 { ERROR(%q) }\nvar %s unsafe.Pointer\nif len(%s) > 0 { %s = C.CBytes(%s); defer C.free(%s) }\n", name, "C buffer exceeds 64 MiB", cName, name, cName, name, cName)
			cArguments = append(cArguments, "("+cty+")("+cName+")")
			if kind == "inout_bytes" {
				outputs = append(outputs, output{goTypes: []string{"[]byte"}, rawTypes: []string{"ffi::RawSlice[u8]"}, gomType: "Bytes", goValues: []string{"C.GoBytes(" + cName + ", C.int(len(" + name + ")))"}, gomValue: fmt.Sprintf("Bytes::from_vec(ffi::slice_to_vec_copy(r%d))", rawCount)})
				rawCount++
			}
		case "length":
			rep, err := g.w.representation(ty, "")
			if err != nil || rep.handle != "" || !strings.Contains("iu", rep.gom[:1]) {
				return fmt.Errorf("%s: length requires an integer parameter", fn.Name)
			}
			g.abi(ty, rep)
			buffer := -1
			for j, param := range fn.Parameters {
				if param.Name == mapping.Of {
					buffer = j
				}
			}
			expression := fmt.Sprintf("len(a%d)", buffer)
			fmt.Fprintf(&before, "if uint64(%s(%s)) != uint64(%s) { ERROR(%q) }\n", rep.cgo, expression, expression, "buffer length does not fit C parameter")
			cArguments = append(cArguments, rep.cgo+"("+expression+")")
		case "out":
			if !strings.HasSuffix(canonical, " *") {
				return fmt.Errorf("%s: out requires a pointer", fn.Name)
			}
			pointee := strings.TrimSpace(strings.TrimSuffix(canonical, "*"))
			rep, err := g.w.representation(pointee, mapping.Type)
			if err != nil {
				return fmt.Errorf("%s output %s: %w", fn.Name, mapping.Name, err)
			}
			cty := rep.cgo
			if rep.handle != "" {
				for _, opaque := range g.w.Project.Config.Types {
					if opaque.Name == rep.handle {
						cty, err = cgoType(opaque.CType)
					}
				}
				if err != nil {
					return err
				}
			}
			g.abi(pointee, rep)
			fmt.Fprintf(&before, "var %s %s\n", cName, cty)
			cArguments = append(cArguments, "&"+cName)
			outputs = append(outputs, g.outputValue(rep, cName, rawCount))
			rawCount++
		case "out_string":
			if canonical != "char * *" {
				return fmt.Errorf("%s: out_string requires char **", fn.Name)
			}
			fmt.Fprintf(&before, "var %s *C.char\n", cName)
			cArguments = append(cArguments, "&"+cName)
			outputs = append(outputs, g.stringOutput(cName, rawCount))
			rawCount += 2
			copies = append(copies, stringCopy{cName, mapping.Release, mapping.MaxBytes})
		}
	}
	goTypes, rawTypes, gomTypes, goValues, gomValues := []string{}, []string{}, []string{}, []string{}, []string{}
	for _, output := range outputs {
		goTypes = append(goTypes, output.goTypes...)
		rawTypes = append(rawTypes, output.rawTypes...)
		gomTypes = append(gomTypes, output.gomType)
		goValues = append(goValues, output.goValues...)
		gomValues = append(gomValues, output.gomValue)
	}
	failure := []string{}
	for _, ty := range goTypes {
		failure = append(failure, goZero(ty))
	}
	fail := func(message string) string {
		return "return " + strings.Join(append(append([]string{}, failure...), "fmt.Errorf("+strconv.Quote(message)+")"), ", ")
	}
	expandErrors := func(source string) string {
		for {
			start := strings.Index(source, "ERROR(")
			if start < 0 {
				return source
			}
			end := strings.Index(source[start:], ")") + start
			message, _ := strconv.Unquote(source[start+6 : end])
			source = source[:start] + fail(message) + source[end+1:]
		}
	}
	for _, copy := range copies {
		release, err := g.release(copy.release, copy.expression)
		if err != nil {
			return err
		}
		after.WriteString(release)
	}
	for _, copy := range copies {
		limit := copy.limit
		if limit == 0 {
			limit = 1 << 20
		}
		fmt.Fprintf(&after, "var %sText string\nif %s != nil { n := C.goml_c_bounded_length(%s, %d); if n > %d { %s }; %sText = C.GoStringN(%s, C.int(n)) }\n",
			copy.expression, copy.expression, copy.expression, limit+1, limit, fail("C string exceeds configured copy limit"), copy.expression, copy.expression)
	}
	fmt.Fprintf(&g.native, "func GomlC_%s(%s) (%s) {\n%s%sC.%s(%s)\n%sreturn %s\n}\n",
		fn.Name, strings.Join(nativeParameters, ", "), strings.Join(append(goTypes, "error"), ", "), expandErrors(before.String()), callPrefix, fn.Symbol, strings.Join(cArguments, ", "), after.String(), strings.Join(append(goValues, "nil"), ", "))
	retSignature := strings.Join(append(rawTypes, "ffi::Error"), ", ")
	if len(rawTypes) > 0 {
		retSignature = "(" + retSignature + ")"
	}
	fmt.Fprintf(&g.gom, "#[go_ffi(%q, %q)]\nextern fn goml_c_%s(%s) -> %s;\n", g.w.Project.GoImport, "GomlC_"+fn.Name, fn.Name, strings.Join(rawParameters, ", "), retSignature)
	bindings := []string{}
	for i := 0; i < rawCount; i++ {
		bindings = append(bindings, fmt.Sprintf("r%d", i))
	}
	bindings = append(bindings, "error")
	binding := strings.Join(bindings, ", ")
	if rawCount > 0 {
		binding = "(" + binding + ")"
	}
	fmt.Fprintf(&g.gom, "pub fn %s(%s) -> Result[%s, c::Error] {\nlet %s = goml_c_%s(%s);\nc::check_error(error)?;\nResult::Ok(%s)\n}\n", fn.Name, strings.Join(publicParameters, ", "), tuple(gomTypes), binding, fn.Name, strings.Join(goArguments, ", "), tuple(gomValues))
	signatureParams := []string{}
	for _, param := range params {
		signatureParams = append(signatureParams, param.Type.Qual)
	}
	if len(signatureParams) == 0 {
		signatureParams = append(signatureParams, "void")
	}
	fmt.Fprintf(&g.assertions, "typedef %s (*goml_c_signature_%s)(%s);\n_Static_assert(__builtin_types_compatible_p(__typeof__(&%s), goml_c_signature_%s), \"C signature changed; regenerate GoML bindings\");\n", returnType, fn.Name, strings.Join(signatureParams, ", "), fn.Symbol, fn.Name)
	return nil
}

func tuple(values []string) string {
	if len(values) == 1 {
		return values[0]
	}
	return "(" + strings.Join(values, ", ") + ")"
}

func cgoFlag(p Project, flag string) string {
	if strings.HasPrefix(flag, "-I") || strings.HasPrefix(flag, "-L") {
		directory := includeDirectory(p, flag[2:])
		if relative, err := filepath.Rel(filepath.Dir(p.GoFile), directory); err == nil && (directory == p.Root || strings.HasPrefix(directory, p.Root+string(filepath.Separator))) {
			flag = flag[:2] + "${SRCDIR}/" + filepath.ToSlash(relative)
		} else {
			flag = flag[:2] + directory
		}
	}
	if strings.ContainsAny(flag, " \t") {
		return strconv.Quote(flag)
	}
	return flag
}

func Generate(w *World) (generated, error) {
	g := &generator{w: w}
	fmt.Fprintf(&g.gom, "package %s;\nuse std::ffi;\nuse std::c;\nuse std::bytes::{Bytes};\n", w.Project.Config.Package)
	for _, ty := range w.Project.Config.Types {
		cty, err := cgoType(ty.CType)
		if err != nil {
			return generated{}, err
		}
		fmt.Fprintf(&g.native, "type %s struct { raw %s }\nfunc GomlC_Null%s() %s { return %s{} }\nfunc GomlC_IsNull%s(value %s) bool { return value.raw == nil }\nfunc GomlC_Equal%s(left, right %s) bool { return left.raw == right.raw }\n", ty.Name, cty, ty.Name, ty.Name, ty.Name, ty.Name, ty.Name, ty.Name, ty.Name)
		fmt.Fprintf(&g.gom, "#[go_type(%q, %q)]\nextern type GomlC%s;\npub struct %s { raw: GomlC%s }\n", w.Project.GoImport, ty.Name, ty.Name, ty.Name, ty.Name)
		fmt.Fprintf(&g.gom, "#[go_ffi(%q, %q)]\nextern fn goml_c_null_%s() -> GomlC%s;\n", w.Project.GoImport, "GomlC_Null"+ty.Name, ty.Name, ty.Name)
		fmt.Fprintf(&g.gom, "#[go_ffi(%q, %q)]\nextern fn goml_c_is_null_%s(value: GomlC%s) -> bool;\n", w.Project.GoImport, "GomlC_IsNull"+ty.Name, ty.Name, ty.Name)
		fmt.Fprintf(&g.gom, "#[go_ffi(%q, %q)]\nextern fn goml_c_equal_%s(left: GomlC%s, right: GomlC%s) -> bool;\n", w.Project.GoImport, "GomlC_Equal"+ty.Name, ty.Name, ty.Name, ty.Name)
		fmt.Fprintf(&g.gom, "impl %s { pub fn null() -> %s { %s { raw: goml_c_null_%s() } } pub fn is_null(self: %s) -> bool { goml_c_is_null_%s(self.raw) } pub fn same_as(self: %s, other: %s) -> bool { goml_c_equal_%s(self.raw, other.raw) } }\n", ty.Name, ty.Name, ty.Name, ty.Name, ty.Name, ty.Name, ty.Name, ty.Name, ty.Name)
	}
	for _, fn := range w.Project.Config.Functions {
		if err := g.function(fn); err != nil {
			return generated{}, err
		}
	}
	for index, constant := range w.Project.Config.Constants {
		ty := w.Constants[fmt.Sprintf("goml_c_probe_constant_%d", index)]
		rep, err := w.representation(ty, "")
		if err != nil || rep.handle != "" || rep.gom == "()" {
			return generated{}, fmt.Errorf("constant %s requires an integer C type", constant.Symbol)
		}
		value := w.Values[fmt.Sprintf("goml_c_probe_value_%d", index)]
		if value == "" || !strings.Contains("iu", rep.gom[:1]) {
			return generated{}, fmt.Errorf("constant %s requires an integer constant expression", constant.Symbol)
		}
		g.abi(ty, rep)
		literal := value + "ULL"
		if strings.HasPrefix(value, "-") {
			literal = "(-" + strings.TrimPrefix(value, "-") + "ULL)"
		}
		fmt.Fprintf(&g.assertions, "_Static_assert((%s) == (%s)%s, \"C constant changed; regenerate GoML bindings\");\n", constant.Symbol, ty, literal)
		fmt.Fprintf(&g.native, "func GomlC_%s() %s { return %s(C.%s) }\n", constant.Name, rep.goType, rep.goType, constant.Symbol)
		fmt.Fprintf(&g.gom, "pub const %s: %s = %s;\n", constant.Name, rep.gom, value)
	}
	var source strings.Builder
	fmt.Fprintf(&source, "package %s\n\n/*\n", w.Project.Config.GoPackage)
	flags := []string{"-std=c11", "-I" + w.Project.Directory}
	for _, dir := range w.Project.Config.IncludeDirs {
		flags = append(flags, "-I"+includeDirectory(w.Project, dir))
	}
	flags = append(flags, w.Project.Config.CFlags...)
	for i, flag := range flags {
		flags[i] = cgoFlag(w.Project, flag)
	}
	fmt.Fprintf(&source, "#cgo CFLAGS: %s\n", strings.Join(flags, " "))
	flags = []string{}
	for _, flag := range w.Project.Config.LDFlags {
		flags = append(flags, cgoFlag(w.Project, flag))
	}
	if len(flags) > 0 {
		fmt.Fprintf(&source, "#cgo LDFLAGS: %s\n", strings.Join(flags, " "))
	}
	if len(w.Project.Config.PkgConfig) > 0 {
		fmt.Fprintf(&source, "#cgo pkg-config: %s\n", strings.Join(w.Project.Config.PkgConfig, " "))
	}
	source.WriteString("#include <stdlib.h>\n#include <stdint.h>\n#include <stdbool.h>\n")
	for _, header := range w.Project.Config.Headers {
		fmt.Fprintf(&source, "#include <%s>\n", header)
	}
	source.WriteString(g.assertions.String())
	source.WriteString("static size_t goml_c_bounded_length(const char *p, size_t limit) { size_t n = 0; while (n < limit && p[n] != 0) { n++; } return n; }\n*/\nimport \"C\"\n")
	body := g.native.String()
	imports := []string{}
	for _, item := range []struct{ prefix, name string }{{"fmt.", "fmt"}, {"strings.", "strings"}, {"unsafe.", "unsafe"}} {
		if strings.Contains(body, item.prefix) {
			imports = append(imports, strconv.Quote(item.name))
		}
	}
	if len(imports) > 0 {
		source.WriteString("import (" + strings.Join(imports, "\n") + ")\n")
	}
	source.WriteString(body)
	native, err := format.Source([]byte(source.String()))
	if err != nil {
		return generated{}, fmt.Errorf("invalid generated Go source: %w", err)
	}
	return generated{Goml: []byte(g.gom.String()), Go: native}, nil
}
