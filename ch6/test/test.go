package main

import (
	"fmt"
	"gobook/ch6/geometry"
)

func main() {
	perim := geometry.Path{{X: 1, Y: 1}, {X: 5, Y: 1}, {X: 5, Y: 4}, {X: 1, Y: 1}}
	fmt.Println(perim.Distance())
	for i := range perim {
		perim[i].ScaleBy(10)
	}
	fmt.Println(perim.Distance())
}
