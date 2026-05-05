package models

import "time"

type Rol struct {
	IdRol             int        `json:"id_rol"`
	NombreRol         string     `json:"nombre_rol"`
	Descripcion       *string    `json:"descripcion"`
	Activo            bool       `json:"activo"`
	FechaCreacion     time.Time  `json:"fecha_creacion"`
	FechaModificacion *time.Time `json:"fecha_modificacion"`
}
