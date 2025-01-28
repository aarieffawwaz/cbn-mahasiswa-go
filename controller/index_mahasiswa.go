package controller

import (
	"database/sql"
	"html/template"
	"net/http"
	"path/filepath"
)

type Mahasiswa struct {
	ID     string
	Name   string
	NIM    string
	Alamat string
}

func NewIndexMahasiswa(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT id, name, nim, alamat FROM mahasiswa")
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var mahasiswas []Mahasiswa
		for rows.Next() {
			var mahasiswa Mahasiswa

			err = rows.Scan(
				&mahasiswa.ID,
				&mahasiswa.Name,
				&mahasiswa.NIM,
				&mahasiswa.Alamat,
			)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			mahasiswas = append(mahasiswas, mahasiswa)
		}

		fp := filepath.Join("views", "index.html")
		tmpl, err := template.ParseFiles(fp)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		data := make(map[string]any)
		data["mahasiswas"] = mahasiswas

		err = tmpl.Execute(w, data)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
