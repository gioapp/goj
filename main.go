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
	"github.com/gioapp/goj/player"
	"image"
	"image/color"
)

func main() {
	gj := player.NewGoJoy()
	//gj.NewPlayer()
	go func() {
		for e := range gj.Window.Events() {
			if e, ok := e.(system.FrameEvent); ok {
				gj.Context.Reset(e.Config, e.Size)
				DrawRectangle(gj.Context, gj.Context.Constraints.Width.Max, gj.Context.Constraints.Height.Max, HexARGB("ff303030"), [4]float32{0, 0, 0, 0}, unit.Dp(0))
				gj.Layouts.Main.Layout(gj.Context,
					layout.Flexed(1, func() {
						gj.Layouts.Body.Layout(gj.Context,
							layout.Flexed(0.5, func() {
								if gj.Player.Playing != nil {
									gj.Layouts.TrackInfo.Layout(gj.Context,
										layout.Rigid(func() {
											layout.Flex{Axis: layout.Vertical}.Layout(gj.Context,
												layout.Rigid(func() {
													gj.Theme.H5("Filename: ").Layout(gj.Context)
												}),
												layout.Rigid(func() {
													gj.Theme.H6("Artist: ").Layout(gj.Context)
												}),
												layout.Rigid(func() {
													gj.Theme.H6("Title: ").Layout(gj.Context)
												}),
												layout.Rigid(func() {
													gj.Theme.H6("Album: ").Layout(gj.Context)
												}),
												layout.Rigid(func() {
													gj.Theme.H6("Track: ").Layout(gj.Context)
												}),
												layout.Rigid(func() {
													gj.Theme.H6("Genre: ").Layout(gj.Context)
												}),
												layout.Rigid(func() {
													gj.Theme.H6("Year: ").Layout(gj.Context)
												}),
											)
										}),
										layout.Flexed(1, func() {

											layout.Flex{Axis: layout.Vertical}.Layout(gj.Context,
												layout.Rigid(func() {
													gj.Theme.Body1(gj.Player.Playing.Filename).Layout(gj.Context)
												}),
												layout.Rigid(func() {
													if gj.Player.Playing.Artist != "" {
														gj.Theme.Body1(gj.Player.Playing.Artist).Layout(gj.Context)
													}
												}),
												layout.Rigid(func() {
													if gj.Player.Playing.Title != "" {
														gj.Theme.Body1(gj.Player.Playing.Title).Layout(gj.Context)
													}
												}),
												layout.Rigid(func() {
													if gj.Player.Playing.Image != nil {
														i, _ := material.NewIcon(gj.Player.Playing.Image)
														//sz := gj.Context.Constraints.Width.Min
														//img := image.NewRGBA(image.Rectangle{Max: image.Point{X: sz, Y: sz}})
														//draw.ApproxBiLinear.Scale(img, img.Bounds(), img, img.Bounds(), draw.Src, nil)
														//addrQR := paint.NewImageOp(img)
														//gj.Theme.Image(addrQR)
														i.Layout(gj.Context, unit.Dp(50))
													}
												}),
												layout.Rigid(func() {
													if gj.Player.Playing.Album != "" {
														gj.Theme.Body1(gj.Player.Playing.Album).Layout(gj.Context)
													}
												}),
												layout.Rigid(func() {
													if gj.Player.Playing.Track != "" {
														gj.Theme.Body1(gj.Player.Playing.Track).Layout(gj.Context)
													}
												}),
												layout.Rigid(func() {
													if gj.Player.Playing.Genre != "" {
														gj.Theme.Body1(gj.Player.Playing.Genre).Layout(gj.Context)
													}
												}),
												layout.Rigid(func() {
													if gj.Player.Playing.Year != "" {
														gj.Theme.Body1(gj.Player.Playing.Year).Layout(gj.Context)
													}
												}),
											)
										}),
									)
								}
							}),
							layout.Flexed(0.5, func() {
								if gj.Player.Playlist.Tracks != nil {
									gj.Layouts.Playlist.Layout(gj.Context, len(gj.Player.Playlist.Tracks), func(i int) {
										track := gj.Player.Playlist.Tracks[i]
										for gj.Player.Playlist.Buttons[track.Id].Clicked(gj.Context) {
											gj.Player.Playing = &track
										}

										b := gj.Theme.Button(track.Filename)
										b.Layout(gj.Context, gj.Player.Playlist.Buttons[track.Id])

										//fmt.Println(song.Path)
									})
								}
							}),
						)
					}),
					layout.Rigid(gj.Menu.Layout(gj.Context, gj.Theme, gj.Layouts.Menu)),
				)
				e.Frame(gj.Context.Ops)
			}
		}
	}()
	app.Main()
}

func DrawRectangle(gtx *layout.Context, w, h int, color color.RGBA, borderRadius [4]float32, inset unit.Value) {
	in := layout.UniformInset(inset)
	in.Layout(gtx, func() {
		//cs := gj.Context.Constraints
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
