package player

import (
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op/paint"
	"github.com/faiface/beep"
	"github.com/faiface/beep/effects"
	"github.com/gioapp/gel"
	"github.com/gioapp/gelook"
	"github.com/gioapp/goj/pkg/wavreader"
	"image"
	"os"
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

	seekElement   *ScrollerGauge
	volumeElement *ScrollerGauge
	Sound         *Sound
}
type Sound struct {
	supportedFormats []string
	mainCtrl         *beep.Ctrl
	s                beep.StreamSeekCloser
	format           beep.Format
	volume           *effects.Volume
}

type Playlist struct {
	TracksNumber int
	Buttons      map[int]*gel.Button
	Tracks       map[int]Track
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
	f          *os.File
	w          *wavreader.Reader
	wavImage   *image.NRGBA
	wavImageOp paint.ImageOp

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
