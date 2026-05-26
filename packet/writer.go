package packet

import (
	"encoding/binary"
	"fmt"
	"math"
)

type Writer struct {
	buf []byte
}

func NewWriter() *Writer {
	return &Writer{buf: make([]byte, 0, 32)}
}

func (w *Writer) Reset() *Writer {
	w.buf = w.buf[:0]
	return w
}

// Variable Length

func (w *Writer) VarInt(v int32) *Writer {
	u := uint32(v)
	for u >= 0x80 {
		w.buf = append(w.buf, byte(u&0x7F|0x80))
		u >>= 7
	}
	w.buf = append(w.buf, byte(u))
	return w
}

func (w *Writer) String(v string) *Writer {
	w.VarInt(int32(len(v)))
	w.buf = append(w.buf, v...)
	return w
}

// Fixed Lenght Signed

func (w *Writer) Byte(v byte) *Writer {
	w.buf = append(w.buf, v)
	return w
}

func (w *Writer) Short(v int16) *Writer {
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, uint16(v))
	w.buf = append(w.buf, b...)
	return w
}

func (w *Writer) Int(v int32) *Writer {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, uint32(v))
	w.buf = append(w.buf, b...)
	return w
}

func (w *Writer) Long(v int64) *Writer {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	w.buf = append(w.buf, b...)
	return w
}

// Fixed Length Unsigned

func (w *Writer) UByte(v uint8) *Writer {
	w.buf = append(w.buf, v)
	return w
}

func (w *Writer) UShort(v uint16) *Writer {
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, v)
	w.buf = append(w.buf, b...)
	return w
}

func (w *Writer) UInt(v uint32) *Writer {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, v)
	w.buf = append(w.buf, b...)
	return w
}

// Floating Point

func (w *Writer) Float(v float32) *Writer {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, math.Float32bits(v))
	w.buf = append(w.buf, b...)
	return w
}

func (w *Writer) Double(v float64) *Writer {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, math.Float64bits(v))
	w.buf = append(w.buf, b...)
	return w
}

func (w *Writer) Bytes(v []byte) *Writer {
	w.buf = append(w.buf, v...)
	return w
}

func (w *Writer) Build(packetID int32) []byte {
	bodyWriter := NewWriter().VarInt(packetID).Bytes(w.buf)
	bodyLen := int32(len(bodyWriter.buf))

	finalWriter := NewWriter()
	finalWriter.VarInt(bodyLen)
	finalWriter.Bytes(bodyWriter.buf)

	return finalWriter.buf
}

func (w *Writer) Raw() []byte {
	return w.buf
}

func (w *Writer) Hex() string {
	return fmt.Sprintf("%x", w.buf)
}
