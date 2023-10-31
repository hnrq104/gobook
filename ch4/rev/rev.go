package main

import (
	"fmt"
	"unicode/utf8"
)

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func reverseb(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// exercise 4.3
func reverse_array(s *[6]int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// exercise 4.4
func rotate(s []int, r int) {
	var cop []int = make([]int, len(s))
	copy(cop, s)
	for i := 0; i < len(s); i++ {
		n := (i + r) % len(s)
		if n < 0 {
			n = len(s) + n
		}
		s[i] = cop[n]
	}
}

// as utf-8 characters have different sizes, I cant do it in place
// exercise 4.7 TRASH
func reverseutf8(bytes []byte) {
	cop := make([]byte, 0, len(bytes))
	for i := 0; i < len(bytes); {
		_, size := utf8.DecodeRune(bytes[i:])
		cop = append(cop, bytes[i:i+size]...)
		i += size
	}

	for i, j := 0, len(cop); j > 0 && i < len(cop); {
		_, size := utf8.DecodeLastRune(cop[:j])
		copy(bytes[i:i+size], cop[j-size:j])
		i += size
		j -= size
	}
}

func UTFreverse(b []byte) {
	for i := 0; i < len(b); {
		_, s := utf8.DecodeRune(b[i:])
		reverseb(b[i : i+s])
		i += s
	}
	reverseb(b)
}

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5}
	reverse(a[:])
	fmt.Println(a)

	s := []int{0, 1, 2, 3, 4, 5}
	reverse(s[:2])
	reverse(s[2:])
	reverse(s)
	fmt.Println(s)

	b := [...]int{0, 1, 2, 3, 4, 5}
	reverse_array(&b)
	fmt.Println(b)

	rotate(b[:], -2)
	fmt.Println(b)

	t := []byte("abcdefgh\u4e16")
	fmt.Println(t)
	reverseutf8(t)
	fmt.Println(string(t))

	UTFreverse(t)
	fmt.Println(string(t))

}
