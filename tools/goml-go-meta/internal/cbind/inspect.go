package cbind

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"sort"
	"strconv"
	"strings"
)

type astType struct {
	Qual      string `json:"qualType"`
	Desugared string `json:"desugaredQualType"`
}

type astNode struct {
	StorageClass string    `json:"storageClass"`
	Kind         string    `json:"kind"`
	Name         string    `json:"name"`
	Type         astType   `json:"type"`
	Variadic     bool      `json:"variadic"`
	Value        string    `json:"value"`
	Inner        []astNode `json:"inner"`
}

type World struct {
	EnumTypes   map[string]string
	Project     Project
	Clang       string
	Aliases     map[string]string
	Functions   map[string]astNode
	Constants   map[string]string
	Values      map[string]string
	Handles     map[string]string
	Sizes       map[string]int
	Fingerprint string
}

type limitedBuffer struct {
	bytes.Buffer
	limit int
}

func (b *limitedBuffer) Write(data []byte) (int, error) {
	if len(data) > b.limit-b.Len() {
		return 0, fmt.Errorf("native tool output exceeds %d bytes", b.limit)
	}
	return b.Buffer.Write(data)
}

func run(ctx context.Context, directory, input, program string, arguments ...string) ([]byte, error) {
	command := exec.CommandContext(ctx, program, arguments...)
	command.Dir = directory
	command.Stdin = strings.NewReader(input)
	command.Env = append(os.Environ(), "LC_ALL=C", "GOWORK=off", "GOTOOLCHAIN=local", "GOPROXY=off", "GOSUMDB=off")
	var stdout, stderr limitedBuffer
	stdout.limit, stderr.limit = 64<<20, 1<<20
	command.Stdout, command.Stderr = &stdout, &stderr
	if err := command.Run(); err != nil {
		return nil, fmt.Errorf("%s: %w\n%s", program, err, stderr.String())
	}
	return stdout.Bytes(), nil
}

func clangProgram() (string, error) {
	if configured := os.Getenv("GOML_CLANG"); configured != "" {
		return exec.LookPath(configured)
	}
	for _, name := range []string{"clang", "clang-20", "clang-19", "clang-18", "clang-17", "clang-16", "clang-15"} {
		if found, err := exec.LookPath(name); err == nil {
			return found, nil
		}
	}
	return "", fmt.Errorf("Clang is required; set GOML_CLANG to its executable")
}

func includeDirectory(p Project, value string) string {
	if !filepath.IsAbs(value) {
		return filepath.Join(p.Directory, value)
	}
	return value
}

func nativeFlags(p Project, flags []string) []string {
	result := []string{}
	for _, flag := range flags {
		if strings.HasPrefix(flag, "-I") || strings.HasPrefix(flag, "-L") {
			flag = flag[:2] + includeDirectory(p, flag[2:])
		}
		result = append(result, flag)
	}
	return result
}

