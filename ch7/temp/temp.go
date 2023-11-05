package main

import (
	"flag"
	"fmt"
	"gobook/ch2/conv/tempconv"
)

var c = tempconv.CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*c)
}
