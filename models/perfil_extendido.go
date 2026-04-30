package models

import "time"

type PerfilExtendido struct {
	IdPerfil           int        `json:"id_perfil"`
	IdUsuario          int        `json:"id_usuario"`
	Direccion          *string    `json:"direccion"`
	Ciudad             *string    `json:"ciudad"`
	Intereses          *string    `json:"intereses"`
	RedesSociales      *string    `json:"redes_sociales"`
	FechaCreacion      time.Time  `json:"fecha_creacion"`
	FechaModificacion  time.Time  `json:"fecha_modificacion"`
}
