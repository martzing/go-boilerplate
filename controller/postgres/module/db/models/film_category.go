package models

import "time"

type FilmCategory struct {
	FilmId      	int			`gorm:"primarykey" json:"film_id"`
	CategoryId 		uint8 		`gorm:"primarykey" json:"category_id"`
	LastUpdate		time.Time 	`gorm:"column:last_update" json:"last_update"`
	Category		Category	`gorm:"foreignKey:CategoryId" json:"category"`
}


