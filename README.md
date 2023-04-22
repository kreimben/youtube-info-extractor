# youtube-info-extractor
YouTube video info extractor module written in Go.

## Installation
```
go get github.com/kreimben/youtube-info-extractor
```

## Example
```go
import ytex "github.com/kreimben/youtube-info-extractor"

func main() {
    ytex.YtDlpPath = "/opt/homebrew/bin/yt-dlp" // can overwrite or just use it.
    videoCh := make(chan *video.Video) // make channel to get a results.
    for {
        go video.SearchOneVideoKeyword("keyword you want to search like using youtube", videoCh)
        if v, ok := <-videoCh; ok {
            log.Println("video: ", v)
        } else if !ok {
            break
        }
    }
}
```
