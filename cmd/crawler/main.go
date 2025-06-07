package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"paper-app-backend/internal/conference"
	"paper-app-backend/internal/crawler"
	"paper-app-backend/internal/model"
	"strings"
	"time"
)

func main() {
	crawlerICML := conference.NewICMLConferenceCrawler(2024)
	crawlerSemanticScholar := crawler.NewSemanticScholarConferenceCrawler()
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

		rawpaper := model.PaperObjectInDB{
			Title:      paper_title,
			Authors:    paper_author_joined,
			Conference: paper_Conference,
			Year:       paper_year,
		}

		fmt.Printf("Crawled paper: %s, Authors: %s, Conference: %s, Year: %d\n",
			rawpaper.Title, rawpaper.Authors, rawpaper.Conference, rawpaper.Year)

		// Semantic Scholar APIにPOSTリクエストを送信
		paperInDB, err := crawlerSemanticScholar.Crawl(&rawpaper, nil)
		if err != nil {
			fmt.Printf("Error during crawling Semantic Scholar for paper %s: %v\n", paper_title, err)
			continue
		}

		jsonBytes, _ := json.Marshal(paperInDB)

		//jsonBytesの内容を確認
		fmt.Printf("JSON Payload: %s\n", jsonBytes)

		req, _ := http.NewRequest("POST", `https://paperresearch-production.up.railway.app/api/papers`, bytes.NewBuffer(jsonBytes))
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("Error sending request for paper %s: %v\n", paper_title, err)
			continue
		}
		if resp.StatusCode != http.StatusOK {
			fmt.Printf("Failed to add paper %s, status code: %d\n", paper_title, resp.StatusCode)
			// レスポンスボディを読み取る
			responseBody, _ := io.ReadAll(resp.Body)
			fmt.Printf("Response body: %s\n", responseBody)
		}
		defer resp.Body.Close()

		time.Sleep(2 * time.Second) // 1秒のスリープを追加
	}
}
