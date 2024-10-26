package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:], " "), rand.Float64())
	fmt.Println(os.Args)
}
