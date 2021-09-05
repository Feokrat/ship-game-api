package postgresql

import (
	"fmt"
	"ship-game-api/pkg/config"

	"github.com/jmoiron/sqlx"
)

func NewPostgresDB(cfg config.PGConfig) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres",
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBName, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, err
}
