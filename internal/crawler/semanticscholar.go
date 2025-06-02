package crawler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"paper-app-backend/internal/model"
	"strings"

	"gorm.io/gorm"
	// "strings"
)

type SemanticScholarConferenceCrawler struct {
}

func (c *SemanticScholarConferenceCrawler) Crawl(paper *model.PaperObjectInDB, db *gorm.DB) ([]model.PaperObjectInDB, error) {
	BASE_URL := "https://api.semanticscholar.org/graph/v1/paper/search"

	req_url := fmt.Sprintf(`%s?query=%s&fields=title,year,venue,abstract,citationCount,authors,url,citationStyles&limit=1`, BASE_URL, url.QueryEscape(paper.Title))

	req, err := http.NewRequest("GET", req_url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(resp.Body) // Read the body to avoid resource leak
		return nil, fmt.Errorf("request failed with status: %s, error message: %s", resp.Status, string(b))
	}

	var response struct {
		Total  int `json:"total"`
		Offset int `json:"offset"`
		Next   int `json:"next"`
		Data   []struct {
			PaperID       string `json:"paperId"`
			URL           string `json:"url"`
			Title         string `json:"title"`
			Venue         string `json:"venue"`
			Year          int    `json:"year"`
			CitationCount int    `json:"citationCount"`
			OpenAccessPdf struct {
				URL        string `json:"url"`
				Status     any    `json:"status"`
				License    any    `json:"license"`
				Disclaimer string `json:"disclaimer"`
			} `json:"openAccessPdf"`
			CitationStyles struct {
				Bibtex string `json:"bibtex"`
			} `json:"citationStyles"`
			Authors []struct {
				AuthorID string `json:"authorId"`
				Name     string `json:"name"`
			} `json:"authors"`
			Abstract string `json:"abstract"`
		} `json:"data"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	var papers []model.PaperObjectInDB

	for i, paperData := range response.Data {
		// 取得した全てのデータを書き出す
		fmt.Printf("Paper %d: ID: %s, Title: %s, Venue: %s, Year: %d, Abstract: %s, CitationCount: %d, URL: %s, OpenAccessPDF: %s, Bibtex: %s\n",
			i+1, paperData.PaperID, paperData.Title, paperData.Venue, paperData.Year,
			paperData.Abstract, paperData.CitationCount, paperData.URL,
			paperData.OpenAccessPdf.URL, paperData.CitationStyles.Bibtex)
		authors := make([]string, len(paperData.Authors))
		for j, author := range paperData.Authors {
			authors[j] = author.Name
		}
		authors_joined := strings.Join(authors, ", ")
		paperInDB := model.PaperObjectInDB{
			Title:         paperData.Title,
			Conference:    paperData.Venue,
			Year:          paperData.Year,
			Authors:       authors_joined,
			Abstract:      paperData.Abstract,
			URL:           paperData.URL,
			CitationCount: paperData.CitationCount,
			Bibtex:        paperData.CitationStyles.Bibtex,
			PDFURL:        paperData.OpenAccessPdf.URL,
		}
		papers = append(papers, paperInDB)
	}

	return papers, nil
}

func NewSemanticScholarConferenceCrawler() *SemanticScholarConferenceCrawler {
	return &SemanticScholarConferenceCrawler{}
}
