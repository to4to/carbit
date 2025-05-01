package models

import (
	"errors"

	"github.com/google/uuid"
)

type Engine struct {
	EngineId       uuid.UUID `json:"engine_id"`
	Displacement   int64     `json:"displacement"`
	NoOfCyclenders int64     `json:"noOfCyclenders"`
	CarRange       int64     `json:"carRange"`
}

type EngineRequest struct {
	Displacement   int64 `json:"displacement"`
	NoOfCyclenders int64 `json:"noOfCyclenders"`
	CarRange       int64 `json:"carRange"`
}

func EngineValidate(engine EngineRequest) error {
	if err := validateDisplacement(engine.Displacement); err != nil {
		return err
	}
	if err := validateNoOfCyclenders(engine.NoOfCyclenders); err != nil {
		return err
	}
	if err := validateCarRange(engine.CarRange); err != nil {
		return err
	}
	return nil
}
func validateDisplacement(displacement int64) error {
	if displacement <= 0 {
		return errors.New("displacement must be greater than 0")
	}
	return nil
}

func validateNoOfCyclenders(noOfCyclenders int64) error {
	if noOfCyclenders <= 0 {
		return errors.New("noOfCyclenders must be greater than 0")
	}
	return nil
}

func validateCarRange(carRange int64) error {
	if carRange <= 0 {
		return errors.New("carRange must be greater than 0")
	}
	return nil
}
