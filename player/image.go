package player

import (
	"flag"
	"fmt"
	"github.com/gioapp/goj/pkg/waveform"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
	"image"
	"image/color"
	"image/draw"
)

var (
	wavStart = flag.Float64("start", 0.0, "set start (in seconds)")
	wavEnd   = flag.Float64("end", -1.0, "set end (in seconds)")
	margin   = flag.Int("margin", 50, "set margins in pixels")
	segLen   = flag.Float64("seg", 0.1, "set segment length (in seconds)")
	rangeL   = flag.Float64("range-start", -1.0, "set range left (in seconds)")
	rangeR   = flag.Float64("range-end", -1.0, "set range right (in seconds)")
)

func (t *Track) processWav() error {

	n := t.w.Len()
	s := t.w.Duration().Seconds()
	if *wavStart > s || *wavStart < 0 {
		return fmt.Errorf("invalid start %.2fs (duration: %.2fs)", *wavStart, s)
	}
	if *wavEnd > s {
		return fmt.Errorf("invalid end %.2fs (duration: %.2fs)", *wavEnd, s)
	}
	e0 := s
	if *wavEnd > 0 {
		e0 = *wavEnd
	}
	s0 := *wavStart
	if e0 < s0 {
		return fmt.Errorf("end < start")
	}

	r0 := t.w.Rate()
	n0 := uint64(s0 * float64(n) / s)
	n1 := uint64(e0 * float64(n) / s)
	fmt.Printf("r0: %d, n: %d - s0: %.2f, e0: %.2f - n0: %d, n1: %d\n", r0, n, s0, e0, n0, n1)
	w1, err := t.w.Slice(n0, n1)
	if err != nil {
		return err
	}

	// ---

	t.im = waveform.MinMax(w1, &waveform.Options{
		Width:   1800,
		Height:  400,
		Zoom:    1.7,
		Half:    false,
		MarginL: *margin,
		MarginR: *margin,
		MarginT: *margin,
		MarginB: *margin,
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

	// ---
	rc := t.im.Bounds()
	idx := rc.Dx()
	idy := rc.Dy()

	img := image.NewNRGBA(rc)

	// fill with checkerboard
	for y := 0; y < idy; y++ {
		for x := 0; x < idx; x++ {
			c := color.NRGBA{
				R: 20,
				G: 20,
				B: 20,
				A: 255,
			}
			nx := x / 10
			ny := y / 10
			if (nx+ny)%2 == 0 {
				c = color.NRGBA{
					R: 30,
					G: 30,
					B: 30,
					A: 255,
				}
			}
			img.SetNRGBA(x, y, c)
		}
	}

	dx := idx - *margin*2
	dy := idy - *margin*2

	t1 := w1.Duration().Seconds()
	fmt.Printf("sample-duration: %.3fs\n", t1)
	fmt.Printf("sample-rate:     %d\n", w1.Rate())
	fmt.Printf("pixels:          %d\n", dx)

	draw.Draw(img, rc, t.im, image.ZP, draw.Over)

	if *rangeL >= 0 || *rangeR >= 0 {
		if *rangeL > *rangeR {
			return fmt.Errorf("rangeL > rangeR")
		}

		if *rangeL > t1 {
			return fmt.Errorf("rangeL > end")
		}

		if *rangeR > t1 {
			return fmt.Errorf("rangeR > end")
		}

		rng := image.NewNRGBA(rc)

		fmt.Printf("range-start:     %.3fs\n", *rangeL)
		fmt.Printf("range-end:       %.3fs\n", *rangeR)

		x0 := int((*rangeL / t1) * float64(dx))
		x1 := int((*rangeR / t1) * float64(dx))

		c := color.NRGBA{
			R: 50,
			G: 150,
			B: 150,
			A: 100,
		}
		for x := x0; x < x1; x++ {
			for y := 0; y < dy; y++ {
				rng.SetNRGBA(x+*margin, y+*margin, c)
			}
		}

		draw.Draw(img, rc, rng, image.ZP, draw.Over)

		col := color.NRGBA{250, 150, 100, 220}
		addLabel(img, col, x0+*margin, *margin-5, fmt.Sprintf("%.3f", *rangeL))
		addLabel(img, col, x1+*margin, *margin-5, fmt.Sprintf("%.3f", *rangeR))
	}

	if *segLen > 0 {
		s1 := t1 / *segLen
		tx := int(float64(dx) / s1)
		// fmt.Printf("%d samples per 10ms\n", 10*w1.Rate()/1000)
		// fmt.Printf("%d pixels per 10ms\n", tx)

		for x := 0; x < dx; x++ {
			c := color.NRGBA{
				R: 100,
				G: 100,
				B: 100,
				A: 255,
			}
			if (x/tx)%2 != 0 {
				c = color.NRGBA{
					R: 200,
					G: 200,
					B: 200,
					A: 255,
				}
			}
			for y := dy - 3; y < dy; y++ {
				img.SetNRGBA(x+*margin, y+*margin, c)
			}
		}
	}

	if *segLen > 0 {
		col := color.NRGBA{250, 120, 200, 200}
		s1 := t1 / *segLen
		tx := int(float64(dx) / s1)

		if tx > 50 {
			s := 0.0
			for x := 0; x <= dx; x += tx {
				addLabel(img, col, x+*margin, *margin+dy+20, fmt.Sprintf("%.3f", s))

				for y := 0; y < 5; y++ {
					img.SetNRGBA(x+*margin, *margin+dy+y, col)
				}
				s += *segLen
			}
		}

		for y := -1; y <= dy; y++ {
			img.SetNRGBA(*margin, *margin+y, col)
			if (y+1)%(dy/10) == 0 {
				for x := 0; x < 5; x++ {
					img.SetNRGBA(*margin-x, *margin+y, col)
				}
			}
		}
	}

	return nil
}

func addLabel(img *image.NRGBA, col color.NRGBA, x, y int, label string) {
	point := fixed.Point26_6{
		X: fixed.Int26_6(x * 64),
		Y: fixed.Int26_6(y * 64),
	}

	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(col),
		Face: basicfont.Face7x13,
		Dot:  point,
	}
	adv := d.MeasureString(label)
	d.Dot.X -= adv / 2
	d.DrawString(label)
}
