package model

import "time"

type PaperObjectInDB struct {
	ID            int       `gorm:"primaryKey" json:"id"`
	Title         string    `json:"title"`
	Conference    string    `json:"conference"`
	Year          int       `json:"year"`
	Abstract      string    `json:"abstract"`
	URL           string    `json:"url"`
	CitationCount int       `json:"citation_count"`
	Bibtex        string    `json:"bibtex"`
	PDFURL        string    `json:"pdf_url"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type PaperObjectWithAuthors struct {
	ID            int       `gorm:"primaryKey" json:"id"`
	Title         string    `json:"title"`
	Conference    string    `json:"conference"`
	Year          int       `json:"year"`
	Authors       []string  `json:"authors"`
	Abstract      string    `json:"abstract"`
	URL           string    `json:"url"`
	CitationCount int       `json:"citation_count"`
	Bibtex        string    `json:"bibtex"`
	PDFURL        string    `json:"pdf_url"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
