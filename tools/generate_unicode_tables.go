package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

func render(family string) string {
	if family == "case" || family == "simple_fold" {
		var records []string
		for value := rune(0); value <= unicode.MaxRune; value++ {
			if value >= 0xd800 && value <= 0xdfff {
				continue
			}
			if family == "case" {
				upper, lower, title := unicode.ToUpper(value), unicode.ToLower(value), unicode.ToTitle(value)
				if upper != value || lower != value || title != value {
					records = append(records, fmt.Sprintf("%06x%06x%06x%06x", value, upper, lower, title))
				}
			} else if next := unicode.SimpleFold(value); next != value {
				records = append(records, fmt.Sprintf("%06x%06x", value, next))
			}
		}
		return fmt.Sprintf("package unicode;\n\nconst %s_DATA: string = %q;\n", strings.ToUpper(family), strings.Join(records, ""))
	}
	families := map[string]map[string]*unicode.RangeTable{
		"category":      unicode.Categories,
		"script":        unicode.Scripts,
		"property":      unicode.Properties,
		"fold_category": unicode.FoldCategory,
		"fold_script":   unicode.FoldScript,
	}
	tables, ok := families[family]
	if !ok {
		panic("unknown family")
	}
	var names []string
	for name := range tables {
		names = append(names, name)
	}
	sort.Strings(names)
	var arms strings.Builder
	for _, name := range names {
		var data strings.Builder
		for _, value := range tables[name].R16 {
			fmt.Fprintf(&data, "%06x%06x%08x", value.Lo, value.Hi, value.Stride)
		}
		for _, value := range tables[name].R32 {
			fmt.Fprintf(&data, "%06x%06x%08x", value.Lo, value.Hi, value.Stride)
		}
		fmt.Fprintf(&arms, "        %q => Option::Some(RangeTable { data: %q }),\n", name, data.String())
	}
	for index, name := range names {
		names[index] = strconv.Quote(name)
	}
	return fmt.Sprintf("package unicode;\n\npub fn %s_names() -> Vec[string] {\n    Vec::from_array([%s])\n}\n\npub fn %s(name: string) -> Option[RangeTable] {\n    match name {\n%s        _ => Option::None,\n    }\n}\n", family, strings.Join(names, ", "), family, arms.String())
}

func main() {
	if unicode.Version != "15.0.0" {
		panic("require Unicode 15.0.0")
	}
	if len(os.Args) != 2 || (os.Args[1] != "--check" && os.Args[1] != "--write") {
		panic("usage: --check | --write")
	}
	formatter := os.Getenv("GOMLFMT")
	if formatter == "" {
		formatter = "stage2/bin/gomlfmt"
	}
	for _, family := range []string{"category", "script", "property", "fold_category", "fold_script", "case", "simple_fold"} {
		command := exec.Command(formatter)
		command.Stdin = strings.NewReader(render(family))
		command.Stderr = os.Stderr
		expected, err := command.Output()
		if err != nil {
			panic(err)
		}
		path := filepath.Join("lib", "std", "unicode", "tables_"+family+".goml")
		if os.Args[1] == "--write" {
			if err := os.WriteFile(path, expected, 0644); err != nil {
				panic(err)
			}
		} else {
			actual, err := os.ReadFile(path)
			if err != nil {
				panic(err)
			}
			if !bytes.Equal(actual, expected) {
				panic("Unicode table differs: " + path)
			}
		}
		fmt.Println(path)
	}
}
