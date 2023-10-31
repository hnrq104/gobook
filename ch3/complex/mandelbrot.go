package main

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"math/cmplx"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	mandelbrotPng(w)
}

func mandelbrotPng(out io.Writer) {
	const (
		xmin, xmax, ymin, ymax = -2, 2, -2, 2
		width, height          = 1024, 1024
	)

	var colormap [width + 1][height + 1]color.Color

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height+1; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width+1; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			colormap[px][py] = mandelbrot(z)
			// img.Set(px, py, c)
		}
	}

	for py := 0; py < height; py++ {
		for px := 0; px < width; px++ {
			r1, g1, b1, _ := colormap[px][py].RGBA()
			r2, g2, b2, _ := colormap[px][py+1].RGBA()
			r3, g3, b3, _ := colormap[px+1][py].RGBA()
			r4, g4, b4, _ := colormap[px+1][py+1].RGBA()

			r := (r1 + r2 + r3 + r4) / 4
			g := (g1 + g2 + g3 + g4) / 4
			b := (b1 + b2 + b3 + b4) / 4

			c := color.RGBA{uint8(r), uint8(g), uint8(b), 0xff}
			img.Set(px, py, c)

		}
	}

	png.Encode(out, img)
}

func mandelbrot(z complex128) color.Color {
	const (
		iterations = 200
		contrast   = 15
	)
	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.RGBA{0, 0,
				uint8(n * contrast), 0xff}
		}
	}
	return color.Black
}
