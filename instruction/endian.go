package instruction

import (
	"encoding/binary"
	"io"
)

var endian = binary.LittleEndian

func readF32(r io.Reader) ([4]byte, error) {
	var b [4]byte
	if _, err := io.ReadFull(r, b[:]); err != nil {
		return b, err
	}

	return b, nil
}

func readF64(r io.Reader) ([8]byte, error) {
	var b [8]byte
	if _, err := io.ReadFull(r, b[:]); err != nil {
		return b, err
	}

	return b, nil
}
