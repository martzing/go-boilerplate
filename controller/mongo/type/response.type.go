package mongoType

import (
	"time"

	"github.com/martzing/go-boilerplate/controller/postgres/module/db/models"
)

type CustomerType struct {
	ID         string    `json:"id"`
	StoreId    int       `json:"store_id"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	Email      string    `json:"email"`
	AddressId  int       `json:"address_id"`
	Activebool bool      `json:"active_bool"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type ListCustomerResponse struct {
	Count     int64          `json:"count"`
	Customers []CustomerType `json:"customers"`
}

type FilmType struct {
	ID           int                   `json:"id"`
	Title        string                `json:"title"`
	Description  string                `json:"description"`
	ReleaseYear  string                `json:"release_year"`
	UpdatedAt    time.Time             `json:"updated_at"`
	FilmCategory []models.FilmCategory `json:"film_category"`
}

type ListFilmResponse struct {
	Count int64         `json:"count"`
	Films []models.Film `json:"films"`
}

type DeleteCustumerResponse struct {
	Status string `json:"status"`
}
