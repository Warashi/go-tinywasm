package execution

func (r *Runtime) Store() *Store {
	return r.store
}

func (s *Store) Memories() []MemoryInst {
	return s.memories
}

func (m *MemoryInst) Data() []byte {
	return m.data
}
