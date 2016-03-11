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
	"math/big"
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
			//			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			//			img.Set(px, py, mandelbrot128(z))
			k := mandelbrotRat(big.NewRat(0, 1).SetFloat64(x), big.NewRat(0, 1).SetFloat64(y))
			fmt.Printf("%v,%v,%v\n", py, px, k)
			img.Set(px, py, k)
		}
	}
	fmt.Printf("Create\n")
	out, _ := os.OpenFile("out8.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer out.Close()
	png.Encode(out, img)
	fmt.Printf("Done\n")
}

func mandelbrotRat(r, i *big.Rat) color.Color {
	const iterations = 200
	const contrast = 15

	a := big.NewRat(0, 1) //実部の初期値
	b := big.NewRat(0, 1) //虚部の初期値
	for n := uint8(0); n < iterations; n++ {
		a = addR(subR(mulR(a, a), mulR(b, b)), r)
		b = addR(mulR(big.NewRat(2, 1), mulR(a, b)), i)
		if addR(mulR(a, a), mulR(b, b)).Cmp(big.NewRat(4, 1)) == 1 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

func mulR(a, b *big.Rat) *big.Rat {
	r := big.NewRat(0, 1)
	r.Mul(a, b)
	return r
}

func addR(a, b *big.Rat) *big.Rat {
	r := big.NewRat(0, 1)
	r.Add(a, b)
	return r
}

func subR(a, b *big.Rat) *big.Rat {
	r := big.NewRat(0, 1)
	r.Sub(a, b)
	return r
}

func mandelbrotFloat(r, i *big.Float) color.Color {
	const iterations = 200
	const contrast = 15

	a := big.NewFloat(0) //実部の初期値
	b := big.NewFloat(0) //虚部の初期値
	for n := uint8(0); n < iterations; n++ {
		a = addF(subF(mulF(a, a), mulF(b, b)), r)
		b = addF(mulF(big.NewFloat(2), mulF(a, b)), i)
		if addF(mulF(a, a), mulF(b, b)).Cmp(big.NewFloat(4)) == 1 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

func mulF(a, b *big.Float) *big.Float {
	r := big.NewFloat(0)
	r.Mul(a, b)
	return r
}

func addF(a, b *big.Float) *big.Float {
	r := big.NewFloat(0)
	r.Add(a, b)
	return r
}

func subF(a, b *big.Float) *big.Float {
	r := big.NewFloat(0)
	r.Sub(a, b)
	return r
}

func mandelbrot128(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

func mandelbrot64(z complex64) color.Color {
	return mandelbrot128(complex128(z))
}
