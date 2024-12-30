package execution

import (
	"encoding/binary"
	"fmt"
	"os"
)

type WasiSnapshotPreview1 struct {
	fileTable []*os.File
}

func NewWasiSnapshotPreview1() *WasiSnapshotPreview1 {
	return &WasiSnapshotPreview1{
		fileTable: []*os.File{
			os.Stdin,
			os.Stdout,
			os.Stderr,
		},
	}
}

func (w *WasiSnapshotPreview1) invoke(store *Store, fn string, args ...Value) ([]Value, error) {
	switch fn {
	case "fd_write":
		return w.fdWrite(store, args...)
	default:
		return nil, fmt.Errorf("unknown function: %s", fn)
	}
}

func (w *WasiSnapshotPreview1) fdWrite(store *Store, args ...Value) ([]Value, error) {
	fd, ok := args[0].(ValueI32)
	if !ok {
		return nil, fmt.Errorf("unexpected type: %T", args[0])
	}

	iovs, ok := args[1].(ValueI32)
	if !ok {
		return nil, fmt.Errorf("unexpected type: %T", args[1])
	}

	iovsLen, ok := args[2].(ValueI32)
	if !ok {
		return nil, fmt.Errorf("unexpected type: %T", args[2])
	}

	rp, ok := args[3].(ValueI32)
	if !ok {
		return nil, fmt.Errorf("unexpected type: %T", args[3])
	}

	if fd < 0 || len(w.fileTable) <= int(fd) {
		return nil, fmt.Errorf("invalid file descriptor: %d", fd)
	}

	file := w.fileTable[fd]

	nwritten := 0
	for range iovsLen {
		start, err := memory_read(store.memories[0].data, iovs)
		iovs += 4

		len, err := memory_read(store.memories[0].data, iovs)
		iovs += 4

		end := start + len
		n, err := file.Write(store.memories[0].data[start:end])
		if err != nil {
			return nil, fmt.Errorf("failed to write: %w", err)
		}
		nwritten += n
	}

	if err := memory_write(store.memories[0].data, rp, int32(nwritten)); err != nil {
		return nil, fmt.Errorf("failed to write: %w", err)
	}

	return []Value{ValueI32(0)}, nil
}

type ints interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

func memory_read[T ints](buf []byte, start T) (int32, error) {
	end := start + 4

	var v int32
	_, err := binary.Decode(buf[start:end], binary.LittleEndian, &v)

	return v, err
}

func memory_write[T ints](buf []byte, start T, v int32) error {
	end := start + 4

	_, err := binary.Encode(buf[start:end], binary.LittleEndian, int32(v))

	return err
}
