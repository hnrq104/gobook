package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

var b = flag.Uint64("b", 256, "set amount of bits for hashing: 384, 512  and defaut 256")

func main() {
	flag.Parse()

	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		switch *b {
		case 384:
			fmt.Printf("%x\n", sha512.Sum384([]byte(input.Text())))
		case 512:
			fmt.Printf("%x\n", sha512.Sum512([]byte(input.Text())))
		default:
			fmt.Printf("%x\n", sha256.Sum256([]byte(input.Text())))
		}

	}

}
