package lissajous

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
)

var pallete = []color.Color{
	color.RGBA{0xFF, 0xFF, 0xFF, 0xFF}, // Branco
	color.RGBA{0xFF, 0x00, 0x00, 0xFF}, // Vermelho
	color.RGBA{0x00, 0xFF, 0x00, 0xFF}, // Verde
	color.RGBA{0x00, 0x00, 0xFF, 0xFF}, // Azul
	color.RGBA{0xFF, 0xA5, 0x00, 0xFF}, // Laranja
}

type LissajousParams struct {
	Cycles  int
	Res     float64
	Size    int
	Nframes int
	Delay   int
}

func NewLissajousParams(cycles int, res float64, size, nframes, delay int) LissajousParams {
	return LissajousParams{
		cycles, res, size, nframes, delay,
	}
}

func Lissajous(out io.Writer, params LissajousParams) {
	const (
		defaultCycles  = 5     // número de revoluções completas do oscilador X
		defaultRes     = 0.001 // resolução angular
		defaultSize    = 100   // canvas da imagem cobre de (-size..+size)
		defaultNframes = 64    // número de quadros da animação
		defaultDelay   = 8     // tempo entre quadros em unidades de 10ms, ou seja, 80ms
	)

	cycles := params.Cycles
	if cycles == 0 {
		cycles = defaultCycles
	}

	res := params.Res
	if res == 0 {
		res = defaultRes
	}

	size := params.Size
	if size == 0 {
		size = defaultSize
	}

	nframes := params.Nframes
	if nframes == 0 {
		nframes = defaultNframes
	}

	delay := params.Delay
	if delay == 0 {
		delay = defaultDelay
	}

	// frequência relativa do oscilador y
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, pallete)
		colorIndex := rand.Intn(len(pallete)-1) + 1
		fmt.Printf("color: %d\n", colorIndex)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), uint8(colorIndex))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
