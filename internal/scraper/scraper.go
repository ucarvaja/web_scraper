package scraper

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly/v2"
)

type VideoInfo struct {
	Title           string `json:"title"`
	ChannelName     string `json:"channel_name"`
	ThumbnailURL    string `json:"thumbnail_url"`
	Views           string `json:"views"`
	PublicationDate string `json:"publication_date"`
}

func EnterLink(filePath string) {

	var url string
	fmt.Print("Enter the YouTube video URL: ")
	fmt.Scan(&url)

	c := colly.NewCollector()

	videoInfo := VideoInfo{}

	c.OnHTML("meta[name='title']", func(e *colly.HTMLElement) {
		videoInfo.Title = e.Attr("content")
	})

	c.OnHTML("meta[itemprop='author']", func(e *colly.HTMLElement) {
		videoInfo.ChannelName = e.Attr("content")
	})

	c.OnHTML("link[itemprop='thumbnailUrl']", func(e *colly.HTMLElement) {
		videoInfo.ThumbnailURL = e.Attr("href")
	})

	c.OnHTML("meta[itemprop='interactionCount']", func(e *colly.HTMLElement) {
		videoInfo.Views = e.Attr("content")
	})

	c.OnHTML("meta[itemprop='datePublished']", func(e *colly.HTMLElement) {
		videoInfo.PublicationDate = e.Attr("content")
	})

	err := c.Visit(url)
	if err != nil {
		log.Fatalf("Error visiting URL: %v", err)
	}

	var videoInfos []VideoInfo
	file, err := os.Open(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			videoInfos = []VideoInfo{}
		} else {
			log.Fatalf("Error opening JSON file: %v", err)
		}
	} else {
		defer file.Close()
		decoder := json.NewDecoder(file)
		err = decoder.Decode(&videoInfos)
		if err != nil {

			if err.Error() == "json: cannot unmarshal object into Go value of type []scraper.VideoInfo" {
				videoInfos = []VideoInfo{}
			} else {
				log.Fatalf("Error decoding JSON: %v", err)
			}
		}
	}

	videoInfos = append(videoInfos, videoInfo)

	file, err = os.Create(filePath)
	if err != nil {
		log.Fatalf("Error creating JSON file: %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(videoInfos)
	if err != nil {
		log.Fatalf("Error encoding JSON: %v", err)
	}

	fmt.Println("Video information saved in", filePath)
}
