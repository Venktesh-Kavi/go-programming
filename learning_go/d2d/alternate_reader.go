package main

import (
	"fmt"
	"io"
)

type AlternateReader struct {
	buf    []byte
	done   bool
	offset int // offset is required to indicate how many bytes were read.
}

/*
AlternateReader Alternate Reader implements the reader interface.
Read method of a reader interface provides implementation of reading into a user provided buffer.
Since calls to methods io.ReadFull and we might want to pass the reader to other functions that expect a reader interface. Which can access the read method we would a pointer type.
*/
func NewAlternateReader(buf []byte) *AlternateReader {
	return &AlternateReader{
		buf: buf,
	}
}

// we need a pointer type here to indicate that we have completely read the buffer
func (a *AlternateReader) Read(p []byte) (int, error) {
	if a.done {
		return 0, io.EOF
	}

	remaining := len(a.buf) - a.offset
	copyLen := min(remaining, len(p))
	copy(p, a.buf[a.offset:a.offset+copyLen])
	a.offset += copyLen
	if a.offset == len(a.buf) {
		a.done = true
	}
	return a.offset, nil
}

func main() {
	nr := AlternateReader{
		buf: []byte("foo bar"),
	}
	newBuf := make([]byte, len(nr.buf))

	readIdx, err := io.ReadFull(&nr, newBuf) // read all the bytes to the length on the newBuffer
	fmt.Println("read till index: ", readIdx)

	if err != nil {
		fmt.Println(err)
	}
}
