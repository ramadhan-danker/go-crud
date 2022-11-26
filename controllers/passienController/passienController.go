package passienController

import (
	"github.com/ramadhan-danker/go-crud/entities"
	"github.com/ramadhan-danker/go-crud/models"
	"html/template"
	"net/http"
)

func Index(response http.ResponseWriter, request *http.Request) {
	temp, err := template.ParseFiles("views/pasien/index.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(response, nil)
}

var pasienModel = models.NewPasienModel()

func Add(response http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		temp, err := template.ParseFiles("views/pasien/add.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(response, nil)
	} else if request.Method == http.MethodPost {
		request.ParseForm()

		var pasien entities.Pasien
		pasien.NamaLengkap = request.Form.Get("nama_lengkap")
		pasien.NIK = request.Form.Get("nik")
		pasien.JenisKelamin = request.Form.Get("jenis_kelamin")
		pasien.TempatLahir = request.Form.Get("tempat_lahir")
		pasien.TanggalLahir = request.Form.Get("tanggal_lahir")
		pasien.Alamat = request.Form.Get("alamat")
		pasien.NoHp = request.Form.Get("no_hp")

		pasienModel.Create(pasien)
		data := map[string]interface{}{
			"pesan": "Data pasien berhasil di simpan",
		}
		temp, _ := template.ParseFiles("views/pasien/add.html")
		temp.Execute(response, data)
	}

}
func Edit(response http.ResponseWriter, reques *http.Request) {

}
func Delete(response http.ResponseWriter, reques *http.Request) {

}
