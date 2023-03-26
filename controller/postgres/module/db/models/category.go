package models

import "time"

type Category struct {
	CategoryId      uint8 		`gorm:"primarykey" json:"category_id"`
	Name      		string 		`json:"name"`
	LastUpdate		time.Time 	`gorm:"column:last_update" json:"last_update"`
}
