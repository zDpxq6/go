// © 2016 zDpxq6
// Code by Alan A. A. Donovan & Brian W. Kernighan/Adapted.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Server3 is an "echo" server that displays request parameters.
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strings"
	"sync"
)

const (
	ff uint8 = 255
)

var mu sync.Mutex

var palette = []color.Color{
	color.Black,
	color.RGBA{ff, 0, 0, ff},  //red
	color.RGBA{0, ff, 0, ff},  //green
	color.RGBA{0, 0, ff, ff},  //blue
	color.RGBA{ff, ff, 0, ff}, //yellow
	color.RGBA{ff, 0, ff, ff}, //magenta
	color.RGBA{0, ff, ff, ff}, //cyan
	color.White,
}

var counter uint8 = 1 // next color in palette

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "?counter=") {
			counter++
		}
		lissajous(w)
	})
		http.HandleFunc("/counter=", func(w http.ResponseWriter, r *http.Request) {
			counter++
			lissajous(w)
		})
	log.Fatal(http.ListenAndServe("localhost:8001", nil))
}

func lissajous(out io.Writer) {
	const ( // number of complete x oscillator revolutions
		cycles  = 5
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				counter%7)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
