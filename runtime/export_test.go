package runtime

import "github.com/Warashi/wasmium/types/runtime"

func (r *Runtime) Store() *Store {
	return r.store
}

func (s *Store) Memories() []runtime.MemoryInst {
	return s.memories
}
