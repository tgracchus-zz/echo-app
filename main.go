package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/tgracchus/echo-app/config"
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
	connectionString := echoConfig.Db.BuildConnectionString()

	fmt.Println(connectionString)
	//SQL
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	// Routes
	e.GET("/", hello)
	e.GET("/db", newDbInfoHandler(db))

	// Start server
	e.Logger.Debug(e.Start(":1323"))
}

func newDbInfoHandler(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Connect and check the server version
		var version string
		row := db.QueryRow("SELECT VERSION()")
		err := row.Scan(&version)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		fmt.Println("Connected to:", version)
		return c.String(http.StatusOK, fmt.Sprintf("Connected to %s:", version))
	}
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
