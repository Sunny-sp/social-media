package main

import (
	"log"
	"social/internal/api"
	"social/internal/api/handlers"
	"social/internal/domain/auth"
	"social/internal/domain/user"
	"social/internal/infra/config"
	"social/internal/infra/db"
	"social/internal/infra/repository"
)

func main() {
	// Load config
	cfg := config.Loadenv()

	// DB connection
	pool, err := db.ConnectPool(cfg.DB.ConnString())
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}
	defer pool.Close()

	// Repositories
	userRepo := repository.NewUserRepo(pool)
	authRepo := repository.NewAuthRepo(pool)

	// Services
	userSvc := user.NewUserService(userRepo)
	authSvc := auth.NewAuthService(authRepo, userRepo, []byte(cfg.JWT.Secret), cfg.JWT.Expiration)

	// Handlers
	userHandler := handlers.NewUserHandler(userSvc)
	authHandler := handlers.NewAuthHandler(authSvc)

	// Server
	app := api.NewServer(cfg.Server, userHandler, authHandler)
	mux := app.Mount()
	if err := app.Run(mux); err != nil {
		log.Fatal(err)
	}
}
