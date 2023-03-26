package models

import "time"

type Customer struct {
	CustomerId      int 		`gorm:"primarykey" json:"customer_id"`
	StoreId 		int			`json:"store_id"`
	FirstName 		string		`json:"first_name"`
	LastName 		string		`json:"last_name"`
	Email 			string		`json:"email"`
	AddressId 		int			`json:"address_id"`
	Activebool 		bool		`json:"active_bool"`
	CreateDate 		time.Time	`json:"created_date" gorm:"autoCreateTime"`
	UpdatedAt 		time.Time 	`gorm:"column:last_update;autoCreateTime" json:"update_at"`
}
