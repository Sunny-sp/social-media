package main

import (
	"context"
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
	"social/internal/infra/adapters"
	"social/internal/infra/db"
	"social/internal/infra/repository"
	"social/internal/storage"

	s3config "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
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
	// storage

	awscfg, err := s3config.LoadDefaultConfig(context.Background(),
		s3config.WithRegion(cfg.AWS.Region),
		s3config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			cfg.AWS.AccessKey,
			cfg.AWS.SecretKey,
			cfg.AWS.SessionToken,
		),
		),
	)

	if err != nil {
		log.Fatalf("failed to load AWS SDK config: %v", err)
	}
	s3Client := s3.NewFromConfig(awscfg)

	storageSvc := storage.NewS3Storage(s3Client, cfg.AWS.Bucket)

	// Repositories
	userRepo := repository.NewUserRepo(pool)
	authRepo := repository.NewAuthRepo(pool)
	postReo := repository.NewPostRepo(pool)

	// provider
	postProvider := adapters.NewPostProviderAdapter(postReo)

	// Services
	// method 1 injecting Post Interfcae Cause couple so not using
	// userSvc := user.NewUserService(userRepo, postReo)

	// method 2 injecting Provider Interfcae no cross domain reference
	// its an ACL(Anti Curreption Layer)
	userSvc := user.NewUserService(userRepo, postProvider)

	// testSecret := "testscret"
	// log.Println("testSecret:", testSecret)

	authSvc := auth.NewAuthService(authRepo, userRepo, []byte(cfg.JWT.Secret), cfg.JWT.Expiration)
	postSrv := post.NewPostService(postReo, storageSvc)

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
