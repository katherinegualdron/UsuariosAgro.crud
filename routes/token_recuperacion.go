package routes

import (
	"USUARIOS/controllers"

	"github.com/gorilla/mux"
)

func RegistrarRutasTokenRecuperacion(router *mux.Router) {
	router.HandleFunc("/token_recuperacion", controllers.ObtenerTokensRecuperacion).Methods("GET")
	router.HandleFunc("/token_recuperacion/{id}", controllers.ObtenerTokenRecuperacionPorID).Methods("GET")
	router.HandleFunc("/token_recuperacion", controllers.CrearTokenRecuperacion).Methods("POST")
	router.HandleFunc("/token_recuperacion/{id}", controllers.ActualizarTokenRecuperacion).Methods("PUT")
	router.HandleFunc("/token_recuperacion/{id}", controllers.EliminarTokenRecuperacion).Methods("DELETE")
}
