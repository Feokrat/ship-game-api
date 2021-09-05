package repository

import (
	"ship-game-api/pkg/models"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type FleetRepositoryPostgres struct {
	db *sqlx.DB
}

func NewFleetRepositoryPostgres(db *sqlx.DB) *FleetRepositoryPostgres {
	return &FleetRepositoryPostgres{db}
}

func (s *FleetRepositoryPostgres) Create(fleet models.Fleet) (uuid.UUID, error) {
	return uuid.New(), nil
}

func (s *FleetRepositoryPostgres) GetById(id uuid.UUID) (*models.Fleet, error) {
	return nil, nil
}

func (s *FleetRepositoryPostgres) DeleteById(uuid.UUID) error {
	return nil
}
