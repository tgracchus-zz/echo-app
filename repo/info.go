package repo

import (
	"database/sql"
)

func Info(db *sql.DB) (string, error) {
	// Connect and check the server version
	var version string
	row := db.QueryRow("SELECT VERSION()")
	err := row.Scan(&version)
	if err != nil {
		return "", err
	}
	return version, nil
}
