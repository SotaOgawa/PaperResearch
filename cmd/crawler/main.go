package main

import (
	"fmt"
	"paper-app-backend/internal/conference"
	"strings"
)

func main() {
	crawlerICML := conference.NewICMLConferenceCrawler(2023)
	paperlist, err := crawlerICML.Crawl()
	if err != nil {
		fmt.Println("Etrror during crawling: %w", err)
		return
	}
	for _, paper := range paperlist {
		println("Title:", paper.Title)
		authors := strings.Join(paper.Authors, ", ")
		println("Authors:", authors)
		println("Venue:", paper.Venue)
		println("Year:", paper.Year)
	}
	println("Crawling completed successfully")
	println("Crawled", len(paperlist), "papers")
}
