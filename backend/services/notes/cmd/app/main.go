package main

import (
	"log"

	"github.com/Ixorlive/tw_vk/backend/services/notes/config"
	"github.com/Ixorlive/tw_vk/backend/services/notes/internal/app"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)
}
