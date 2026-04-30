package models

import "time"

type VerificacionDosPasos struct {
	IdVerificacion    int        `json:"id_verificacion"`
	IdUsuario         int        `json:"id_usuario"`
	TipoMetodo        string     `json:"tipo_metodo"`
	CodigoOtp         *string    `json:"codigo_otp"`
	ExpiracionOtp     *time.Time `json:"expiracion_otp"`
	Activo            bool       `json:"activo"`
	FechaCreacion     time.Time  `json:"fecha_creacion"`
	FechaModificacion time.Time  `json:"fecha_modificacion"`
}
