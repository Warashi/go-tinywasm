package binary

import (
	"encoding/binary"
	"fmt"
	"io"
)

type Module struct {
	magic   string
	version uint32
}

func NewModule(r io.Reader) (*Module, error) {
	return decode(r)
}

func decode(r io.Reader) (*Module, error) {
	magic, version, err := decodePreamble(r)
	if err != nil {
		return nil, err
	}

	return &Module{
		magic:   magic,
		version: version,
	}, nil
}

func decodePreamble(r io.Reader) (string, uint32, error) {
	var (
		magic   [4]byte
		version [4]byte
	)
	if _, err := io.ReadFull(r, magic[:]); err != nil {
		return "", 0, fmt.Errorf("failed to read magic binary: %w", err)
	}
	if string(magic[:]) != "\x00asm" {
		return "", 0, fmt.Errorf("invalid magic header: %x", magic[:])
	}
	if _, err := io.ReadFull(r, version[:]); err != nil {
		return "", 0, fmt.Errorf("failed to read version: %w", err)
	}

	return string(magic[:]), binary.LittleEndian.Uint32(version[:]), nil
}
