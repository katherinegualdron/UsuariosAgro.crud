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

func ObtenerUsuarios(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query(`SELECT id_usuario, nombre_completo, correo, telefono, id_rol, verificacion_dos_pasos, avatar, activo, fecha_creacion, fecha_modificacion FROM "Usuario" ORDER BY id_usuario`)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "error al consultar usuario")
		return
	}
	defer rows.Close()
	items := []models.Usuario{}
	for rows.Next() {
		item, err := scanUsuario(rows.Scan)
		if err != nil {
			writeError(w, http.StatusInternalServerError, "error al leer usuario")
			return
		}
		items = append(items, item)
	}
	writeJSON(w, http.StatusOK, items)
}

func ObtenerUsuarioPorID(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromRequest(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, "id invalido")
		return
	}
	row := config.DB.QueryRow(`SELECT id_usuario, nombre_completo, correo, telefono, id_rol, verificacion_dos_pasos, avatar, activo, fecha_creacion, fecha_modificacion FROM "Usuario" WHERE id_usuario = $1`, id)
	item, err := scanUsuario(row.Scan)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			writeError(w, http.StatusNotFound, "usuario no encontrado")
			return
		}
		writeError(w, http.StatusInternalServerError, "error al consultar usuario")
		return
	}
	writeJSON(w, http.StatusOK, item)
}

func CrearUsuario(w http.ResponseWriter, r *http.Request) {
	var item models.Usuario
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		writeError(w, http.StatusBadRequest, "json invalido")
		return
	}
	if ok, err := existeRegistro(`"Rol"`, "id_rol", item.IdRol); err != nil {
		writeError(w, http.StatusBadRequest, "error al validar id_rol")
		return
	} else if !ok {
		writeError(w, http.StatusBadRequest, "id_rol no existe")
		return
	}
	row := config.DB.QueryRow(`INSERT INTO "Usuario" (nombre_completo, correo, telefono, id_rol, verificacion_dos_pasos, avatar, activo) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id_usuario, nombre_completo, correo, telefono, id_rol, verificacion_dos_pasos, avatar, activo, fecha_creacion, fecha_modificacion`,
		item.NombreCompleto, item.Correo, item.Telefono, item.IdRol, item.VerificacionDosPasos, item.Avatar, item.Activo)
	item, err := scanUsuario(row.Scan)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			writeError(w, http.StatusConflict, "correo ya existe")
			return
		}
		writeError(w, http.StatusInternalServerError, "error al crear usuario")
		return
	}
	writeJSON(w, http.StatusCreated, item)
}

func ActualizarUsuario(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromRequest(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, "id invalido")
		return
	}
	var item models.Usuario
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		writeError(w, http.StatusBadRequest, "json invalido")
		return
	}
	if ok, err := existeRegistro(`"Rol"`, "id_rol", item.IdRol); err != nil {
		writeError(w, http.StatusBadRequest, "error al validar id_rol")
		return
	} else if !ok {
		writeError(w, http.StatusBadRequest, "id_rol no existe")
		return
	}
	row := config.DB.QueryRow(`UPDATE "Usuario" SET nombre_completo = $1, correo = $2, telefono = $3, id_rol = $4, verificacion_dos_pasos = $5, avatar = $6, activo = $7 WHERE id_usuario = $8 RETURNING id_usuario, nombre_completo, correo, telefono, id_rol, verificacion_dos_pasos, avatar, activo, fecha_creacion, fecha_modificacion`,
		item.NombreCompleto, item.Correo, item.Telefono, item.IdRol, item.VerificacionDosPasos, item.Avatar, item.Activo, id)
	item, err = scanUsuario(row.Scan)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			writeError(w, http.StatusNotFound, "usuario no encontrado")
			return
		}
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			writeError(w, http.StatusConflict, "correo ya existe")
			return
		}
		writeError(w, http.StatusInternalServerError, "error al actualizar usuario")
		return
	}
	writeJSON(w, http.StatusOK, item)
}

func EliminarUsuario(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromRequest(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, "id invalido")
		return
	}
	result, err := config.DB.Exec(`DELETE FROM "Usuario" WHERE id_usuario = $1`, id)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23503" {
			writeError(w, http.StatusConflict, "no se puede eliminar usuario porque tiene registros relacionados")
			return
		}
		writeError(w, http.StatusInternalServerError, "error al eliminar usuario")
		return
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		writeError(w, http.StatusNotFound, "usuario no encontrado")
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"message": "usuario eliminado correctamente"})
}

func scanUsuario(scan func(dest ...any) error) (models.Usuario, error) {
	var item models.Usuario
	var telefono sql.NullString
	var avatar sql.NullString
	err := scan(&item.IdUsuario, &item.NombreCompleto, &item.Correo, &telefono, &item.IdRol, &item.VerificacionDosPasos, &avatar, &item.Activo, &item.FechaCreacion, &item.FechaModificacion)
	if err != nil {
		return models.Usuario{}, err
	}
	item.Telefono = nullStringToPointer(telefono)
	item.Avatar = nullStringToPointer(avatar)
	return item, nil
}
