// comma inserts comma in non-negative decimal integer string
package main

import (
	"bytes"
	"fmt"
	"unicode"
)

func main() {
	s := "1234567"

	fmt.Println(comma(s))
	fmt.Println(nonrecursive(s))

	e := "-12.554.22555" // => -1,2.55,4.22,555
	fmt.Println(enhancedcomma(e))
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

// exercise 3.10
// Write a non-recursive version of comma using bytes.Buffer
// instead of concatenation

func nonrecursive(s string) string {
	var buf bytes.Buffer
	var i int

	for i = 0; i < len(s)%3; i++ {
		buf.WriteByte(s[i])
	}
	for j := 0; j < len(s)-i; j++ {
		if j%3 == 0 {
			buf.WriteByte(',')
		}
		buf.WriteByte(s[j+i])
	}
	return buf.String()
}

func enhancedcomma(s string) string {
	var buf bytes.Buffer
	var digits int
	for _, c := range s {
		if unicode.IsDigit(c) {
			digits++
		}
	}

	var i, j int
	for i = 0; j < digits%3; i++ {
		buf.WriteByte(s[i])
		if unicode.IsDigit(rune(s[i])) {
			j++
		}
	}

	for j = 0; i < len(s); i++ {
		if j%3 == 0 {
			buf.WriteByte(',')
		}
		buf.WriteByte(s[i])
		if unicode.IsDigit(rune(s[i])) {
			j++
		}
	}
	return buf.String()
}
