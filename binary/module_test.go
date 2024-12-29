package binary

import (
	"bytes"
	"os"
	"testing"
)

func TestParsePreamble(t *testing.T) {
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
