package player

import (
	"bytes"
	"gioui.org/widget"
	"github.com/bogem/id3v2"
	"github.com/mitchellh/go-homedir"
	"image/jpeg"
	"log"
	"os"
	"path/filepath"
)

func LoadPlaylist() *Playlist {
	var songDir string
	var err error
	if len(os.Args) > 1 {
		songDir = os.Args[1]
	} else {
		songDir, err = homedir.Expand("~/Music/")
		if err != nil {
			log.Fatal("Can't open ~/Music directory")
		}
	}
	fileList, err := getSongList(songDir)
	if err != nil {
		log.Fatal("Can't get song list")
	}
	tracks := make(map[int]Track)
	buttons := make(map[int]*widget.Button)

	for trackNum, fileName := range fileList {
		currentFile, err := os.Open(fileName)
		if err == nil {
			metadata, err := id3v2.ParseReader(currentFile, id3v2.Options{Parse: true})
			if err != nil {
			}
			track := Track{
				//Metadata: make(map[string]interface{}),
				//Image:metadata.Raw()[""].(string),
				//Path:metadata.Raw()[""].(string),
				Path: fileName,
			}
			track.Id = trackNum
			if metadata != nil {
				//track.Metadata = metadata.Raw()
			}
			track.Filename = filepath.Base(track.Path)
			pictures := metadata.GetFrames(metadata.CommonID("Attached picture"))
			for _, f := range pictures {
				pic, ok := f.(id3v2.PictureFrame)
				if !ok {
					log.Fatal("Couldn't assert picture frame")
				}
				loadedImage, err := jpeg.Decode(bytes.NewReader(pic.Picture))
				if err != nil {
					// Handle error
				}
				track.CoverImage = loadedImage
			}
			track.Artist = metadata.Artist()
			track.Title = metadata.Title()
			track.Album = metadata.Album()
			track.Genre = metadata.Genre()
			track.Year = metadata.Year()

			tracks[track.Id] = track
			buttons[track.Id] = new(widget.Button)
		}
		currentFile.Close()
	}
	if len(tracks) == 0 {
		log.Fatal("Could find any songs to play")
	}

	return &Playlist{
		Buttons: buttons,
		Tracks:  tracks,
	}
}
