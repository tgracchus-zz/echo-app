package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/tgracchus/echo-app/config"
	"github.com/tgracchus/echo-app/handlers"
	"github.com/tgracchus/echo-app/repo"
)

func main() {

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	echoConfig, err := config.NewEchoConfig()
	if err != nil {
		panic(err)
	}

	db, err := repo.NewDBInstance(echoConfig.Db, "mysql")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Routes
	e.GET("/", handlers.HelloHandler)
	e.GET("/db", handlers.NewInfoHandler(db))

	// Start server
	e.Logger.Debug(e.Start(":1323"))
}