func Inspect(ctx context.Context, p Project) (*World, error) {
	goEnvironment, err := run(ctx, p.Root, "", "go", "env", "GOOS", "GOARCH", "CGO_ENABLED")
	if err != nil {
		return nil, err
	}
	expected := runtime.GOOS + "\n" + runtime.GOARCH + "\n1\n"
	if p.Config.Backend == "dynamic" {
		expected = "linux\namd64\n0\n"
		version, versionError := run(ctx, p.Root, "", "go", "env", "GOVERSION")
		if versionError != nil {
			return nil, versionError
		}
		if runtime.GOOS != "linux" || runtime.GOARCH != "amd64" || !strings.HasPrefix(string(version), "go1.26.") {
			return nil, fmt.Errorf("dynamic C bindings require a Linux amd64 host and Go 1.26.x")
		}
	}
	if string(goEnvironment) != expected {
		return nil, fmt.Errorf("C binding generation requires the host Go target and CGO_ENABLED=%s for this backend", strings.TrimSpace(strings.Split(expected, "\n")[2]))
	}
	clang, err := clangProgram()
	if err != nil {
		return nil, err
	}
	w := &World{Project: p, Clang: clang, Aliases: map[string]string{}, Functions: map[string]astNode{}, Constants: map[string]string{}, Values: map[string]string{}, Handles: map[string]string{}, Sizes: map[string]int{}}
	var probe strings.Builder
	probe.WriteString("#include <stddef.h>\n#include <stdint.h>\n#include <stdbool.h>\n")
	for _, header := range p.Config.Headers {
		fmt.Fprintf(&probe, "#include <%s>\n", header)
	}
	for index, ty := range p.Config.Types {
		fmt.Fprintf(&probe, "typedef %s goml_c_probe_type_%d;\n", ty.CType, index)
	}
	for index, constant := range p.Config.Constants {
		fmt.Fprintf(&probe, "static const __typeof__(%s) goml_c_probe_constant_%d = (%s);\n", constant.Symbol, index, constant.Symbol)
		fmt.Fprintf(&probe, "enum { goml_c_probe_value_%d = (%s) };\n", index, constant.Symbol)
	}
	for _, ty := range []string{"char", "short", "int", "long", "long long", "float", "double", "size_t"} {
		fmt.Fprintf(&probe, "typedef char goml_c_probe_size_%s[sizeof(%s)];\n", strings.ReplaceAll(ty, " ", "_"), ty)
	}
	probe.WriteString("typedef char goml_c_probe_size_char_signed[((char)-1 < 0) ? 1 : 2];\n")
	flags := []string{"-x", "c"}
	nativeEnvironment, err := run(ctx, p.Root, "", "go", "env", "-json", "CGO_CPPFLAGS", "CGO_CFLAGS", "PKG_CONFIG")
	if err != nil {
		return nil, err
	}
	var cgoFlags map[string]string
	if err := json.Unmarshal(nativeEnvironment, &cgoFlags); err != nil {
		return nil, fmt.Errorf("invalid cgo environment: %w", err)
	}
	for _, variable := range []string{"CGO_CPPFLAGS", "CGO_CFLAGS"} {
		value := cgoFlags[variable]
		if strings.ContainsAny(value, "\"'\\") {
			return nil, fmt.Errorf("quoted %s is unsupported during C inspection; use include_dirs for paths containing spaces", variable)
		}
	}
	flags = append(flags, strings.Fields(cgoFlags["CGO_CPPFLAGS"])...)
	if len(p.Config.PkgConfig) > 0 {
		args := append([]string{"--cflags"}, p.Config.PkgConfig...)
		output, err := run(ctx, p.Directory, "", cgoFlags["PKG_CONFIG"], args...)
		if err != nil {
			return nil, err
		}
		if strings.ContainsAny(string(output), "\"'\\") {
			return nil, fmt.Errorf("quoted pkg-config flags are unsupported")
		}
		flags = append(flags, strings.Fields(string(output))...)
	}
	flags = append(flags, strings.Fields(cgoFlags["CGO_CFLAGS"])...)
	flags = append(flags, "-std=c11", "-I"+p.Directory)
	for _, directory := range p.Config.IncludeDirs {
		flags = append(flags, "-I"+includeDirectory(p, directory))
	}
	flags = append(flags, nativeFlags(p, p.Config.CFlags)...)
	args := append(append([]string{}, flags...), "-Xclang", "-ast-dump=json", "-fsyntax-only", "-")
	output, err := run(ctx, p.Directory, probe.String(), clang, args...)
	if err != nil {
		return nil, err
	}
	var tree astNode
	if err := json.Unmarshal(output, &tree); err != nil {
		return nil, fmt.Errorf("invalid Clang AST: %w", err)
	}
	for _, node := range tree.Inner {
		switch node.Kind {
		case "EnumDecl":
			for _, constant := range node.Inner {
				if strings.HasPrefix(constant.Name, "goml_c_probe_value_") {
					w.Values[constant.Name] = constantValue(constant)
				}
			}
		case "TypedefDecl":
			ty := node.Type.Desugared
			if ty == "" || ty == node.Name {
				ty = node.Type.Qual
			}
			w.Aliases[node.Name] = ty
			if strings.HasPrefix(node.Name, "goml_c_probe_size_") {
				start := strings.LastIndex(ty, "[")
				end := strings.LastIndex(ty, "]")
				if start < 0 || end <= start {
					return nil, fmt.Errorf("cannot determine C ABI size for %s", node.Name)
				}
				size, err := strconv.Atoi(ty[start+1 : end])
				if err != nil {
					return nil, err
				}
				w.Sizes[strings.TrimPrefix(node.Name, "goml_c_probe_size_")] = size
			}
		case "FunctionDecl":
			w.Functions[node.Name] = node
		case "VarDecl":
			if strings.HasPrefix(node.Name, "goml_c_probe_constant_") {
				ty := node.Type.Desugared
				if ty == "" {
					ty = node.Type.Qual
				}
				w.Constants[node.Name] = ty
			}
		}
	}
	if w.Sizes["char"] != 1 || w.Sizes["float"] != 4 || w.Sizes["double"] != 8 {
		return nil, fmt.Errorf("unsupported C scalar ABI")
	}
	if p.Config.Backend == "dynamic" {
		if w.Sizes["short"] != 2 || w.Sizes["int"] != 4 || w.Sizes["long"] != 8 || w.Sizes["long_long"] != 8 || w.Sizes["size_t"] != 8 {
			return nil, fmt.Errorf("dynamic backend requires the Linux amd64 LP64 C ABI")
		}
		w.EnumTypes = map[string]string{}
		usedEnums := map[string]bool{}
		for _, fn := range p.Config.Functions {
			node := w.Functions[fn.Symbol]
			usedEnums[w.normalize(resultType(node))] = true
			for index, parameter := range parameters(node) {
				canonical := w.normalize(parameter.Type.Qual)
				if index < len(fn.Parameters) && fn.Parameters[index].Kind == "out" {
					canonical = strings.TrimSpace(strings.TrimSuffix(canonical, "*"))
				}
				usedEnums[canonical] = true
			}
		}
		spellings := map[string]string{}
		for _, node := range tree.Inner {
			if node.Kind == "EnumDecl" && node.Name != "" && usedEnums["enum "+node.Name] {
				spellings["enum "+node.Name] = "enum " + node.Name
			}
			if node.Kind == "TypedefDecl" {
				canonical := w.normalize(node.Name)
				if usedEnums[canonical] && strings.HasPrefix(canonical, "enum ") && !strings.ContainsAny(canonical, "*()[]") {
					spellings[canonical] = node.Name
				}
			}
		}
		enumNames := []string{}
		for name := range spellings {
			enumNames = append(enumNames, name)
		}
		sort.Strings(enumNames)
		var enumProbe strings.Builder
		enumProbe.WriteString(probe.String())
		for index, name := range enumNames {
			fmt.Fprintf(&enumProbe, "enum { goml_c_enum_size_%d = sizeof(%s), goml_c_enum_sign_%d = ((%s)-1 < 0) };\n", index, spellings[name], index, spellings[name])
		}
		output, err := run(ctx, p.Directory, enumProbe.String(), clang, args...)
		if err != nil {
			return nil, err
		}
		var enums astNode
		if err := json.Unmarshal(output, &enums); err != nil {
			return nil, err
		}
		values := map[string]string{}
		for _, node := range enums.Inner {
			if node.Kind == "EnumDecl" {
				for _, constant := range node.Inner {
					values[constant.Name] = constantValue(constant)
				}
			}
		}
		for index, name := range enumNames {
			size, _ := strconv.Atoi(values[fmt.Sprintf("goml_c_enum_size_%d", index)])
			if size != 1 && size != 2 && size != 4 {
				continue
			}
			prefix := "uint"
			if values[fmt.Sprintf("goml_c_enum_sign_%d", index)] == "1" {
				prefix = "int"
			}
			w.EnumTypes[name] = fmt.Sprintf("%s%d", prefix, size*8)
		}
	}
	for index, ty := range p.Config.Types {
		canonical := w.normalize(w.Aliases[fmt.Sprintf("goml_c_probe_type_%d", index)])
		if !strings.HasSuffix(canonical, "*") || strings.Contains(canonical, "(") {
			return nil, fmt.Errorf("%s must name an opaque C pointer type", ty.Name)
		}
		w.Handles[ty.Name] = canonical
	}
	for index, fn := range p.Config.Functions {
		node, found := w.Functions[fn.Symbol]
		if !found || node.Variadic || strings.HasSuffix(node.Type.Qual, "()") {
			return nil, fmt.Errorf("%s must be a declared, non-variadic C function", fn.Symbol)
		}
		if p.Config.Backend == "dynamic" {
			if err := dynamicSignature(node); err != nil {
				return nil, err
			}
		}
		params := parameters(node)
		if fn.Parameters == nil {
			for n := range params {
				fn.Parameters = append(fn.Parameters, Parameter{Name: fmt.Sprintf("arg%d", n)})
			}
			w.Project.Config.Functions[index] = fn
		}
		if len(fn.Parameters) != len(params) {
			return nil, fmt.Errorf("%s: expected %d parameter mappings, found %d", fn.Symbol, len(params), len(fn.Parameters))
		}
	}
	preprocessed, err := run(ctx, p.Directory, probe.String(), clang, append(append([]string{}, flags...), "-E", "-P", "-dD", "-")...)
	if err != nil {
		return nil, err
	}
	version, err := run(ctx, p.Directory, "", clang, "--version")
	if err != nil {
		return nil, err
	}
	configuration, _ := json.Marshal(p.Config)
	hash := sha256.New()
	for _, part := range [][]byte{[]byte("goml-c-bind-v1\x00"), configuration, preprocessed, version, goEnvironment} {
		hash.Write(part)
		hash.Write([]byte{0})
	}
	w.Fingerprint = hex.EncodeToString(hash.Sum(nil))
	return w, nil
}

