package player

import (
	"fmt"
	"gioui.org/layout"
	"github.com/gioapp/goj/pkg/gel"
	"github.com/gioapp/goj/pkg/gelook"
)

var (
	icon *gelook.DuoUIicon
)

// MenuBar is the component that define the menu bar.
type MenuBar struct {
	PlayPause *gel.Button
	Stop      *gel.Button
	Forward   *gel.Button
	Backward  *gel.Button
	Next      *gel.Button
	Back      *gel.Button
	Quit      *gel.Button
}

func (g *GoJoy) MenuBar() *MenuBar {
	return &MenuBar{
		PlayPause: new(gel.Button),
		Stop:      new(gel.Button),
		Forward:   new(gel.Button),
		Backward:  new(gel.Button),
		Next:      new(gel.Button),
		Back:      new(gel.Button),
		Quit:      new(gel.Button),
	}
}

func (g *GoJoy) MenuBarLayout(gtx *layout.Context, th *gelook.DuoUItheme, ly *layout.Flex) func() {
	return func() {
		ly.Layout(gtx,
			layout.Flexed(0.25, g.Menu.menuButton(gtx, th, "Play/Pause", "Run", g.Menu.PlayPause, func() {
				if g.state == Playing {
					g.trackPos++
					if g.trackLen != 0 {
						g.seek.Value = int(float32(g.trackPos) / float32(g.trackLen) * 100)
						g.seek.Label = fmt.Sprintf("%d:%.2d / %d:%.2d", g.trackPos/60, g.trackPos%60, g.trackLen/60, g.trackLen%60)
						if g.seek.Value >= 100 {
							g.trackNum++
							if g.trackNum >= len(g.Playlist.Tracks) {
								g.trackNum = 0
							}
							g.playSong(g.trackNum)
						}
					}
				} else if g.state == Stopped {
					g.trackPos = 0
				}

				//
				//err := g.getSong()
				//if err != nil {
				//}
				//g.playSong()
				//
				//g.seek.Value = g.TrackTotal

				//th.Caption(fmt.Sprint(i)).Layout(gtx)
				//fmt.Println("addad", i)

			})),
			layout.Flexed(0.15, g.Menu.menuButton(gtx, th, "Stop", "Stop", g.Menu.Stop, func() {})),
			layout.Flexed(0.15, g.Menu.menuButton(gtx, th, "Backward", "Backward", g.Menu.Backward, func() {})),
			layout.Flexed(0.15, g.Menu.menuButton(gtx, th, "Forward", "Forward", g.Menu.Forward, func() {})),
			layout.Flexed(0.15, g.Menu.menuButton(gtx, th, "Back", "Back", g.Menu.Back, func() {})),
			layout.Flexed(0.15, g.Menu.menuButton(gtx, th, "Next", "Next", g.Menu.Next, func() {})),

			//layout.Flexed(0.15, g.Menu.menuButton(gtx, th, "SEND", "sendIcon", g.Menu.Next),

		)
		//iconPlay, _ := material.NewIcon(icons.AVPlayArrow)
		//iconStop, _ := material.NewIcon(icons.AVStop)
		//iconPause, _ := material.NewIcon(icons.AVPause)

	}
}

//func (m *MenuBar) menuButton(gtx *layout.Context, th *gelook.DuoUItheme, label string, button *gel.Button, action func()) func() {
func (m *MenuBar) menuButton(gtx *layout.Context, th *gelook.DuoUItheme, label, icon string, button *gel.Button, action func()) func() {
	return func() {
		//for button.Clicked(gtx) {
		//	action()
		//}
		//b := th.Button(label)
		//
		//b.Layout(gtx, button)

		menuItem := th.DuoUIbutton(th.Fonts["Secondary"],
			label, th.Colors["Light"],
			th.Colors["LightGrayII"],
			th.Colors["LightGrayII"],
			th.Colors["Dark"], icon,
			th.Colors["Primary"],
			24, 48, gtx.Constraints.Width.Max, 96, 0, 0, 0, 0,
			//nav.TextSize, nav.IconSize,
			//nav.Width, nav.Height,
			//nav.PaddingVertical, nav.PaddingHorizontal, nav.PaddingVertical, nav.PaddingHorizontal,
		)
		for button.Clicked(gtx) {
			action()
		}
		menuItem.MenuLayout(gtx, button)
	}
}
