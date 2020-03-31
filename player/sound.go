package player

import (
	"fmt"
	"github.com/faiface/beep"
	"github.com/faiface/beep/flac"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
	"github.com/gioapp/goj/pkg/wavreader"
	"os"
	"path/filepath"
	"time"
)

func (g *GoJoy) playSong() (int, error) {
	g.seekElement.Value = 0
	var err error

	f, err := os.Open(g.Playing.Path)
	if err != nil {
	}
	defer f.Close()

	g.Playing.f = f
	wr, err := wavreader.New(f)
	if err != nil {
	}

	g.Playing.w = wr

	fmt.Println("sss", g.Playing.w)
	//g.Playing.processWav()
	g.trackLen, err = g.OnSelect(*g.Playing)
	fmt.Println("rPlaylistackNum", g.trackNum)
	if err == nil {
		g.state = Playing

		g.renderSong()
		g.renderStatus()
	}

	switch fileExt := filepath.Ext(g.Playing.Path); fileExt {
	case ".mp3":
		g.Sound.s, g.Sound.format, err = mp3.Decode(g.Playing.f)
	case ".wav":
		g.Sound.s, g.Sound.format, err = wav.Decode(g.Playing.f)
	case ".flac":
		g.Sound.s, g.Sound.format, err = flac.Decode(g.Playing.f)
	}
	if err != nil {
		return 0, err
	}
	g.Sound.volume.Streamer = g.Sound.s
	g.Sound.mainCtrl = &beep.Ctrl{Streamer: g.Sound.volume}
	speaker.Init(g.Sound.format.SampleRate, g.Sound.format.SampleRate.N(time.Second/10))
	speaker.Play(g.Sound.mainCtrl)
	return int(float32(g.Sound.s.Len()) / float32(g.Sound.format.SampleRate)), nil
}

func (g *GoJoy) pauseSong(state bool) {
	speaker.Lock()
	g.Sound.mainCtrl.Paused = state
	speaker.Unlock()
}

func (g *GoJoy) seek(pos int) {
	speaker.Lock()
	_ = g.Sound.s.Seek(pos * int(g.Sound.format.SampleRate))
	speaker.Unlock()
}

func (g *GoJoy) setVolue(percent int) {
	if percent == 0 {
		g.Sound.volume.Silent = true
	} else {
		g.Sound.volume.Silent = false
		g.Sound.volume.Volume = -float64(100-percent) / 100.0 * 5
	}
}
