package cbind

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type ownership struct {
	Version int               `json:"version"`
	Files   map[string]string `json:"files"`
}

func digest(data []byte) string {
	value := sha256.Sum256(data)
	return hex.EncodeToString(value[:])
}

func owned(p Project) (map[string][]byte, error) {
	data, err := os.ReadFile(p.Manifest)
	if os.IsNotExist(err) {
		for _, file := range []string{p.GomlFile, p.GoFile} {
			if _, err := os.Lstat(file); !os.IsNotExist(err) {
				return nil, fmt.Errorf("refusing unowned output %s", file)
			}
		}
		return map[string][]byte{}, nil
	}
	if err != nil {
		return nil, err
	}
	var manifest ownership
	if err := decode(data, &manifest); err != nil {
		return nil, fmt.Errorf("invalid ownership manifest: %w", err)
	}
	if manifest.Version != 1 || len(manifest.Files) != 2 {
		return nil, fmt.Errorf("invalid ownership manifest")
	}
	result := map[string][]byte{p.Manifest: data}
	for _, file := range []string{p.GomlFile, p.GoFile} {
		relative, _ := filepath.Rel(p.Directory, file)
		data, err := os.ReadFile(file)
		if err != nil || digest(data) != manifest.Files[filepath.ToSlash(relative)] {
			return nil, fmt.Errorf("modified or missing generated output %s; restore it before regenerating", file)
		}
		result[file] = data
	}
	return result, nil
}

func formatGoml(ctx context.Context, p Project, source []byte, formatter string) ([]byte, error) {
	output, err := run(ctx, p.Root, string(source), formatter)
	if err != nil {
		return nil, fmt.Errorf("GoML formatter: %w", err)
	}
	return output, nil
}

func validateNative(ctx context.Context, p Project, generated generated, temporary string) error {
	missing := []string{}
	for directory := filepath.Dir(p.GoFile); directory != p.Root; directory = filepath.Dir(directory) {
		if _, err := os.Lstat(directory); os.IsNotExist(err) {
			missing = append(missing, directory)
		} else if err != nil {
			return err
		} else {
			break
		}
	}
	defer func() {
		for _, directory := range missing {
			os.Remove(directory)
		}
	}()
	if err := os.MkdirAll(filepath.Dir(p.GoFile), 0755); err != nil {
		return err
	}
	file := filepath.Join(temporary, "generated.go")
	if err := os.WriteFile(file, generated.Go, 0600); err != nil {
		return err
	}
	data, _ := json.Marshal(map[string]any{"Replace": map[string]string{p.GoFile: file}})
	overlay := filepath.Join(temporary, "overlay.json")
	if err := os.WriteFile(overlay, data, 0600); err != nil {
		return err
	}
	var source strings.Builder
	fmt.Fprintf(&source, "package main\nimport (\"fmt\"; native %q)\nfunc main() {\n", p.GoImport)
	for _, ty := range p.Config.Types {
		fmt.Fprintf(&source, "fmt.Printf(\"%%v\", native.GomlC_Null%s)\n", ty.Name)
	}
	for _, fn := range p.Config.Functions {
		fmt.Fprintf(&source, "fmt.Printf(\"%%v\", native.GomlC_%s)\n", fn.Name)
	}
	for _, constant := range p.Config.Constants {
		fmt.Fprintf(&source, "fmt.Printf(\"%%v\", native.GomlC_%s)\n", constant.Name)
	}
	source.WriteString("}\n")
	main := filepath.Join(temporary, "main.go")
	if err := os.WriteFile(main, []byte(source.String()), 0600); err != nil {
		return err
	}
	_, err := run(ctx, p.Root, "", "go", "build", "-mod=readonly", "-overlay="+overlay, "-o", filepath.Join(temporary, "probe"), main)
	if err != nil {
		return fmt.Errorf("native compile/link validation: %w", err)
	}
	return nil
}