func constantValue(node astNode) string {
	if node.Kind == "ConstantExpr" {
		return node.Value
	}
	for _, child := range node.Inner {
		if value := constantValue(child); value != "" {
			return value
		}
	}
	return ""
}

func parameters(node astNode) []astNode {
	result := []astNode{}
	for _, child := range node.Inner {
		if child.Kind == "ParmVarDecl" {
			result = append(result, child)
		}
	}
	return result
}

func resultType(node astNode) string {
	index := strings.Index(node.Type.Qual, "(")
	if index < 0 {
		return ""
	}
	return strings.TrimSpace(node.Type.Qual[:index])
}

func (w *World) normalize(value string) string {
	for iteration := 0; iteration < 64; iteration++ {
		value = strings.ReplaceAll(value, "*", " * ")
		words := strings.Fields(value)
		result := []string{}
		for _, word := range words {
			if word != "const" && word != "volatile" && word != "restrict" && word != "__restrict" {
				result = append(result, word)
			}
		}
		next := strings.Join(result, " ")
		if len(result) > 0 {
			if alias, exists := w.Aliases[result[0]]; exists && alias != result[0] {
				next = alias + " " + strings.Join(result[1:], " ")
			}
		}
		next = strings.TrimSpace(next)
		if next == value || next == strings.Join(words, " ") {
			return next
		}
		value = next
	}
	return value
}
