package runtime

import "github.com/Warashi/go-tinywasm/types/runtime"

type ImportFunc func(*Store, ...runtime.Value) ([]runtime.Value, error)
type Import map[string]map[string]ImportFunc
