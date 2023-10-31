package main

import (
	"crypto/sha256"
	"fmt"
	"time"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))

	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	fmt.Println("different bits = ", popDiff([32]byte{0: 5}, [32]byte{0: 4}))

	start := time.Now()
	d := popDiff(c1, c2)
	fmt.Printf("d = %d, took : %vns\n", d, time.Since(start).Nanoseconds())
}

// 2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db02258717921a4881
// 4b68ab3847feda7d6c62c1fbcbeebfa35eab7351ed5e78f4ddadea5df64b8015
// false
// [32]uint8

// exercise 4.1 wirte a function that counts the number
// of bits that are different in two sha256 hashes

func popClear(a byte) int {
	var pop int
	for a != 0 {
		pop++
		a &= a - 1
	}
	return pop
}

func popDiff(a, b [32]byte) int {
	var pop int
	for i := 0; i < 32; i++ {
		pop += popClear(a[i] ^ b[i])
	}
	return pop
}
