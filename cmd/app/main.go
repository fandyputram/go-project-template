package main

import (
	"log"
	"net/http"

	"github.com/fandyputram/go-project-template/config"
	httpHandler "github.com/fandyputram/go-project-template/internal/delivery/http"
	"github.com/fandyputram/go-project-template/internal/repository/mysql"
	"github.com/fandyputram/go-project-template/internal/usecase"
	"github.com/fandyputram/go-project-template/pkg/database"
)

func main() {
	cfg := config.LoadConfig()

	db, err := database.NewMySQL(cfg.Database.DSN)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize repository
	repo := mysql.NewMySQLRepository(db)

	// Initialize usecase
	uc := usecase.NewUsecase(repo)

	// Initialize HTTP handler
	handler := httpHandler.NewHandler(uc)

	// Start HTTP server
	log.Printf("Starting server on %s", cfg.Server.Address)
	if err := http.ListenAndServe(cfg.Server.Address, handler); err != nil {
		log.Fatal(err)
	}
}
