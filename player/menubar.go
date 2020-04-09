package player

import (
	"gioui.org/layout"
	"github.com/gioapp/gel"
	"github.com/gioapp/gelook"
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
				if g.trackNum != -1 {
					if g.state == Playing {
						g.OnPause(true)
						g.state = Paused
					} else {
						g.OnPause(false)
						g.state = Playing

					}
					g.renderStatus()
				}
			})),
			layout.Flexed(0.15, g.Menu.menuButton(gtx, th, "Stop", "Stop", g.Menu.Stop, func() {
				g.playSong()
				g.OnPause(true)
				g.state = Stopped
				g.seekElement.Value = 0
				g.seekElement.Label = "0:00 / 0:00"
				g.renderStatus()
			})),
			layout.Flexed(0.15, g.Menu.menuButton(gtx, th, "Backward", "Backward", g.Menu.Backward, func() {
				if g.trackNum != -1 {
					g.trackPos -= 10
					if g.trackPos < 0 {
						g.trackPos = 0
					}
					g.OnSeek(g.trackPos)
				}
			})),
			layout.Flexed(0.15, g.Menu.menuButton(gtx, th, "Forward", "Forward", g.Menu.Forward, func() {
				if g.trackNum != -1 {
					g.trackPos += 10
					g.OnSeek(g.trackPos)
				}
			})),
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

		//menuItem := th.DuoUIbutton(th.Fonts["Secondary"],
		//	label, th.Colors["Light"],
		//	th.Colors["LightGrayII"],
		//	th.Colors["LightGrayII"],
		//	th.Colors["Dark"], icon,
		//	th.Colors["Primary"],
		//	24, 48, gtx.Constraints.Width.Max, 96, 0, 0, 0, 0,
		//	//nav.TextSize, nav.IconSize,
		//	//nav.Width, nav.Height,
		//	//nav.PaddingVertical, nav.PaddingHorizontal, nav.PaddingVertical, nav.PaddingHorizontal,
		//)
		for button.Clicked(gtx) {
			action()
		}
		//menuItem.MenuLayout(gtx, button)
	}
}
