package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
)

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	buf := bufio.NewScanner(bytes.NewReader(p))
	buf.Split(bufio.ScanWords)
	var nbytes int
	for buf.Scan() {
		nbytes += len(buf.Bytes())
		*c += 1
	}
	return nbytes, nil
}

type LineCounter int

func (c *LineCounter) Write(p []byte) (int, error) {
	buf := bufio.NewScanner(bytes.NewReader(p))
	var nbytes int
	for buf.Scan() {
		nbytes += len(buf.Bytes())
		*c += 1
	}
	return nbytes, nil
}

type CountingW struct {
	written    int64
	trueWriter io.Writer
}

func (c *CountingW) Write(p []byte) (int, error) {
	n, err := c.trueWriter.Write(p)
	c.written += int64(n)
	return n, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	c := CountingW{written: 0, trueWriter: w}
	return &c, &c.written

}

func main() {
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c)

	c = 0
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c)

	var w WordCounter
	fmt.Fprintf(&w, "primeira segunda terceira palavra %s\n sexta", name)
	fmt.Println(w)

	var l LineCounter
	fmt.Fprintf(&l, "primeira segunda terceira palavra %s\n sexta", name)
	fmt.Println(l)

}
