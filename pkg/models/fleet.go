package models

import "github.com/google/uuid"

type Fleet struct {
	Id          uuid.UUID   `json:"fleet_id" db:"id"`
	Name        string      `json:"fleet_name" db:"name"`
	CommanderId uuid.UUID   `json:"commander_id" db:"commander_id"`
	ShipIds     []uuid.UUID `json:"ship_ids" db:"ship_ids"`
}
