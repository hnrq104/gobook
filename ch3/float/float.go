package main

import (
	"fmt"
	"math"
)

const Avogrado = 6.02214129e23
const Planck = 6.62606957e-34

func main() {
	for x := 0; x < 8; x++ {
		fmt.Printf("x = %d e^x = %8f\n", x, math.Exp(float64(x)))
	}

	var z float64
	fmt.Println(z, -z, 1/z, -1/z, z/z) // 0 -0 +Inf -Inf NaN

	nan := math.NaN()
	fmt.Println(nan == nan, nan < nan, nan > nan) //false false false

}
