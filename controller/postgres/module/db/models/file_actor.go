package models

import "time"

type FilmActor struct {
	ActorId      		uint8 `gorm:"primarykey"`
	FilmId      		uint8 `gorm:"primarykey"`
	LastUpdate			time.Time `gorm:"column:last_update"`
}
