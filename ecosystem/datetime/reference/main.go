package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
)

func mustNumber(text string) int64 {
	number, err := strconv.ParseInt(text, 10, 64)
	if err != nil {
		panic(err)
	}
	return number
}

func main() {
	zones := map[string]*time.Location{}
	scanner := bufio.NewScanner(os.Stdin)
	output := bufio.NewWriter(os.Stdout)
	defer output.Flush()
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), "\t")
		name := fields[len(fields)-1]
		zone := zones[name]
		if zone == nil {
			data, err := os.ReadFile(filepath.Join(os.Args[1], name))
			if err != nil {
				panic(err)
			}
			zone, err = time.LoadLocationFromTZData(name, data)
			if err != nil {
				panic(err)
			}
			zones[name] = zone
		}
		if fields[0] == "instant" {
			value := time.Unix(mustNumber(fields[1]), mustNumber(fields[2]))
			local := value.In(zone)
			abbreviation, offset := local.Zone()
			fmt.Fprintf(output, "%sZ\t%s\t%d\t%s\t%t\n", value.UTC().Format("2006-01-02T15:04:05.999999999"), local.Format("2006-01-02T15:04:05.999999999"), offset, abbreviation, local.IsDST())
		} else {
			local, err := time.Parse("2006-01-02T15:04:05.999999999", fields[1])
			if err != nil {
				panic(err)
			}
			candidates := []int64{}
			for _, offset := range []int64{-14400, -10800, 0, 3600} {
				candidate := local.Add(-time.Duration(offset) * time.Second)
				if candidate.In(zone).Format("2006-01-02T15:04:05.999999999") == fields[1] {
					candidates = append(candidates, candidate.Unix())
				}
			}
			sort.Slice(candidates, func(i, j int) bool { return candidates[i] < candidates[j] })
			switch len(candidates) {
			case 0:
				fmt.Fprintln(output, "gap")
			case 1:
				fmt.Fprintf(output, "unique\t%d\t%d\n", candidates[0], local.Nanosecond())
			case 2:
				fmt.Fprintf(output, "fold\t%d\t%d\t%d\n", candidates[0], candidates[1], local.Nanosecond())
			default:
				panic("unexpected number of candidates")
			}
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
