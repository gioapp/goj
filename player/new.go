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
		Theme: material.NewTheme(),
		//Playlist: play.New(),

		Layouts: layouts,
	}

	//userInterface, err := NewUi(songs, len(songDir))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//g.OnSelect = playSong
	//g.OnPause = pauseSong
	//g.OnSeek = seek
	//g.OnVolume = setVolue

	//g.songs = songList
	g.songNum = -1

	//g.songNames = make([]string, len(g.songs))
	//for i, v := range g.songs {
	//	if v.Metadata != nil {
	//		g.songNames[i] = fmt.Sprintf("[%d] %s - %s", i+1, v.Metadata["Artist"].(string), v.Metadata["Title"].(string))
	//	} else {
	//		g.songNames[i] = fmt.Sprintf("[%d] %s", i+1, v.Path[pathPrefix:])
	//	}
	//}
	//g.Playlist.Tracks = g.songNames
	g.setSong(0, false)

	g.Context = layout.NewContext(g.Window.Queue())
	g.Menu = g.MenuBar()
	g.Playlist = LoadPlaylist()

	g.seek = g.scrollerGauge(seek)
	g.volume = g.scrollerGauge(setVolue)
	g.volume.Value = 100
	//p, _ := NewPlayer(gojoy.Player.Playlist.Tracks, 0)
	//fmt.Println("tet",p)
	//gojoy.Player = p
	return g
}
