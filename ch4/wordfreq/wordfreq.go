package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var file = flag.String("file", "", "Specify file to read")

func main() {
	flag.Parse()

	var buf *bufio.Scanner
	if *file != "" {
		f, err := os.Open(*file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "wordfreq: %v\n", err)
			os.Exit(1)
		}
		buf = bufio.NewScanner(f)
	} else {
		buf = bufio.NewScanner(os.Stdin)
	}

	buf.Split(bufio.ScanWords)

	words := make(map[string]int)

	for buf.Scan() {
		if err := buf.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "wordfreq: %v\n", err)
		}
		words[buf.Text()]++
	}

	fmt.Print("words\tfreq\n")
	for w, n := range words {
		fmt.Printf("%q\t%d\n", w, n)
	}

}
