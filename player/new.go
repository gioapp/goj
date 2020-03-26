package player

import (
	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"
)

func NewGoJoy() *GoJoy {
	gofont.Register()
	layouts := &Layouts{
		Main: &layout.Flex{
			Axis:      layout.Vertical,
			Spacing:   layout.SpaceBetween,
			Alignment: layout.Middle,
		},
		Menu: &layout.Flex{
			Axis:      layout.Horizontal,
			Spacing:   layout.SpaceBetween,
			Alignment: layout.Middle,
		},
		Body: &layout.Flex{
			Axis:      layout.Horizontal,
			Spacing:   layout.SpaceBetween,
			Alignment: layout.Middle,
		},
		Info: &layout.Flex{
			Axis:      layout.Horizontal,
			Spacing:   layout.SpaceBetween,
			Alignment: layout.Middle,
		},
		Playlist: &layout.List{
			Axis: layout.Vertical,
		},
		TrackInfo: &layout.Flex{
			Axis: layout.Horizontal,
		},
	}
	gojoy := &GoJoy{
		Window: app.NewWindow(
			app.Size(unit.Dp(400), unit.Dp(800)),
			app.Title("ParallelCoin"),
		),
		Theme: material.NewTheme(),
		//Playlist: play.New(),
		Layouts: layouts,
		Player:  NewPlayer(),
	}
	gojoy.Context = layout.NewContext(gojoy.Window.Queue())
	gojoy.Menu = gojoy.Player.Menu()
	//p, _ := NewPlayer(gojoy.Player.Playlist.Tracks, 0)
	//fmt.Println("tet",p)
	//gojoy.Player = p
	return gojoy
}
