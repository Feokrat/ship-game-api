package repository

import (
	"ship-game-api/pkg/models"
	repository "ship-game-api/pkg/repository/postgresql"

	"github.com/google/uuid"

	"github.com/jmoiron/sqlx"
)

type Ship interface {
	Create(models.Ship) (uuid.UUID, error)
	GetById(uuid.UUID) (*models.Ship, error)
	DeleteById(uuid.UUID) error
}

type Fleet interface {
	Create(models.Fleet) (uuid.UUID, error)
	GetById(uuid.UUID) (*models.Fleet, error)
	DeleteById(uuid.UUID) error
}

type Commander interface {
	Create(models.Commander) (uuid.UUID, error)
	GetById(uuid.UUID) (*models.Commander, error)
	DeleteById(uuid.UUID) error
}

type Repositories struct {
	Ship      Ship
	Fleet     Fleet
	Commander Commander
}

func NewRepositories(db *sqlx.DB) *Repositories {
	return &Repositories{
		Ship:      repository.NewShipRepositoryPostgres(db),
		Fleet:     repository.NewFleetRepositoryPostgres(db),
		Commander: repository.NewCommanderRepositoryPostgres(db),
	}
}
