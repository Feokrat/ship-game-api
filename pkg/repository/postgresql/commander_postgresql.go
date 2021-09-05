package repository

import (
	"ship-game-api/pkg/models"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type CommanderRepositoryPostgres struct {
	db *sqlx.DB
}

func NewCommanderRepositoryPostgres(db *sqlx.DB) *CommanderRepositoryPostgres {
	return &CommanderRepositoryPostgres{db}
}

func (s *CommanderRepositoryPostgres) Create(commander models.Commander) (uuid.UUID, error) {
	return uuid.New(), nil
}

func (s *CommanderRepositoryPostgres) GetById(id uuid.UUID) (*models.Commander, error) {
	return nil, nil
}

func (s *CommanderRepositoryPostgres) DeleteById(uuid.UUID) error {
	return nil
}
