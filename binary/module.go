package binary

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"

	"github.com/Warashi/go-tinywasm/leb128"
)

type Module struct {
	magic           string
	version         uint32
	memorySection   []Memory
	dataSection     []Data
	typeSection     []FuncType
	functionSection []uint32
	codeSection     []Function
	exportSection   []Export
	importSection   []Import
}

func NewModule(r io.Reader) (*Module, error) {
	return decode(r)
}

func (m *Module) MemorySection() []Memory   { return m.memorySection }
func (m *Module) DataSection() []Data       { return m.dataSection }
func (m *Module) TypeSection() []FuncType   { return m.typeSection }
func (m *Module) FunctionSection() []uint32 { return m.functionSection }
func (m *Module) CodeSection() []Function   { return m.codeSection }
func (m *Module) ExportSection() []Export   { return m.exportSection }
func (m *Module) ImportSection() []Import   { return m.importSection }

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

	return string(magic[:]), binary.LittleEndian.Uint32(version[:]), nil
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

func decodeTypeSection(r io.Reader) ([]FuncType, error) {
	count, err := leb128.Uint32(r)
	if err != nil {
		return nil, fmt.Errorf("failed to read type section count: %w", err)
	}

	funcTypes := make([]FuncType, 0, count)
	for range count {
		f, err := readByte(r)
		if err != nil {
			return nil, fmt.Errorf("failed to read function type: %w", err)
		}
		if f != 0x60 {
			return nil, fmt.Errorf("unsupported function type: %x", f)
		}

		paramCount, err := leb128.Uint32(r)
		if err != nil {
			return nil, fmt.Errorf("failed to read parameter count: %w", err)
		}

		params := make([]ValueType, 0, paramCount)
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

		results := make([]ValueType, 0, resultCount)
		for range resultCount {
			v, err := decodeValueType(r)
			if err != nil {
				return nil, fmt.Errorf("failed to decode result type: %w", err)
			}
			results = append(results, v)
		}

		funcTypes = append(funcTypes, FuncType{params: params, results: results})
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

func decodeCodeSection(r io.Reader) ([]Function, error) {
	count, err := leb128.Uint32(r)
	if err != nil {
		return nil, fmt.Errorf("failed to read function count: %w", err)
	}

	functions := make([]Function, 0, count)
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

func decodeValueType(r io.Reader) (ValueType, error) {
	b, err := readByte(r)
	if err != nil {
		return 0, fmt.Errorf("failed to read value type: %w", err)
	}
	return ValueType(b), nil
}

func decodeFunctionBody(r io.Reader) (Function, error) {
	count, err := leb128.Uint32(r)
	if err != nil {
		return Function{}, fmt.Errorf("failed to read local count: %w", err)
	}

	locals := make([]FunctionLocal, 0, count)
	for range count {
		typeCount, err := leb128.Uint32(r)
		if err != nil {
			return Function{}, fmt.Errorf("failed to read type count: %w", err)
		}
		valueType, err := decodeValueType(r)
		if err != nil {
			return Function{}, fmt.Errorf("failed to decode value type: %w", err)
		}
		locals = append(locals, FunctionLocal{typeCount: typeCount, valueType: valueType})
	}

	instructions, err := decodeInstructions(r)
	if err != nil {
		return Function{}, fmt.Errorf("failed to decode instructions: %w", err)
	}

	return Function{locals: locals, code: instructions}, nil
}

func decodeInstructions(r io.Reader) ([]Instruction, error) {
	var (
		instructions []Instruction
	)
	for {
		opcode, err := readByte(r)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return nil, fmt.Errorf("failed to read opcode: %w", err)
		}

		switch Opcode(opcode) {
		case OpcodeEnd:
			instructions = append(instructions, InstructionEnd{})
		case OpcodeCall:
			i := new(InstructionCall)
			if err := i.ReadFrom(r); err != nil {
				return nil, fmt.Errorf("failed to read call instruction: %w", err)
			}
			instructions = append(instructions, i)
		case OpcodeLocalGet:
			i := new(InstructionLocalGet)
			if err := i.ReadFrom(r); err != nil {
				return nil, fmt.Errorf("failed to read local.get instruction: %w", err)
			}
			instructions = append(instructions, i)
		case OpcodeLocalSet:
			i := new(InstructionLocalSet)
			if err := i.ReadFrom(r); err != nil {
				return nil, fmt.Errorf("failed to read local.set instruction: %w", err)
			}
			instructions = append(instructions, i)
		case OpcodeI32Store:
			i := new(InstructionI32Store)
			if err := i.ReadFrom(r); err != nil {
				return nil, fmt.Errorf("failed to read i32.store instruction: %w", err)
			}
			instructions = append(instructions, i)
		case OpcodeI32Const:
			i := new(InstructionI32Const)
			if err := i.ReadFrom(r); err != nil {
				return nil, fmt.Errorf("failed to read i32.const instruction: %w", err)
			}
			instructions = append(instructions, i)
		case OpcodeI32Add:
			instructions = append(instructions, InstructionI32Add{})
		default:
			return nil, fmt.Errorf("unsupported opcode: %x", opcode)
		}
	}

	return instructions, nil
}

func decodeExportSection(r io.Reader) ([]Export, error) {
	count, err := leb128.Uint32(r)
	if err != nil {
		return nil, fmt.Errorf("failed to read export count: %w", err)
	}

	exports := make([]Export, 0, count)

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
			exports = append(exports, Export{name: name, desc: ExportDescFunc{index: index}})
		default:
			return nil, fmt.Errorf("unsupported export kind: %x", kind)
		}
	}

	return exports, nil
}

func decodeImportSection(r io.Reader) ([]Import, error) {
	count, err := leb128.Uint32(r)
	if err != nil {
		return nil, fmt.Errorf("failed to read import count: %w", err)
	}

	imports := make([]Import, 0, count)

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
			imports = append(imports, Import{module: module, field: name, desc: ImportDescFunc{index: index}})
		default:
			return nil, fmt.Errorf("unsupported import kind: %x", kind)
		}
	}

	return imports, nil
}

func decodeMemorySection(r io.Reader) ([]Memory, error) {
	count, err := leb128.Uint32(r)
	if err != nil {
		return nil, fmt.Errorf("failed to read memory count: %w", err)
	}

	memories := make([]Memory, 0, count)

	for range count {
		m := Memory{}
		m.limits, err = decodeLimits(r)
		if err != nil {
			return nil, fmt.Errorf("failed to decode memory limits: %w", err)
		}
		memories = append(memories, m)
	}

	return memories, nil
}

func decodeLimits(r io.Reader) (Limits, error) {
	hasMax, err := leb128.Uint32(r)
	if err != nil {
		return Limits{}, fmt.Errorf("failed to read hasMax: %w", err)
	}

	min, err := leb128.Uint32(r)
	if err != nil {
		return Limits{}, fmt.Errorf("failed to read min: %w", err)
	}

	if hasMax == 0 {
		return Limits{min: min, max: 0}, nil
	}

	max, err := leb128.Uint32(r)
	if err != nil {
		return Limits{}, fmt.Errorf("failed to read max: %w", err)
	}

	return Limits{min: min, max: max}, nil
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

func decodeDataSection(r io.Reader) ([]Data, error) {
	count, err := leb128.Uint32(r)
	if err != nil {
		return nil, fmt.Errorf("failed to read data count: %w", err)
	}

	data := make([]Data, 0, count)

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
		data = append(data, Data{memoryIndex: memoryIndex, offset: offset, init: init})
	}

	return data, nil
}
