package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"paper-app-backend/internal/conference"
	"paper-app-backend/internal/crawler"
	"strings"
)

func main() {
	crawlerICML := conference.NewICMLConferenceCrawler(2023)
	paperlist, err := crawlerICML.Crawl()
	if err != nil {
		fmt.Println("Error during crawling: %w", err)
		return
	}
	for _, paper := range paperlist {
		paper_title := paper.Title
		paper_authors := paper.Authors
		paper_Conference := paper.Venue
		paper_year := paper.Year

		paper_author_joined := strings.Join(paper_authors, ", ")

		paperInDB := crawler.RawPaperInDB{
			Title:      paper_title,
			Authors:    paper_author_joined,
			Conference: paper_Conference,
			Year:       paper_year,
		}

		fmt.Printf("Crawled paper: %s, Authors: %s, Conference: %s, Year: %d\n",
			paperInDB.Title, paperInDB.Authors, paperInDB.Conference, paperInDB.Year)

		jsonBytes, _ := json.Marshal(paperInDB)

		//jsonBytesの内容を確認
		fmt.Printf("JSON Payload: %s\n", jsonBytes)

		req, _ := http.NewRequest("POST", `http://localhost:8080/api/papers`, bytes.NewBuffer(jsonBytes))
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("Error sending request for paper %s: %v\n", paper_title, err)
			continue
		}
		defer resp.Body.Close()
	}
}
