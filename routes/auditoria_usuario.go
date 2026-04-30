package routes

import (
	"USUARIOS/controllers"

	"github.com/gorilla/mux"
)

func RegistrarRutasAuditoriaUsuario(router *mux.Router) {
	router.HandleFunc("/auditoria_usuario", controllers.ObtenerAuditoriasUsuario).Methods("GET")
	router.HandleFunc("/auditoria_usuario/{id}", controllers.ObtenerAuditoriaUsuarioPorID).Methods("GET")
	router.HandleFunc("/auditoria_usuario", controllers.CrearAuditoriaUsuario).Methods("POST")
	router.HandleFunc("/auditoria_usuario/{id}", controllers.ActualizarAuditoriaUsuario).Methods("PUT")
	router.HandleFunc("/auditoria_usuario/{id}", controllers.EliminarAuditoriaUsuario).Methods("DELETE")
}
