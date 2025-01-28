package main

import (
	"net/http"

	"github.com/aarieffawwaz/cbn-mahasiswa-go/database"
	"github.com/aarieffawwaz/cbn-mahasiswa-go/routes"
)

func main() {
	db := database.InitDatabase()

	server := http.NewServeMux()

	routes.MapRoutes(server, db)

	http.ListenAndServe(":8080", server)
}
