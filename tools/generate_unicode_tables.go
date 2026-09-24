package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

type output struct {
	Names []string
	Arms  string
	Data  string
}

func bounds(length, shard int) (int, int) {
	size := (length + 15) / 16
	return min(length, shard*size), min(length, (shard+1)*size)
}

func generate(family string, shard int) output {
	result := output{}
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
		start, end := bounds(len(records), shard)
		result.Data = strings.Join(records[start:end], "")
	} else {
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
		start, end := bounds(len(names), shard)
		result.Names = names[start:end]
		var arms strings.Builder
		for _, name := range result.Names {
			var data strings.Builder
			for _, value := range tables[name].R16 {
				fmt.Fprintf(&data, "%06x%06x%08x", value.Lo, value.Hi, value.Stride)
			}
			for _, value := range tables[name].R32 {
				fmt.Fprintf(&data, "%06x%06x%08x", value.Lo, value.Hi, value.Stride)
			}
			fmt.Fprintf(&arms, "        %q => Option::Some(RangeTable { data: %q }),\n", name, data.String())
		}
		result.Arms = arms.String()
	}
	return result
}

func render(family string) string {
	var names []string
	var arms, data strings.Builder
	for shard := 0; shard < 16; shard++ {
		part := generate(family, shard)
		for _, name := range part.Names {
			names = append(names, strconv.Quote(name))
		}
		arms.WriteString(part.Arms)
		data.WriteString(part.Data)
	}
	if family == "case" || family == "simple_fold" {
		return fmt.Sprintf("package unicode;\n\nconst %s_DATA: string = %q;\n", strings.ToUpper(family), data.String())
	}
	return fmt.Sprintf("package unicode;\n\npub fn %s_names() -> Vec[string] {\n    Vec::from_array([%s])\n}\n\npub fn %s(name: string) -> Option[RangeTable] {\n    match name {\n%s        _ => Option::None,\n    }\n}\n", family, strings.Join(names, ", "), family, arms.String())
}

func main() {
	if unicode.Version != "15.0.0" {
		panic("require Unicode 15.0.0")
	}
	if len(os.Args) == 2 && (os.Args[1] == "--check" || os.Args[1] == "--write") {
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
			path := filepath.Join("lib", "std", "unicode", "tables_"+family+".gom")
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
		return
	}
	if len(os.Args) != 3 {
		panic("usage: --check | --write | family shard (0..15)")
	}
	shard, err := strconv.Atoi(os.Args[2])
	if err != nil || shard < 0 || shard >= 16 {
		panic("invalid shard")
	}
	if err := json.NewEncoder(os.Stdout).Encode(generate(os.Args[1], shard)); err != nil {
		panic(err)
	}
}
