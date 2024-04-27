package app

import (
	"fmt"

	"github.com/Ixorlive/tw_vk/backend/services/notes/config"
	"github.com/Ixorlive/tw_vk/backend/services/notes/internal/controller/http"
	"github.com/Ixorlive/tw_vk/backend/services/notes/internal/usecase"
	"github.com/Ixorlive/tw_vk/backend/services/notes/internal/usecase/repo"
	"github.com/Ixorlive/tw_vk/backend/services/notes/pkg/postgres"
)

func Run(cfg *config.Config) {
	pg, err := postgres.New(cfg.PG.URL)
	if err != nil {
		fmt.Printf("Error to try connect to postgres: %s", err.Error())
		return
	}
	defer pg.Close()
	noteService := usecase.NewNoteService(repo.NewPGNoteRepo(pg))
	router := http.NewRouter(noteService)
	router.Run(cfg.HTTP.Port)
}
