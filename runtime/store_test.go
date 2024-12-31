package runtime

import (
	"bytes"
	"os"
	"testing"

	"github.com/Warashi/go-tinywasm/binary"
)

func TestInitMemory(t *testing.T) {
	t.Parallel()

	b, err := os.ReadFile("../testdata/memory.wasm")
	if err != nil {
		t.Errorf("failed to load testdata: %v", err)
		t.FailNow()
	}

	module, err := binary.NewModule(bytes.NewReader(b))
	if err != nil {
		t.Errorf("failed to parse wasm: %v", err)
		t.FailNow()
	}

	store, err := NewStore(module)
	if err != nil {
		t.Errorf("failed to create store: %v", err)
		t.FailNow()
	}

	if len(store.memories) != 1 {
		t.Errorf("unexpected number of memories: %d", len(store.memories))
	}
	if len(store.memories[0].Data) != 65536 {
		t.Errorf("unexpected memory size: %d", len(store.memories[0].Data))
	}
	if store.memories[0].Max != 0 {
		t.Errorf("unexpected max memory size: %d", store.memories[0].Max)
	}
	if string(store.memories[0].Data[0:5]) != "hello" {
		t.Errorf("unexpected memory content: %s", string(store.memories[0].Data[0:5]))
	}
	if string(store.memories[0].Data[5:10]) != "world" {
		t.Errorf("unexpected memory content: %s", string(store.memories[0].Data[5:11]))
	}
}
