package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/spf13/viper"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//Env
	configPath := flag.String("configPath", "config/", "Path to the configuration directory")
	env := flag.String("env", "local", "Environment")
	flag.Parse()

	fmt.Printf("Starting with config file: %s \n", *configPath)
	fmt.Printf("Starting with env: %s\n ", *env)

	viper.AddConfigPath(*configPath)
	viper.SetConfigType("yaml")
	viper.SetConfigName("echo-" + *env)
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	viper.SetDefault("DATABASE_HOST", "127.0.0.1")
	viper.SetDefault("DATABASE_PORT", "3306")
	viper.SetDefault("DATABASE_NAME", "echo")
	viper.SetDefault("DATABASE_USER", "user")
	viper.SetDefault("DATABASE_PASSWORD", "password")

	databaseHost := viper.Get("DATABASE_HOST")
	databasePort := viper.Get("DATABASE_PORT")
	database := viper.Get("DATABASE_NAME")
	databaseUser := viper.Get("DATABASE_USER")
	databasePassword := viper.Get("DATABASE_PASSWORD")

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		databaseUser, databasePassword, databaseHost, databasePort, database)

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
