package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
)

var pallete = []color.Color{color.Black, color.RGBA{0, 0xff, 0, 0xff},
	color.RGBA{0xff, 0, 0, 0xff}, color.RGBA{0, 0, 0xff, 0xff}} //exercise 1.5

const (
	blackIndex = 0 //first color in pallete
	greenIndex = 1 //second color in pallete
	redIndex   = 2 // third
	blueIndex  = 3 //fourth
)

func parametrized_lj(out io.Writer, p map[string]int) {
	const (
		cycles  = 5     //number of complete x oscillator revolutions
		res     = 0.001 //angular resolution
		size    = 100   //image canvas [-size..+size]
		nframes = 64    //number of animation frames
		delay   = 8     //delay between frames in 10ms units
	)

	freq := rand.Float64() * 3.0 //relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 //phase difference

	cyc, s, n, d := cycles, size, nframes, delay

	if v, ok := p["size"]; ok {
		s = v
	}
	if v, ok := p["cycles"]; ok {
		cyc = v
	}
	if v, ok := p["nframes"]; ok {
		n = v
	}
	if v, ok := p["delay"]; ok {
		d = v
	}

	for i := 0; i < n; i++ {
		rect := image.Rect(0, 0, 2*s, 2*s)
		img := image.NewPaletted(rect, pallete)
		// var c uint8 = 1
		for t := 0.0; t < float64(cyc)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)

			//exercise 1.6
			c := color.RGBA{uint8(rand.Uint32()), uint8(rand.Uint32()), uint8(rand.Uint32()), 0xff}

			// img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),c)
			// c = c%3 + 1

			img.Set(s+int(x*float64(s)+0.5), s+int(float64(s)*y+0.5), c)

		}

		phase += 0.1
		anim.Delay = append(anim.Delay, d)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(out, &anim) //note : ignoring encoding errors
}

func lj(out io.Writer) {
	const (
		cycles  = 5     //number of complete x oscillator revolutions
		res     = 0.001 //angular resolution
		size    = 100   //image canvas [-size..+size]
		nframes = 64    //number of animation frames
		delay   = 8     //delay between frames in 10ms units
	)

	freq := rand.Float64() * 3.0 //relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 //phase difference

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, pallete)
		// var c uint8 = 1
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)

			//exercise 1.6
			c := color.RGBA{uint8(rand.Uint32()), uint8(rand.Uint32()), uint8(rand.Uint32()), 0xff}

			// img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),c)
			// c = c%3 + 1

			img.Set(size+int(x*size+0.5), size+int(y*size+0.5), c)

		}

		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(out, &anim) //note : ignoring encoding errors
}
