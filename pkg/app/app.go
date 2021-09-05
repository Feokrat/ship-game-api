package app

import (
	"log"
	"ship-game-api/pkg/config"
)

func Run(configPath string) {
	cfg, err := config.Init(configPath)
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Println(cfg)
}
