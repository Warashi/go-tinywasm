package binary

import (
	"errors"
	"fmt"
	"io"

	"github.com/Warashi/wasmium/leb128"
	"github.com/Warashi/wasmium/opcode"
	"github.com/Warashi/wasmium/types/binary"
)

type Module struct {
	magic           string
	version         uint32
	memorySection   []binary.Memory
	dataSection     []binary.Data
	typeSection     []binary.FuncType
	functionSection []uint32
	codeSection     []binary.Function
	exportSection   []binary.Export
	importSection   []binary.Import
}

func NewModule(r io.Reader) (*Module, error) {
	return decode(r)
}

func (m *Module) MemorySection() []binary.Memory { return m.memorySection }
func (m *Module) DataSection() []binary.Data     { return m.dataSection }
func (m *Module) TypeSection() []binary.FuncType { return m.typeSection }
func (m *Module) FunctionSection() []uint32      { return m.functionSection }
func (m *Module) CodeSection() []binary.Function { return m.codeSection }
func (m *Module) ExportSection() []binary.Export { return m.exportSection }
func (m *Module) ImportSection() []binary.Import { return m.importSection }

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

		sectionContents, err := take(size)(r)
		if err != nil {
			return nil, fmt.Errorf("failed to take section contents: %w", err)
		}

		switch code {
		case SectionCodeCustom:
			// ignore custom sections
		case SectionCodeType:
			module.typeSection, err = decodeTypeSection(sectionContents)
			if err != nil {
				return nil, fmt.Errorf("failed to decode type section: %w", err)
			}
		case SectionCodeImport:
			module.importSection, err = decodeImportSection(sectionContents)
			if err != nil {
				return nil, fmt.Errorf("failed to decode import section: %w", err)
			}
		case SectionCodeFunction:
			module.functionSection, err = decodeFunctionSection(sectionContents)
			if err != nil {
				return nil, fmt.Errorf("failed to decode function section: %w", err)
			}
		case SectionCodeMemory:
			module.memorySection, err = decodeMemorySection(sectionContents)
			if err != nil {
				return nil, fmt.Errorf("failed to decode memory section: %w", err)
			}
		case SectionCodeExport:
			module.exportSection, err = decodeExportSection(sectionContents)
			if err != nil {
				return nil, fmt.Errorf("failed to decode export section: %w", err)
			}
		case SectionCodeCode:
			module.codeSection, err = decodeCodeSection(sectionContents)
			if err != nil {
				return nil, fmt.Errorf("failed to decode code section: %w", err)
			}
		case SectionCodeData:
			module.dataSection, err = decodeDataSection(sectionContents)
			if err != nil {
				return nil, fmt.Errorf("failed to decode data section: %w", err)
			}
		default:
			return nil, fmt.Errorf("unsupported section code: %d", code)
		}
	}

	return module, nil
}

func readByte(r io.Reader) (byte, error) {
	var (
		b [1]byte
	)
	if _, err := io.ReadFull(r, b[:]); err != nil {
		return 0, fmt.Errorf("failed to read byte: %w", err)
	}
	return b[0], nil
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

	return string(magic[:]), endian.Uint32(version[:]), nil
}

func decodeSectionHeader(r io.Reader) (SectionCode, uint32, error) {
	code, err := readByte(r)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to read section code: %w", err)
	}

	size, err := leb128.Uint32(r)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to read section size: %w", err)
	}

	return SectionCode(code), size, nil
}

