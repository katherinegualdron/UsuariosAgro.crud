package models

import "time"

type AuditoriaUsuario struct {
	IdAuditoria   int        `json:"id_auditoria"`
	IdUsuario     int        `json:"id_usuario"`
	TipoEvento    string     `json:"tipo_evento"`
	Descripcion   *string    `json:"descripcion"`
	IpOrigen      *string    `json:"ip_origen"`
	FechaEvento   time.Time  `json:"fecha_evento"`
}
