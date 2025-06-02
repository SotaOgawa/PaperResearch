package crawler

import (
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"io"
	"net/http"
	"net/url"
	"paper-app-backend/internal/model"
	// "strings"
)

type SemanticScholarConferenceCrawler struct {
}

func (c *SemanticScholarConferenceCrawler) Crawl(paper *model.PaperObjectInDB, db *gorm.DB) ([]model.PaperObjectInDB, error) {
	BASE_URL := "https://api.semanticscholar.org/graph/v1/paper/search"

	req_url := fmt.Sprintf("%s?query=%s&fields=title,year,venue,abstract,referenceCount,citationCount,influentialCitationCount,isOpenAccess,fieldsOfStudy,authors&limit=10", BASE_URL, url.QueryEscape(paper.Title))

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
			PaperID                  string   `json:"paperId"`
			Title                    string   `json:"title"`
			Venue                    string   `json:"venue"`
			Abstract                 string   `json:"abstract"`
			Year                     int      `json:"year"`
			ReferenceCount           int      `json:"referenceCount"`
			CitationCount            int      `json:"citationCount"`
			InfluentialCitationCount int      `json:"influentialCitationCount"`
			IsOpenAccess             bool     `json:"isOpenAccess"`
			FieldsOfStudy            []string `json:"fieldsOfStudy"`
			Authors                  []struct {
				AuthorID string `json:"authorId"`
				Name     string `json:"name"`
			} `json:"authors"`
		} `json:"data"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	for i, paperData := range response.Data {
		// 取得した全てのデータを書き出す
		fmt.Printf("Paper %d: ID: %s, Title: %s, Venue: %s, Year: %d, Abstract: %s, ReferenceCount: %d, CitationCount: %d, InfluentialCitationCount: %d, IsOpenAccess: %t, FieldsOfStudy: %v\n",
			i, paperData.PaperID, paperData.Title, paperData.Venue, paperData.Year, paperData.Abstract,
			paperData.ReferenceCount, paperData.CitationCount, paperData.InfluentialCitationCount, paperData.IsOpenAccess, paperData.FieldsOfStudy)
	}

	return nil, nil
}

func NewSemanticScholarConferenceCrawler() *SemanticScholarConferenceCrawler {
	return &SemanticScholarConferenceCrawler{}
}
