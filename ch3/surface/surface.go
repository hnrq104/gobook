// Surface computes an SVG rendering of a 3-D surface function

package main

import (
	"fmt"
	"image/color"
	"math"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30                  // axis range (-xyrange .. + xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        //pixels per z unit
	angle         = math.Pi / 6         // angle of x,y axes (=45º)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30º), cos(30º)

func main() {
	fmt.Printf("<svg xmlns = 'http://www.w3.org/2000/svg' "+
		"style = 'stroke: grey; fill: black; stroke-width: 0.7' "+
		"width = '%d' height = '%d'>", width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, ac, ak := corner(i+1, j)
			bx, by, _, bk := corner(i, j)
			cx, cy, _, ck := corner(i, j+1)
			dx, dy, _, dk := corner(i+1, j+1)

			if !(ak && bk && ck && dk) {
				continue
			}

			fmt.Printf("<polygon points= '%g,%g %g,%g %g,%g %g,%g' "+
				"style= 'fill:#%02x00%02x'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, ac.R, ac.B)
		}
	}
	fmt.Println("</svg>")
}

// Modified according to 5.6
func corner(i, j int) (sx, sy float64, col color.RGBA, ok bool) {
	//find points(x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	//compute surface height z.
	z := f(x, y)
	if math.IsInf(math.Abs(z), 1) || math.IsNaN(z) {
		return math.NaN(), math.NaN(), color.RGBA{0, 0, 0, 0}, false
	}

	var red uint8
	var blue uint8
	var c uint8 = uint8(math.Abs(z) * float64(255))
	if z < 0 {
		blue = c
	} else {
		red = c
	}

	sx = width/2 + (x-y)*cos30*xyscale
	sy = height/2 + (x+y)*sin30*xyscale - z*zscale
	col = color.RGBA{red, 0, blue, 0xff}
	ok = true
	return
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from the origin
	return math.Sin(r) / r
}

//testing other functions

func g(x, y float64) float64 { //saddle
	return (x*x - y*y) / 5
}

func h(x, y float64) float64 { // eggbox
	return (-math.Abs(math.Sin(x)+math.Cos(y)) / 10)

}
