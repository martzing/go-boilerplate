package mongoType

type ListCustomerParams struct {
  Search string
  Offset int64
  Limit  int64
}

type CreateCustomerParams struct {
  StoreId 		int			`json:"store_id"`
  FirstName 	string	`json:"first_name"`
  LastName 		string	`json:"last_name"`
  Email 			string	`json:"email"`
  AddressId 	int	    `json:"address_id"`
  Activebool 	bool		`json:"active_bool"`
}

type UpdateCustomerParams struct {
  ID         string    `json:"customer_id"`
  StoreId    int       `json:"store_id"`
  FirstName  string    `json:"first_name"`
  LastName   string    `json:"last_name"`
  Email      string    `json:"email"`
  AddressId  int       `json:"address_id"`
  Activebool *bool     `json:"active_bool"`
}

type DeleteCustomerParams struct {
  ID string             `json:"customer_id"`
}
