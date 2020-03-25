package player

import (
	"fmt"
	"gioui.org/widget"
	"github.com/dhowden/tag"
	"github.com/mitchellh/go-homedir"
	"log"
	"os"
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
	tracks := make([]Track, 0, len(fileList))
	buttons := make(map[string]*widget.Button)

	for _, fileName := range fileList {
		currentFile, err := os.Open(fileName)
		if err == nil {
			metadata, _ := tag.ReadFrom(currentFile)
			fmt.Println("meta:", metadata)
			track := Track{
				Metadata: make(map[string]interface{}),
				//Image:metadata.Raw()[""].(string),
				//Path:metadata.Raw()[""].(string),
				Path: fileName,
			}
			if metadata != nil {
				track.Metadata = metadata.Raw()
			}

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

			tracks = append(tracks, track)
			buttons[fileName] = new(widget.Button)
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
	//userInterface, err := NewUi(songs, len(songDir))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//userInterface.OnSelect = playSong
	//userInterface.OnPause = pauseSong
	//userInterface.OnSeek = seek
	//userInterface.OnVolume = setVolue
}
