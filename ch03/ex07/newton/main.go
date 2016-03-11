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
			r1, g1, b1, a1 := newton(complex(x1, y1)).RGBA()
			r2, g2, b2, a2 := newton(complex(x1, y2)).RGBA()
			r3, g3, b3, a3 := newton(complex(x2, y1)).RGBA()
			r4, g4, b4, a4 := newton(complex(x2, y2)).RGBA()
			c := color.RGBA{
				uint8(r1+r2+r3+r4) / 4,
				uint8(g1+g2+g3+g4) / 4,
				uint8(b1+b2+b3+b4) / 4,
				uint8(a1+a2+a3+a4) / 4,
			}
			img.Set(px, py, c)
		}
	}
	out, _ := os.OpenFile("out7.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer out.Close()
	png.Encode(out, img)
	fmt.Println("** ** ** ** DONE ** ** ** **")
}

func newton(z complex128) color.Color {
	const iterations = 200
	const contrast = 15
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z-1) < 0.01 {
			return color.Gray{255 - contrast*i}
		}
	}
	return color.Black
}
