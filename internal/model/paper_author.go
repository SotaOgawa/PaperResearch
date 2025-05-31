package model

type PaperAuthorConnection struct {
	PaperID     int `gorm:"primaryKey" json:"paper_id"`
	AuthorID    int `gorm:"primaryKey" json:"author_id"`
	AuthorOrder int `json:"author_order"`
}
