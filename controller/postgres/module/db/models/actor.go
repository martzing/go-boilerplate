package models

import "time"

type Actor struct {
	ActorId      	int `gorm:"primarykey"`
	FirstName 		string
	LastName 		string
	UpdatedAt 		time.Time `gorm:"column:last_update"`
}
