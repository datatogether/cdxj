package cdxj

import (
	"bufio"
	"bytes"
)

// A Reader reads records from a CSV-encoded file.
//
// As returned by NewReader, a Reader expects input conforming to RFC 4180.
// The exported fields can be changed to customize the details before the
// first call to Read or ReadAll.
//
//
type Reader struct {
	// record counter
	record int
	s      *bufio.Scanner
}

// NewReader returns a new Reader that reads from r.
func NewReader(r io.Reader) *Reader {
	return Reader{
		s: bufio.NewScanner(r),
	}
}

// Read reads a record from the reader
func (r *Reader) Read() (*Record, error) {
	rec := &Record{}
	// scan until we have a non-header record
	for {
		r.s.Scan()
		if bytes.HasPrefix(r.s.Bytes(), []byte("!")) {
			continue
		}
		break
	}
	if err := rec.UnmarshalCDXJ(r.s.Bytes()); err != nil {
		return nil, err
	}

	r.record++
	return rec, nil
}
