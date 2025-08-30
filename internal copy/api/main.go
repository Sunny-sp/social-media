	package main

	import (
		"log"
		"social/internal/api/controller"
		"social/internal/api/db"
		"social/internal/api/repository"
		"social/internal/api/service"
		"social/internal/env"
	)

	func main() {
		//load env
		envCfg := env.Loadenv()

		// DB connect
		dsn := envCfg.DB.ConnString()
		pool, err := db.ConnectPool(dsn)

		if err != nil {
			log.Fatalf("failed to connect to db: %v", err)
		}

		// repository
		repo := repository.NewRepository(pool)

		//service
		srv := service.NewService(repo)

		//controller
		ctrl := controller.NewController(srv)

		// create application
		app := &application{
			config: envCfg.Server,
			ctrl:   ctrl,
		}

		//add all routes
		mux := app.mount()
		// run the server
		err = app.run(mux)
		log.Fatal(err)
	}
