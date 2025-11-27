package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

var pallete = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var n_cycles int
		var n_size int

		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}

		for k, v := range r.Form {
			if k == "cycles" && len(v) == 1 {
				converted_int, err := strconv.Atoi(v[0])
				if err != nil {
					fmt.Fprintf(w, "Something unexpected happened: %s", err)
				}
				n_cycles = int(converted_int)
			} else if k == "size" && len(v) == 1 {
				converted_size, err := strconv.Atoi(v[0])
				if err != nil {
					fmt.Fprintf(w, "Something unexpected happened: %s", err)
				}
				n_size = int(converted_size)
			}

		}

		lissajous(w, n_cycles, n_size)
	})

	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}

func lissajous(out io.Writer, n_cycles int, size int) {

	const (
		res     = 0.001
		nframes = 64
		delay   = 8
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0

	for i := 0; i < nframes; i++ {

		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, pallete)

		for t := 0.0; t < float64(n_cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)

			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), blackIndex)
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
