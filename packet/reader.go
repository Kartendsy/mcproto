package packet

import (
	"encoding/binary"
	"errors"
	"math"
)

type Reader struct {
	buf []byte
	off int
	err error
}

func NewReader(data []byte) *Reader {
	return &Reader{buf: data}
}

func (r *Reader) VarInt() int32 {
	if r.err != nil {
		return 0
	}
	var val uint32
	var size int

	for {
		if r.off >= len(r.buf) {
			r.err = errors.New("eof")
			return 0
		}

		b := r.buf[r.off]
		r.off++
		val |= uint32(b&0x7F) << (size * 7)
		size++
		if b&0x80 == 0 {
			break
		}
	}

	return int32(val)
}

func (r *Reader) String() string {
	if r.err != nil {
		return ""
	}
	length := int(r.VarInt())
	if r.off+length > len(r.buf) {
		r.err = errors.New("string length overflow")
		return ""
	}

	res := string(r.buf[r.off : r.off+length])
	r.off += length
	return res
}

func (r *Reader) Byte() int8 {
	if r.err != nil {
		return 0
	}
	if r.off >= len(r.buf) {
		r.err = errors.New("eof byte")
		return 0
	}

	val := int8(r.buf[r.off])
	r.off++
	return val
}

func (r *Reader) Short() int16 {
	return int16(r.UShort())
}

func (r *Reader) Int() int32 {
	return int32(r.UInt())
}

func (r *Reader) Long() int64 {
	if r.err != nil {
		return 0
	}
	if r.off+8 > len(r.buf) {
		r.err = errors.New("eof long")
		return 0
	}

	val := binary.BigEndian.Uint64(r.buf[r.off : r.off+8])
	r.off += 8
	return int64(val)
}

func (r *Reader) UByte() uint8 {
	return uint8(r.Byte())
}

func (r *Reader) UShort() uint16 {
	if r.err != nil {
		return 0
	}
	if r.off+2 > len(r.buf) {
		r.err = errors.New("eof ushort")
		return 0
	}
	val := binary.BigEndian.Uint16(r.buf[r.off : r.off+2])
	r.off += 2
	return val
}

func (r *Reader) UInt() uint32 {
	if r.err != nil {
		return 0
	}
	if r.off+4 > len(r.buf) {
		r.err = errors.New("eof uint")
		return 0
	}
	val := binary.BigEndian.Uint32(r.buf[r.off : r.off+4])
	r.off += 4
	return val
}

func (r *Reader) Float() float32 {
	return math.Float32frombits(r.UInt())
}

func (r *Reader) Double() float64 {
	if r.err != nil {
		return 0
	}
	if r.off+8 > len(r.buf) {
		r.err = errors.New("eof double")
		return 0
	}

	val := binary.BigEndian.Uint64(r.buf[r.off : r.off+8])
	r.off += 8
	return math.Float64frombits(val)
}

func (r *Reader) Bytes(n int) []byte {
	if r.err != nil {
		return nil
	}
	if r.off+n > len(r.buf) {
		r.err = errors.New("eof bytes")
		return nil
	}
	val := r.buf[r.off : r.off+n]
	r.off += n
	return val
}

func (r *Reader) Offset() int {
	return r.off
}

func (r *Reader) Error() error {
	return r.err
}
