// © 2016 zDpxq6
// Code by Alan A. A. Donovan & Brian W. Kernighan/Adapted.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	out, _ := os.OpenFile("out5.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer out.Close()
	png.Encode(out, img)
	fmt.Println("** ** ** ** DONE ** ** ** **")
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			i := 255 - contrast*n
			switch {
			case n%8 == 0:
				return color.NRGBA{0, 0, 0, i}
			case n%8 == 1:
				return color.NRGBA{0, 0, i, i}
			case n%8 == 2:
				return color.NRGBA{0, i, 0, i}
			case n%8 == 3:
				return color.NRGBA{0, i, i, i}
			case n%8 == 4:
				return color.NRGBA{i, 0, 0, i}
			case n%8 == 5:
				return color.NRGBA{i, 0, i, i}
			case n%8 == 6:
				return color.NRGBA{i, i, 0, i}
			case n%8 == 7:
				return color.NRGBA{i, i, i, i}
			}
		}
	}
	return color.Black
}