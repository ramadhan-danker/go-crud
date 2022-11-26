package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func DBConection() (*sql.DB, error) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPassowrd := ""
	dbName := "go_crud"

	db, err := sql.Open(dbDriver, dbUser+":"+dbPassowrd+"@/"+dbName)
	return db, err
}
