package main

import (
	_ "embed"

	"flag"
	"log"
	"os"
	"runtime/pprof"

	"github.com/Warashi/wasmium/runtime"
	"github.com/Warashi/wasmium/wasip1"
)

func main() {
	var prof bool
	flag.BoolVar(&prof, "prof", false, "record cpuprofile with profile.out")
	flag.Parse()

	if prof {
		f, err := os.Create("profile.out")
		if err != nil {
			log.Fatalln(err)
		}
		defer func() {
			if err := f.Close(); err != nil {
				log.Fatalln(err)
			}
		}()
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatalln(err)
		}
		defer pprof.StopCPUProfile()
	}

	f, err := os.Open(flag.Arg(0))
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	r, err := runtime.New(f)
	if err != nil {
		log.Fatalln(err)
	}

	wasip1.NewWasiPreview1().Register(r)

	if _, err := r.Call("_start"); err != nil {
		log.Fatalln(err)
	}
}
