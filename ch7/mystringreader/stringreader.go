package main

import (
	"fmt"
	"io"
)

type MyReader struct {
	s     string
	index int
}

func (mr *MyReader) Read(b []byte) (int, error) {
	if mr.index >= len(mr.s) {
		return 0, io.EOF
	}

	n := copy(b, mr.s[mr.index:])
	mr.index += n
	return n, nil
}

func NewMyReader(str string) *MyReader {
	return &MyReader{str, 0}
}

type LimitedReader struct {
	r     io.Reader
	max   int64
	index int64
}

func (lr *LimitedReader) Read(b []byte) (int, error) {
	if lr.index >= lr.max {
		return 0, io.EOF
	}

	n, err := lr.r.Read(b)
	if err != nil {
		return 0, err
	}

	lr.index += int64(n)
	return n, nil
}

func LimitReader(r io.Reader, max int64) io.Reader {
	return &LimitedReader{r, max, 0}
}

func main() {
	r := NewMyReader("henrique , como diria o profeta")
	b := make([]byte, 10)
	for {
		_, err := r.Read(b)
		if err != nil {
			break
		}
		fmt.Println(string(b), r.index)
	}
}
