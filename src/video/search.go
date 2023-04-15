package video

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type options struct {
	Type  string
	Limit uint32
}

func SearchVideo(keyword string) {
	// results := make(any, 0)
	// details := make(any, 0)
	// fetched := false
	// ops := options{Type: "video", Limit: 0}

	url := fmt.Sprintf("https://youtube.com/results?q=%s&hl=kr", url.QueryEscape(keyword))
	res, err := http.Get(url)
	if err != nil {
		log.Fatalln("get: ", err)
	}
	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalln("read all: ", err)
	}
	body := string(b)
	data := strings.Split(body, "ytInitialData =")[1]
	data = strings.Split(data, ";</script>")[0]
	// log.Println("Second data: ", data)

	ioutil.WriteFile("data.json", []byte(data), 0644)
}
