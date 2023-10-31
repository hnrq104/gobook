package main

import "fmt"

func main() {
	//binary, uint and int
	fmt.Println("Printing binary numbers and operations:")
	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2

	fmt.Printf("%08b\n", x) // "00100010"
	fmt.Printf("%08b\n", y) // "00000110"
	fmt.Printf("%08b\n", x&y)
	fmt.Printf("%08b\n", x|y)
	fmt.Printf("%08b\n", x^y)
	fmt.Printf("%08b\n", x&^y)
	fmt.Printf("%08b\n", x<<1)
	fmt.Printf("%08b\n", x>>1)
	fmt.Println()

	//float
	fmt.Println("Printing floating point conversions and operations:")
	f := 1e100
	fmt.Printf("%g\n", f)
	fmt.Printf("%f\n", f)
	fmt.Printf("%X\n", int(f)) //depends on implementation
	fmt.Println()

	fmt.Println("Printing in different bases:")
	o := 0666
	fmt.Printf("%d %[1]o %#[1]o\n", o) // 438 666 0666
	x1 := int64(0xdeadbeef)
	fmt.Printf("%d %[1]x %#[1]x %#[1]X\n", x1)
	fmt.Println()

	//runes
	fmt.Println("Priting runes:")
	ascii := 'a'
	unicode := rune(0x56fd)
	newline := '\n'
	fmt.Printf("%d %[1]c %[1]q\n", ascii)
	fmt.Printf("%d %[1]c %[1]q\n", unicode)
	fmt.Printf("%d %[1]q\n", newline)
	fmt.Println()

	fmt.Println("Modular operations:")
	fmt.Printf("10%%3 = %d : -10%%3 = %d : 10%%-3 = %d\n", 10%3, -10%3, 10%-3) //keeps sign of dividend
	fmt.Println()

}
