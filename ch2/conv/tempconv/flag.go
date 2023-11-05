package tempconv

import (
	"flag"
	"fmt"
)

type celsiusFlag struct{ Celsius }

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit) // no error check needed
	switch unit {
	case "C", "Cº":
		f.Celsius = Celsius(value)
		return nil
	case "F", "Fº":
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	case "K", "Kº":
		f.Celsius = KToC(Kelvin(value))
		return nil
	}

	return fmt.Errorf("invalid temperature %q", s)
}

/*
Celsius flag defines a Celsius flag with specified name,
default value, and usage, and returns the address of the flag variable.
The flag argument must have a quantity and a unit, e.g., "100C"
*/
func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}