func decodeTypeSection(r io.Reader) ([]binary.FuncType, error) {
	count, err := leb128.Uint32(r)
	if err != nil {
		return nil, fmt.Errorf("failed to read type section count: %w", err)
	}

	funcTypes := make([]binary.FuncType, 0, count)
	for range count {
		f, err := readByte(r)
		if err != nil {
			return nil, fmt.Errorf("failed to read function type: %w", err)
		}
		if f != 0x60 {
			return nil, fmt.Errorf("unsupported function type: %2x", f)
		}

		paramCount, err := leb128.Uint32(r)
		if err != nil {
			return nil, fmt.Errorf("failed to read parameter count: %w", err)
		}

		params := make([]binary.ValueType, 0, paramCount)
		for range paramCount {
			v, err := decodeValueType(r)
			if err != nil {
				return nil, fmt.Errorf("failed to decode parameter type: %w", err)
			}
			params = append(params, v)
		}

		resultCount, err := leb128.Uint32(r)
		if err != nil {
			return nil, fmt.Errorf("failed to read result count: %w", err)
		}

		results := make([]binary.ValueType, 0, resultCount)
		for range resultCount {
			v, err := decodeValueType(r)
			if err != nil {
				return nil, fmt.Errorf("failed to decode result type: %w", err)
			}
			results = append(results, v)
		}

		funcTypes = append(funcTypes, binary.FuncType{Params: params, Results: results})
	}

	return funcTypes, nil
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

func decodeCodeSection(r io.Reader) ([]binary.Function, error) {
	count, err := leb128.Uint32(r)
	if err != nil {
		return nil, fmt.Errorf("failed to read function count: %w", err)
	}

	functions := make([]binary.Function, 0, count)
	for range count {
		size, err := leb128.Uint32(r)
		if err != nil {
			return nil, fmt.Errorf("failed to read function size: %w", err)
		}
		body, err := take(size)(r)
		if err != nil {
			return nil, fmt.Errorf("failed to take function body: %w", err)
		}
		f, err := decodeFunctionBody(body)
		if err != nil {
			return nil, fmt.Errorf("failed to decode function body: %w", err)
		}

		functions = append(functions, f)
	}

	return functions, nil
}

func decodeValueType(r io.Reader) (binary.ValueType, error) {
	b, err := readByte(r)
	if err != nil {
		return 0, fmt.Errorf("failed to read value type: %w", err)
	}
	return binary.ValueType(b), nil
}

func decodeFunctionBody(r io.Reader) (binary.Function, error) {
	count, err := leb128.Uint32(r)
	if err != nil {
		return binary.Function{}, fmt.Errorf("failed to read local count: %w", err)
	}

	locals := make([]binary.FunctionLocal, 0, count)
	for range count {
		typeCount, err := leb128.Uint32(r)
		if err != nil {
			return binary.Function{}, fmt.Errorf("failed to read type count: %w", err)
		}
		valueType, err := decodeValueType(r)
		if err != nil {
			return binary.Function{}, fmt.Errorf("failed to decode value type: %w", err)
		}
		locals = append(locals, binary.FunctionLocal{TypeCount: typeCount, ValueType: valueType})
	}

	instructions, err := decodeInstructions(r)
	if err != nil {
		return binary.Function{}, fmt.Errorf("failed to decode instructions: %w", err)
	}

	return binary.Function{Locals: locals, Code: instructions}, nil
}

func decodeInstructions(r io.Reader) ([]binary.Instruction, error) {
	var (
		instructions []binary.Instruction
	)
	for {
		b, err := readByte(r)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return nil, fmt.Errorf("failed to read opcode: %w", err)
		}

		instruction, err := fromOpcode(opcode.Opcode(b))
		if err != nil {
			return nil, fmt.Errorf("failed to create instruction: %w", err)
		}
		if err := instruction.ReadOperandsFrom(r); err != nil {
			return nil, fmt.Errorf("failed to read operands: %w", err)
		}
		instructions = append(instructions, instruction)
	}

	return instructions, nil
}

func decodeExportSection(r io.Reader) ([]binary.Export, error) {
	count, err := leb128.Uint32(r)
	if err != nil {
		return nil, fmt.Errorf("failed to read export count: %w", err)
	}

	exports := make([]binary.Export, 0, count)

	for range count {
		name, err := decodeName(r)
		if err != nil {
			return nil, fmt.Errorf("failed to decode export name: %w", err)
		}
		kind, err := readByte(r)
		if err != nil {
			return nil, fmt.Errorf("failed to read export kind: %w", err)
		}
		index, err := leb128.Uint32(r)
		if err != nil {
			return nil, fmt.Errorf("failed to read export index: %w", err)
		}

		switch kind {
		case 0x00:
			exports = append(exports, binary.Export{Name: name, Desc: binary.ExportDescFunc{Index: index}})
		default:
			return nil, fmt.Errorf("unsupported export kind: %2x", kind)
		}
	}

	return exports, nil
}

