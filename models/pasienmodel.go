package models

import (
	"database/sql"
)

type PasienModel struct {
	conn *sql.DB
}

//func NewPasienModel() *PasienModel {
//	conn, err := config.DBConnection()
//}
