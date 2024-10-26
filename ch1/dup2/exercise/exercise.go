// Dup1 prints the text of each line that appears more than
// once in the input. I reads from stdin or from a list of named files.

package main

//test

import (
	"bufio"
	"fmt"
	"os"
)

func countLines(f *os.File, arg string, counts map[string]int, files map[string]map[string]bool) {
	input := bufio.NewScanner(f)

	for input.Scan() {
		counts[input.Text()]++
		if files[input.Text()] == nil {
			files[input.Text()] = make(map[string]bool)
		}
		files[input.Text()][arg] = true
	}
	//note ignoring potential erros from input.err()
}

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	filenames := make(map[string]map[string]bool)

	if len(files) == 0 {
		countLines(os.Stdin, "stdin", counts, filenames)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2 new: %v\n", err)
			}
			countLines(f, arg, counts, filenames)
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t", n, line)
			s, sep := "", ""
			for file := range filenames[line] {
				s += sep + file
				sep = " "
			}
			fmt.Println(s)
		}
	}
}
