package video

import (
	"encoding/json"
	"log"
	"os/exec"
)

var (
	YtDlpPath = "yt-dlp" // If you want to set other program or anything, overwrite it.
)

type Video struct {
	Id           string `json:"id"`
	Title        string `json:"title"`
	ThumbnailUrl string `json:"thumbnail"`
	Duration     uint32 `json:"duration"` // in seconds
	PlayUrl      string `json:"urls"`     // googlevideo.com domain.
}

// SearchOneVideoKeyword
// This is only for searching video on YouTube.
// Support for only one searching result.
func SearchOneVideoKeyword(keyword string, output chan *Video) {
	ytCommand := exec.Command(
		YtDlpPath,
		"--default-search=ytsearch",
		string("ytsearch1:"+keyword),
		"--skip-download",
		"--no-playlist",
		"--dump-json",
	)

	res, err := ytCommand.Output()
	if err != nil {
		log.Fatalln("Error during getting output: ", err)
	}

	video := &Video{}
	err = json.Unmarshal(res, video)
	if err != nil {
		log.Fatalln("Error on converting json: ", err)
	}
	output <- video
	close(output)
}
