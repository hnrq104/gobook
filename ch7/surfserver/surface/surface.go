// Surface computes an SVG rendering of a 3-D surface function
package surface

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

func Surface(out io.Writer, f func(a, b float64) float64, c string) {
	fmt.Fprintf(out, "<svg xmlns = 'http://www.w3.org/2000/svg' "+
		"style = 'stroke: grey; fill: black; stroke-width: 0.7' "+
		"width = '%d' height = '%d'>", width, height)

	if c == "" {
		c = "#000000"
	}

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j, f)
			bx, by := corner(i, j, f)
			cx, cy := corner(i, j+1, f)
			dx, dy := corner(i+1, j+1, f)

			fmt.Fprintf(out, "<polygon points= \"%g,%g %g,%g %g,%g %g,%g\" "+
				"style= \"fill:%s\"/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, c)
		}
	}
	fmt.Fprintf(out, "</svg>\n")
}

func corner(i, j int, f func(a, b float64) float64) (float64, float64) {
	//find points(x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	//compute surface height z.
	z := f(x, y)
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale

	return sx, sy
}
