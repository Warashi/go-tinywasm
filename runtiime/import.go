package runtime

type ImportFunc func(*Store, ...Value) ([]Value, error)
type Import map[string]map[string]ImportFunc
