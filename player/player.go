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

type Player struct {
	Playlist *Playlist
	Playing  *Track
	infoList []string
	//scrollerGauge *termp.GaugeTheme
	//volumeGauge   *termp.GaugeTheme
	//controlsPar   *termp.ParagraphTheme

	songs     []Track
	songNames []string

	volume int

	songNum int

	songSel int
	songPos int
	songLen int

	OnSelect selectCallback
	OnPause  pauseCallback
	OnSeek   seekCallback
	OnVolume volumeCallback

	state playerState
}

func NewPlayer() *Player {

	p := new(Player)
	p.Playlist = LoadPlaylist()
	p.volume = 100

	//p.songs = songList
	p.songNum = -1

	//p.songNames = make([]string, len(p.songs))
	//for i, v := range p.songs {
	//	if v.Metadata != nil {
	//		p.songNames[i] = fmt.Sprintf("[%d] %s - %s", i+1, v.Metadata["Artist"].(string), v.Metadata["Title"].(string))
	//	} else {
	//		p.songNames[i] = fmt.Sprintf("[%d] %s", i+1, v.Path[pathPrefix:])
	//	}
	//}
	//p.Playlist.Tracks = p.songNames
	p.setSong(0, false)

	return p
}

func (p *Player) PlaySong(number int) {
	p.songPos = 0
	var err error
	p.songLen, err = p.OnSelect(p.Playlist.Tracks[number])
	if err == nil {
		p.state = Playing
		p.renderSong()
		p.renderStatus()
	}
}

func (p *Player) renderSong() {
	if p.songSel != -1 {
		lyrics := p.songs[p.songSel].Metadata["Lyrics"].(string)
		trackNum, _ := p.songs[p.songSel].Metadata["Track"].(int)
		p.infoList = []string{
			"[Artist:](fg-green) " + p.songs[p.songSel].Metadata["Artist"].(string),
			"[Title:](fg-green)  " + p.songs[p.songSel].Metadata["Title"].(string),
			"[Album:](fg-green)  " + p.songs[p.songSel].Metadata["Album"].(string),
			fmt.Sprintf("[Track:](fg-green)  %d", trackNum),
			"[Genre:](fg-green)  " + p.songs[p.songSel].Metadata["Genre"].(string),
			fmt.Sprintf("[Year:](fg-green)   %d", p.songs[p.songSel].Metadata["Year"].(string)),
		}
		if lyrics != "" {
			p.infoList = append(p.infoList, "Lyrics:  "+lyrics)
		}
	} else {
		p.infoList = []string{}
	}
}

func (p *Player) renderStatus() {
	//var status string
	//switch p.state {
	//case Playing:
	//	status = "[(Playing)](fg-black,bg-green)"
	//case Paused:
	//	status = "[(Paused)](fg-black,bg-yellow)"
	//case Stopped:
	//	status = "[(Stopped)](fg-black,bg-red)"
	//}
	//p.scrollerGauge.BorderLabel = status

}

//Song selection

func (p *Player) songDown() {
	if p.songSel < len(p.songNames)-1 {
		p.setSong(p.songSel+1, true)
	}
}

func (p *Player) songUp() {
	if p.songSel > 0 {
		p.setSong(p.songSel-1, true)
	}
}

func (p *Player) volumeUp() {
	if p.volume < 100 {
		p.volume += 5
	}
	//p.volumeGauge.Percent = p.volume
	p.OnVolume(p.volume)

}

func (p *Player) volumeDown() {
	if p.volume > 0 {
		p.volume -= 5
	}
	//p.volumeGauge.Percent = p.volume
	p.OnVolume(p.volume)

}

func (p *Player) setSong(num int, unset bool) {
	//skip := 0
	//for num-skip >= p.playList.Height-2 {
	//	skip += p.playList.Height - 2
	//}
	if unset {
		p.songNames[p.songSel] = p.songNames[p.songSel][1 : len(p.songNames[p.songSel])-20]
	}
	p.songSel = num
	//p.songNames[num] = fmt.Sprintf("[%s](fg-black,bg-green)", p.songNames[num])
	//p.Playlist.Tracks = p.songNames[skip:]
}
