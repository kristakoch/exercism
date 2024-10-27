package circular

import (
	"errors"
)

var (
	errBufferEmpty = errors.New("circular buffer is empty")
	errBufferFull  = errors.New("circular buffer is full")
)

// Buffer is a type for circular buffers.
type Buffer struct {
	buf    []byte
	oldest int
	newest int
	items  int
}

// NewBuffer creates and returns a new buffer.
func NewBuffer(size int) *Buffer {
	var b Buffer
	b.buf = make([]byte, size)
	b.newest = -1
	return &b
}

// ReadByte reads from the buffer and returns an error
// if there's nothing to read.
func (b *Buffer) ReadByte() (byte, error) {
	if b.items == 0 {
		return 0, errBufferEmpty
	}

	ret := b.buf[b.oldest]
	b.buf[b.oldest] = 0
	b.items--

	b.oldest = (b.oldest + 1) % cap(b.buf)

	return ret, nil
}

// WriteByte performs a write to the buffer, returning
// an error if the buffer is full.
func (b *Buffer) WriteByte(c byte) error {
	if b.items == cap(b.buf) {
		return errBufferFull
	}

	b.newest = (b.newest + 1) % cap(b.buf)
	b.buf[b.newest] = c
	b.items++

	return nil
}

// Overwrite performs a write to the buffer, overwriting
// if the buffer is full.
func (b *Buffer) Overwrite(c byte) {
	if b.items < cap(b.buf) {
		b.WriteByte(c)
		return
	}

	b.buf[b.oldest] = c
	b.oldest = (b.oldest + 1) % cap(b.buf)
}

// Reset clears the buffer to its original state.
func (b *Buffer) Reset() {
	b.buf = make([]byte, cap(b.buf))
	b.newest = -1
	b.items = 0
	b.oldest = 0
}
