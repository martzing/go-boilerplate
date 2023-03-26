package postgresType

type ListCustomerValidate struct {
  Search string `query:"search" validate:"omitempty"`
  Offset string `query:"offset" validate:"omitempty,number,min=1"`
  Limit  string `query:"limit" validate:"omitempty,number,min=1"`
}

type ListFilmValidate struct {
  Search string `query:"search" validate:"omitempty"`
  Offset string `query:"offset" validate:"omitempty,number,min=1"`
  Limit  string `query:"limit" validate:"omitempty,number,min=1"`
}

type CreateCustomerValidate struct {
  StoreId 		int			`json:"store_id" form:"store_id" validate:"number,min=1,required"`
  FirstName 	string	`json:"first_name" form:"first_name" validate:"required"`
  LastName 		string	`json:"last_name" form:"last_name" validate:"required"`
  Email 			string	`json:"email" form:"email" validate:"email,required"`
  AddressId 	int	    `json:"address_id" form:"address_id" validate:"number,required"`
  Activebool 	bool		`json:"active_bool" form:"active_bool" validate:"required"`
}

type UpdateCustomerValidate struct {
  CustomerId int      `json:"customer_id" validate:"number,min=1,required"`
  StoreId 		int			`json:"store_id" form:"store_id" validate:"number"`
  FirstName 	string	`json:"first_name" form:"first_name" validate:""`
  LastName 		string	`json:"last_name" form:"last_name" validate:""`
  Email 			string	`json:"email" form:"email" validate:"email"`
  AddressId 	int	    `json:"address_id" form:"address_id" validate:"number"`
  Activebool 	*bool		`json:"active_bool" form:"active_bool" validate:""`
}

type DeleteCustomerValidate struct {
  CustomerId int       `json:"customer_id" validate:"number,min=1,required"`
}

