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
	tableSection    []binary.TableType
	globalSection   []binary.Global
	startSection    *uint32
}

func NewModule(r io.Reader) (*Module, error) {
	return decode(r)
}

func (m *Module) MemorySection() []binary.Memory   { return m.memorySection }
func (m *Module) DataSection() []binary.Data       { return m.dataSection }
func (m *Module) TypeSection() []binary.FuncType   { return m.typeSection }
func (m *Module) FunctionSection() []uint32        { return m.functionSection }
func (m *Module) TableSection() []binary.TableType { return m.tableSection }
func (m *Module) CodeSection() []binary.Function   { return m.codeSection }
func (m *Module) ExportSection() []binary.Export   { return m.exportSection }
func (m *Module) ImportSection() []binary.Import   { return m.importSection }
func (m *Module) GlobalSection() []binary.Global   { return m.globalSection }

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
		case SectionCodeTable:
			module.tableSection, err = decodeTableSection(sectionContents)
			if err != nil {
				return nil, fmt.Errorf("failed to decode table section: %w", err)
			}
		case SectionCodeMemory:
			module.memorySection, err = decodeMemorySection(sectionContents)
			if err != nil {
				return nil, fmt.Errorf("failed to decode memory section: %w", err)
			}
		case SectionCodeGlobal:
			module.globalSection, err = decodeGlobalSection(sectionContents)
			if err != nil {
				return nil, fmt.Errorf("failed to decode global section: %w", err)
			}
		case SectionCodeExport:
			module.exportSection, err = decodeExportSection(sectionContents)
			if err != nil {
				return nil, fmt.Errorf("failed to decode export section: %w", err)
			}
		case SectionCodeStart:
			module.startSection, err = decodeStartSection(sectionContents)
			if err != nil {
				return nil, fmt.Errorf("failed to decode start section: %w", err)
			}
		case SectionCodeElement:
			// TODO
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
		case SectionCodeDataCount:
			// TODO
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
		case 0x01:
			exports = append(exports, binary.Export{Name: name, Desc: binary.ExportDescTable{Index: index}})
		case 0x02:
			exports = append(exports, binary.Export{Name: name, Desc: binary.ExportDescMemory{Index: index}})
		case 0x03:
			exports = append(exports, binary.Export{Name: name, Desc: binary.ExportDescGlobal{Index: index}})
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

func decodeExprValue(r io.Reader) (binary.ExprValue, error) {
	b, err := readByte(r)
	if err != nil {
		return nil, fmt.Errorf("failed to read expr opcode: %w", err)
	}
	var value binary.ExprValue
	switch opcode.Opcode(b) {
	case opcode.OpcodeI32Const:
		v, err := leb128.Int32(r)
		if err != nil {
			return nil, fmt.Errorf("failed to read i32.const value: %w", err)
		}
		value = binary.ExprValueConstI32(v)
	case opcode.OpcodeI64Const:
		v, err := leb128.Int64(r)
		if err != nil {
			return nil, fmt.Errorf("failed to read i64.const value: %w", err)
		}
		value = binary.ExprValueConstI64(v)
	case opcode.OpcodeF32Const:
		v, err := readF32(r)
		if err != nil {
			return nil, fmt.Errorf("failed to read f32.const value: %w", err)
		}
		value = binary.ExprValueConstF32(v)
	case opcode.OpcodeF64Const:
		v, err := readF64(r)
		if err != nil {
			return nil, fmt.Errorf("failed to read f64.const value: %w", err)
		}
		value = binary.ExprValueConstF64(v)
	default:
		return nil, fmt.Errorf("unsupported expr opcode: %v", opcode.Opcode(b))
	}

	b, err = readByte(r)
	if err != nil {
		return nil, fmt.Errorf("failed to read expr end opcode: %w", err)
	}

	if opcode.Opcode(b) != opcode.OpcodeEnd {
		return nil, fmt.Errorf("invalid expr end opcode: %v", opcode.Opcode(b))
	}

	return value, nil
}

