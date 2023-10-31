package main

import (
	"fmt"
	"gobook/ch2/popcount"
	"os"
	"strconv"
)

func main() {
	for _, args := range os.Args[1:] {
		x, err := strconv.ParseUint(args, 10, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "testpopcount: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Pop %d: %d %d %d %d\n", x,
			popcount.PopCount(x), popcount.PopCountLoop(x),
			popcount.PopCountShift(x), popcount.PopCountClears(x))
	}
}
