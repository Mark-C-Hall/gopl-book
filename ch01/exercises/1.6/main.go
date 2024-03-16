package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var palette = []color.Color{
	color.RGBA{0x00, 0x00, 0x00, 0xff}, // black
	color.RGBA{0x00, 0x80, 0x00, 0xff}, // green
	color.RGBA{0xff, 0x00, 0x00, 0xff}, // red
	color.RGBA{0x00, 0x00, 0xff, 0xff}, // blue
	color.RGBA{0xff, 0xff, 0x00, 0xff}, // yellow
	color.RGBA{0xff, 0x00, 0xff, 0xff}, // magenta
	color.RGBA{0x00, 0xff, 0xff, 0xff}, // cyan
	color.RGBA{0xff, 0x80, 0x00, 0xff}, // orange
	color.RGBA{0x80, 0x00, 0xff, 0xff}, // purple
	color.RGBA{0x80, 0x80, 0x80, 0xff}, // gray
}

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*3.14; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			colorIndex := uint8(i % len(palette))
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), colorIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
