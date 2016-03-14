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
	for py := 0; py < height*2; py++ {
		y1 := float64(py)/height*(ymax-ymin) + ymin
		y2 := float64(py+1)/height*(ymax-ymin) + ymin
		for px := 0; px < width*2; px++ {
			x1 := float64(px)/width*(xmax-xmin) + xmin
			x2 := float64(px+1)/width*(xmax-xmin) + xmin
			r1, g1, b1, a1 := mandelbrot(complex(x1, y1)).RGBA()
			r2, g2, b2, a2 := mandelbrot(complex(x1, y2)).RGBA()
			r3, g3, b3, a3 := mandelbrot(complex(x2, y1)).RGBA()
			r4, g4, b4, a4 := mandelbrot(complex(x2, y2)).RGBA()
			c := color.RGBA{
				uint8(r1+r2+r3+r4) / 4,
				uint8(g1+g2+g3+g4) / 4,
				uint8(b1+b2+b3+b4) / 4,
				uint8(a1+a2+a3+a4) / 4,
			}
			img.Set(px, py, c)
		}
	}
	out, _ := os.OpenFile("out6.png", os.O_WRONLY|os.O_CREATE, 0600)
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
