package player

import (
	"fmt"
	"gioui.org/layout"
	"github.com/gioapp/gel"
)

var (
	tracksListPanelElement = gel.NewPanel()
)

func (g *GoJoy) TrackList() func() {
	return func() {
		if g.Playlist.Tracks != nil {

			tracksListPanelElement.PanelObject = g.Playlist.Tracks
			tracksListPanelElement.PanelObjectsNumber = len(g.Playlist.Tracks)
			tracksListPanel := g.Theme.DuoUIpanel()
			tracksListPanel.ScrollBar = g.Theme.ScrollBar(0)

			tracksListPanel.Layout(g.Context, tracksListPanelElement, func(i int, in interface{}) {
				//g.Layouts.Playlist.Layout(g.Context, len(g.Playlist.Tracks), func(i int) {
				track := g.Playlist.Tracks[i]
				for g.Playlist.Buttons[track.Id].Clicked(g.Context) {
					g.Playing = &track
					//g.setSong(g.trackNum, true)
					//g.trackNum = g.trackSel

					g.playSong()
				}
				b := g.Theme.DuoUIbutton("", "", "", "ff888888", "", "", "", "", 0, 0, 0, 0, 0, 0, 0, 0)
				b.InsideLayout(g.Context, g.Playlist.Buttons[track.Id], func() {
					layout.Flex{Axis: layout.Vertical}.Layout(g.Context,
						layout.Rigid(func() {
							layout.Flex{Spacing: layout.SpaceBetween}.Layout(g.Context,
								layout.Rigid(func() {
									g.Theme.DuoUIcontainer(2, g.Theme.Colors["Primary"]).Layout(g.Context, layout.Center, func() {
										g.Theme.Body1(fmt.Sprint(track.Id)).Layout(g.Context)
									})
								}),
								layout.Flexed(1, func() {
									name := g.Theme.Body1(track.Filename)
									name.Font.Typeface = g.Theme.Fonts["Primary"]
									name.Layout(g.Context)
								}))
						}),
						layout.Rigid(g.Theme.DuoUIline(g.Context, 1, 0, 1, g.Theme.Colors["Dark"])),
					)
				})
			})
		}
	}
}
