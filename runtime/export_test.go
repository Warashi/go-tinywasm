package runtime

import "github.com/Warashi/go-tinywasm/types/runtime"

func (r *Runtime) Store() *Store {
	return r.store
}

func (s *Store) Memories() []runtime.MemoryInst {
	return s.memories
}
