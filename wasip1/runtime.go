package wasip1

import (
	"encoding/binary"
	"fmt"
	"os"

	runtime "github.com/Warashi/wasmium/runtime"
	tr "github.com/Warashi/wasmium/types/runtime"
)

type Runtime interface {
	AddImport(module string, name string, fn runtime.ImportFunc)
}

type WasiSnapshotPreview1 struct {
	fileTable []*os.File
}

func NewWasiPreview1() *WasiSnapshotPreview1 {
	return &WasiSnapshotPreview1{
		fileTable: []*os.File{
			os.Stdin,
			os.Stdout,
			os.Stderr,
		},
	}
}

func (w *WasiSnapshotPreview1) Register(runtime Runtime) {
	runtime.AddImport("wasi_snapshot_preview1", "fd_write", w.FdWrite)
}

func (w *WasiSnapshotPreview1) FdWrite(store *runtime.Store, args ...tr.Value) ([]tr.Value, error) {
	fd, ok := args[0].(tr.ValueI32)
	if !ok {
		return nil, fmt.Errorf("unexpected type: %T", args[0])
	}

	iovs, ok := args[1].(tr.ValueI32)
	if !ok {
		return nil, fmt.Errorf("unexpected type: %T", args[1])
	}

	iovsLen, ok := args[2].(tr.ValueI32)
	if !ok {
		return nil, fmt.Errorf("unexpected type: %T", args[2])
	}

	rp, ok := args[3].(tr.ValueI32)
	if !ok {
		return nil, fmt.Errorf("unexpected type: %T", args[3])
	}

	if fd < 0 || len(w.fileTable) <= int(fd) {
		return nil, fmt.Errorf("invalid file descriptor: %d", fd)
	}

	file := w.fileTable[fd]

	memory, err := store.Memory(0)
	read := func(addr tr.ValueI32) (int32, error) {
		var buf [4]byte
		if _, err := memory.ReadAt(buf[:], int64(addr)); err != nil {
			return 0, err
		}

		var v int32
		if _, err := binary.Decode(buf[:], binary.LittleEndian, &v); err != nil {
			return 0, err
		}

		return v, nil
	}

	if err != nil {
		return nil, fmt.Errorf("failed to get memory: %w", err)
	}
	nwritten := 0
	for range iovsLen {
		start, err := read(iovs)
		iovs += 4

		len, err := read(iovs)
		iovs += 4

		buf := make([]byte, len)

		if _, err := memory.ReadAt(buf, int64(start)); err != nil {
			return nil, fmt.Errorf("failed to read: %w", err)
		}

		n, err := file.Write(buf)
		if err != nil {
			return nil, fmt.Errorf("failed to write: %w", err)
		}
		nwritten += n
	}

	var buf [4]byte
	if _, err := binary.Encode(buf[:], binary.LittleEndian, int32(nwritten)); err != nil {
		return nil, fmt.Errorf("failed to write: %w", err)
	}

	if _, err := memory.WriteAt(buf[:], int64(rp)); err != nil {
		return nil, fmt.Errorf("failed to write: %w", err)
	}

	return []tr.Value{tr.ValueI32(0)}, nil
}
