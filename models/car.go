package models

import (
	"errors"
	"strconv"
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

type CarRequest struct {
	Name     string `json:"name"`
	Year     int    `json:"year"`
	Brand    string `json:"brand"`
	FuelType string `json:"fuel_type"`

	Engine Engine  `json:"engine"`
	Price  float64 `json:"price"`
}



func ValidateRequest(carReq CarRequest)error{




	
}







func validateName(name string) error {
	if name == "" {
		return errors.New("Enter Name")
	}
	return nil
}

func validateYear(year string) error {
	if year == "" {
		return errors.New("Year is Required")
	}

	_, err := strconv.Atoi(year)

	if err != nil {
		return errors.New("Year must be a Valid Number")
	}

	currentYear := time.Now().Year()
	yearInInt, _ := strconv.Atoi(year)

	if yearInInt < 1886 || yearInInt > currentYear {
		return errors.New("Year Must be Between 1886 and Current Year")
	}

	return nil
}

func validateBrand(brand string) error {

	if brand == "" {
		return errors.New("Enter Brand")
	}
	return nil

}

func validateFuelType(fuelType string) error {
	validateFuelTypes := []string{"Petrol", "Diesel", "Electric", "Hybrid"}

	for _, valid := range validateFuelTypes {
		if fuelType == valid {
			return nil
		}
	}

	return errors.New("Wrong Fuel type Provided options include Petrol, Diesel, Electric, Hybrid")
}

func validateEngine(engine Engine) error {

	if engine.EngineId == uuid.Nil {
		return errors.New("Engine UUid Required")
	}

	if engine.Displacement <= 0 {
		errors.New("Displacement must be GREATER than ZERO")
	}

	if engine.NoOfCyclenders > 0 {
		errors.New("No of Cyclenders Greater than zero")
	}

	if engine.CarRange <= 0 {
		errors.New("Car Range Greater than zero")
	}
	return nil
}

func validPrice(price float64) error {

	if price <= 0 {
		return errors.New("Price must be greater than zero")
	}

	return nil
}
