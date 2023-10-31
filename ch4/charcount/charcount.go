// charcount computes counts of Unicode characters
package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int)    // counts of Unicode characters.
	var utflen [utf8.UTFMax + 1]int // count lenghts of UTF-8 characters.
	invalid := 0                    //count of invalid UTF-8 characters.

	in := bufio.NewReader(os.Stdin)

	c := make(map[string]int)

	for {
		r, n, err := in.ReadRune()

		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			break
		}

		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}

		for cat, rt := range unicode.Categories {
			if unicode.Is(rt, r) {
				c[cat]++
			}
		}

		counts[r]++
		utflen[n]++
	}

	fmt.Print("rune\tcount\n")
	for c, v := range counts {
		fmt.Printf("%q\t%d\n", c, v)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		fmt.Printf("%d\t%d\n", i, n)
	}
	fmt.Print("\ncat\tcount\n")
	for i, n := range c {
		if n > 0 {
			fmt.Printf("%q\t%d\n", i, n)
		}
	}

	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}

}
