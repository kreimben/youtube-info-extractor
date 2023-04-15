package video

import (
	"testing"
)

func TestSearchVideoOne(t *testing.T) {
	videoCh := make(chan *Video, 1)
	go SearchOneVideoKeyword("위고 백지헌 직캠", videoCh)
	if v, ok := <-videoCh; ok {
		t.Log("No Playlist Video: ", v.Title)
	} else if !ok {
		t.Fatal("Error to get video")
		return
	}
}

func BenchmarkSearchVideoOne(b *testing.B) {
	videoCh := make(chan *Video, 1)
	for i := 0; i < b.N; i++ {
		go SearchOneVideoKeyword("위고 백지헌 직캠", videoCh)
		if v, ok := <-videoCh; ok {
			b.Log("No Playlist Video: ", v.Title)
		} else if !ok {
			b.Fatal("Error to get video")
			return
		}
	}
}
