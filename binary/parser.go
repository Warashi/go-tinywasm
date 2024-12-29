package binary

import (
	"bytes"
	"fmt"
	"io"
)

type ints interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

func take[T ints](n T) func(r io.Reader) (io.Reader, error) {
	return func(r io.Reader) (io.Reader, error) {
		b := make([]byte, int(n))
		if _, err := io.ReadFull(r, b); err != nil {
			return nil, fmt.Errorf("failed to read %d bytes: %w", n, err)
		}
		return bytes.NewReader(b), nil
	}
}