func publish(p Project, generated generated, previous map[string][]byte, temporary string) error {
	files := map[string][]byte{p.GoFile: generated.Go, p.GomlFile: generated.Goml}
	manifest := ownership{Version: 1, Files: map[string]string{}}
	for file, data := range files {
		relative, _ := filepath.Rel(p.Directory, file)
		manifest.Files[filepath.ToSlash(relative)] = digest(data)
	}
	encoded, _ := json.MarshalIndent(manifest, "", "  ")
	files[p.Manifest] = append(encoded, '\n')
	order := []string{p.GoFile, p.GomlFile, p.Manifest}
	recovery := map[string]string{}
	for index, file := range order {
		if old, exists := previous[file]; exists {
			name := "previous-" + strconv.Itoa(index)
			if err := os.WriteFile(filepath.Join(temporary, name), old, 0600); err != nil {
				return err
			}
			recovery[file] = name
		}
	}
	backup, _ := json.MarshalIndent(recovery, "", "  ")
	if err := os.WriteFile(filepath.Join(temporary, "previous.json"), backup, 0600); err != nil {
		return err
	}
	for index, file := range order {
		if err := os.MkdirAll(filepath.Dir(file), 0755); err != nil {
			return err
		}
		if err := os.WriteFile(filepath.Join(temporary, strconv.Itoa(index)), files[file], 0644); err != nil {
			return err
		}
	}
	for _, file := range order {
		current, err := os.ReadFile(file)
		old, existed := previous[file]
		if existed && (err != nil || !bytes.Equal(current, old)) || !existed && !os.IsNotExist(err) {
			return fmt.Errorf("output changed during generation: %s", file)
		}
	}
	for index, file := range order {
		if err := os.Rename(filepath.Join(temporary, strconv.Itoa(index)), file); err != nil {
			for _, written := range order[:index] {
				if old, existed := previous[written]; existed {
					if restore := os.WriteFile(written, old, 0644); restore != nil {
						return fmt.Errorf("publication failed: %v; rollback %s: %w", err, written, restore)
					}
				} else if restore := os.Remove(written); restore != nil {
					return fmt.Errorf("publication failed: %v; rollback: %w", err, restore)
				}
			}
			return err
		}
	}
	return nil
}

func Execute(ctx context.Context, args []string, formatter string, output io.Writer) error {
	file := ""
	dry, check, positional := false, false, false
	for _, arg := range args {
		if !positional {
			switch arg {
			case "--help", "-h":
				fmt.Fprintln(output, "Usage: goml bind-c <CONFIG> [--dry-run | --check]\n\nGenerate allowlisted C bindings using Clang and cgo.\n\n  --dry-run  validate paths and print outputs without writing\n  --check    verify generated sources and print the C input fingerprint")
				return nil
			case "--":
				positional = true
				continue
			case "--dry-run":
				if dry || check {
					return fmt.Errorf("expected one mode")
				}
				dry = true
				continue
			case "--check":
				if dry || check {
					return fmt.Errorf("expected one mode")
				}
				check = true
				continue
			}
			if strings.HasPrefix(arg, "-") {
				return fmt.Errorf("unknown option %s", arg)
			}
		}
		if file != "" {
			return fmt.Errorf("expected one configuration file")
		}
		file = arg
	}
	if file == "" {
		return fmt.Errorf("expected a configuration file; see goml bind-c --help")
	}
	p, err := Load(file)
	if err != nil {
		return err
	}
	configuration, err := os.ReadFile(p.File)
	if err != nil {
		return err
	}
	previous, err := owned(p)
	if err != nil {
		return err
	}
	if dry {
		fmt.Fprintf(output, "%s\n%s\n%s\n", p.GomlFile, p.GoFile, p.Manifest)
		return nil
	}
	if check && len(previous) == 0 {
		return fmt.Errorf("C bindings have not been generated; run goml bind-c %s", file)
	}
	lock := filepath.Join(p.Root, ".goml-bind-c-lock")
	if !check {
		if err := os.Mkdir(lock, 0700); err != nil {
			return fmt.Errorf("cannot acquire generator lock %s: %w", lock, err)
		}
		defer os.Remove(lock)
		previous, err = owned(p)
		if err != nil {
			return err
		}
	}
	w, err := Inspect(ctx, p)
	if err != nil {
		return err
	}
	generated, err := Generate(w)
	if err != nil {
		return err
	}
	generated.Goml, err = formatGoml(ctx, p, generated.Goml, formatter)
	if err != nil {
		return err
	}
	if check {
		if !bytes.Equal(previous[p.GomlFile], generated.Goml) || !bytes.Equal(previous[p.GoFile], generated.Go) {
			return fmt.Errorf("C declarations or binding configuration changed; run goml bind-c %s", file)
		}
		fmt.Fprintln(output, w.Fingerprint)
		return nil
	}
	temporary, err := os.MkdirTemp(p.Root, ".goml-bind-c-stage-")
	if err != nil {
		return err
	}
	retain := false
	defer func() {
		if !retain {
			os.RemoveAll(temporary)
		}
	}()
	if err := validateNative(ctx, p, generated, temporary); err != nil {
		return err
	}
	current, err := os.ReadFile(p.File)
	if err != nil || !bytes.Equal(current, configuration) {
		return fmt.Errorf("configuration changed during generation")
	}
	if _, err := Load(file); err != nil {
		return err
	}
	if err := publish(p, generated, previous, temporary); err != nil {
		retain = true
		return fmt.Errorf("%w; preserved recovery files in %s", err, temporary)
	}
	fmt.Fprintf(output, "Generated %s and %s\n", p.GomlFile, p.GoFile)
	return nil
}
