package main

import (
	"encoding/json"
	"fmt"
	"net/http"
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

	result := make([]map[string]interface{}, 0)
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		fmt.Println("Error decoding response:", err)
		return
	}
	fmt.Println("Papers count:", len(result))
	for i, paper := range result {
		fmt.Printf("Paper %d: ID=%v, Title=%v, Authors=%v, Venue=%v, Year=%v\n",
			i+1,
			paper["id"],
			paper["title"],
			paper["authors"],
			paper["venue"],
			paper["year"])

		if i >= 10 { // 最初の10件だけ表示
			break
		}
	}
}
