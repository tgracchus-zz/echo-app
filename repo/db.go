package repo

import (
	"database/sql"

	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/tgracchus/echo-app/config"
)

func NewDBInstance(cfg config.DatabaseConfig, driver string) (*sql.DB, error) {

	connectionString := cfg.BuildConnectionString()

	fmt.Println(connectionString)
	//SQL
	return sql.Open(driver, connectionString)

}
