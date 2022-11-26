package main

import (
	"github.com/ramadhan-danker/go-crud/controllers/passienController"
	"net/http"
)

func main() {
	http.HandleFunc("/", passienController.Index)
	http.HandleFunc("/pasien", passienController.Index)
	http.HandleFunc("/pasien/index", passienController.Index)
	http.HandleFunc("/pasien/add", passienController.Add)
	http.HandleFunc("/pasien/edit", passienController.Edit)
	http.HandleFunc("/pasien/delete", passienController.Delete)

	http.ListenAndServe(":3000", nil)
}
