package leb128

import "io"

func readByte(r io.Reader) (byte, error) {
	var b [1]byte
	_, err := r.Read(b[:])
	return b[0], err
}

func Int32(r io.Reader) (int32, error) {
	const size = 32
	var x int32
	var s uint
	var b byte
	for {
		var err error
		b, err = readByte(r)
		if err != nil {
			return 0, err
		}
		x |= int32(b&0x7f) << s
		s += 7
		if b&0x80 == 0 {
			break
		}
	}
	if s < size && b&0x40 != 0 {
		x |= (^0 << s)
	}
	return x, nil
}

func Int64(r io.Reader) (int64, error) {
	const size = 64
	var x int64
	var s uint
	var b byte
	for {
		var err error
		b, err = readByte(r)
		if err != nil {
			return 0, err
		}
		x |= int64(b&0x7f) << s
		s += 7
		if b&0x80 == 0 {
			break
		}
	}
	if s < size && b&0x40 != 0 {
		x |= (^0 << s)
	}
	return x, nil
}

func Uint32(r io.Reader) (uint32, error) {
	var x uint32
	var s uint
	for {
		b, err := readByte(r)
		if err != nil {
			return 0, err
		}
		x |= uint32(b&0x7f) << s
		if b&0x80 == 0 {
			break
		}
		s += 7
	}
	return x, nil
}

func Uint64(r io.Reader) (uint64, error) {
	var x uint64
	var s uint
	for {
		b, err := readByte(r)
		if err != nil {
			return 0, err
		}
		x |= uint64(b&0x7f) << s
		if b&0x80 == 0 {
			break
		}
		s += 7
	}
	return x, nil
}
