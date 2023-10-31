// Cf2 reads from Stdin or Argument list and performs type conversions.
package main

import (
	"bufio"
	"fmt"
	"gobook/ch2/conv/lenghtconv"
	"gobook/ch2/conv/tempconv"
	"gobook/ch2/conv/weightconv"
	"os"
	"strconv"
)

func main() {
	// If arguments were passed
	if len(os.Args) > 1 {
		for _, args := range os.Args[1:] {
			v, err := strconv.ParseFloat(args, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "cf2: %v\n", err)
				continue
			}
			printConversions(v)
		}

		return
	}

	// Read from Stdin
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		v, err := strconv.ParseFloat(input.Text(), 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf2: %v\n", err)
			continue
		}
		printConversions(v)
	}

}

func printConversions(v float64) {
	//Temperature
	c := tempconv.Celsius(v)
	f := tempconv.Fahrenheit(v)
	fmt.Printf("Temperature: %s = %s, %s = %s\n",
		c, tempconv.CToF(c), f, tempconv.FToC(f))

	//Lenght
	m := lenghtconv.Metre(v)
	fe := lenghtconv.Feet(v)
	fmt.Printf("Lenght: %s = %s, %s = %s\n",
		m, lenghtconv.MToF(m), fe, lenghtconv.FToM(fe))

	//Weight
	k := weightconv.Kilo(v)
	lb := weightconv.Pound(v)
	fmt.Printf("Weight: %s = %s, %s = %s\n",
		k, weightconv.KgToLb(k), lb, weightconv.LbToKg(lb))
}
