package controller

import (
	"database/sql"
	"html/template"
	"net/http"
	"path/filepath"
)

func NewUpdateMahasiswaController(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			id := r.URL.Query().Get("id")
			r.ParseForm()

			name := r.Form["name"][0]
			nim := r.Form["nim"][0]
			alamat := r.Form["alamat"][0]
			_, err := db.Exec("UPDATE mahasiswa SET name=?, nim=?, alamat=? WHERE id=?", name, nim, alamat, id)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			http.Redirect(w, r, "/mahasiswa", http.StatusMovedPermanently)

			return
		} else if r.Method == "GET" {
			id := r.URL.Query().Get("id")

			row := db.QueryRow("SELECT name, nim, alamat FROM mahasiswa WHERE id = ?", id)
			if row.Err() != nil {
				w.Write([]byte(row.Err().Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			var mahasiswa Mahasiswa
			err := row.Scan(
				&mahasiswa.Name,
				&mahasiswa.NIM,
				&mahasiswa.Alamat,
			)
			mahasiswa.ID = id
			if err != nil {
				w.Write([]byte(row.Err().Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			fp := filepath.Join("views", "update.html")
			tmpl, err := template.ParseFiles(fp)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			data := make(map[string]any)
			data["mahasiswa"] = mahasiswa

			err = tmpl.Execute(w, data)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
	}
}
