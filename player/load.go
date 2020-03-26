package player

import (
	"gioui.org/widget"
	"github.com/dhowden/tag"
	"github.com/mitchellh/go-homedir"
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
			metadata, _ := tag.ReadFrom(currentFile)
			track := Track{
				Metadata: make(map[string]interface{}),
				//Image:metadata.Raw()[""].(string),
				//Path:metadata.Raw()[""].(string),
				Path: fileName,
			}
			track.Id = trackNum
			if metadata != nil {
				track.Metadata = metadata.Raw()
			}

			track.Filename = filepath.Base(track.Path)

			if track.Metadata["TPE1"] != nil {
				track.Artist = track.Metadata["TPE1"].(string)
			}
			if track.Metadata["TIT2"] != nil {
				track.Title = track.Metadata["TIT2"].(string)
			}
			if track.Metadata["TALB"] != nil {
				track.Album = track.Metadata["TALB"].(string)
			}
			if track.Metadata["TRCK"] != nil {
				track.Track = track.Metadata["TRCK"].(string)
			}
			if track.Metadata["TCON"] != nil {
				track.Genre = track.Metadata["TCON"].(string)
			}
			if track.Metadata["TYER"] != nil {
				track.Year = track.Metadata["TYER"].(string)
			}

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
