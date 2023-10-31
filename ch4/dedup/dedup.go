package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	buf := bufio.NewScanner(os.Stdin)
	seen := make(map[string]bool)

	for buf.Scan() {
		line := buf.Text()
		if !seen[line] {
			seen[line] = true
			fmt.Println(line)
		}

		if err := buf.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
			os.Exit(1)
		}
	}

}
