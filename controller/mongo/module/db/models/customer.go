package models

import (
	"time"
)

type Customer struct {
  ID          string      `json:"id" bson:"_id"`
  StoreId 		int			    `json:"store_id"`
  FirstName 	string		  `json:"first_name"`
  LastName 		string		  `json:"last_name"`
  Email 			string		  `json:"email"`
  AddressId 	int			    `json:"address_id"`
  Activebool 	bool		    `json:"active_bool"`
  CreatedAt   time.Time   `json:"created_at,omitempty" bson:"created_at"`
  UpdatedAt   time.Time   `json:"updated_at,omitempty" bson:"updated_at"`
}

type CustomerCreate struct {
  StoreId 		int			    `json:"store_id"`
  FirstName 	string		  `json:"first_name"`
  LastName 		string		  `json:"last_name"`
  Email 			string		  `json:"email"`
  AddressId 	int			    `json:"address_id"`
  Activebool 	bool		    `json:"active_bool"`
  CreatedAt   time.Time   `json:"created_at,omitempty" bson:"created_at"`
  UpdatedAt   time.Time   `json:"updated_at,omitempty" bson:"updated_at"`
}
