package handler

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var dbConnector *sql.DB

func SetDB(db *sql.DB) {
	dbConnector = db
}
