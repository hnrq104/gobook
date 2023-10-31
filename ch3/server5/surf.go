// Surface computes an SVG rendering of a 3-D surface function
package main

import (
	"fmt"
	"io"
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

func ProduceSVG(out io.Writer, c string) {
	fmt.Fprintf(out, "<svg xmlns = 'http://www.w3.org/2000/svg' "+
		"style = 'stroke: grey; fill: black; stroke-width: 0.7' "+
		"width = '%d' height = '%d'>", width, height)

	if c == "" {
		c = "#000000"
	}

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)

			fmt.Fprintf(out, "<polygon points= \"%g,%g %g,%g %g,%g %g,%g\" "+
				"style= \"fill:%s\"/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, c)
		}
	}
	fmt.Fprintf(out, "</svg>\n")
}

func corner(i, j int) (float64, float64) {
	//find points(x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	//compute surface height z.
	z := f(x, y)
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale

	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from the origin
	return math.Sin(r) / r
}
