package models

import (
	"database/sql"
	"fmt"
	"github.com/ramadhan-danker/go-crud/config"
	"github.com/ramadhan-danker/go-crud/entities"
)

type PasienModel struct {
	conn *sql.DB
}

func NewPasienModel() *PasienModel {
	conn, err := config.DBConection()
	if err != nil {
		panic(err)
	}
	return &PasienModel{
		conn: conn,
	}
}
func (p *PasienModel) Create(pasien entities.Pasien) bool {
	result, err := p.conn.Exec("insert into pasien (nama_lengkap,nik,jenis_kelamin,tempat_lahir,tanggal_lahir,alamat,no_hp) values (?,?,?,?,?,?,?)", pasien.NamaLengkap, pasien.NIK, pasien.JenisKelamin, pasien.TempatLahir, pasien.TanggalLahir, pasien.Alamat, pasien.NoHp)

	if err != nil {
		fmt.Println(err)
		return false
	}
	lastInsertId, _ := result.LastInsertId()
	return lastInsertId > 0
}
