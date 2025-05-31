package handler

import (
	"gorm.io/gorm"
)

type PaperQuery struct {
	ID 			int   	`form:"ID"`
	Title      	string 	`form:"title"`
	Conference 	string 	`form:"conference"`
	Year       	int    	`form:"year"`
}

func (q *PaperQuery) Apply(db *gorm.DB) *gorm.DB {
	if q.ID != 0 {
		db = db.Where("id = ?", q.ID)
	}
	if q.Title != "" {
		db = db.Where("title LIKE ?", "%"+q.Title+"%")
	}
	if q.Conference != "" {
		db = db.Where("conference LIKE ?", "%"+q.Conference+"%")
	}
	if q.Year != 0 {
		db = db.Where("year = ?", q.Year)
	}
	return db
}