package model

type AuthorObjectInDB struct {
	ID    int    `gorm:"primaryKey" json:"id"`
	Name  string `json:"name"`
	ORCID string `json:"orcid"`
}
