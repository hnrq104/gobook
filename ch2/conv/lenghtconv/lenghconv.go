// package lenghtconv performs Metre and Feet lenght computations.
package lenghtconv

import "fmt"

type Metre float64
type Feet float64

const (
	Kilometre Metre = 1000
	Milimetre Metre = 0.001
)

func (m Metre) String() string { return fmt.Sprintf("%gm", m) }
func (f Feet) String() string  { return fmt.Sprintf("%gf", f) }
