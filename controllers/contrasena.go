package controllers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"USUARIOS/config"
	"USUARIOS/models"
)

func ObtenerContrasenas(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query(`SELECT id_contrasena, id_usuario, contrasena_hash, activa, fecha_creacion, fecha_modificacion FROM "Contrasena" ORDER BY id_contrasena`)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "error al consultar contrasena")
		return
	}
	defer rows.Close()
	items := []models.Contrasena{}
	for rows.Next() {
		var item models.Contrasena
		if err := rows.Scan(&item.IdContrasena, &item.IdUsuario, &item.ContrasenaHash, &item.Activa, &item.FechaCreacion, &item.FechaModificacion); err != nil {
			writeError(w, http.StatusInternalServerError, "error al leer contrasena")
			return
		}
		items = append(items, item)
	}
	writeJSON(w, http.StatusOK, items)
}

func ObtenerContrasenaPorID(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromRequest(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, "id invalido")
		return
	}
	var item models.Contrasena
	err = config.DB.QueryRow(`SELECT id_contrasena, id_usuario, contrasena_hash, activa, fecha_creacion, fecha_modificacion FROM "Contrasena" WHERE id_contrasena = $1`, id).
		Scan(&item.IdContrasena, &item.IdUsuario, &item.ContrasenaHash, &item.Activa, &item.FechaCreacion, &item.FechaModificacion)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			writeError(w, http.StatusNotFound, "contrasena no encontrada")
			return
		}
		writeError(w, http.StatusInternalServerError, "error al consultar contrasena")
		return
	}
	writeJSON(w, http.StatusOK, item)
}

func CrearContrasena(w http.ResponseWriter, r *http.Request) {
	var item models.Contrasena
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
	err := config.DB.QueryRow(`INSERT INTO "Contrasena" (id_usuario, contrasena_hash, activa) VALUES ($1, $2, $3) RETURNING id_contrasena, id_usuario, contrasena_hash, activa, fecha_creacion, fecha_modificacion`,
		item.IdUsuario, item.ContrasenaHash, item.Activa).Scan(&item.IdContrasena, &item.IdUsuario, &item.ContrasenaHash, &item.Activa, &item.FechaCreacion, &item.FechaModificacion)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "error al crear contrasena")
		return
	}
	writeJSON(w, http.StatusCreated, item)
}

func ActualizarContrasena(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromRequest(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, "id invalido")
		return
	}
	var item models.Contrasena
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
	err = config.DB.QueryRow(`UPDATE "Contrasena" SET id_usuario = $1, contrasena_hash = $2, activa = $3 WHERE id_contrasena = $4 RETURNING id_contrasena, id_usuario, contrasena_hash, activa, fecha_creacion, fecha_modificacion`,
		item.IdUsuario, item.ContrasenaHash, item.Activa, id).Scan(&item.IdContrasena, &item.IdUsuario, &item.ContrasenaHash, &item.Activa, &item.FechaCreacion, &item.FechaModificacion)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			writeError(w, http.StatusNotFound, "contrasena no encontrada")
			return
		}
		writeError(w, http.StatusInternalServerError, "error al actualizar contrasena")
		return
	}
	writeJSON(w, http.StatusOK, item)
}

func EliminarContrasena(w http.ResponseWriter, r *http.Request) {
	eliminarGenericoUsuarios(w, r, `"Contrasena"`, "id_contrasena", "contrasena")
}
