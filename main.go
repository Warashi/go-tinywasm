package main

import (
	_ "embed"

	"flag"
	"log"
	"os"

	"github.com/Warashi/wasmium/runtime"
	"github.com/Warashi/wasmium/wasip1"
)

func main() {
	flag.Parse()

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

	if _, err := r.Call("_start", nil); err != nil {
		log.Fatalln(err)
	}
}
