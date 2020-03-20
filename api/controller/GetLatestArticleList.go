package main

import (
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/labstack/echo"
	"github.com/mmcdole/gofeed"
)

func imageFromFeed(feed string) string {
	reader := strings.NewReader(feed)
	doc, _ := goquery.NewDocumentFromReader(reader)
	ImageURL, _ := doc.Find("img").Attr("src")
	return ImageURL
}

func main() {
	e := echo.New()
	e.GET("/", feedFactory)
	e.Logger.Fatal(e.Start(":8770"))
}

// feedFactory is factory of feed json
func feedFactory(c echo.Context) error {
	feedparser := gofeed.NewParser()
	rssUrls := []string{
		"http://www.vsnp.net/index.rdf",
		"http://blog.livedoor.jp/dqnplus/index.rdf",
	}
	var feedArray []map[string]interface{}

	for _, rssURL := range rssUrls {
		feed, _ := feedparser.ParseURL(rssURL)
		items := feed.Items
		for _, item := range items {
			feedmap := map[string]interface{}{
				"title":       item.Title,
				"url":         item.Link,
				"image":       imageFromFeed(item.Content),
				"publishedAt": "tbd",
			}
			feedArray = append(feedArray, feedmap)
		}
	}
	return c.JSON(http.StatusOK, feedArray)
}
