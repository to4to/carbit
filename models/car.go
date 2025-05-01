package models

import (
	"errors"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type Car struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Year     string    `json:"year"`
	Brand    string    `json:"brand"`
	FuelType string    `json:"fuelType"`
	Engine   Engine    `json:"engine"`
	Price    float64   `json:"price"`
	// CreatedAt time.Time `json:"created_at"`
	// UpdatedAt time.Time `json:"updated_at"`
}

type CarRequest struct {
	Name     string  `json:"name"`
	Year     string  `json:"year"`
	Brand    string  `json:"brand"`
	FuelType string  `json:"fuelType"`
	Engine   Engine  `json:"engine"`
	Price    float64 `json:"price"`
}

func ValidateCarRequest(carRequest CarRequest) error {
	if err := validateName(carRequest.Name); err != nil {
		return err
	}
	if err := validateYear(carRequest.Year); err != nil {
		return err
	}
	if err := validateBrand(carRequest.Brand); err != nil {
		return err
	}
	if err := validateFuelType(carRequest.FuelType); err != nil {
		return err
	}
	if err := validateEngine(carRequest.Engine); err != nil {
		return err
	}
	if err := validatePrice(carRequest.Price); err != nil {
		return err
	}
	return nil
}

func validateName(name string) error {
	if name == "" {
		return errors.New("name is required")
	}
	return nil
}

func validateYear(year string) error {
	if year == "" {
		return errors.New("year is required")
	}

	_, err := strconv.Atoi(year)
	if err != nil {
		return errors.New("year must be a number")
	}
	curentYear := time.Now().Year()
	yearInt, _ := strconv.Atoi(year)

	if yearInt > curentYear || yearInt < 1900 {
		return errors.New("year must be less than or equal to the current year and greater than 1900")
	}
	return nil
}

func validateBrand(brand string) error {
	if brand == "" {
		return errors.New("brand is required")
	}
	return nil
}

func validateFuelType(fuelType string) error {
	if fuelType == "" {
		return errors.New("fuelType is required")
	}

	validateFuelTypes := []string{"Petrol", "Diesel", "Electric", "Hybrid"}
	for _, fuel := range validateFuelTypes {
		if fuel == fuelType {
			return nil
		}
	}
	return errors.New("fuelType must be one of the following: Petrol, Diesel, Electric, Hybrid")
}

func validateEngine(engine Engine) error {
	if engine.EngineId == uuid.Nil {
		return errors.New("engine is required")
	}

	if engine.Displacement <= 0 {
		return errors.New("engineName is required")
	}

	if engine.NoOfCyclenders <= 0 {
		return errors.New("No Of cylinders is required")
	}

	if engine.CarRange <= 0 {
		return errors.New("Car Range is required")
	}
	return nil
}

func validatePrice(price float64) error {
	if price <= 0 {
		return errors.New("price must be greater than 0")
	}
	return nil
}
