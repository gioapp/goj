package player

import (
	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/layout"
	"gioui.org/unit"
	"github.com/faiface/beep/effects"
	"github.com/gioapp/gelook"
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
		Track: &layout.Flex{
			Axis: layout.Vertical,
		},
		TrackInfo: &layout.Flex{
			Axis: layout.Horizontal,
		},
	}
	g := &GoJoy{
		Window: app.NewWindow(
			app.Size(unit.Dp(400), unit.Dp(800)),
			app.Title("ParallelCoin"),
		),
		Theme: gelook.NewDuoUItheme(),
		//Playlist: play.New(),

		Layouts: layouts,
	}

	s := &Sound{
		supportedFormats: []string{".mp3", ".wav", ".flac"},
		volume: &effects.Volume{
			Base: 2,
		},
	}
	g.Sound = s
	//
	//g.OnSelect = playSong
	//g.OnPause = pauseSong
	//g.OnSeek = seek
	//g.OnVolume = setVolue

	//g.songs = songList
	g.trackNum = -1

	//g.songNames = make([]string, len(g.songs))
	//for i, v := range g.songs {
	//	if v.Metadata != nil {
	//		g.songNames[i] = fmt.Sprintf("[%d] %s - %s", i+1, v.Metadata["Artist"].(string), v.Metadata["Title"].(string))
	//	} else {
	//		g.songNames[i] = fmt.Sprintf("[%d] %s", i+1, v.Path[pathPrefix:])
	//	}
	//}
	//g.Playlist.Tracks = g.songNames

	g.Context = layout.NewContext(g.Window.Queue())
	g.Menu = g.MenuBar()

	g.Playlist = g.LoadPlaylist()

	g.seekElement = g.scrollerGauge(seek)
	g.seekElement.body.CursorHeight = 64
	g.volumeElement = g.scrollerGauge(setVolue)
	g.volumeElement.Value = 100

	g.setSong(0, false)
	//p, _ := NewPlayer(gojoy.Player.Playlist.Tracks, 0)
	//fmt.Println("tet",p)
	//gojoy.Player = p

	g.OnSelect = playSong
	g.OnPause = pauseSong
	g.OnSeek = seek
	g.OnVolume = setVolue
	//g.Start()
	//defer g.Close()

	return g
}
