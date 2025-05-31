package conference

import (
	"paper-app-backend/cmd/crawler"
	"strconv"
)

type ACLConferenceCrawler struct {
	year int
}

func NewACLConferenceCrawler(year int) *ACLConferenceCrawler {
	return &ACLConferenceCrawler{year: year}
}

func (c *ACLConferenceCrawler) Year() int {
	return c.year
}
func (c *ACLConferenceCrawler) Name(year int) string {
	return "ACL Conference Crawler " + strconv.Itoa(year)
}

func (c *ACLConferenceCrawler) Crawl() ([]crawler.RawPaper, error) {
	// for test
	return []crawler.RawPaper{
		{
			Title:   "Sample Paper Title",
			Authors: []string{"Author One", "Author Two"},
			Venue:   "ACL Conference",
			Year:    2023,
		},
	}, nil
}