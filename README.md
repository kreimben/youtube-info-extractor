# youtube-info-extractor
YouTube video info extractor module written in Go.

## Installation
```
go get github.com/kreimben/youtube-info-extractor
```

## Example
```go
video.YtDlpPath = "/opt/homebrew/bin/yt-dlp" // can overwrite

videoCh := make(chan *video.Video)
for {
    go video.SearchOneVideoKeyword("위고 백지헌 직캠", videoCh)
    if v, ok := <-videoCh; ok {
        log.Println("No Playlist Video: ", v)
    } else if !ok {
        log.Println("EOF")
        break
    }
}
```