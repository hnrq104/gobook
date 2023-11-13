package mythumbnail

import (
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func Image(src image.Image) image.Image {
	xs := src.Bounds().Size().X
	ys := src.Bounds().Size().Y

	width, height := 256, 256

	if aspect := xs / ys; aspect < 1.0 {
		width = int(width * aspect)
	} else {
		height = int(height / aspect)
	}

	xscale := float64(xs) / float64(width)
	yscale := float64(ys) / float64(height)

	dst := image.NewRGBA(image.Rect(0, 0, width, height))

	//very slow
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			dst.Set(x, y, src.At(xs*int(xscale), ys*int(yscale)))
		}
	}

	return dst
}

func ImageStream(w io.Writer, r io.Reader) error {
	src, _, err := image.Decode(r)
	if err != nil {
		return err
	}

	dst := Image(src)
	return jpeg.Encode(w, dst, nil)
}

// imagefile2 reads an image from infile and writes
// the thumbnail scaled version to outfile.

func ImageFile2(outfile, infile string) (err error) {
	in, err := os.Open(infile)
	if err != nil {
		return err
	}

	defer in.Close()

	out, err := os.Create(outfile)
	if err != nil {
		return err
	}

	if err := ImageStream(out, in); err != nil {
		out.Close()
		fmt.Printf("scaling %s to %s: %s", infile, outfile, err)
		return err
	}

	return out.Close()
}

func ImageFile(infile string) (string, error) {
	ext := filepath.Ext(infile)
	outfile := strings.TrimSuffix(infile, ext) + ".thumb" + ext
	return outfile, ImageFile2(outfile, infile)
}
