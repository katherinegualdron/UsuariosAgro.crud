package routes

import (
	"USUARIOS/controllers"

	"github.com/gorilla/mux"
)

func RegistrarRutasPerfilExtendido(router *mux.Router) {
	router.HandleFunc("/perfil_extendido", controllers.ObtenerPerfilesExtendidos).Methods("GET")
	router.HandleFunc("/perfil_extendido/{id}", controllers.ObtenerPerfilExtendidoPorID).Methods("GET")
	router.HandleFunc("/perfil_extendido", controllers.CrearPerfilExtendido).Methods("POST")
	router.HandleFunc("/perfil_extendido/{id}", controllers.ActualizarPerfilExtendido).Methods("PUT")
	router.HandleFunc("/perfil_extendido/{id}", controllers.EliminarPerfilExtendido).Methods("DELETE")
}
