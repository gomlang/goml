package cbind

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/token"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type Type struct {
	Name  string `json:"name"`
	CType string `json:"c_type"`
}

type Parameter struct {
	Name     string `json:"name"`
	Kind     string `json:"kind,omitempty"`
	Of       string `json:"of,omitempty"`
	Type     string `json:"type,omitempty"`
	Release  string `json:"release,omitempty"`
	MaxBytes int    `json:"max_bytes,omitempty"`
}

type Return struct {
	Kind     string `json:"kind,omitempty"`
	Type     string `json:"type,omitempty"`
	Release  string `json:"release,omitempty"`
	MaxBytes int    `json:"max_bytes,omitempty"`
}

type Function struct {
	Name       string      `json:"name"`
	Symbol     string      `json:"symbol"`
	Parameters []Parameter `json:"parameters"`
	Return     Return      `json:"return,omitempty"`
}

type Constant struct {
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

type Config struct {
	Version     int        `json:"version"`
	Package     string     `json:"package"`
	Output      string     `json:"output"`
	GoPackage   string     `json:"go_package"`
	GoOutput    string     `json:"go_output"`
	Headers     []string   `json:"headers"`
	IncludeDirs []string   `json:"include_dirs,omitempty"`
	CFlags      []string   `json:"cflags,omitempty"`
	LDFlags     []string   `json:"ldflags,omitempty"`
	PkgConfig   []string   `json:"pkg_config,omitempty"`
	Types       []Type     `json:"types,omitempty"`
	Functions   []Function `json:"functions,omitempty"`
	Constants   []Constant `json:"constants,omitempty"`
}

type Project struct {
	Config    Config
	File      string
	Root      string
	Directory string
	GoImport  string
	GoFile    string
	GomlFile  string
	Manifest  string
}

var identifier = regexp.MustCompile(`^[A-Za-z_][A-Za-z_0-9]*$`)
var cTypeName = regexp.MustCompile(`^(struct )?[A-Za-z_][A-Za-z_0-9]*(\s*\*)*$`)
var headerName = regexp.MustCompile(`^[A-Za-z_0-9./+-]+$`)
var gomlKeywords = strings.Fields("fn let mut if else match for in while loop break continue return defer use pub package module struct enum trait impl extern type const comptime dyn as where true false self Self")

func nameOK(name string, upper bool) bool {
	if !identifier.MatchString(name) || name == "_" || token.Lookup(name).IsKeyword() || strings.HasPrefix(name, "goml_c_") || strings.HasPrefix(name, "GomlC") {
		return false
	}
	for _, word := range gomlKeywords {
		if name == word {
			return false
		}
	}
	if upper {
		return name[0] >= 'A' && name[0] <= 'Z'
	}
	return name[0] >= 'a' && name[0] <= 'z' || name[0] == '_'
}

func decode(data []byte, value any) error {
	if len(data) > 1<<20 {
		return fmt.Errorf("configuration exceeds 1 MiB")
	}
	decoder := json.NewDecoder(bytes.NewReader(data))
	if err := uniqueJSON(decoder); err != nil {
		return err
	}
	if _, err := decoder.Token(); err != io.EOF {
		return fmt.Errorf("trailing configuration data")
	}
	decoder = json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	return decoder.Decode(value)
}

func uniqueJSON(decoder *json.Decoder) error {
	return uniqueJSONDepth(decoder, 0)
}

func uniqueJSONDepth(decoder *json.Decoder, depth int) error {
	if depth > 64 {
		return fmt.Errorf("configuration nesting exceeds 64")
	}
	value, err := decoder.Token()
	if err != nil {
		return err
	}
	if delimiter, ok := value.(json.Delim); ok {
		seen := map[string]bool{}
		for decoder.More() {
			if delimiter == '{' {
				key, err := decoder.Token()
				if err != nil {
					return err
				}
				name, ok := key.(string)
				if !ok || seen[name] {
					return fmt.Errorf("duplicate or invalid field %v", key)
				}
				seen[name] = true
			}
			if err := uniqueJSONDepth(decoder, depth+1); err != nil {
				return err
			}
		}
		_, err = decoder.Token()
	}
	return err
}

func validate(c Config) error {
	if c.Version != 1 || !nameOK(c.Package, false) || !nameOK(c.GoPackage, false) {
		return fmt.Errorf("expected version 1 and valid package names")
	}
	if len(c.Headers) == 0 || len(c.Types)+len(c.Functions)+len(c.Constants) == 0 || len(c.Types)+len(c.Functions)+len(c.Constants) > 4096 {
		return fmt.Errorf("expected headers and between 1 and 4096 bindings")
	}
	for _, header := range c.Headers {
		if !headerName.MatchString(header) || filepath.IsAbs(header) || strings.Contains(header, "..") {
			return fmt.Errorf("invalid header %q", header)
		}
	}
	for _, dir := range c.IncludeDirs {
		if dir == "" || strings.ContainsAny(dir, "\x00\r\n\"") || strings.Contains(dir, "*/") {
			return fmt.Errorf("invalid include directory %q", dir)
		}
	}
	for _, flag := range append(append([]string{}, c.CFlags...), c.LDFlags...) {
		if flag == "" || strings.ContainsAny(flag, "\x00\r\n\t \"'") || strings.Contains(flag, "*/") {
			return fmt.Errorf("invalid native flag %q", flag)
		}
	}
	for _, flag := range c.CFlags {
		if !strings.HasPrefix(flag, "-D") && !strings.HasPrefix(flag, "-U") && !strings.HasPrefix(flag, "-I") && !strings.HasPrefix(flag, "-std=") {
			return fmt.Errorf("unsupported C flag %q", flag)
		}
	}
	for _, flag := range c.LDFlags {
		if !strings.HasPrefix(flag, "-L") && !strings.HasPrefix(flag, "-l") && flag != "-pthread" {
			return fmt.Errorf("unsupported linker flag %q", flag)
		}
	}
	for _, pkg := range c.PkgConfig {
		if !regexp.MustCompile(`^[A-Za-z_0-9.+-]+$`).MatchString(pkg) || strings.HasPrefix(pkg, "-") {
			return fmt.Errorf("invalid pkg-config package %q", pkg)
		}
	}
	names := map[string]bool{"c": true, "ffi": true, "Bytes": true, "C": true, "Result": true, "Option": true}
	for _, ty := range c.Types {
		if !nameOK(ty.Name, true) || !cTypeName.MatchString(ty.CType) || names[ty.Name] {
			return fmt.Errorf("invalid or duplicate C type binding %q", ty.Name)
		}
		names[ty.Name] = true
	}
	for _, fn := range c.Functions {
		if !nameOK(fn.Name, false) || !identifier.MatchString(fn.Symbol) || names[fn.Name] || len(fn.Parameters) > 128 {
			return fmt.Errorf("invalid or duplicate function %q", fn.Name)
		}
		names[fn.Name] = true
		params := map[string]string{}
		for _, param := range fn.Parameters {
			if !nameOK(param.Name, false) || params[param.Name] != "" {
				return fmt.Errorf("invalid or duplicate parameter in %s", fn.Name)
			}
			kind := param.Kind
			if kind == "" {
				kind = "value"
			}
			switch kind {
			case "value", "cstring", "bytes", "inout_bytes", "length", "out", "out_string":
			default:
				return fmt.Errorf("unsupported parameter kind %q", kind)
			}
			params[param.Name] = kind
			if param.Release != "" && (kind != "out_string" || !identifier.MatchString(param.Release)) || param.MaxBytes < 0 || param.MaxBytes > 1<<30 {
				return fmt.Errorf("invalid string release or copy limit")
			}
			if param.Type != "" && kind != "value" && kind != "out" {
				return fmt.Errorf("type override requires value or out")
			}
			if param.MaxBytes != 0 && kind != "out_string" {
				return fmt.Errorf("max_bytes is only valid on copied output strings")
			}
		}
		for _, param := range fn.Parameters {
			if param.Kind == "length" {
				if params[param.Of] != "bytes" && params[param.Of] != "inout_bytes" {
					return fmt.Errorf("length parameter must name a byte buffer")
				}
			} else if param.Of != "" {
				return fmt.Errorf("of is only valid on length parameters")
			}
		}
		if fn.Return.Kind != "" && fn.Return.Kind != "value" && fn.Return.Kind != "cstring" {
			return fmt.Errorf("unsupported return kind %q", fn.Return.Kind)
		}
		if fn.Return.Release != "" && (fn.Return.Kind != "cstring" || !identifier.MatchString(fn.Return.Release)) || fn.Return.MaxBytes < 0 || fn.Return.MaxBytes > 1<<30 {
			return fmt.Errorf("invalid return release or copy limit")
		}
		if fn.Return.MaxBytes != 0 && fn.Return.Kind != "cstring" || fn.Return.Type != "" && fn.Return.Kind == "cstring" {
			return fmt.Errorf("invalid return mapping options")
		}
	}
	for _, constant := range c.Constants {
		if !(nameOK(constant.Name, false) || nameOK(constant.Name, true)) || !identifier.MatchString(constant.Symbol) || names[constant.Name] {
			return fmt.Errorf("invalid or duplicate constant %q", constant.Name)
		}
		names[constant.Name] = true
	}
	return nil
}

func checkedPath(root, base, relative string) (string, error) {
	if relative == "" || filepath.IsAbs(relative) {
		return "", fmt.Errorf("output must be module-relative")
	}
	for _, part := range strings.Split(filepath.ToSlash(relative), "/") {
		if part == ".." || strings.HasPrefix(part, ".goml-bind-") {
			return "", fmt.Errorf("parent traversal and generator recovery paths are unsupported")
		}
	}
	file := filepath.Clean(filepath.Join(base, relative))
	if !strings.HasPrefix(file, root+string(filepath.Separator)) {
		return "", fmt.Errorf("output is outside module")
	}
	rel, _ := filepath.Rel(root, file)
	current := root
	for _, part := range strings.Split(rel, string(filepath.Separator)) {
		current = filepath.Join(current, part)
		info, err := os.Lstat(current)
		if err != nil && !os.IsNotExist(err) {
			return "", err
		}
		if err == nil && (info.Mode()&os.ModeSymlink != 0 || current == file && !info.Mode().IsRegular() || current != file && !info.IsDir()) {
			return "", fmt.Errorf("unsafe output path %s", current)
		}
		if current != file {
			for _, manifest := range []string{"goml.toml", "go.mod"} {
				if _, err := os.Lstat(filepath.Join(current, manifest)); err == nil {
					return "", fmt.Errorf("output crosses nested module %s", current)
				}
			}
		}
	}
	return file, nil
}

func Load(file string) (Project, error) {
	file, err := filepath.Abs(file)
	if err != nil {
		return Project{}, err
	}
	directory := filepath.Dir(file)
	root := directory
	for {
		if _, err := os.Stat(filepath.Join(root, "goml.toml")); err == nil {
			break
		}
		parent := filepath.Dir(root)
		if parent == root {
			return Project{}, fmt.Errorf("expected a GoML module")
		}
		root = parent
	}
	root, err = filepath.EvalSymlinks(root)
	if err != nil {
		return Project{}, err
	}
	relative, err := filepath.Rel(root, file)
	if err != nil {
		return Project{}, err
	}
	if _, err := checkedPath(root, root, relative); err != nil {
		return Project{}, err
	}
	info, err := os.Stat(file)
	if err != nil || !info.Mode().IsRegular() || info.Size() > 1<<20 {
		return Project{}, fmt.Errorf("configuration must be a regular file of at most 1 MiB")
	}
	data, err := os.ReadFile(file)
	if err != nil {
		return Project{}, err
	}
	var c Config
	if err := decode(data, &c); err != nil {
		return Project{}, err
	}
	if err := validate(c); err != nil {
		return Project{}, err
	}
	goml, err := checkedPath(root, directory, c.Output)
	if err != nil || filepath.Ext(goml) != ".gom" {
		return Project{}, fmt.Errorf("invalid GoML output: %v", err)
	}
	goFile, err := checkedPath(root, directory, c.GoOutput)
	if err != nil || filepath.Ext(goFile) != ".go" || filepath.Dir(goFile) == root {
		return Project{}, fmt.Errorf("native output requires a .go file in a module subdirectory: %v", err)
	}
	mod, err := os.ReadFile(filepath.Join(root, "go.mod"))
	if err != nil {
		return Project{}, fmt.Errorf("expected module-root go.mod: %w", err)
	}
	module := ""
	for _, line := range strings.Split(string(mod), "\n") {
		fields := strings.Fields(line)
		if len(fields) >= 2 && fields[0] == "module" {
			module = strings.Trim(fields[1], "\"")
		}
	}
	if module == "" || strings.ContainsAny(module, "\"\r\n\\") {
		return Project{}, fmt.Errorf("invalid Go module identity")
	}
	rel, _ := filepath.Rel(root, filepath.Dir(goFile))
	manifest, err := checkedPath(root, directory, filepath.Base(file)+".goml-c-bind.json")
	if err != nil {
		return Project{}, err
	}
	if file == goml || file == goFile {
		return Project{}, fmt.Errorf("configuration and outputs must differ")
	}
	return Project{Config: c, File: file, Root: root, Directory: directory, GoImport: module + "/" + filepath.ToSlash(rel), GoFile: goFile, GomlFile: goml, Manifest: manifest}, nil
}
