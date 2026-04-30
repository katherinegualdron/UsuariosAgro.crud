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

func ObtenerPerfilesExtendidos(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query(`SELECT id_perfil, id_usuario, direccion, ciudad, intereses, redes_sociales, fecha_creacion, fecha_modificacion FROM "PerfilExtendido" ORDER BY id_perfil`)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "error al consultar perfil_extendido")
		return
	}
	defer rows.Close()
	items := []models.PerfilExtendido{}
	for rows.Next() {
		item, err := scanPerfilExtendido(rows.Scan)
		if err != nil {
			writeError(w, http.StatusInternalServerError, "error al leer perfil_extendido")
			return
		}
		items = append(items, item)
	}
	writeJSON(w, http.StatusOK, items)
}

func ObtenerPerfilExtendidoPorID(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromRequest(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, "id invalido")
		return
	}
	row := config.DB.QueryRow(`SELECT id_perfil, id_usuario, direccion, ciudad, intereses, redes_sociales, fecha_creacion, fecha_modificacion FROM "PerfilExtendido" WHERE id_perfil = $1`, id)
	item, err := scanPerfilExtendido(row.Scan)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			writeError(w, http.StatusNotFound, "perfil_extendido no encontrado")
			return
		}
		writeError(w, http.StatusInternalServerError, "error al consultar perfil_extendido")
		return
	}
	writeJSON(w, http.StatusOK, item)
}

func CrearPerfilExtendido(w http.ResponseWriter, r *http.Request) {
	var item models.PerfilExtendido
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
	row := config.DB.QueryRow(`INSERT INTO "PerfilExtendido" (id_usuario, direccion, ciudad, intereses, redes_sociales) VALUES ($1, $2, $3, $4, $5) RETURNING id_perfil, id_usuario, direccion, ciudad, intereses, redes_sociales, fecha_creacion, fecha_modificacion`,
		item.IdUsuario, item.Direccion, item.Ciudad, item.Intereses, item.RedesSociales)
	item, err := scanPerfilExtendido(row.Scan)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			writeError(w, http.StatusConflict, "ya existe un perfil_extendido para ese usuario")
			return
		}
		writeError(w, http.StatusInternalServerError, "error al crear perfil_extendido")
		return
	}
	writeJSON(w, http.StatusCreated, item)
}

func ActualizarPerfilExtendido(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromRequest(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, "id invalido")
		return
	}
	var item models.PerfilExtendido
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
	row := config.DB.QueryRow(`UPDATE "PerfilExtendido" SET id_usuario = $1, direccion = $2, ciudad = $3, intereses = $4, redes_sociales = $5 WHERE id_perfil = $6 RETURNING id_perfil, id_usuario, direccion, ciudad, intereses, redes_sociales, fecha_creacion, fecha_modificacion`,
		item.IdUsuario, item.Direccion, item.Ciudad, item.Intereses, item.RedesSociales, id)
	item, err = scanPerfilExtendido(row.Scan)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			writeError(w, http.StatusNotFound, "perfil_extendido no encontrado")
			return
		}
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			writeError(w, http.StatusConflict, "ya existe un perfil_extendido para ese usuario")
			return
		}
		writeError(w, http.StatusInternalServerError, "error al actualizar perfil_extendido")
		return
	}
	writeJSON(w, http.StatusOK, item)
}

func EliminarPerfilExtendido(w http.ResponseWriter, r *http.Request) {
	eliminarGenericoUsuarios(w, r, `"PerfilExtendido"`, "id_perfil", "perfil_extendido")
}

func scanPerfilExtendido(scan func(dest ...any) error) (models.PerfilExtendido, error) {
	var item models.PerfilExtendido
	var direccion sql.NullString
	var ciudad sql.NullString
	var intereses sql.NullString
	var redes sql.NullString
	err := scan(&item.IdPerfil, &item.IdUsuario, &direccion, &ciudad, &intereses, &redes, &item.FechaCreacion, &item.FechaModificacion)
	if err != nil {
		return models.PerfilExtendido{}, err
	}
	item.Direccion = nullStringToPointer(direccion)
	item.Ciudad = nullStringToPointer(ciudad)
	item.Intereses = nullStringToPointer(intereses)
	item.RedesSociales = nullStringToPointer(redes)
	return item, nil
}

func eliminarGenericoUsuarios(w http.ResponseWriter, r *http.Request, tabla string, columna string, nombre string) {
	id, err := getIDFromRequest(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, "id invalido")
		return
	}
	result, err := config.DB.Exec("DELETE FROM "+tabla+" WHERE "+columna+" = $1", id)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "error al eliminar "+nombre)
		return
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		writeError(w, http.StatusNotFound, nombre+" no encontrado")
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"message": nombre + " eliminado correctamente"})
}
