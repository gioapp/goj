package player

import (
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type GoJoy struct {
	Window  *app.Window
	Context *layout.Context
	Theme   *material.Theme
	Player  *Player
	//Navigation map[string]*theme.DuoUIthemeNav
	Menu    *MenuBar
	Layouts *Layouts
}
type Playlist struct {
	Buttons map[int]*widget.Button
	Tracks  map[int]Track
}
type Layouts struct {
	Main      *layout.Flex
	Menu      *layout.Flex
	Body      *layout.Flex
	Info      *layout.Flex
	Playlist  *layout.List
	TrackInfo *layout.Flex
}

type Track struct {
	Metadata map[string]interface{}
	Id       int
	Filename string
	Artist   string
	Title    string
	Album    string
	Track    string
	Genre    string
	Year     string
	Image    []byte
	Path     string
}

type scrollerGauge struct {
	Label   string
	Percent int
}
