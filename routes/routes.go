package routes

import "github.com/gorilla/mux"

func RegistrarRutas(router *mux.Router) {
	RegistrarRutasRol(router)
	RegistrarRutasUsuario(router)
	RegistrarRutasContrasena(router)
	RegistrarRutasTokenRecuperacion(router)
	RegistrarRutasVerificacionDosPasos(router)
	RegistrarRutasAuditoriaUsuario(router)
	RegistrarRutasPerfilExtendido(router)
}
