package player

import (
	"fmt"
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

var (
	icon *material.Icon
)

// MenuBar is the component that define the menu bar.
type MenuBar struct {
	PlayPause *widget.Button
	Stop      *widget.Button
	Forward   *widget.Button
	Backward  *widget.Button
	Next      *widget.Button
	Back      *widget.Button
	Quit      *widget.Button
}

func (g *GoJoy) MenuBar() *MenuBar {
	return &MenuBar{
		PlayPause: new(widget.Button),
		Stop:      new(widget.Button),
		Forward:   new(widget.Button),
		Backward:  new(widget.Button),
		Next:      new(widget.Button),
		Back:      new(widget.Button),
		Quit:      new(widget.Button),
	}
}

func (g *GoJoy) MenuBarLayout(gtx *layout.Context, th *material.Theme, ly *layout.Flex) func() {
	return func() {
		ly.Layout(gtx,
			layout.Flexed(0.25, g.Menu.menuButton(gtx, th, "Play/Pause", g.Menu.PlayPause, func() {
				//
				//if m.player.state == Playing {
				//	m.player.songPos++
				//	if m.player.songLen != 0 {
				//		m.player.scrollerGauge.Percent = int(float32(m.player.songPos) / float32(m.player.songLen) * 100)
				//		m.player.scrollerGauge.Label = fmt.Sprintf("%d:%.2d / %d:%.2d", m.player.songPos/60, m.player.songPos%60, m.player.songLen/60, m.player.songLen%60)
				//		if m.player.scrollerGauge.Percent >= 100 {
				//			m.player.songNum++
				//			if m.player.songNum >= len(m.player.Playlist.Tracks) {
				//				m.player.songNum = 0
				//			}
				//			m.player.playSong(m.player.songNum)
				//		}
				//	}
				//} else if m.player.state == Stopped {
				//	m.player.songPos = 0
				//}
				i, err := playSong(g.Playing)
				if err != nil {
				}
				g.seek.Value = i
				th.Caption(fmt.Sprint(i)).Layout(gtx)
				fmt.Println("addad", i)

			})),
			layout.Flexed(0.15, g.Menu.menuButton(gtx, th, "Stop", g.Menu.Stop, func() {})),
			layout.Flexed(0.15, g.Menu.menuButton(gtx, th, "Backward", g.Menu.Backward, func() {})),
			layout.Flexed(0.15, g.Menu.menuButton(gtx, th, "Forward", g.Menu.Forward, func() {})),
			layout.Flexed(0.15, g.Menu.menuButton(gtx, th, "Back", g.Menu.Back, func() {})),
			layout.Flexed(0.15, g.Menu.menuButton(gtx, th, "Next", g.Menu.Next, func() {})),
		)
		//iconPlay, _ := material.NewIcon(icons.AVPlayArrow)
		//iconStop, _ := material.NewIcon(icons.AVStop)
		//iconPause, _ := material.NewIcon(icons.AVPause)

	}
}

func (m *MenuBar) menuButton(gtx *layout.Context, th *material.Theme, label string, button *widget.Button, action func()) func() {
	return func() {
		for button.Clicked(gtx) {
			action()
		}
		b := th.Button(label)

		b.Layout(gtx, button)
	}
}
