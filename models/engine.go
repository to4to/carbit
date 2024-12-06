package models

import "github.com/google/uuid"


type Engine struct {
	EngineId       uuid.UUID `json:"engine_id"`
	Displacement   int64     `json:"displacement"`
	NoOfCyclenders int64     `json:"noOfCyclenders"`
	CarRange       int64     `json:"carRange"`
}
