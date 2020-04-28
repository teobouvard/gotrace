package main

import (
	"flag"
	"log"
	"os"
	"runtime/pprof"

	"github.com/teobouvard/gotrace"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	flag.Parse()

	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	scene := gotrace.LightMarbleScene()
	scene.Render()
}
