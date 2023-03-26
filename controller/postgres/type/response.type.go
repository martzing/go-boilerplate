package postgresType

import (
	"time"

	"github.com/martzing/go-boilerplate/controller/postgres/module/db/models"
)

type CustomerType struct {
	ID         int       `json:"id"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	FullName   string    `json:"full_name"`
	Email      string    `json:"email"`
	CreateDate time.Time `json:"created_date"`
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
