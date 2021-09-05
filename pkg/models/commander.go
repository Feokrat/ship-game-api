package models

import "github.com/google/uuid"

type Commander struct {
	Id   uuid.UUID `json:"commander_id" db:"id"`
	Name string    `json:"commander_name" db:"name"`
	Age  int       `json:"age" db:"age"`
}
