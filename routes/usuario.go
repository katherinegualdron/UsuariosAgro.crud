package routes

import (
	"USUARIOS/controllers"

	"github.com/gorilla/mux"
)

func RegistrarRutasUsuario(router *mux.Router) {
	router.HandleFunc("/usuario", controllers.ObtenerUsuarios).Methods("GET")
	router.HandleFunc("/usuario/{id}", controllers.ObtenerUsuarioPorID).Methods("GET")
	router.HandleFunc("/usuario", controllers.CrearUsuario).Methods("POST")
	router.HandleFunc("/usuario/{id}", controllers.ActualizarUsuario).Methods("PUT")
	router.HandleFunc("/usuario/{id}", controllers.EliminarUsuario).Methods("DELETE")
}
