package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"net/http"
	"strconv"
)

const (
	height = 1024
	width  = 1024
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
	)
	ydif, err := toInt(r.URL.Query().Get("y"), 0)
	if err != nil {
		log.Println(err)
		http.Error(w, "y should be integer", http.StatusBadRequest)
		return
	}
	xdif, err := toInt(r.URL.Query().Get("x"), 0)
	if err != nil {
		log.Println(err)
		http.Error(w, "x should be integer", http.StatusBadRequest)
		return
	}
	zoom, err := toInt(r.URL.Query().Get("zoom"), 1)
	if err != nil {
		log.Println(err)
		http.Error(w, "zoom should be integer", http.StatusBadRequest)
		return
	}
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/float64(height)*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/float64(width)*(xmax-xmin) + xmin
			z := complex(x/float64(zoom), y/float64(zoom))
			// Image point (px, py) represents complex value z.
			img.Set(px-xdif*zoom, py+ydif*zoom, mandelbrot(z))
		}
	}
	png.Encode(w, img)
}

func mandelbrot(z complex128) color.Color {
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

func toInt(param string, d int) (int, error) {
	if len(param) == 0 {
		return d, nil
	}
	h, err := strconv.Atoi(param)
	if err != nil {
		return 0, err
	}
	return h, nil
}
