package models

import (
	"time"

	"github.com/google/uuid"
)

type Car struct {
	Id       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Year     int       `json:"year"`
	Brand    string    `json:"brand"`
	FuelType string    `json:"fuel_type"`

	Engine    Engine    `json:"engine"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_At"`
}





