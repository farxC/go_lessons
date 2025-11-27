package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
	color         = "grey"
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {

	mux := http.NewServeMux()

	srv := &http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	blue := "#0000ff"

	mux.HandleFunc("/svg", func(w http.ResponseWriter, r *http.Request) {

		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}
		localW := width
		localH := height
		localIC := color
		localOC := blue
		localxyscale := xyscale
		localzscale := zscale

		qW := r.FormValue("width")

		if qW != "" {
			if val, err := strconv.Atoi(qW); err == nil {
				localW = val
				localxyscale = float64(localW) / 2 / xyrange
			} else {
				log.Fatal(err)
			}
		}
		qH := r.FormValue("height")

		if qH != "" {
			if val, err := strconv.Atoi(qH); err == nil {
				localH = val
				localzscale = float64(localH) * 0.4
			} else {
				log.Fatal(err)
			}
		}

		qIC := r.FormValue("innerColor")

		if qIC != "" {
			localIC = qIC
		}

		qOC := r.FormValue("outerColor")
		if qOC != "" {
			localOC = qOC
		}

		w.Header().Set("Content-Type", "image/svg+xml; charset=utf-8")
		fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
			"style='stroke: %s; fill: white; stroke-width: 0.7' "+
			"width='%d' height='%d'>", localIC, localW, localH)
		for i := 0; i < cells; i++ {
			for j := 0; j < cells; j++ {

				ax, ay := corner(i+1, j, localxyscale, localzscale)
				bx, by := corner(i, j, localxyscale, localzscale)
				cx, cy := corner(i, j+1, localxyscale, localzscale)
				dx, dy := corner(i+1, j+1, localxyscale, localzscale)

				fmt.Fprintf(w, `<polygon style="fill:%s;stroke:%s" points='%g,%g %g,%g %g,%g %g,%g'/>\n`,
					localIC, localOC, ax, ay, bx, by, cx, cy, dx, dy)
			}
		}
		fmt.Fprintln(w, "</svg>")
	})

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		if err := r.ParseForm(); err != nil {
			log.Fatal(err)
		}

		qs := r.Form.Encode()

		src := "/svg"

		if qs != "" {
			src = src + "?" + qs
		}
		fmt.Fprintln(w, `<!doctype html><html><head><meta charset="utf-8"><title>Surface</title></head><body style="margin:0">`)
		fmt.Fprintln(w, `<div style="display:flex;justify-content:center;align-items:center;min-height:100vh;flex-direction:column">`)
		fmt.Fprintln(w, `<h1 style="margin:0 0 1rem 0">Surface</h1>`)
		fmt.Fprintf(w, `<img src="%s" alt="surface svg" style="max-width:100%%;height:auto;display:block"/>`, src)
		fmt.Fprintln(w, `</div></body></html>`)
	})

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}

func corner(i, j int, xyscale, zscale float64) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

//!-
