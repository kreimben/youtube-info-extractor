package video

import (
	"bytes"
	"context"
	"log"
	"os/exec"
)

func SearchVideo(keyword string, amount uint8, isPlaylist bool) (string, error) {
	ytCommand, err := exec.CommandContext(context.Background(),
		"yt-dlp",
		"--use-extractors=\"youtube:search\"",
		string("ytsearch"+amount+":"+keyword),
		"--skip-download",
		"--dump-json",
	)
	if err != nil {
		log.Fatalln("Error while creating yt-dlp command: ", err)
	}

	// make a new empty buffer
	var out bytes.Buffer
	ytCommand.Stdout = &out

	if isPlaylist {
		ytCommand.Args = append(ytCommand.Args, "--yes-playlist")
	} else {
		ytCommand.Args = append(ytCommand.Args, "--no-playlist")
	}

	err = ytCommand.Run()
	if err != nil {
		return nil, err
	}

	if err = ytCommand.Wait(); err != nil {
		return nil, err
	}
}
