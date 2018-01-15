package main

import (
	"time"

	"github.com/Hendra-Huang/go-standard-layout"
	"github.com/Hendra-Huang/go-standard-layout/errorutil"
	"github.com/Hendra-Huang/go-standard-layout/log"
	"github.com/Hendra-Huang/go-standard-layout/mysql"
	"github.com/Hendra-Huang/go-standard-layout/router"
	"github.com/Hendra-Huang/go-standard-layout/server"
	"github.com/Hendra-Huang/go-standard-layout/server/handler"
)

func main() {
	// initialize database
	db, err := mysql.New(mysql.Options{
		DBHost:     "127.0.0.1",
		DBPort:     "3307",
		DBUser:     "myapp",
		DBPassword: "myapp",
		DBName:     "myapp",
	})
	errorutil.CheckWithErrorHandler(err, func(err error) {
		log.Errors(err)
		log.Fatal("Failed to initialize database")
	})

	// initialize repository
	userRepository := mysql.NewUserRepository(db, db)
	articleRepository := mysql.NewArticleRepository(db, db)

	// initialize service
	userService := myapp.NewUserService(userRepository)
	articleService := myapp.NewArticleService(articleRepository)

	// initialize handler
	pingHandler := handler.NewPingHandler()
	userHandler := handler.NewUserHandler(userService)
	articleHandler := handler.NewArticleHandler(articleService)

	server := server.New(server.Options{
		ListenAddress: ":5555",
	})

	r := router.New(router.Options{
		Timeout: 5 * time.Second,
	})

	router.RegisterRoute(r, pingHandler, userHandler, articleHandler)
	err = server.Serve(r)
	if err != nil {
		log.Fatalf("Error starting webserver. %s\n", err.Error())
	}
}
