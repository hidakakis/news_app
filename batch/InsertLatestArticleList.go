package main

import (
	"database/sql"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/mmcdole/gofeed"
)

// SiteInfo is metainfomation of RSS Site
type SiteInfo struct {
	ID         int
	title      string
	rssURL     string
	latestDate string
}

// SiteRecord is site infomation on DB
type SiteRecord struct {
	title       string
	URL         string
	image       string
	publishedAt string
}

// Db is PostgreSQL Instance
var Db *sql.DB

func imageFromFeed(feed string) string {
	reader := strings.NewReader(feed)
	doc, _ := goquery.NewDocumentFromReader(reader)
	ImageURL, _ := doc.Find("img").Attr("src")
	return ImageURL
}

func initDB() {
	var err error
	Db, err = sql.Open("postgres", "user=postgres dbname=sample password=password sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func insertLatestArticleToDB(Article string) {

}

func getSiteInfoList() []SiteInfo {
	siteinfo := SiteInfo{}
	siteinfolist := []SiteInfo{}
	sql01_01 := "SELECT /* sql01_01 */ ID, title, rssURL, latestDate FROM site_tbl"

	getSiteInfoListFromDB, err := Db.Query(sql01_01)
	if err != nil {
		panic(err)
	}
	defer getSiteInfoListFromDB.Close()
	for getSiteInfoListFromDB.Next() {
		if err := getSiteInfoListFromDB.Scan(
			&siteinfo.ID,
			&siteinfo.title,
			&siteinfo.rssURL,
			&siteinfo.latestDate,
		); err != nil {
			panic(err)
		}
		newsiteinfo := SiteInfo{
			ID:         siteinfo.ID,
			title:      siteinfo.title,
			rssURL:     siteinfo.rssURL,
			latestDate: siteinfo.latestDate,
		}
		siteinfolist = append(siteinfolist, newsiteinfo)
	}
	return siteinfolist
}

func getUpdateDate() {

}

func main() {
	siteinfolist := getSiteInfoList()
	feedparser := gofeed.NewParser()
	feedArray := []SiteRecord{}
	for _, siteinfo := range siteinfolist {
		feed, _ := feedparser.ParseURL(siteinfo.rssURL)
		items := feed.Items
		for _, item := range items {
			feedmap := SiteRecord{
				title:       item.Title,
				URL:         item.Link,
				image:       imageFromFeed(item.Content),
				publishedAt: "tbd",
			}
			feedArray = append(feedArray, feedmap)
		}
	}
}
