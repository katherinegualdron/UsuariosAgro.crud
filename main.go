package main

import (
	"log"
	"net/http"

	"USUARIOS/config"
	"USUARIOS/routes"

	"github.com/gorilla/mux"
)

func main() {
	config.ConnectDB()

	router := mux.NewRouter()
	routes.RegistrarRutas(router)

	log.Println("Servidor USUARIOS escuchando en :8091")
	log.Fatal(http.ListenAndServe(":8091", router))
}
