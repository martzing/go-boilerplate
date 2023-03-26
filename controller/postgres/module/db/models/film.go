package models

import "time"

type Film struct {
	FilmId      		int 			`gorm:"primarykey" json:"film_id"`
	Title 				string			`json:"title"`
	Description 		string			`json:"description"`
	ReleaseYear 		string			`json:"release_year"`
	LanguageId 			uint8			`json:"language_id"`
	RentalDuration		uint8			`json:"rental_duration"`
	RentalRate			float32			`json:"rental_rate"`
	Length				uint8			`json:"length"`
	ReplacementCost		float32			`json:"replacement_cost"`
	Rating				string			`json:"rating"`
	LastUpdate			time.Time 		`gorm:"column:last_update" json:"last_update"`
	Fulltext 			string			`json:"full_text"`
	FilmCategory		[]FilmCategory	`gorm:"foreignKey:FilmId" json:"film_category"`
}
