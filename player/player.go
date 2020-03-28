package player

type playerState int

const (
	Stopped playerState = iota
	Playing
	Paused
)

type selectCallback func(Track) (int, error)
type pauseCallback func(bool)
type seekCallback func(int)
type volumeCallback func(int)

func (g *GoJoy) playSong(number int) {
	g.seek.Value = 0
	var err error
	g.trackLen, err = g.OnSelect(g.Playlist.Tracks[number])
	if err == nil {
		g.state = Playing
		g.renderSong()
		g.renderStatus()
	}
}

func (g *GoJoy) renderSong() {
	if g.trackSel != -1 {
		//lyrics := g.Playlist.Tracks[g.trackSel].Metadata["Lyrics"].(string)
		//trackNum, _ := g.Playlist.Tracks[g.trackSel].Metadata["Track"].(int)
		//g.infoList = []string{
		//	"[Artist:](fg-green) " + g.Playlist.Tracks[g.trackSel].Metadata["Artist"].(string),
		//	"[Title:](fg-green)  " + g.Playlist.Tracks[g.trackSel].Metadata["Title"].(string),
		//	"[Album:](fg-green)  " + g.Playlist.Tracks[g.trackSel].Metadata["Album"].(string),
		//	fmt.Sprintf("[Track:](fg-green)  %d", trackNum),
		//	"[Genre:](fg-green)  " + g.Playlist.Tracks[g.trackSel].Metadata["Genre"].(string),
		//	fmt.Sprintf("[Year:](fg-green)   %d", g.Playlist.Tracks[g.trackSel].Metadata["Year"].(string)),
		//}
		//if lyrics != "" {
		//	g.infoList = append(g.infoList, "Lyrics:  "+lyrics)
		//}
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
	if g.trackSel < len(g.Playlist.Tracks)-1 {
		g.setSong(g.trackSel+1, true)
	}
}

func (g *GoJoy) songUp() {
	if g.trackSel > 0 {
		g.setSong(g.trackSel-1, true)
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
		//g.Playlist.Tracks[g.trackSel] = g.songNames[g.trackSel][1 : len(g.songNames[g.trackSel])-20]
	}
	g.trackSel = num
	//g.songNames[num] = fmt.Sprintf("[%s](fg-black,bg-green)", g.songNames[num])
	//g.Playlist.Tracks = g.songNames[skip:]
}
