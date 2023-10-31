// package weightconv performs Pounds and Kilograms mass computations.
package weightconv

import "fmt"

type Kilo float64
type Pound float64

const (
	Gram Kilo = 0.001
)

func (k Kilo) String() string   { return fmt.Sprintf("%gKg", k) }
func (lb Pound) String() string { return fmt.Sprintf("%glb", lb) }
