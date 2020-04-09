package player

import (
	"github.com/gioapp/goj/pkg/waveform"
	"image"
	"image/color"
)

func (g *GoJoy) loadWav(width, height int, zoom float32) *image.NRGBA {
	return waveform.MinMax(g.Playing.w, &waveform.Options{
		Width:   width,
		Height:  height,
		Zoom:    zoom,
		Half:    false,
		MarginL: 0,
		MarginR: 0,
		MarginT: 0,
		MarginB: 0,
		Front: &color.NRGBA{
			R: 255,
			G: 128,
			B: 0,
			A: 150,
		},
		Back: &color.NRGBA{
			A: 0,
		},
	})

}
