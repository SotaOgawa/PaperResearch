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

type OpenReviewResponseNote struct {
	Content struct {
		Title struct {
			Value string `json:"value"`
		} `json:"title"`
		Venue struct {
			Value string `json:"value"`
		} `json:"venue"`
		Authors struct {
			Value []string `json:"value"`
		} `json:"authors"`
		Authorids struct {
			Value []string `json:"value"`
		} `json:"authorids"`
		Venueid struct {
			Value string `json:"value"`
		} `json:"venueid"`
		Bibtex struct {
			Value string `json:"value"`
		} `json:"_bibtex"`
		HTML struct {
			Value string `json:"value"`
		} `json:"html"`
		Abstract struct {
			Value string `json:"value"`
		} `json:"abstract"`
		Pdf struct {
			Value string `json:"value"`
		} `json:"pdf"`
		Paperhash struct {
			Value string `json:"value"`
		} `json:"paperhash"`
	} `json:"content"`
	ID          string   `json:"id"`
	Forum       string   `json:"forum"`
	License     string   `json:"license"`
	Signatures  []string `json:"signatures"`
	Readers     []string `json:"readers"`
	Writers     []string `json:"writers"`
	Number      int      `json:"number"`
	Invitations []string `json:"invitations"`
	Domain      string   `json:"domain"`
	Tcdate      int64    `json:"tcdate"`
	Cdate       int64    `json:"cdate"`
	Tmdate      int64    `json:"tmdate"`
	Mdate       int64    `json:"mdate"`
	Pdate       int64    `json:"pdate"`
	Version     int      `json:"version"`
}

type OpenReviewResponse struct {
	Notes []OpenReviewResponseNote `json:"notes"`
	Count int `json:"count"`
}