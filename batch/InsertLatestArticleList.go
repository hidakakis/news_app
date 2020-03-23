package main

import (
	"database/sql"
	"fmt"
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

// SiteRecord is article infomation for DB
type SiteRecord struct {
	title      string
	URL        string
	image      string
	updateDate string
	siteID     int
}

// EsRecord is article infomation for ElasticSearch
type EsRecord struct {
	ID    int
	title string
}

// Db is PostgreSQL Instance
var Db *sql.DB

func initDB() {
	var err error
	Db, err = sql.Open("postgres", "user=postgres dbname=sample password=password sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func getImageFromFeed(feed string) string {
	reader := strings.NewReader(feed)
	doc, _ := goquery.NewDocumentFromReader(reader)
	ImageURL, _ := doc.Find("img").Attr("src")
	return ImageURL
}

func registerLatestArticleToDB(articleList []SiteRecord) []EsRecord {
	sql01_02 := "INSERT INTO /* sql01_02 */ article_tbl (title, url, update_date, click, site_id) VALUES ($1, $2, $3, $4, $5) RETURNING (id, title)"
	stmt, err := Db.Prepare(sql01_02)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	esRecordList := []EsRecord{}
	var esRecord EsRecord
	for _, article := range articleList {
		err = stmt.QueryRow(article.title, article.URL, article.updateDate, 0, article.siteID).Scan(&esRecord.ID, &esRecord.title)
		esRecordList = append(esRecordList, esRecord)
	}
	return esRecordList
	/*
		--past python code--
		SQL_02_03 = 'SELECT max(id) FROM article_tbl'
		cur.execute(SQL_02_03)
		max_id = cur.fetchone()
		SQL_02_04 = 'INSERT INTO article_tbl (title, url, update_date, click, site_id) VALUES %s'
		extras.execute_values(cur, SQL_02_04, update_article_list)
		SQL_02_05 = 'SELECT id, title FROM article_tbl where id > %s'
		cur.execute(SQL_02_05, (max_id, ))
		update_es_list = cur.fetchall()
		save_elasticsearch(update_es_list)
		--past python code--
	*/
}

func registerLatestArticleToES(articleList []EsRecord) {
	for _, article := range articleList {
		fmt.Println("registerLatestArticleToES:", article.ID, article.title)
		// TBD
	}
}

func updateLatestDate(siteID int, updateDate string) {
	// TBD
}

func getSiteInfoList() []SiteInfo {
	siteinfo := SiteInfo{}
	siteinfolist := []SiteInfo{}
	sql01_01 := "SELECT /* sql01_01 */ ID, title, rssURL, latestDate FROM site_tbl"

	selectSiteInfoList, err := Db.Query(sql01_01)
	if err != nil {
		panic(err)
	}
	defer selectSiteInfoList.Close()
	for selectSiteInfoList.Next() {
		if err := selectSiteInfoList.Scan(
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

func main() {
	siteinfolist := getSiteInfoList()
	feedparser := gofeed.NewParser()
	feedArray := []SiteRecord{}
	var isVisit map[int]bool
	for _, siteinfo := range siteinfolist {
		isVisit[siteinfo.ID] = false
		feed, _ := feedparser.ParseURL(siteinfo.rssURL)
		items := feed.Items
		for _, item := range items {
			feedmap := SiteRecord{
				title:      item.Title,
				URL:        item.Link,
				image:      getImageFromFeed(item.Content),
				updateDate: item.Updated,
				siteID:     siteinfo.ID,
			}
			if feedmap.updateDate > siteinfo.latestDate {
				feedArray = append(feedArray, feedmap)
				if !isVisit[siteinfo.ID] {
					updateLatestDate(siteinfo.ID, feedmap.updateDate)
					isVisit[siteinfo.ID] = true
				}
			}
		}
	}
	esRecord := registerLatestArticleToDB(feedArray)
	registerLatestArticleToES(esRecord)
}
