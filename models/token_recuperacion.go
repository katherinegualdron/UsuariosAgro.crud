package models

import "time"

type TokenRecuperacion struct {
	IdToken          int       `json:"id_token"`
	Token            string    `json:"token"`
	IdUsuario        int       `json:"id_usuario"`
	Expiracion       time.Time `json:"expiracion"`
	Usado            bool      `json:"usado"`
	FechaCreacion    time.Time `json:"fecha_creacion"`
}
