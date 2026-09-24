package cbind

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

func RuntimeDirectory() (string, error) {
	executable, err := os.Executable()
	if err != nil {
		return "", err
	}
	executable, err = filepath.EvalSymlinks(executable)
	if err != nil {
		return "", err
	}
	directory := filepath.Join(filepath.Dir(executable), "..", "lib", "cabi")
	if _, err := os.Stat(filepath.Join(directory, "go.mod")); err != nil {
		return "", fmt.Errorf("GoML C ABI runtime is missing from this toolchain: %w", err)
	}
	return filepath.Clean(directory), nil
}

func dynamicSignature(node astNode) error {
	if node.StorageClass == "static" {
		return fmt.Errorf("%s: dynamic backend requires an exported C function, not static or static inline", node.Name)
	}
	if node.Variadic || strings.HasSuffix(node.Type.Qual, "()") || strings.Contains(node.Type.Qual, "__attribute__") || len(parameters(node)) > 128 {
		return fmt.Errorf("%s: dynamic backend requires a fixed System V signature with at most 128 scalar parameters", node.Name)
	}
	for _, child := range node.Inner {
		if strings.HasSuffix(child.Kind, "Attr") && (strings.Contains(child.Kind, "ABI") || strings.Contains(child.Kind, "Call") || child.Kind == "RegparmAttr" || child.Kind == "PCSAttr") {
			return fmt.Errorf("%s: unsupported calling convention %s", node.Name, child.Kind)
		}
	}
	return nil
}

func (g *generator) dynamic() bool { return g.w.Project.Config.Backend == "dynamic" }

func (g *generator) ctype(value string) (string, error) {
	if !g.dynamic() {
		return cgoType(value)
	}
	canonical := g.w.normalize(value)
	if strings.HasSuffix(canonical, " *") && !strings.ContainsAny(canonical, "()[]") {
		return "unsafe.Pointer", nil
	}
	if strings.HasPrefix(canonical, "enum ") {
		if ty := g.w.EnumTypes[canonical]; ty != "" {
			return ty, nil
		}
		return "", fmt.Errorf("unsupported C enum ABI: %s", canonical)
	}
	rep, err := g.w.representation(value, "")
	if err != nil {
		return "", err
	}
	if rep.handle != "" {
		return "unsafe.Pointer", nil
	}
	return rep.goType, nil
}

func (g *generator) representation(value, requested string) (representation, error) {
	rep, err := g.w.representation(value, requested)
	if err == nil && g.dynamic() {
		rep.cgo, err = g.ctype(value)
	}
	return rep, err
}

func (g *generator) call(symbol string) string {
	if g.dynamic() {
		return "gomlDynamic_" + symbol
	}
	return "C." + symbol
}

func (g *generator) dynamicSupport() (string, error) {
	symbols := map[string]bool{}
	for _, fn := range g.w.Project.Config.Functions {
		symbols[fn.Symbol] = true
		if fn.Return.Release != "" {
			symbols[fn.Return.Release] = true
		}
		for _, param := range fn.Parameters {
			if param.Release != "" {
				symbols[param.Release] = true
			}
		}
	}
	ordered := []string{}
	for symbol := range symbols {
		ordered = append(ordered, symbol)
	}
	sort.Strings(ordered)
	var source strings.Builder
	libraries := []string{}
	for _, library := range g.w.Project.Config.Libraries {
		libraries = append(libraries, strconv.Quote(library))
	}
	fmt.Fprintf(&source, "var gomlCLibrary = cabi.Library{Names: []string{%s}}\nvar gomlCOnce sync.Once\nvar gomlCError error\nvar gomlCAddresses [%d]uint64\nfunc gomlCLoad() error { gomlCOnce.Do(func() {\n", strings.Join(libraries, ","), len(ordered))
	for index, symbol := range ordered {
		fmt.Fprintf(&source, "gomlCAddresses[%d], gomlCError = gomlCLibrary.Symbol(%q); if gomlCError != nil { return }\n", index, symbol)
	}
	source.WriteString("}); return gomlCError }\n")
	for index, symbol := range ordered {
		node, exists := g.w.Functions[symbol]
		if !exists {
			return "", fmt.Errorf("C function %s is unavailable", symbol)
		}
		if err := dynamicSignature(node); err != nil {
			return "", err
		}
		result, err := g.ctype(resultType(node))
		if err != nil {
			return "", err
		}
		names, assignments := []string{}, []string{}
		integerCount, floatCount, stackCount := 0, 0, 0
		for i, param := range parameters(node) {
			ty, err := g.ctype(param.Type.Qual)
			if err != nil {
				return "", err
			}
			name := fmt.Sprintf("p%d", i)
			names = append(names, name+" "+ty)
			value, destination := "uint64("+name+")", ""
			switch ty {
			case "unsafe.Pointer":
				value = "uint64(uintptr(" + name + "))"
			case "float32":
				value = "uint64(math.Float32bits(" + name + "))"
			case "float64":
				value = "math.Float64bits(" + name + ")"
			case "bool":
				value = name + "Bits"
				assignments = append(assignments, fmt.Sprintf("var %s uint64; if %s { %s = 1 }", value, name, value))
			}
			if strings.HasPrefix(ty, "float") {
				if floatCount < 8 {
					destination = fmt.Sprintf("f.Float[%d]", floatCount)
					floatCount++
				}
			} else if integerCount < 6 {
				destination = fmt.Sprintf("f.Integer[%d]", integerCount)
				integerCount++
			}
			if destination == "" {
				destination = fmt.Sprintf("f.Stack[%d]", stackCount)
				stackCount++
			}
			assignments = append(assignments, destination+" = "+value)
		}
		fmt.Fprintf(&source, "func gomlDynamic_%s(%s) %s {\nf := cabi.Frame{Function: gomlCAddresses[%d], StackCount: %d}\n%s\nf.Call()\n", symbol, strings.Join(names, ","), result, index, stackCount, strings.Join(assignments, "\n"))
		switch result {
		case "":
		case "unsafe.Pointer":
			source.WriteString("return cabi.Pointer(f.Result)\n")
		case "float32":
			source.WriteString("return math.Float32frombits(uint32(f.FloatResult))\n")
		case "float64":
			source.WriteString("return math.Float64frombits(f.FloatResult)\n")
		case "bool":
			source.WriteString("return uint8(f.Result) != 0\n")
		default:
			fmt.Fprintf(&source, "return %s(f.Result)\n", result)
		}
		source.WriteString("}\n")
	}
	return source.String(), nil
}
