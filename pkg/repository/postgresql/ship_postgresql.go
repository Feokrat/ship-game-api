package repository

import (
	"ship-game-api/pkg/models"

	"github.com/google/uuid"

	"github.com/jmoiron/sqlx"
)

type ShipRepositoryPostgres struct {
	db *sqlx.DB
}

func NewShipRepositoryPostgres(db *sqlx.DB) *ShipRepositoryPostgres {
	return &ShipRepositoryPostgres{db}
}

func (s *ShipRepositoryPostgres) Create(ship models.Ship) (uuid.UUID, error) {
	return uuid.New(), nil
}

func (s *ShipRepositoryPostgres) GetById(id uuid.UUID) (*models.Ship, error) {
	return nil, nil
}

func (s *ShipRepositoryPostgres) DeleteById(uuid.UUID) error {
	return nil
}
