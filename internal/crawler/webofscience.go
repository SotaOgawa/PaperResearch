package crawler

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"net/http"
	"net/url"
	"os"
	"paper-app-backend/internal/model"
	"strings"
)

type WebOfScienceConferenceCrawler struct {
}

func (c *WebOfScienceConferenceCrawler) Crawl(paper *model.PaperObjectInDB, db *gorm.DB) ([]model.PaperObjectInDB, error) {
	BASE_URL := "https://api.clarivate.com/apis/wos-starter/v1"

	req_url := fmt.Sprintf("%s/documents?db=WOS&q=TI=(%s)", BASE_URL, url.PathEscape(paper.Title))

	req, err := http.NewRequest("GET", req_url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	// 環境変数からapiキーを取得
	err = godotenv.Load(".env.local")

	if err != nil {
		return nil, fmt.Errorf("failed to load .env file: %w", err)
	}
	apiKey := os.Getenv("WOS_API_KEY")

	req.Header.Set("X-ApiKey", apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed with status: %s", resp.Status)
	}
	var response struct {
		Data []struct {
			Title string `json:"title"`
			Names struct {
				Authors []struct {
					Name string `json:"displayName"`
				} `json:"authors"`
			} `json:"names"`
		} `json:"hits"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	fmt.Printf("Found %d papers for title: %s\n", len(response.Data), paper.Title)
	var papers []model.PaperObjectInDB
	for _, item := range response.Data {
		authors := make([]string, len(item.Names.Authors))
		for i, author := range item.Names.Authors {
			authors[i] = author.Name
		}
		paperInDB := model.PaperObjectInDB{
			Title:      item.Title,
			Authors:    strings.Join(authors, ", "),
			Conference: paper.Conference,
			Year:       paper.Year,
			URL:        fmt.Sprintf("%s/documents/%s", BASE_URL, item.Title), // Assuming title is unique for URL
		}
		papers = append(papers, paperInDB)
	}

	fmt.Printf("Crawled %d papers from Web of Science for conference %s in year %d\n", len(papers), paper.Conference, paper.Year)

	return papers, nil
}

func NewWebOfScienceConferenceCrawler() *WebOfScienceConferenceCrawler {
	return &WebOfScienceConferenceCrawler{}
}
