package youporn

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

const apiURL = "http://www.youporn.com/api/webmasters/"
const APITimeout = 5

func SearchVideos(search string) YoupornSearchResult {
	timeout := time.Duration(APITimeout * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	resp, _ := client.Get(fmt.Sprintf(apiURL+"search?search=%s&thumbsize=all", url.QueryEscape(search)))
	b, _ := ioutil.ReadAll(resp.Body)
	var result YoupornSearchResult
	err := json.Unmarshal(b, &result)
	if err != nil {
		log.Println(err)
	}
	return result
}

func GetVideoByID(ID string) YoupornSingleVideo {
	timeout := time.Duration(APITimeout * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	resp, _ := client.Get(fmt.Sprintf(apiURL+"video_by_id/?video_id=%s", ID))
	b, _ := ioutil.ReadAll(resp.Body)
	var result YoupornSingleVideo
	err := json.Unmarshal(b, &result)
	if err != nil {
		log.Println(err)
	}
	return result
}

func GetVideoEmbedCode(ID string) YoupornEmbedCode {
	timeout := time.Duration(APITimeout * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	resp, _ := client.Get(fmt.Sprintf(apiURL+"video_embed_code/?video_id=%s", ID))
	b, _ := ioutil.ReadAll(resp.Body)
	var result YoupornEmbedCode
	err := json.Unmarshal(b, &result)
	if err != nil {
		log.Println(err)
	}
	return result
}
