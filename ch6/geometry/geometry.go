package geometry

import "math"

type Point struct{ X, Y float64 }

// Traditional function
func Distance(p, q Point) float64 {
	return math.Hypot((p.X - q.X), (p.Y - q.Y))
}

// same thing, but as a method of Point type
func (p Point) Distance(q Point) float64 {
	return math.Hypot((p.X - q.X), (p.Y - q.Y))
}

// A Path is a journey connecting the points with straight lines
type Path []Point

func (path Path) Distance() float64 {
	sum := 0.
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}
