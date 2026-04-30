package models

import "time"

type Usuario struct {
	IdUsuario              int        `json:"id_usuario"`
	NombreCompleto         string     `json:"nombre_completo"`
	Correo                 string     `json:"correo"`
	Telefono               *string    `json:"telefono"`
	IdRol                  int        `json:"id_rol"`
	VerificacionDosPasos   bool       `json:"verificacion_dos_pasos"`
	Avatar                 *string    `json:"avatar"`
	Activo                 bool       `json:"activo"`
	FechaCreacion          time.Time  `json:"fecha_creacion"`
	FechaModificacion      time.Time  `json:"fecha_modificacion"`
}
