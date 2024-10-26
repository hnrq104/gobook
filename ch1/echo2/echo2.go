package main

import (
	"fmt"
	"os"
)

func main() {

	// s:= "" -> use this (only usable within functions)
	// var s string -> or this
	// var s = ""
	// var s string = ""

	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
