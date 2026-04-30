package routes

import (
	"USUARIOS/controllers"

	"github.com/gorilla/mux"
)

func RegistrarRutasContrasena(router *mux.Router) {
	router.HandleFunc("/contrasena", controllers.ObtenerContrasenas).Methods("GET")
	router.HandleFunc("/contrasena/{id}", controllers.ObtenerContrasenaPorID).Methods("GET")
	router.HandleFunc("/contrasena", controllers.CrearContrasena).Methods("POST")
	router.HandleFunc("/contrasena/{id}", controllers.ActualizarContrasena).Methods("PUT")
	router.HandleFunc("/contrasena/{id}", controllers.EliminarContrasena).Methods("DELETE")
}
