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
	player    *Player
}

func (p *Player) Menu() *MenuBar {
	return &MenuBar{
		PlayPause: new(widget.Button),
		Stop:      new(widget.Button),
		Forward:   new(widget.Button),
		Backward:  new(widget.Button),
		Next:      new(widget.Button),
		Back:      new(widget.Button),
		Quit:      new(widget.Button),
		player:    p,
	}
}

func (m *MenuBar) Layout(gtx *layout.Context, th *material.Theme, ly *layout.Flex) func() {
	return func() {
		ly.Layout(gtx,
			layout.Flexed(0.25, m.menuButton(gtx, th, "Play/Pause", m.PlayPause, func() {
				i, err := playSong(m.player.Playing)
				if err != nil {
				}
				th.Caption(fmt.Sprint(i)).Layout(gtx)
			})),
			layout.Flexed(0.15, m.menuButton(gtx, th, "Stop", m.Stop, func() {})),
			layout.Flexed(0.15, m.menuButton(gtx, th, "Backward", m.Backward, func() {})),
			layout.Flexed(0.15, m.menuButton(gtx, th, "Forward", m.Forward, func() {})),
			layout.Flexed(0.15, m.menuButton(gtx, th, "Back", m.Back, func() {})),
			layout.Flexed(0.15, m.menuButton(gtx, th, "Next", m.Next, func() {})),
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
