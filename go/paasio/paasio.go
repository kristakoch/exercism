package paasio

import (
	"io"
	"sync"
)

type ioWriteCounter struct {
	w            io.Writer
	bytesWritten int64
	numWrites    int
	mutex        sync.Mutex
}

type ioReadCounter struct {
	r         io.Reader
	bytesRead int64
	numReads  int
	mutex     sync.Mutex
}

type ioReadWriteCounter struct {
	ioWriteCounter
	ioReadCounter
}

// NewWriteCounter is a factory function for WriteCounters.
func NewWriteCounter(w io.Writer) WriteCounter {
	return &ioWriteCounter{w: w, mutex: sync.Mutex{}}
}

// NewReadCounter is a factory function for ReadCounters.
func NewReadCounter(r io.Reader) ReadCounter {
	return &ioReadCounter{r: r, mutex: sync.Mutex{}}
}

// NewReadWriteCounter is a factory function
// for ReadWriteCounters.
func NewReadWriteCounter(rw io.ReadWriter) ReadWriteCounter {
	return &ioReadWriteCounter{
		ioWriteCounter{w: rw},
		ioReadCounter{r: rw},
	}
}

// Read is a wrapper method for Read on ioReadCounter that
// records bytes read and number of read ops.
func (rc *ioReadCounter) Read(p []byte) (n int, err error) {
	rc.mutex.Lock()
	defer rc.mutex.Unlock()

	if n, err = rc.r.Read(p); nil != err {
		return n, err
	}

	rc.bytesRead += int64(n)
	rc.numReads++

	return n, err
}

// ReadCount returns the number of reads made
// by an ioReadCounter.
func (rc *ioReadCounter) ReadCount() (n int64, nops int) {
	return rc.bytesRead, rc.numReads
}

// Write is a wrapper method for Write ioWriteCounter that
// records bytes written and number of write ops.
func (wc *ioWriteCounter) Write(p []byte) (n int, err error) {
	wc.mutex.Lock()
	defer wc.mutex.Unlock()

	if n, err = wc.w.Write(p); nil != err {
		return n, err
	}

	wc.bytesWritten += int64(n)
	wc.numWrites++

	return n, err
}

// WriteCount returns the number of writes made
// by an ioWriteCounter.
func (wc *ioWriteCounter) WriteCount() (n int64, nops int) {
	return wc.bytesWritten, wc.numWrites
}