func decodeImportSection(r io.Reader) ([]binary.Import, error) {
	count, err := leb128.Uint32(r)
	if err != nil {
		return nil, fmt.Errorf("failed to read import count: %w", err)
	}

	imports := make([]binary.Import, 0, count)

	for range count {
		module, err := decodeName(r)
		if err != nil {
			return nil, fmt.Errorf("failed to decode import module name: %w", err)
		}
		name, err := decodeName(r)
		if err != nil {
			return nil, fmt.Errorf("failed to decode import name: %w", err)
		}
		kind, err := readByte(r)
		if err != nil {
			return nil, fmt.Errorf("failed to read import kind: %w", err)
		}
		switch kind {
		case 0x00:
			index, err := leb128.Uint32(r)
			if err != nil {
				return nil, fmt.Errorf("failed to read import index: %w", err)
			}
			imports = append(imports, binary.Import{Module: module, Field: name, Desc: binary.ImportDescFunc{Index: index}})
		default:
			return nil, fmt.Errorf("unsupported import kind: %x", kind)
		}
	}

	return imports, nil
}

func decodeMemorySection(r io.Reader) ([]binary.Memory, error) {
	count, err := leb128.Uint32(r)
	if err != nil {
		return nil, fmt.Errorf("failed to read memory count: %w", err)
	}

	memories := make([]binary.Memory, 0, count)

	for range count {
		m := binary.Memory{}
		m.Limits, err = decodeLimits(r)
		if err != nil {
			return nil, fmt.Errorf("failed to decode memory limits: %w", err)
		}
		memories = append(memories, m)
	}

	return memories, nil
}

func decodeLimits(r io.Reader) (binary.Limits, error) {
	hasMax, err := leb128.Uint32(r)
	if err != nil {
		return binary.Limits{}, fmt.Errorf("failed to read hasMax: %w", err)
	}

	min, err := leb128.Uint32(r)
	if err != nil {
		return binary.Limits{}, fmt.Errorf("failed to read min: %w", err)
	}

	if hasMax == 0 {
		return binary.Limits{Min: min, Max: 0}, nil
	}

	max, err := leb128.Uint32(r)
	if err != nil {
		return binary.Limits{}, fmt.Errorf("failed to read max: %w", err)
	}

	return binary.Limits{Min: min, Max: max}, nil
}

func decodeName(r io.Reader) (string, error) {
	size, err := leb128.Uint32(r)
	if err != nil {
		return "", fmt.Errorf("failed to read name size: %w", err)
	}

	name := make([]byte, size)
	if _, err := io.ReadFull(r, name); err != nil {
		return "", fmt.Errorf("failed to read name: %w", err)
	}

	return string(name), nil
}

// TODO: implement expr decoding
func decodeExpr(r io.Reader) (uint32, error) {
	if _, err := leb128.Uint32(r); err != nil {
		return 0, fmt.Errorf("failed to read expr size: %w", err)
	}
	offset, err := leb128.Uint32(r)
	if err != nil {
		return 0, fmt.Errorf("failed to read expr offset: %w", err)

	}
	if _, err := leb128.Uint32(r); err != nil {
		return 0, fmt.Errorf("failed to read expr end: %w", err)
	}
	return offset, nil
}

func decodeDataSection(r io.Reader) ([]binary.Data, error) {
	count, err := leb128.Uint32(r)
	if err != nil {
		return nil, fmt.Errorf("failed to read data count: %w", err)
	}

	data := make([]binary.Data, 0, count)

	for range count {
		memoryIndex, err := leb128.Uint32(r)
		if err != nil {
			return nil, fmt.Errorf("failed to read memory index: %w", err)
		}
		offset, err := decodeExpr(r)
		if err != nil {
			return nil, fmt.Errorf("failed to read offset: %w", err)
		}
		size, err := leb128.Uint32(r)
		if err != nil {
			return nil, fmt.Errorf("failed to read data size: %w", err)
		}
		rest, err := take(size)(r)
		if err != nil {
			return nil, fmt.Errorf("failed to take data: %w", err)
		}
		init, err := io.ReadAll(rest)
		if err != nil {
			return nil, fmt.Errorf("failed to read data: %w", err)
		}
		data = append(data, binary.Data{MemoryIndex: memoryIndex, Offset: offset, Init: init})
	}

	return data, nil
}

func decodeBlock(r io.Reader) (binary.Block, error) {
	b, err := readByte(r)
	if err != nil {
		return binary.Block{}, fmt.Errorf("failed to read block type: %w", err)
	}

	switch b {
	case 0x40:
		return binary.Block{BlockType: binary.BlockTypeVoid{}}, nil
	default:
		return binary.Block{BlockType: binary.BlockTypeValue{ValueTypes: []binary.ValueType{binary.ValueType(b)}}}, nil
	}
}
