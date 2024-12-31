package runtime

import "fmt"

var (
	ErrOutOfBounds       = fmt.Errorf("out of bounds")
	ErrMemoryOutOfBounds = fmt.Errorf("memory out of bounds")
	ErrInvalidValue      = fmt.Errorf("invalid value")
)
