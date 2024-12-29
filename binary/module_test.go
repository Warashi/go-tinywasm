package binary

import (
	"bytes"
	"os"
	"testing"
)

func TestDecodePreamble(t *testing.T) {
	b, err := os.ReadFile("../testdata/minimal.wasm")
	if err != nil {
		t.Errorf("failed to load testdata: %v", err)
		t.FailNow()
	}

	m, err := NewModule(bytes.NewReader(b))
	if err != nil {
		t.Errorf("failed to parse wasm: %v", err)
	}

	if m.magic != "\x00asm" {
		t.Errorf("wrong magic bytes: %x", m.magic)
	}
	if m.version != 1 {
		t.Errorf("wrong version: %d", m.version)
	}
}

func TestDecodeMinimalFunc(t *testing.T) {
	b, err := os.ReadFile("../testdata/minimal_func.wasm")
	if err != nil {
		t.Errorf("failed to load testdata: %v", err)
		t.FailNow()
	}

	m, err := NewModule(bytes.NewReader(b))
	if err != nil {
		t.Errorf("failed to parse wasm: %v", err)
	}

	if len(m.typeSection) != 0 {
		t.Errorf("wrong type section length: %d", len(m.typeSection))
	}
	if len(m.functionSection) != 1 {
		t.Errorf("wrong function section length: %d", len(m.functionSection))
	}
	if len(m.codeSection) != 0 {
		t.Errorf("wrong code section length: %d", len(m.codeSection))
	}
}
