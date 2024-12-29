package binary

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"

	"github.com/Warashi/go-tinywasm/leb128"
)

type Module struct {
	magic           string
	version         uint32
	typeSection     []FuncType
	functionSection []uint32
}

func NewModule(r io.Reader) (*Module, error) {
	return decode(r)
}

func decode(r io.Reader) (*Module, error) {
	var (
		err    error
		module = new(Module)
	)

	module.magic, module.version, err = decodePreamble(r)
	if err != nil {
		return nil, err
	}

	for {
		code, size, err := decodeSectionHeader(r)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return nil, fmt.Errorf("failed to decode section header: %w", err)
		}

		sectionContents := make([]byte, size)
		if _, err := io.ReadFull(r, sectionContents); err != nil {
			return nil, fmt.Errorf("failed to read section: %w", err)
		}

		section := bytes.NewReader(sectionContents)

		switch code {
		case SectionCodeType:
			module.typeSection, err = decodeTypeSection(section)
			if err != nil {
				return nil, fmt.Errorf("failed to decode type section: %w", err)
			}
		case SectionCodeFunction:
			module.functionSection, err = decodeFunctionSection(section)
			if err != nil {
				return nil, fmt.Errorf("failed to decode function section: %w", err)
			}
		default:
			return nil, fmt.Errorf("unsupported section code: %d", code)
		}
	}

	return module, nil
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

func decodeSectionHeader(r io.Reader) (SectionCode, uint32, error) {
	var (
		code [1]byte
	)
	if _, err := io.ReadFull(r, code[:]); err != nil {
		return 0, 0, fmt.Errorf("failed to read section code: %w", err)
	}

	size, err := leb128.Uint32(r)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to read section size: %w", err)
	}

	return SectionCode(code[0]), size, nil
}

func decodeTypeSection(r io.Reader) ([]FuncType, error) {
	return nil, nil
}

func decodeFunctionSection(r io.Reader) ([]uint32, error) {
	count, err := leb128.Uint32(r)
	if err != nil {
		return nil, fmt.Errorf("failed to read function count: %w", err)
	}

	idxs := make([]uint32, 0, count)

	for range count {
		idx, err := leb128.Uint32(r)
		if err != nil {
			return nil, fmt.Errorf("failed to read function index: %w", err)
		}
		idxs = append(idxs, idx)
	}

	return idxs, nil
}
