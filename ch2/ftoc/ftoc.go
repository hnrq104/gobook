package main

import "fmt"

func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%gºF = %gºC\n", freezingF, fToC(freezingF)) //32ºF = 0ºC
	fmt.Printf("%gºF = %gºC\n", boilingF, fToC(boilingF))   //212ºF = 100ºC
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
