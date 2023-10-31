package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

// exercise 4.6
// squases unicode spaces into ascii spaces
func squash(s []byte) []byte {
	out := s[:0]
	var last rune
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRune(s[i:])

		if !unicode.IsSpace(r) {
			out = append(out, s[i:i+size]...)
		} else if unicode.IsSpace(r) && !unicode.IsSpace(last) {
			out = append(out, ' ')
		}
		last = r
		i += size
	}
	return out
}

func main() {
	s := []byte("henrique     esteve   aqui\n")
	s = squash(s)
	fmt.Println(string(s))
}
