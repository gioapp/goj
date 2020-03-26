package player

import (
	"fmt"
	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"golang.org/x/image/draw"
	"image"
	"image/color"
)

func (g *GoJoy) View() func() {
	return func() {
		g.Layouts.Body.Layout(g.Context,
			layout.Flexed(0.5, func() {
				if g.Playing != nil {
					g.Layouts.Track.Layout(g.Context,
						layout.Flexed(1, func() {
							g.Layouts.TrackInfo.Layout(g.Context,
								layout.Rigid(func() {
									layout.Flex{Axis: layout.Vertical}.Layout(g.Context,
										layout.Rigid(func() {
											g.Theme.H5("Filename: ").Layout(g.Context)
										}),
										layout.Rigid(func() {
											g.Theme.H6("Artist: ").Layout(g.Context)
										}),
										layout.Rigid(func() {
											g.Theme.H6("Title: ").Layout(g.Context)
										}),
										layout.Rigid(func() {

											g.Theme.H6("Album: ").Layout(g.Context)
										}),
										layout.Rigid(func() {
											g.Theme.H6("Track: ").Layout(g.Context)
										}),
										layout.Rigid(func() {
											g.Theme.H6("Genre: ").Layout(g.Context)
										}),
										layout.Rigid(func() {
											g.Theme.H6("Year: ").Layout(g.Context)
										}),
									)
								}),
								layout.Flexed(1, func() {

									layout.Flex{Axis: layout.Vertical}.Layout(g.Context,
										layout.Rigid(func() {
											g.Theme.Body1(g.Playing.Filename).Layout(g.Context)
										}),
										layout.Rigid(func() {
											if g.Playing.Artist != "" {
												g.Theme.Body1(g.Playing.Artist).Layout(g.Context)
											}
										}),
										layout.Rigid(func() {
											if g.Playing.Title != "" {
												g.Theme.Body1(g.Playing.Title).Layout(g.Context)
											}
										}),
										layout.Rigid(func() {
											if g.Playing.Image != nil {
												sz := g.Context.Constraints.Width.Min
												if g.Playing.imgOp.Size().X != sz {
													img := image.NewRGBA(image.Rectangle{Max: image.Point{X: sz, Y: sz}})
													draw.ApproxBiLinear.Scale(img, img.Bounds(), g.Playing.Image, g.Playing.Image.Bounds(), draw.Src, nil)
													g.Playing.imgOp = paint.NewImageOp(img)
												}
												img := g.Theme.Image(g.Playing.imgOp)
												img.Scale = float32(sz) / float32(g.Context.Px(unit.Dp(float32(sz))))
												img.Layout(g.Context)
											}
										}),
										layout.Rigid(func() {
											if g.Playing.Album != "" {
												g.Theme.Body1(g.Playing.Album).Layout(g.Context)
											}
										}),
										layout.Rigid(func() {
											if g.Playing.Track != "" {
												g.Theme.Body1(g.Playing.Track).Layout(g.Context)
											}
										}),
										layout.Rigid(func() {
											if g.Playing.Genre != "" {
												g.Theme.Body1(g.Playing.Genre).Layout(g.Context)
											}
										}),
										layout.Rigid(func() {
											if g.Playing.Year != "" {
												g.Theme.Body1(g.Playing.Year).Layout(g.Context)
											}
										}),
									)
								}),
							)
						}),
						layout.Rigid(func() {
							g.seek.Layout(g.Context)
						}),
						layout.Rigid(func() {
							g.volume.Layout(g.Context)
						}),
					)
				}
			}),
			layout.Flexed(0.5, g.TrackList()),
		)
	}
}

func (g *GoJoy) TrackList() func() {
	return func() {
		if g.Playlist.Tracks != nil {
			g.Layouts.Playlist.Layout(g.Context, len(g.Playlist.Tracks), func(i int) {
				track := g.Playlist.Tracks[i]
				for g.Playlist.Buttons[track.Id].Clicked(g.Context) {
					g.Playing = &track
				}

				b := g.Theme.Button(track.Filename)
				b.Background = HexARGB("ffcfcf30")
				b.Layout(g.Context, g.Playlist.Buttons[track.Id])

				//fmt.Println(song.Path)
			})
		}
	}
}
func DrawRectangle(gtx *layout.Context, w, h int, color color.RGBA, borderRadius [4]float32, inset unit.Value) {
	in := layout.UniformInset(inset)
	in.Layout(gtx, func() {
		//cs := gtx.Constraints
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
