package main

import (
	"github.com/kreimben/youtube-info-extractor/src/video"
	"log"
)

func main() {
	v, err := video.SearchVideo("위고 백지헌 직캠", 1, false)
	if err != nil {
		log.Println("Error while searching video: ", err)
		return
	}
	log.Println("No Playlist Video: ", v)

	v, err = video.SearchVideo("위고 백지헌 직캠", 1, true)
	if err != nil {
		return
	}
	log.Println("Yes Playlist Video: ", v)
}
