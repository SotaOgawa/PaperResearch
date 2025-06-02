package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"paper-app-backend/internal/crawler"
	"paper-app-backend/internal/model"
	"strconv"
	// "paper-app-backend/internal/model"
)

func main() {
	req, _ := http.NewRequest("GET", "http://localhost:8080/api/papers", nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var result struct {
		Count int                     `json:"count"`
		Data  []model.PaperObjectInDB `json:"papers"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		fmt.Println("Error decoding response:", err)
		return
	}
	fmt.Println("Papers count:", len(result.Data))

	result1 := result.Data[0] //

	fmt.Println("First paper title:", result1.Title)

	crawler := crawler.NewSemanticScholarConferenceCrawler()
	papers, err := crawler.Crawl(&result1, nil)
	if err != nil {
		fmt.Println("Error during crawling:", err)
		return
	}
	for _, paper := range papers {
		paper.ID = result1.ID // Set the ID to the existing paper's ID
		jsonData, _ := json.Marshal(paper)
		req, _ := http.NewRequest("PUT", "http://localhost:8080/api/papers/"+strconv.Itoa(result1.ID), bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		defer resp.Body.Close()

		response_body, _ := io.ReadAll(resp.Body)
		fmt.Println("Response status:", resp.Status)
		fmt.Println("Response body:", string(response_body))
		if err != nil {

			fmt.Println("Error sending request:", err)
			return
		}
	}
}
