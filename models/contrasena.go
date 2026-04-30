package models

import "time"

type Contrasena struct {
	IdContrasena       int        `json:"id_contrasena"`
	IdUsuario          int        `json:"id_usuario"`
	ContrasenaHash     string     `json:"contrasena_hash"`
	Activa             bool       `json:"activa"`
	FechaCreacion      time.Time  `json:"fecha_creacion"`
	FechaModificacion  time.Time  `json:"fecha_modificacion"`
}
