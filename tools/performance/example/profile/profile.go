package profile

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

var count = flag.Int("iterations", 2000000, "number of interpreter steps")
var profilePath = flag.String("profile", "", "CPU profile output")
var profileFile *os.File
var started time.Time
var before runtime.MemStats

func Iterations() int {
	flag.Parse()
	if *count < 1 {
		panic("iterations must be positive")
	}
	return *count
}

func Start() {
	if *profilePath != "" {
		var err error
		profileFile, err = os.Create(*profilePath)
		if err != nil {
			panic(err)
		}
		if err := pprof.StartCPUProfile(profileFile); err != nil {
			panic(err)
		}
	}
	runtime.ReadMemStats(&before)
	started = time.Now()
}

func Finish(checksum uint64, steps int) {
	elapsed := time.Since(started)
	var after runtime.MemStats
	runtime.ReadMemStats(&after)
	if profileFile != nil {
		pprof.StopCPUProfile()
		if err := profileFile.Close(); err != nil {
			panic(err)
		}
	}
	result := struct {
		Checksum       string `json:"checksum"`
		Steps          int    `json:"interpreter_steps"`
		Elapsed        int64  `json:"work_ns"`
		Allocations    uint64 `json:"allocations"`
		AllocatedBytes uint64 `json:"allocated_bytes"`
	}{fmt.Sprintf("%016x", checksum), steps, elapsed.Nanoseconds(), after.Mallocs - before.Mallocs, after.TotalAlloc - before.TotalAlloc}
	if err := json.NewEncoder(os.Stdout).Encode(result); err != nil {
		panic(err)
	}
}
