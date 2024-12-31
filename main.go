package main

import (
	"bytes"
	_ "embed"
	"log"

	"github.com/Warashi/go-tinywasm/runtime"
)

//go:embed testdata/hello_world.wasm
var helloWorld []byte

func main() {
	r, err := runtime.New(bytes.NewReader(helloWorld))
	if err != nil {
		log.Fatalln(err)
	}

	if _, err := r.Call("_start", nil); err != nil {
		log.Fatalln(err)
	}
}
