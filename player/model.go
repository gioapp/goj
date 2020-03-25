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
	Buttons map[string]*widget.Button
	Tracks  []Track
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
	Artist   string
	Title    string
	Album    string
	Track    string
	Genre    string
	Year     string
	Image    []byte
	Path     string
}
