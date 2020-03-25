package player

func Control(songList []Track, pathPrefix int) {

	//termp.Handle("/sys/kbd/p", func(termp.Event) {
	//	if p.songNum != -1 {
	//		if p.state == Playing {
	//			p.OnPause(true)
	//			p.state = Paused
	//		} else {
	//			p.OnPause(false)
	//			p.state = Playing
	//
	//		}
	//		p.renderStatus()
	//	}
	//})
	//termp.Handle("timer/1s", func(termp.Event) {
	//	if p.state == Playing {
	//		p.songPos++
	//		if p.songLen != 0 {
	//			p.scrollerGauge.Percent = int(float32(p.songPos) / float32(p.songLen) * 100)
	//			p.scrollerGauge.Label = fmt.Sprintf("%d:%.2d / %d:%.2d", p.songPos/60, p.songPos%60, p.songLen/60, p.songLen%60)
	//			if p.scrollerGauge.Percent >= 100 {
	//				p.songNum++
	//				if p.songNum >= len(p.songs) {
	//					p.songNum = 0
	//				}
	//				p.playSong(p.songNum)
	//			}
	//			termp.Clear()
	//			termp.Render(termp.Body)
	//		}
	//	} else if p.state == Stopped {
	//		p.songPos = 0
	//	}
	//})
	//
	//termp.Handle("/sys/kbd/<right>", func(termp.Event) {
	//	if p.songNum != -1 {
	//		p.songPos += 10
	//		p.OnSeek(p.songPos)
	//	}
	//})
	//
	//termp.Handle("/sys/kbd/<left>", func(termp.Event) {
	//	if p.songNum != -1 {
	//		p.songPos -= 10
	//		if p.songPos < 0 {
	//			p.songPos = 0
	//		}
	//		p.OnSeek(p.songPos)
	//	}
	//})
	//
	//termp.Handle("/sys/kbd/<escape>", func(termp.Event) {
	//	p.playSong(p.songNum)
	//	p.OnPause(true)
	//	p.state = Stopped
	//	p.scrollerGauge.Percent = 0
	//	p.scrollerGauge.Label = "0:00 / 0:00"
	//	p.renderStatus()
	//})
	//
	//termp.Handle("/sys/kbd/<enter>", func(termp.Event) {
	//	p.songNum = p.songSel
	//	p.playSong(p.songNum)
	//})
	//
	//termp.Handle("/sys/kbd/<up>", func(termp.Event) {
	//	p.songUp()
	//})
	//
	//termp.Handle("/sys/kbd/=", func(termp.Event) {
	//	p.volumeUp()
	//})
	//
	//termp.Handle("/sys/kbd/+", func(termp.Event) {
	//	p.volumeUp()
	//})
	//
	//termp.Handle("/sys/kbd/-", func(termp.Event) {
	//	p.volumeDown()
	//})
	//
	//termp.Handle("/sys/kbd/_", func(termp.Event) {
	//	p.volumeDown()
	//})
	//
	//termp.Handle("/sys/kbd/<down>", func(termp.Event) {
	//	p.songDown()
	//})
	//
	//termp.Handle("/sys/wnd/resize", func(termp.Event) {
	//	p.realign()
	//})
}
