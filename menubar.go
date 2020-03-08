package main

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"golang.org/x/exp/shiny/materialdesign/icons"
)

var (
	icon *material.Icon
)

// MenuBar is the component that define the menu bar.
type MenuBar struct {
	l         *layout.List
	Close     *widget.Button
	Quit      *widget.Button
	PlayPause *widget.Button
	Next      *widget.Button
	Back      *widget.Button
}

func Menu() *MenuBar {
	return &MenuBar{
		l: &layout.List{
			Axis:      layout.Horizontal,
			Alignment: layout.Middle,
		},
		Close:     &widget.Button{},
		Quit:      &widget.Button{},
		PlayPause: &widget.Button{},
		Next:      &widget.Button{},
		Back:      &widget.Button{},
	}
}

func (m *MenuBar) Layout(gtx *layout.Context, th *material.Theme, gojoy *GoJoy) {
	iconPlay, _ := material.NewIcon(icons.AVPlayArrow)
	//iconStop, _ := material.NewIcon(icons.AVStop)
	//iconPause, _ := material.NewIcon(icons.AVPause)
	commands := []func(){
		func() {
			for m.Close.Clicked(gtx) {
				///////
			}
			ic := th.IconButton(iconPlay)
			ic.Size = unit.Dp(float32(48))
			ic.Layout(gtx, m.Close)
		},
		func() {
			for m.Quit.Clicked(gtx) {
				///////
			}
			ic := th.IconButton(iconPlay)
			ic.Size = unit.Dp(float32(48))
			ic.Layout(gtx, m.Quit)
		},
		func() {
			for m.PlayPause.Clicked(gtx) {
				gojoy.Playlist.TogglePause()

			}
			ic := th.IconButton(iconPlay)
			ic.Size = unit.Dp(float32(48))
			ic.Layout(gtx, m.PlayPause)
		},
		func() {
			for m.Next.Clicked(gtx) {
				gojoy.Playlist.Next()
			}
			ic := th.IconButton(iconPlay)
			ic.Size = unit.Dp(float32(48))
			ic.Layout(gtx, m.Next)
		},
		func() {
			for m.Back.Clicked(gtx) {
				gojoy.Playlist.Back()
			}
			ic := th.IconButton(iconPlay)
			ic.Size = unit.Dp(float32(48))
			ic.Layout(gtx, m.Back)
		},
	}
	m.l.Layout(gtx, len(commands), func(i int) {
		layout.UniformInset(unit.Dp(16)).Layout(gtx, commands[i])
	})
}
