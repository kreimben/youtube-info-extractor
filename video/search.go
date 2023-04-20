package video

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os/exec"
	"strings"
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

// SearchVideoKeyword
// This is only for searching video on YouTube.
// Support for only one searching result.
// amount is the amount of searching result.
func SearchVideoKeyword(keyword string, amount int, output chan *Video) {
	if amount <= 0 {
		close(output)
		return
	}

	ytCommand := exec.Command(
		YtDlpPath,
		"--default-search=ytsearch",
		"ytsearch"+fmt.Sprint(amount)+":"+keyword,
		"--skip-download",
		"--no-playlist",
		"--dump-json",
	)

	res, err := ytCommand.Output()
	if err != nil {
		log.Fatalln("Error during getting output: ", err)
	}

	const detect = "\"repository\": \"yt-dlp/yt-dlp\"}}"
	arr := strings.Split(string(res), detect)
	arr = arr[:len(arr)-1]

	for _, raw := range arr {
		video := &Video{}
		err = json.Unmarshal([]byte(raw+detect), video)
		if err != nil {
			log.Fatalln("Error on converting json: ", err)
		}
		output <- video
	}
	close(output)
}

// SearchOneVideoUrl
// This is only for searching video on YouTube using url.
// If url is not valid, just close `output` channel.
// Not support for playlist link.
func SearchOneVideoUrl(rawurl string, output chan *Video) {
	_, err := url.ParseRequestURI(rawurl)
	if err != nil {
		close(output)
		return
	}

	ytCommand := exec.Command(
		YtDlpPath,
		rawurl,
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
