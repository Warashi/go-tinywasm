package main

import (
	"bytes"
	_ "embed"
	"log"

	"github.com/Warashi/go-tinywasm/execution"
)

//go:embed testdata/hello_world.wasm
var helloWorld []byte

func main() {
	runtime, err := execution.NewRuntimeWithWasi(bytes.NewReader(helloWorld))
	if err != nil {
		log.Fatalln(err)
	}

	if _, err := runtime.Call("_start", nil); err != nil {
		log.Fatalln(err)
	}
}
