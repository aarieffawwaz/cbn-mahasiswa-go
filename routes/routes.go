package routes

import (
	"database/sql"
	"net/http"

	"github.com/aarieffawwaz/cbn-mahasiswa-go/controller"
)

func MapRoutes(server *http.ServeMux, db *sql.DB) {
	server.HandleFunc("/", controller.NewHelloWorldController())
	server.HandleFunc("/mahasiswa", controller.NewIndexMahasiswa(db))
	server.HandleFunc("/mahasiswa/create", controller.NewCreateMahasiswaController(db))
	server.HandleFunc("/mahasiswa/update", controller.NewUpdateMahasiswaController(db))
	server.HandleFunc("/mahasiswa/delete", controller.NewDeleteMahasiswaController(db))
}
