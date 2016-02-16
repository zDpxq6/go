// © 2016 zDpxq6
// Code by Alan A. A. Donovan & Brian W. Kernighan/Adapted.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, bx, by, cx, cy, dx, dy, ok := calculatePolygon(i, j)
			if !ok {
				continue
			}
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}

func calculatePolygon(i, j int) (float64, float64, float64, float64, float64, float64, float64, float64, bool) {
	ax, ay, ok := convertAxis(i+1, j)
	if !ok {
		return 0, 0, 0, 0, 0, 0, 0, 0, false
	}
	bx, by, ok := convertAxis(i, j)
	if !ok {
		return 0, 0, 0, 0, 0, 0, 0, 0, false
	}
	cx, cy, ok := convertAxis(i, j+1)
	if !ok {
		return 0, 0, 0, 0, 0, 0, 0, 0, false
	}
	dx, dy, ok := convertAxis(i+1, j+1)
	if !ok {
		return 0, 0, 0, 0, 0, 0, 0, 0, false
	}
	return ax, ay, bx, by, cx, cy, dx, dy, true
}

func convertAxis(i, j int) (float64, float64, bool) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z, ok := calculateHeight(x, y)

	if !ok {
		return 0, 0, false
	}

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, true
}

func calculateHeight(x, y float64) (float64, bool) {
	r := math.Hypot(x, y)
	z := math.Sin(r) / r
	return z, !(math.IsNaN(z) || math.IsInf(z, 0))
}
