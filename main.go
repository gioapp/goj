package main

import (
	"fmt"
	"gioui.org/app"
	"gioui.org/f32"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"github.com/gioapp/goj/player"
	"image"
	"image/color"
)

func main() {
	g := player.NewGoJoy()
	//g.NewPlayer()
	go func() {
		for e := range g.Window.Events() {
			if e, ok := e.(system.FrameEvent); ok {
				g.Context.Reset(e.Config, e.Size)
				DrawRectangle(g.Context, g.Context.Constraints.Width.Max, g.Context.Constraints.Height.Max, HexARGB("ff303030"), [4]float32{0, 0, 0, 0}, unit.Dp(0))
				g.Layouts.Main.Layout(g.Context,
					layout.Flexed(1, g.View()),
					layout.Rigid(g.MenuBarLayout(g.Context, g.Theme, g.Layouts.Menu)),
				)
				e.Frame(g.Context.Ops)
			}
		}
	}()
	app.Main()
}

func DrawRectangle(gtx *layout.Context, w, h int, color color.RGBA, borderRadius [4]float32, inset unit.Value) {
	in := layout.UniformInset(inset)
	in.Layout(gtx, func() {
		//cs := g.Context.Constraints
		square := f32.Rectangle{
			Max: f32.Point{
				X: float32(w),
				Y: float32(h),
			},
		}
		paint.ColorOp{Color: color}.Add(gtx.Ops)

		clip.Rect{Rect: square,
			NE: borderRadius[0], NW: borderRadius[1], SE: borderRadius[2], SW: borderRadius[3]}.Op(gtx.Ops).Add(gtx.Ops) // HLdraw
		paint.PaintOp{Rect: square}.Add(gtx.Ops)
		gtx.Dimensions = layout.Dimensions{Size: image.Point{X: w, Y: h}}
	})
}

func HexARGB(s string) (c color.RGBA) {
	_, _ = fmt.Sscanf(s, "%02x%02x%02x%02x", &c.A, &c.R, &c.G, &c.B)
	return
}
