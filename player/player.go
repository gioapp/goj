package player

import (
	"fmt"
)

type playerState int

const (
	Stopped playerState = iota
	Playing
	Paused
)

type selectCallback func(Track) (int, error)
type pauseCallback func(bool)
type seekCallback func(int) error
type volumeCallback func(int)

func (g *GoJoy) playSong(number int) {
	g.seek.Value = 0
	var err error
	g.songLen, err = g.OnSelect(g.Playlist.Tracks[number])
	if err == nil {
		g.state = Playing
		g.renderSong()
		g.renderStatus()
	}
}

func (g *GoJoy) renderSong() {
	if g.songSel != -1 {
		lyrics := g.Playlist.Tracks[g.songSel].Metadata["Lyrics"].(string)
		trackNum, _ := g.Playlist.Tracks[g.songSel].Metadata["Track"].(int)
		g.infoList = []string{
			"[Artist:](fg-green) " + g.Playlist.Tracks[g.songSel].Metadata["Artist"].(string),
			"[Title:](fg-green)  " + g.Playlist.Tracks[g.songSel].Metadata["Title"].(string),
			"[Album:](fg-green)  " + g.Playlist.Tracks[g.songSel].Metadata["Album"].(string),
			fmt.Sprintf("[Track:](fg-green)  %d", trackNum),
			"[Genre:](fg-green)  " + g.Playlist.Tracks[g.songSel].Metadata["Genre"].(string),
			fmt.Sprintf("[Year:](fg-green)   %d", g.Playlist.Tracks[g.songSel].Metadata["Year"].(string)),
		}
		if lyrics != "" {
			g.infoList = append(g.infoList, "Lyrics:  "+lyrics)
		}
	} else {
		g.infoList = []string{}
	}
}

func (g *GoJoy) renderStatus() {
	//var status string
	//switch g.state {
	//case Playing:
	//	status = "[(Playing)](fg-black,bg-green)"
	//case Paused:
	//	status = "[(Paused)](fg-black,bg-yellow)"
	//case Stopped:
	//	status = "[(Stopped)](fg-black,bg-red)"
	//}
	//g.scrollerGauge.BorderLabel = status

}

//Song selection

func (g *GoJoy) songDown() {
	if g.songSel < len(g.Playlist.Tracks)-1 {
		g.setSong(g.songSel+1, true)
	}
}

func (g *GoJoy) songUp() {
	if g.songSel > 0 {
		g.setSong(g.songSel-1, true)
	}
}

func (g *GoJoy) volumeUp() {
	if g.volume.Value < 100 {
		g.volume.Value += 5
	}
	//g.volumeGauge.Percent = g.volume
	g.OnVolume(g.volume.Value)

}

func (g *GoJoy) volumeDown() {
	if g.volume.Value > 0 {
		g.volume.Value -= 5
	}
	//g.volumeGauge.Percent = g.volume
	g.OnVolume(g.volume.Value)

}

func (g *GoJoy) setSong(num int, unset bool) {
	//skip := 0
	//for num-skip >= g.playList.Height-2 {
	//	skip += g.playList.Height - 2
	//}
	if unset {
		//g.Playlist.Tracks[g.songSel] = g.songNames[g.songSel][1 : len(g.songNames[g.songSel])-20]
	}
	g.songSel = num
	//g.songNames[num] = fmt.Sprintf("[%s](fg-black,bg-green)", g.songNames[num])
	//g.Playlist.Tracks = g.songNames[skip:]
}
