package main

import (
	"database/sql"
)

// SiteInfo is metainfomation of RSS Site
type SiteInfo struct {
	ID         int
	title      string
	rssURL     string
	latestDate string
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

func insertLatestArticleToDB(Article string) {

}

func getSiteInfoList(updateDate string) []SiteInfo {
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
}
