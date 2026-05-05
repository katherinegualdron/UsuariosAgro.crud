package controllers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"USUARIOS/config"
	"USUARIOS/models"

	"github.com/lib/pq"
)

func ObtenerRoles(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query(`SELECT id_rol, nombre_rol, descripcion, activo, fecha_creacion, fecha_modificacion FROM "Usuarios"."Rol" ORDER BY id_rol ASC`)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "error al consultar rol")
		return
	}
	defer rows.Close()

	items := []models.Rol{}
	for rows.Next() {
		item, err := scanRol(rows.Scan)
		if err != nil {
			writeError(w, http.StatusInternalServerError, "error al leer rol")
			return
		}
		items = append(items, item)
	}
	writeJSON(w, http.StatusOK, items)
}

func ObtenerRolPorID(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromRequest(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, "id invalido")
		return
	}
	row := config.DB.QueryRow(`SELECT id_rol, nombre_rol, descripcion, activo, fecha_creacion, fecha_modificacion FROM "Usuarios"."Rol" WHERE id_rol = $1`, id)
	item, err := scanRol(row.Scan)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			writeError(w, http.StatusNotFound, "rol no encontrado")
			return
		}
		writeError(w, http.StatusInternalServerError, "error al consultar rol")
		return
	}
	writeJSON(w, http.StatusOK, item)
}

func CrearRol(w http.ResponseWriter, r *http.Request) {
	var item models.Rol
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		writeError(w, http.StatusBadRequest, "json invalido")
		return
	}
	row := config.DB.QueryRow(`INSERT INTO "Usuarios"."Rol" (nombre_rol, descripcion, activo) VALUES ($1, $2, $3) RETURNING id_rol, nombre_rol, descripcion, activo, fecha_creacion, fecha_modificacion`,
		item.NombreRol, item.Descripcion, item.Activo)
	item, err := scanRol(row.Scan)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "error al crear rol")
		return
	}
	writeJSON(w, http.StatusCreated, item)
}

func ActualizarRol(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromRequest(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, "id invalido")
		return
	}
	var item models.Rol
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		writeError(w, http.StatusBadRequest, "json invalido")
		return
	}
	row := config.DB.QueryRow(`UPDATE "Usuarios"."Rol" SET nombre_rol = $1, descripcion = $2, activo = $3 WHERE id_rol = $4 RETURNING id_rol, nombre_rol, descripcion, activo, fecha_creacion, fecha_modificacion`,
		item.NombreRol, item.Descripcion, item.Activo, id)
	item, err = scanRol(row.Scan)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			writeError(w, http.StatusNotFound, "rol no encontrado")
			return
		}
		writeError(w, http.StatusInternalServerError, "error al actualizar rol")
		return
	}
	writeJSON(w, http.StatusOK, item)
}

func EliminarRol(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromRequest(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, "id invalido")
		return
	}
	result, err := config.DB.Exec(`DELETE FROM "Usuarios"."Rol" WHERE id_rol = $1`, id)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23503" {
			writeError(w, http.StatusConflict, "no se puede eliminar rol porque tiene registros relacionados")
			return
		}
		writeError(w, http.StatusInternalServerError, "error al eliminar rol")
		return
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		writeError(w, http.StatusNotFound, "rol no encontrado")
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"message": "rol eliminado correctamente"})
}

func scanRol(scan func(dest ...any) error) (models.Rol, error) {
	var item models.Rol
	var descripcion sql.NullString
	var fechaModificacion sql.NullTime
	err := scan(&item.IdRol, &item.NombreRol, &descripcion, &item.Activo, &item.FechaCreacion, &fechaModificacion)
	if err != nil {
		return models.Rol{}, err
	}
	item.Descripcion = nullStringToPointer(descripcion)
	item.FechaModificacion = nullTimeToPointer(fechaModificacion)
	return item, nil
}
