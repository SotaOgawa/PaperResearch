package conference

import (
	// "bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"paper-app-backend/internal/crawler"
	"strconv"
)

type ICMLConferenceCrawler struct {
	year int
}

func NewICMLConferenceCrawler(year int) *ICMLConferenceCrawler {
	return &ICMLConferenceCrawler{year: year}
}

func (c *ICMLConferenceCrawler) Year() int {
	return c.year
}
func (c *ICMLConferenceCrawler) Name(year int) string {
	return "ICML Conference Crawler " + strconv.Itoa(year)
}

func (c *ICMLConferenceCrawler) Crawl() ([]crawler.RawPaper, error) {
	limit := 1000 // OpenReview APIの仕様により、最大1000件まで取得可能
	offset := 0 // オフセットの初期値
	
	var papers []crawler.RawPaper // 取得した論文のリスト

	for{
	resp, err := http.Get("https://api2.openreview.net/notes?content.venueid=ICLR.cc%2F2024%2FConference&offset=" + strconv.Itoa(offset))

	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed with status: %s", resp.Status)
	}

	fmt.Println("Request created successfully")

	defer resp.Body.Close()

	// まずGoで JSONをデコードする
	var result crawler.OpenReviewResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	println("Response decoded successfully, count:", result.Count)

	// 次にRawPaperへと変換する
	for _, note := range result.Notes {
		p := crawler.RawPaper{
			Title:   note.Content.Title.Value,
			Authors: note.Content.Authors.Value,
			Venue:   note.Content.Venue.Value,
			Year:    c.Year(),                 // crawlerに保持されてる年
		}
		papers = append(papers, p)
	}

	if len(result.Notes) < limit {break} // 取得件数がlimit未満なら終了


	// var result OpenReviewResponse
	// if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
	// 	return nil, fmt.Errorf("failed to decode response: %w", err)
	// }
	// var papers []crawler.RawPaper
	// for _, note := range result.Notes {
	// 	p := crawler.RawPaper{
	// 		Title:   note.Content.Title,
	// 		Authors: note.Content.Authors,
	// 		Venue:   "ICLR", // 呼び出し元から渡す
	// 		Year:    2025,
	// 	}
	// 	papers = append(papers, p)
	// }
	// return papers, nil
	offset += limit // オフセットを更新
}

	return papers, nil
}