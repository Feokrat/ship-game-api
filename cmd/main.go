package main

import "ship-game-api/pkg/app"

const configPath = "configs/config"

func main() {
	app.Run(configPath)
}
