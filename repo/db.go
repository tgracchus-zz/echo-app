package repo

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/tgracchus/echo-app/config"
)

func NewDBInstance(cfg config.DatabaseConfig, driver string) (*sql.DB, error) {
	connectionString := cfg.BuildConnectionString()
	return sql.Open(driver, connectionString)
}
