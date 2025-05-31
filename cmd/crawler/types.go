package crawler

type RawPaper struct {
	Title	   string `json:"title"`
	Authors []string `json:"authors"`
	Venue 	 string `json:"venue"`
	Year 	   int    `json:"year"`
}

type ConferenceCrawler interface {
	Crawl() ([]RawPaper, error)
	Name() string
	Year() int
}