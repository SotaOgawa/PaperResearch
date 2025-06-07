package model

import "time"

type PaperObjectInDB struct {
	ID            int       `gorm:"primaryKey" json:"id"`
	Title         string    `gorm:"index:idx_unique_paper,unique" json:"title"`
	Conference    string    `gorm:"index:idx_unique_paper,unique" json:"conference"`
	Year          int       `json:"year"`
	Authors       string    `gorm:"index:idx_unique_paper,unique" json:"authors"` // Comma-separated list of authors
	Abstract      string    `json:"abstract"`
	URL           string    `json:"url"`
	CitationCount int       `json:"citation_count"`
	Bibtex        string    `json:"bibtex"`
	PDFURL        string    `json:"pdf_url"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
