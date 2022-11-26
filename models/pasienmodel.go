package models

import (
	"database/sql"
	//"github.com/ramadhan-danker/go-crud/entities"
)

type PasienModel struct {
	conn *sql.DB
}

//func NewPasienModel() *PasienModel {
//	conn, err := config.DBConnection()
//}
