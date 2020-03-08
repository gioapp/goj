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
	"gioui.org/widget/material"
	"github.com/gioapp/goj/play"
	"image"
	"image/color"
)

//
//var (
//    win      app.Windower
//    playlist play.Playlist
//)
//
//func main() {
//    log.Println(os.Args)
//    runtime.GOMAXPROCS(1)
//    if len(os.Args) < 2 {
//        return
//    }
//
//    // Initialize FFT.
//    Init()
//
//    playlist = play.New()
//    playlist.Init(os.Args[1:])
//
//    go func() {
//        playlist.Start()
//    }()
//
//    app.OnLaunch = func() {
//        appMenu := &MenuBar{}
//
//        if menuBar, ok := app.MenuBar(); ok {
//            menuBar.Mount(appMenu)
//        }
//
//        win = newMainWindow()
//        win.Mount(&Player{})
//    }
//
//    app.OnReopen = func() {
//        if win != nil {
//            return
//        }
//        win = newMainWindow()
//        win.Mount(&Player{})
//    }
//
//    app.Run()
//}
//

type GoJoy struct {
	Window   *app.Window
	Context  *layout.Context
	Theme    *material.Theme
	Playlist play.Playlist
	//Navigation map[string]*theme.DuoUIthemeNav
	Menu *MenuBar
}

func NewGoJoy() *GoJoy {
	gojoy := &GoJoy{
		Window: app.NewWindow(
			app.Size(unit.Dp(400), unit.Dp(800)),
			app.Title("ParallelCoin"),
		),
		Theme:    material.NewTheme(),
		Playlist: play.New(),
		Menu:     Menu(),
	}
	gojoy.Context = layout.NewContext(gojoy.Window.Queue())
	return gojoy
}

func main() {
	gojoy := NewGoJoy()

	go func() {

		for e := range gojoy.Window.Events() {
			if e, ok := e.(system.FrameEvent); ok {
				gojoy.Context.Reset(e.Config, e.Size)
				DrawRectangle(gojoy.Context, gojoy.Context.Constraints.Width.Max, gojoy.Context.Constraints.Height.Max, HexARGB("ff303030"), [4]float32{0, 0, 0, 0}, unit.Dp(0))

				layout.Flex{
					Axis: layout.Vertical,
				}.Layout(gojoy.Context,
					layout.Rigid(func() {

						gojoy.Menu.Layout(gojoy.Context, gojoy.Theme, gojoy.Menu)
					}),
					layout.Flexed(1, func() {
						cs := gojoy.Context.Constraints
						DrawRectangle(gojoy.Context, cs.Width.Max, cs.Height.Max, HexARGB("ff303030"), [4]float32{0, 0, 0, 0}, unit.Dp(0))
					}),
				)
				e.Frame(gojoy.Context.Ops)
			}
		}
	}()
	app.Main()
}

func DrawRectangle(gtx *layout.Context, w, h int, color color.RGBA, borderRadius [4]float32, inset unit.Value) {
	in := layout.UniformInset(inset)
	in.Layout(gtx, func() {
		//cs := gojoy.Context.Constraints
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
