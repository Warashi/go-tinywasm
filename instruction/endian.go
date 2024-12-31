package instruction

import (
	"encoding/binary"
	"io"
)

var endian = binary.LittleEndian

func readF32(r io.Reader) (float32, error) {
	var b [4]byte
	if _, err := io.ReadFull(r, b[:]); err != nil {
		return 0, err
	}

	var f float32
	if _, err := binary.Decode(b[:], endian, &f); err != nil {
		return 0, err
	}

	return f, nil
}

func readF64(r io.Reader) (float64, error) {
	var b [8]byte
	if _, err := io.ReadFull(r, b[:]); err != nil {
		return 0, err
	}

	var f float64
	if _, err := binary.Decode(b[:], endian, &f); err != nil {
		return 0, err
	}

	return f, nil
}
