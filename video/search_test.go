package video

import (
	"testing"
)

func TestSearchVideoOne(t *testing.T) {
	videoCh := make(chan *Video, 1)
	go SearchOneVideoKeyword("위고 백지헌 직캠", 15, videoCh)
	for {
		if v, ok := <-videoCh; ok {
			t.Log("No Playlist Video: ", v.Title)
		} else if !ok {
			t.Fatal("Channel closed.")
			return
		}
	}
}

func BenchmarkSearchVideoOne(b *testing.B) {
	videoCh := make(chan *Video, 1)
	for i := 0; i < b.N; i++ {
		go SearchOneVideoKeyword("위고 백지헌 직캠", 3, videoCh)
		if v, ok := <-videoCh; ok {
			b.Log("No Playlist Video: ", v.Title)
		} else if !ok {
			b.Fatal("Error to get video")
			return
		}
	}
}

func TestSearchOneVideoUrlSuccess(t *testing.T) {
	videoCh := make(chan *Video, 1)
	go SearchOneVideoUrl("https://www.youtube.com/watch?v=HM6UpQZvbhY", videoCh)
	if v, ok := <-videoCh; ok {
		t.Log("No Playlist Video: ", v.Title)
	} else if !ok {
		t.Fatal("Error to get video")
		return
	}
}

func TestSearchOneVideoUrlFail(t *testing.T) {
	videoCh := make(chan *Video, 1)
	go SearchOneVideoUrl("fromis_9 we go m/v", videoCh)
	if v, ok := <-videoCh; ok {
		t.Fatal("No Playlist Video: ", v)
	} else if !ok {
		t.Log("Error to get video => Success to test.")
		return
	}
}
