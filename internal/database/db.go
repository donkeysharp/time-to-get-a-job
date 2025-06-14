package database

import (
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
)

func NewDatabaseConnection(username, password, database, hostname string, port int) (*sqlx.DB, error) {
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%v/%s", username, password, hostname, port, database)

	db, err := sqlx.Connect("pgx", connectionString)
	if err != nil {
		log.Errorf("Error connecting to %v:***@%v/%v", username, hostname, database)
		return nil, err
	}
	return db, nil
}
