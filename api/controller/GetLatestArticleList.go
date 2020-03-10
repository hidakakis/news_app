package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/mmcdole/gofeed"
)

func imageFromFeed(feed string) string {
	reader := strings.NewReader(feed)
	doc, _ := goquery.NewDocumentFromReader(reader)
	ImageURL, _ := doc.Find("img").Attr("src")
	return ImageURL
}

func main() {
	fp := gofeed.NewParser()
	rssUrls := []string{
		"http://www.vsnp.net/index.rdf",
		"http://blog.livedoor.jp/dqnplus/index.rdf"}
	var feedArray []map[string]interface{}

	for _, rssURL := range rssUrls {
		feed, _ := fp.ParseURL(rssURL)
		items := feed.Items
		for _, item := range items {
			feedmap := map[string]interface{}{"title": item.Title, "url": item.Link, "image": imageFromFeed(item.Content), "publishedAt": "tbd"}
			feedArray = append(feedArray, feedmap)
		}
	}
	feedJSONIndent, _ := json.MarshalIndent(feedArray, "", "   ")
	fmt.Println(string(feedJSONIndent))
}
