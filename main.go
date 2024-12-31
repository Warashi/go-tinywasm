package main

import (
	"bytes"
	_ "embed"
	"log"

	"github.com/Warashi/go-tinywasm/runtime"
	"github.com/Warashi/go-tinywasm/wasip1"
)

//go:embed testdata/hello_world.wasm
var helloWorld []byte

func main() {
	r, err := runtime.New(bytes.NewReader(helloWorld))
	if err != nil {
		log.Fatalln(err)
	}

	wasip1.NewWasiPreview1().Register(r)

	if _, err := r.Call("_start", nil); err != nil {
		log.Fatalln(err)
	}
}
