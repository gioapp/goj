package player

import (
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op/paint"
	"github.com/gioapp/goj/pkg/gel"
	"github.com/gioapp/goj/pkg/gelook"
	"github.com/gioapp/goj/pkg/wavreader"
	"image"
)

type GoJoy struct {
	Window   *app.Window
	Context  *layout.Context
	Theme    *gelook.DuoUItheme
	Menu     *MenuBar
	Layouts  *Layouts
	Playlist *Playlist
	Playing  *Track
	infoList []string
	trackNum int
	trackSel int
	trackPos int
	trackLen int
	OnSelect selectCallback
	OnPause  pauseCallback
	OnSeek   seekCallback
	OnVolume volumeCallback
	state    playerState

	seek   *ScrollerGauge
	volume *ScrollerGauge
}

type Playlist struct {
	Buttons map[int]*gel.Button
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
	w            *wavreader.Reader
	im           *image.NRGBA
	Id           int
	Filename     string
	Artist       string
	Title        string
	Album        string
	TrackNumber  int
	TrackTotal   int
	Genre        string
	Year         string
	CoverImage   image.Image
	CoverImageOp paint.ImageOp
	Path         string
}