func decodeExpr(r io.Reader) (binary.Expr, error) {
	b, err := readByte(r)
	if err != nil {
		return nil, fmt.Errorf("failed to read expr opcode: %w", err)
	}
	var value binary.Expr
	switch opcode.Opcode(b) {
	case opcode.OpcodeI32Const:
		v, err := leb128.Int32(r)
		if err != nil {
			return nil, fmt.Errorf("failed to read i32.const value: %w", err)
		}
		value = binary.ExprValueConstI32(v)
	case opcode.OpcodeI64Const:
		v, err := leb128.Int64(r)
		if err != nil {
			return nil, fmt.Errorf("failed to read i64.const value: %w", err)
		}
		value = binary.ExprValueConstI64(v)
	case opcode.OpcodeF32Const:
		v, err := readF32(r)
		if err != nil {
			return nil, fmt.Errorf("failed to read f32.const value: %w", err)
		}
		value = binary.ExprValueConstF32(v)
	case opcode.OpcodeF64Const:
		v, err := readF64(r)
		if err != nil {
			return nil, fmt.Errorf("failed to read f64.const value: %w", err)
		}
		value = binary.ExprValueConstF64(v)
	case opcode.OpcodeGlobalGet:
		v, err := leb128.Uint32(r)
		if err != nil {
			return nil, fmt.Errorf("failed to read global index: %w", err)
		}
		value = binary.ExprGlobalIndex(v)
	default:
		return nil, fmt.Errorf("unsupported expr opcode: %s", opcode.Opcode(b))
	}

	b, err = readByte(r)
	if err != nil {
		return nil, fmt.Errorf("failed to read expr end opcode: %w", err)
	}

	if opcode.Opcode(b) != opcode.OpcodeEnd {
		return nil, fmt.Errorf("invalid expr end opcode: %s", opcode.Opcode(b))
	}

	return value, nil
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

func decodeTableSection(r io.Reader) ([]binary.TableType, error) {
	count, err := leb128.Uint32(r)
	if err != nil {
		return nil, fmt.Errorf("failed to read table count: %w", err)
	}

	tables := make([]binary.TableType, 0, count)

	for range count {
		typ, err := readByte(r)
		if err != nil {
			return nil, fmt.Errorf("failed to read element type: %w", err)
		}
		lim, err := decodeLimits(r)
		if err != nil {
			return nil, fmt.Errorf("failed to decode table limits: %w", err)
		}
		tables = append(tables, binary.TableType{ElementType: binary.RefType(typ), Limits: lim})
	}

	return tables, nil
}

func decodeGlobalSection(r io.Reader) ([]binary.Global, error) {
	count, err := leb128.Uint32(r)
	if err != nil {
		return nil, fmt.Errorf("failed to read global count: %w", err)
	}

	globals := make([]binary.Global, 0, count)

	for range count {
		globalType, err := decodeGlobalType(r)
		if err != nil {
			return nil, fmt.Errorf("failed to decode global type: %w", err)
		}

		initExpr, err := decodeExprValue(r)
		if err != nil {
			return nil, fmt.Errorf("failed to decode global init expr: %w", err)
		}

		globals = append(globals, binary.Global{
			Type:     globalType,
			InitExpr: initExpr,
		})
	}

	return globals, nil
}

func decodeGlobalType(r io.Reader) (binary.GlobalType, error) {
	typ, err := readByte(r)
	if err != nil {
		return binary.GlobalType{}, fmt.Errorf("failed to read global type: %w", err)
	}
	mut, err := readByte(r)
	if err != nil {
		return binary.GlobalType{}, fmt.Errorf("failed to read global mutability: %w", err)
	}
	return binary.GlobalType{ValueType: binary.ValueType(typ), Mutable: mut == 0x01}, nil
}

func decodeStartSection(r io.Reader) (*uint32, error) {
	idx, err := leb128.Uint32(r)
	if err != nil {
		return nil, fmt.Errorf("failed to read start index: %w", err)
	}
	return &idx, nil
}
