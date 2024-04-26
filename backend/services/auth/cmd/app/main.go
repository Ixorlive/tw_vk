package main

import (
	"log"

	"github.com/Ixorlive/tw_vk/backend/services/auth/config"
	"github.com/Ixorlive/tw_vk/backend/services/auth/internal/app"
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
