package player

import (
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op/paint"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"image"
)

type GoJoy struct {
	Window   *app.Window
	Context  *layout.Context
	Theme    *material.Theme
	Menu     *MenuBar
	Layouts  *Layouts
	Playlist *Playlist
	Playing  *Track
	infoList []string
	seek     *ScrollerGauge
	volume   *ScrollerGauge
	//songs     []Track
	//songNames []string
	songNum  int
	songSel  int
	songLen  int
	OnSelect selectCallback
	OnPause  pauseCallback
	OnSeek   seekCallback
	OnVolume volumeCallback

	state playerState
}

type Playlist struct {
	Buttons map[int]*widget.Button
	Tracks  map[int]Track
}
type Layouts struct {
	Main *layout.Flex
	Menu *layout.Flex
	Body *layout.Flex
	Info *layout.Flex

	Playlist  *layout.List
	Track     *layout.Flex
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
	Image    image.Image
	imgOp    paint.ImageOp
	Path     string
}
