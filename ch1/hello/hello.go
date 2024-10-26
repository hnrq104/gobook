package main

import (
	"fmt"
	"math/rand"
)

func main() {
	n := rand.Int()
	fmt.Println("Hello Henrique!! :)")
	fmt.Printf("%d %x %o %b\n", n, n, n, n)
}
