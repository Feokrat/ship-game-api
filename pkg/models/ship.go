package models

import "github.com/google/uuid"

type Ship struct {
	Id     uuid.UUID `json:"ship_id" db:"id"`
	Name   string    `json:"ship_name" db:"name"`
	Level  string    `json:"ship_level" db:"level"`
	Health float32   `json:"health" db:"health"`
	Attack float32   `json:"attack" db:"attack"`
	Luck   float32   `json:"luck" db:"luck"`
	Fuel   float32   `json:"fuel" db:"fuel"`
}
