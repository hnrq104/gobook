package main

import (
	"flag"
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"os"
	"sync"
	"time"
)

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const constrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{0xff - constrast*n}
		}
	}

	return color.Black
}

// not parallell
func mandelImage() image.Image {
	const width, height = 1024, 1024
	const xmin, ymin, xmax, ymax = -2, -2, 2, 2

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/height*(xmax-xmin) + xmin
			z := complex(x, y)
			c := mandelbrot(z)
			img.Set(px, py, c)
		}
	}

	return img
}

type pixel struct {
	image.Point
	c color.Color
}

func veryParallelMandelImage() image.Image {
	const width, height = 1024, 1024
	const xmin, ymin, xmax, ymax = -2, -2, 2, 2

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	// ch := make(chan pixel)

	var wg sync.WaitGroup
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/height*(xmax-xmin) + xmin
			z := complex(x, y)

			wg.Add(1)
			go func(px, py int, z complex128) {

				defer wg.Done()
				// ch <- pixel{image.Point{px, py}, mandelbrot(z)}
				img.Set(px, py, mandelbrot(z))
			}(px, py, z)
		}
	}

	//Closer
	// go func() {
	wg.Wait()
	// close(ch)
	// }()

	// for p := range ch {
	// 	img.Set(p.X, p.Y, p.c)
	// }

	return img
}

func okParallelMandelImage() image.Image {
	const width, height = 1024, 1024
	const xmin, ymin, xmax, ymax = -2, -2, 2, 2

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	// ch := make(chan pixel)

	var wg sync.WaitGroup
	for py := 0; py < height; py++ {
		wg.Add(1)
		go func(py int) {
			defer wg.Done()

			y := float64(py)/height*(ymax-ymin) + ymin
			for px := 0; px < width; px++ {
				x := float64(px)/height*(xmax-xmin) + xmin
				z := complex(x, y)
				// ch <- pixel{image.Point{px, py}, mandelbrot(z)}
				img.Set(px, py, mandelbrot(z))
			}

		}(py)
	}

	//Closer
	// go func() {
	wg.Wait()
	// 	close(ch)
	// }()

	// for p := range ch {
	// 	img.Set(p.X, p.Y, p.c)
	// }

	return img
}

// processors has to be a power of 2 less than 1024
func fewProcessorsMandelbrot(processors uint) image.Image {
	const width, height = 1024, 1024
	const xmin, ymin, xmax, ymax = -2, -2, 2, 2

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	// ch := make(chan pixel)

	var wg sync.WaitGroup
	for p := uint(0); p < processors; p++ {
		wg.Add(1)
		go func(p int) {
			defer wg.Done()
			min := p * (height / int(processors))
			max := (p + 1) * (height / int(processors))
			for py := min; py < max; py++ {
				y := float64(py)/height*(ymax-ymin) + ymin
				for px := 0; px < width; px++ {
					x := float64(px)/width*(xmax-xmin) + xmin

					z := complex(x, y)
					// ch <- pixel{image.Point{px, py}, mandelbrot(z)}
					img.Set(px, py, mandelbrot(z))
				}
			}

		}(int(p))

	}

	wg.Wait()

	// for p := range ch {
	// 	img.Set(p.X, p.Y, p.c)
	// }

	return img
}

// I could have channel communication with two counters I guess, one for setting and one for calculating
// but this way seems faster

var p = flag.Uint("p", 1, "number of processors, should be power of 2")
var s = flag.Bool("s", false, "save image to stdout")

func main() {
	flag.Parse()
	start := time.Now()
	img := fewProcessorsMandelbrot(*p)
	// fewProcessorsMandelbrot(*p) //really fun 1:150ms, 2: 75ms, ... 64:22ms
	// okParallelMandelImage() 22 ms
	// mandelImage() 150 ms
	// veryParallelMandelImage() 326 ms
	if !*s {
		log.Print(time.Since(start))
	} else {
		png.Encode(os.Stdout, img)
	}
}
