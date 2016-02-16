// © 2016 zDpxq6
// Code by Alan A. A. Donovan & Brian W. Kernighan/Adapted.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
package surface

import (
	"fmt"
	"math"
	"io"
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

func Draw(w io.Writer) {
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, bx, by, cx, cy, dx, dy, color, ok := calculatePolygon(i, j)
			if !ok {
				continue
			}
			fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='%s'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, color)
		}
	}
	fmt.Fprintf(w, "</svg>")
}

func calculatePolygon(i, j int) (float64, float64, float64, float64, float64, float64, float64, float64, string, bool) {
	tax, tay, taz, ok := convertTo3D(i+1, j)
	if !ok {
		return 0, 0, 0, 0, 0, 0, 0, 0, "#ff0000", false
	}
	tbx, tby, tbz, ok := convertTo3D(i, j)
	if !ok {
		return 0, 0, 0, 0, 0, 0, 0, 0, "#ff0000", false
	}
	tcx, tcy, tcz, ok := convertTo3D(i, j+1)
	if !ok {
		return 0, 0, 0, 0, 0, 0, 0, 0, "#ff0000", false
	}
	tdx, tdy, tdz, ok := convertTo3D(i+1, j+1)
	if !ok {
		return 0, 0, 0, 0, 0, 0, 0, 0, "#ff0000", false
	}
	ax, ay := convertTo2D(tax, tay, taz)
	bx, by := convertTo2D(tbx, tby, tbz)
	cx, cy := convertTo2D(tcx, tcy, tcz)
	dx, dy := convertTo2D(tdx, tdy, tdz)
	c := color((taz + tbz + tcz + tdz) / 4)
	return ax, ay, bx, by, cx, cy, dx, dy, c, true
}

func convertTo3D(i, j int) (float64, float64, float64, bool) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z, ok := calculateHeight(x, y)

	if !ok {
		return 0, 0, 0, false
	} else {
		return x, y, z, true
	}

}

func convertTo2D(x, y, z float64) (float64, float64) {
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func color(z float64) string {
	a := int(math.Floor(z * 255))
	color := "#ffffff"
	if a < 0 {
		color = fmt.Sprintf("#%x%xff", 255+a, 255+a)
	} else if 0 < a {
		color = fmt.Sprintf("#ff%x%x", 255-a, 255-a)
	}
	return color
}

func calculateHeight(x, y float64) (float64, bool) {
	r := math.Hypot(x, y)
	z := math.Sin(r) / r
	return z, !(math.IsNaN(z) || math.IsInf(z, 0))
}
