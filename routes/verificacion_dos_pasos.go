package routes

import (
	"USUARIOS/controllers"

	"github.com/gorilla/mux"
)

func RegistrarRutasVerificacionDosPasos(router *mux.Router) {
	router.HandleFunc("/verificacion_dos_pasos", controllers.ObtenerVerificacionesDosPasos).Methods("GET")
	router.HandleFunc("/verificacion_dos_pasos/{id}", controllers.ObtenerVerificacionDosPasosPorID).Methods("GET")
	router.HandleFunc("/verificacion_dos_pasos", controllers.CrearVerificacionDosPasos).Methods("POST")
	router.HandleFunc("/verificacion_dos_pasos/{id}", controllers.ActualizarVerificacionDosPasos).Methods("PUT")
	router.HandleFunc("/verificacion_dos_pasos/{id}", controllers.EliminarVerificacionDosPasos).Methods("DELETE")
}
