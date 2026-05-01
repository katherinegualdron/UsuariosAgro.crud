package routes

import (
	"USUARIOS/controllers"

	"github.com/gorilla/mux"
)

func RegistrarRutasRol(router *mux.Router) {
	router.HandleFunc("/rol", controllers.ObtenerRoles).Methods("GET")
	router.HandleFunc("/rol/{id}", controllers.ObtenerRolPorID).Methods("GET")
	router.HandleFunc("/rol", controllers.CrearRol).Methods("POST")
	router.HandleFunc("/rol/{id}", controllers.ActualizarRol).Methods("PUT")
	router.HandleFunc("/rol/{id}", controllers.EliminarRol).Methods("DELETE")
}
