package player

import (
	"fmt"
	"gioui.org/layout"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"golang.org/x/image/draw"
	"image"
)

var (
	infoRowsList = &layout.List{
		Axis: layout.Vertical,
	}
)

func (g *GoJoy) Info() func() {
	return func() {
		if g.Playing != nil {
			g.Layouts.Track.Layout(g.Context,
				layout.Flexed(1, func() {
					g.Theme.DuoUIcontainer(16, g.Theme.Colors["Light"]).Layout(g.Context, layout.NW, func() {
						g.Layouts.TrackInfo.Layout(g.Context,
							layout.Flexed(1, func() {
								infoRows := []func(){
									func() {
										layout.Flex{}.Layout(g.Context,
											layout.Rigid(func() {
												g.Theme.H5("Filename: ").Layout(g.Context)
											}),
											layout.Flexed(1, func() {
												g.Theme.Body1(g.Playing.Filename).Layout(g.Context)
											}),
										)
									},
									func() {
										if g.Playing.Artist != "" {
											layout.Flex{}.Layout(g.Context,
												layout.Rigid(func() {
													g.Theme.H5("Artist: ").Layout(g.Context)
												}),
												layout.Flexed(1, func() {
													g.Theme.Body1(g.Playing.Artist).Layout(g.Context)
												}),
											)
										}
									},
									func() {
										if g.Playing.Title != "" {
											layout.Flex{}.Layout(g.Context,
												layout.Rigid(func() {
													g.Theme.H5("Title: ").Layout(g.Context)
												}),
												layout.Flexed(1, func() {
													g.Theme.Body1(g.Playing.Title).Layout(g.Context)
												}),
											)
										}
									},
									func() {
										if g.Playing.Album != "" {
											layout.Flex{}.Layout(g.Context,
												layout.Rigid(func() {
													g.Theme.H5("Album: ").Layout(g.Context)
												}),
												layout.Flexed(1, func() {
													g.Theme.Body1(g.Playing.Album).Layout(g.Context)
												}),
											)
										}
									},
									func() {
										if g.Playing.Album != "" {
											layout.Flex{}.Layout(g.Context,
												layout.Rigid(func() {
													g.Theme.H5("Genre: ").Layout(g.Context)
												}),
												layout.Flexed(1, func() {
													g.Theme.Body1(g.Playing.Genre).Layout(g.Context)
												}),
											)
										}
									},
									func() {
										if g.Playing.Year != "" {
											layout.Flex{}.Layout(g.Context,
												layout.Rigid(func() {
													g.Theme.H5("Year: ").Layout(g.Context)
												}),
												layout.Flexed(1, func() {
													g.Theme.Body1(fmt.Sprint(g.Playing.Year)).Layout(g.Context)
												}),
											)
										}
									},
									func() {
										if g.Playing.CoverImage != nil {
											sz := g.Context.Constraints.Width.Min
											if g.Playing.CoverImageOp.Size().X != sz {
												img := image.NewRGBA(image.Rectangle{Max: image.Point{X: sz, Y: sz}})
												draw.ApproxBiLinear.Scale(img, img.Bounds(), g.Playing.CoverImage, g.Playing.CoverImage.Bounds(), draw.Src, nil)
												g.Playing.CoverImageOp = paint.NewImageOp(img)
											}
											img := g.Theme.Image(g.Playing.CoverImageOp)
											img.Scale = float32(sz) / float32(g.Context.Px(unit.Dp(float32(sz))))
											img.Layout(g.Context)
										}
									},
								}
								infoRowsList.Layout(g.Context, len(infoRows), func(i int) {
									layout.UniformInset(unit.Dp(8)).Layout(g.Context, infoRows[i])
								})
							}),
						)
					})
				}),
				layout.Rigid(func() {
					//cs := g.Context.Constraints
					//sz := g.Context.Constraints.Width.Min
					if g.Playing.wavImageOp.Size().X != 300 {

						g.Playing.wavImageOp = paint.NewImageOp(g.loadWav(300, 64, 1.6))
					}
					imgRender := g.Theme.Image(g.Playing.wavImageOp)
					imgRender.Scale = float32(300) / float32(g.Context.Px(unit.Dp(float32(300))))
					imgRender.Layout(g.Context)
				}),
				layout.Rigid(func() {
					g.Context.Constraints.Height.Max = 64
					g.seekElement.Layout(g.Context)
				}),
				layout.Rigid(func() {
					g.Context.Constraints.Height.Max = 32
					g.volumeElement.Layout(g.Context)
				}),
			)
		}
	}
}
