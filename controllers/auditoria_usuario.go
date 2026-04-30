package controllers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"USUARIOS/config"
	"USUARIOS/models"
)

func ObtenerAuditoriasUsuario(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query(`SELECT id_auditoria, id_usuario, tipo_evento, descripcion, ip_origen, fecha_evento FROM "AuditoriaUsuario" ORDER BY id_auditoria`)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "error al consultar auditoria_usuario")
		return
	}
	defer rows.Close()
	items := []models.AuditoriaUsuario{}
	for rows.Next() {
		item, err := scanAuditoriaUsuario(rows.Scan)
		if err != nil {
			writeError(w, http.StatusInternalServerError, "error al leer auditoria_usuario")
			return
		}
		items = append(items, item)
	}
	writeJSON(w, http.StatusOK, items)
}

func ObtenerAuditoriaUsuarioPorID(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromRequest(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, "id invalido")
		return
	}
	row := config.DB.QueryRow(`SELECT id_auditoria, id_usuario, tipo_evento, descripcion, ip_origen, fecha_evento FROM "AuditoriaUsuario" WHERE id_auditoria = $1`, id)
	item, err := scanAuditoriaUsuario(row.Scan)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			writeError(w, http.StatusNotFound, "auditoria_usuario no encontrado")
			return
		}
		writeError(w, http.StatusInternalServerError, "error al consultar auditoria_usuario")
		return
	}
	writeJSON(w, http.StatusOK, item)
}

func CrearAuditoriaUsuario(w http.ResponseWriter, r *http.Request) {
	var item models.AuditoriaUsuario
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		writeError(w, http.StatusBadRequest, "json invalido")
		return
	}
	if ok, err := existeRegistro(`"Usuario"`, "id_usuario", item.IdUsuario); err != nil {
		writeError(w, http.StatusBadRequest, "error al validar id_usuario")
		return
	} else if !ok {
		writeError(w, http.StatusBadRequest, "id_usuario no existe")
		return
	}
	row := config.DB.QueryRow(`INSERT INTO "AuditoriaUsuario" (id_usuario, tipo_evento, descripcion, ip_origen, fecha_evento) VALUES ($1, $2, $3, $4, $5) RETURNING id_auditoria, id_usuario, tipo_evento, descripcion, ip_origen, fecha_evento`,
		item.IdUsuario, item.TipoEvento, item.Descripcion, item.IpOrigen, item.FechaEvento)
	item, err := scanAuditoriaUsuario(row.Scan)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "error al crear auditoria_usuario")
		return
	}
	writeJSON(w, http.StatusCreated, item)
}

func ActualizarAuditoriaUsuario(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromRequest(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, "id invalido")
		return
	}
	var item models.AuditoriaUsuario
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		writeError(w, http.StatusBadRequest, "json invalido")
		return
	}
	if ok, err := existeRegistro(`"Usuario"`, "id_usuario", item.IdUsuario); err != nil {
		writeError(w, http.StatusBadRequest, "error al validar id_usuario")
		return
	} else if !ok {
		writeError(w, http.StatusBadRequest, "id_usuario no existe")
		return
	}
	row := config.DB.QueryRow(`UPDATE "AuditoriaUsuario" SET id_usuario = $1, tipo_evento = $2, descripcion = $3, ip_origen = $4, fecha_evento = $5 WHERE id_auditoria = $6 RETURNING id_auditoria, id_usuario, tipo_evento, descripcion, ip_origen, fecha_evento`,
		item.IdUsuario, item.TipoEvento, item.Descripcion, item.IpOrigen, item.FechaEvento, id)
	item, err = scanAuditoriaUsuario(row.Scan)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			writeError(w, http.StatusNotFound, "auditoria_usuario no encontrado")
			return
		}
		writeError(w, http.StatusInternalServerError, "error al actualizar auditoria_usuario")
		return
	}
	writeJSON(w, http.StatusOK, item)
}

func EliminarAuditoriaUsuario(w http.ResponseWriter, r *http.Request) {
	eliminarGenericoUsuarios(w, r, `"AuditoriaUsuario"`, "id_auditoria", "auditoria_usuario")
}

func scanAuditoriaUsuario(scan func(dest ...any) error) (models.AuditoriaUsuario, error) {
	var item models.AuditoriaUsuario
	var descripcion sql.NullString
	var ipOrigen sql.NullString
	err := scan(&item.IdAuditoria, &item.IdUsuario, &item.TipoEvento, &descripcion, &ipOrigen, &item.FechaEvento)
	if err != nil {
		return models.AuditoriaUsuario{}, err
	}
	item.Descripcion = nullStringToPointer(descripcion)
	item.IpOrigen = nullStringToPointer(ipOrigen)
	return item, nil
}
