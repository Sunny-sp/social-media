package main

import (
	"log"
	"social/internal/api/auth_api"
	"social/internal/api/middleware"
	"social/internal/api/post_api"
	"social/internal/api/server"
	"social/internal/api/user_api"
	"social/internal/config"
	"social/internal/domain/auth"
	"social/internal/domain/post"
	"social/internal/domain/user"
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
	postReo := repository.NewPostRepo(pool)

	// Services
	userSvc := user.NewUserService(userRepo)
	// testSecret := "testscret"
	// log.Println("testSecret:", testSecret)

	authSvc := auth.NewAuthService(authRepo, userRepo, []byte(cfg.JWT.Secret), cfg.JWT.Expiration)
	postSrv := post.NewPostService(postReo)

	// Handler
	userHandler := user_api.NewUserHandler(userSvc)
	authHandler := auth_api.NewAuthHandler(authSvc)
	postHandler := post_api.NewPosthandler(postSrv)

	// Middleware
	authMiddleware := middleware.NewAuthMiddleware([]byte(cfg.JWT.Secret))

	// Server
	app := server.NewServer(cfg.Server, authMiddleware, userHandler, authHandler, postHandler)
	mux := app.Mount()
	if err := app.Run(mux); err != nil {
		log.Fatal(err)
	}
}
