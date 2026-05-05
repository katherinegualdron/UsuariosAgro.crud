package routes

import (
	"net/http"
	"testing"

	"github.com/gorilla/mux"
)

func TestRegistrarRutas(t *testing.T) {
	router := mux.NewRouter()
	RegistrarRutas(router)

	tests := []struct {
		method string
		path   string
	}{
		{http.MethodGet, "/rol"},
		{http.MethodGet, "/rol/1"},
		{http.MethodPost, "/rol"},
		{http.MethodPut, "/rol/1"},
		{http.MethodDelete, "/rol/1"},
		{http.MethodGet, "/usuario"},
		{http.MethodGet, "/usuario/1"},
		{http.MethodPost, "/usuario"},
		{http.MethodPut, "/usuario/1"},
		{http.MethodDelete, "/usuario/1"},
		{http.MethodGet, "/contrasena"},
		{http.MethodGet, "/contrasena/1"},
		{http.MethodPost, "/contrasena"},
		{http.MethodPut, "/contrasena/1"},
		{http.MethodDelete, "/contrasena/1"},
		{http.MethodGet, "/token_recuperacion"},
		{http.MethodGet, "/token_recuperacion/1"},
		{http.MethodPost, "/token_recuperacion"},
		{http.MethodPut, "/token_recuperacion/1"},
		{http.MethodDelete, "/token_recuperacion/1"},
		{http.MethodGet, "/verificacion_dos_pasos"},
		{http.MethodGet, "/verificacion_dos_pasos/1"},
		{http.MethodPost, "/verificacion_dos_pasos"},
		{http.MethodPut, "/verificacion_dos_pasos/1"},
		{http.MethodDelete, "/verificacion_dos_pasos/1"},
		{http.MethodGet, "/auditoria_usuario"},
		{http.MethodGet, "/auditoria_usuario/1"},
		{http.MethodPost, "/auditoria_usuario"},
		{http.MethodPut, "/auditoria_usuario/1"},
		{http.MethodDelete, "/auditoria_usuario/1"},
		{http.MethodGet, "/perfil_extendido"},
		{http.MethodGet, "/perfil_extendido/1"},
		{http.MethodPost, "/perfil_extendido"},
		{http.MethodPut, "/perfil_extendido/1"},
		{http.MethodDelete, "/perfil_extendido/1"},
	}

	for _, tt := range tests {
		t.Run(tt.method+" "+tt.path, func(t *testing.T) {
			request, err := http.NewRequest(tt.method, tt.path, nil)
			if err != nil {
				t.Fatalf("crear request: %v", err)
			}
			if !router.Match(request, &mux.RouteMatch{}) {
				t.Fatalf("ruta no registrada")
			}
		})
	}
}
