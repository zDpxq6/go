package main

import "io"

// Reader is a io.Reader.
type Reader struct {
	// This implementation is inspired by strings.Reader
	s        string
	i        int64 // current reading index
	prevRune int   // index of previous rune; or < 0
}

func (r *Reader) Read(b []byte) (n int, err error) {
	if len(b) == 0 {
		return 0, nil
	}
	if r.i >= int64(len(r.s)) {
		return 0, io.EOF
	}
	r.prevRune = -1
	n = copy(b, r.s[r.i:])
	r.i += int64(n)
	return
}

//
func NewReader(s string) *Reader { return &Reader{s, 0, -1} }
