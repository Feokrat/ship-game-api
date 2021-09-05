package app

import (
	"log"
	"ship-game-api/pkg/config"
	"ship-game-api/pkg/database/postgresql"
	"ship-game-api/pkg/repository"
)

func Run(configPath string) {
	cfg, err := config.Init(configPath)
	if err != nil {
		log.Fatal(err)
		return
	}

	db, err := postgresql.NewPostgresDB(cfg.Postgresql)
	if err != nil {
		log.Fatal(err)
		return
	}

	repositories := repository.NewRepositories(db)

	log.Printf("Repos: %v\n", repositories)
}
