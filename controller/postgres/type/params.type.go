package postgresType

type ListCustomerParams struct {
  Search string
  Offset int
  Limit  int
}

type ListFilmParams struct {
  Search string
  Offset int
  Limit  int
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
  CustomerId int       `json:"customer_id"`
  StoreId    int       `json:"store_id"`
  FirstName  string    `json:"first_name"`
  LastName   string    `json:"last_name"`
  Email      string    `json:"email"`
  AddressId  int       `json:"address_id"`
  Activebool *bool      `json:"active_bool"`
}

type DeleteCustomerParams struct {
  CustomerId int       `json:"customer_id"`
}

type CreateDemoQueueParams struct {
  Name string
}

type CreateDemoJobParams struct {
  Data string
}

