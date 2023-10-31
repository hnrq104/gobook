package main

import (
	"fmt"
	"strings"
)

func expand(s string, f func(string) string) string {
	return strings.Join(strings.Split(s, "$foo"), f("foo"))
}

func naju(s string) string {
	return "naju"
}

func main() {
	s := "$foo oi$foo$foo henrique 1$foo$foo$foo2$foo 3$foo"
	fmt.Println(expand(s, naju))
}
