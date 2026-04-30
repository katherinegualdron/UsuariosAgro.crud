package controllers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"USUARIOS/config"
	"USUARIOS/models"
)

func ObtenerVerificacionesDosPasos(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query(`SELECT id_verificacion, id_usuario, tipo_metodo, codigo_otp, expiracion_otp, activo, fecha_creacion, fecha_modificacion FROM "VerificacionDosPasos" ORDER BY id_verificacion`)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "error al consultar verificacion_dos_pasos")
		return
	}
	defer rows.Close()
	items := []models.VerificacionDosPasos{}
	for rows.Next() {
		item, err := scanVerificacionDosPasos(rows.Scan)
		if err != nil {
			writeError(w, http.StatusInternalServerError, "error al leer verificacion_dos_pasos")
			return
		}
		items = append(items, item)
	}
	writeJSON(w, http.StatusOK, items)
}

func ObtenerVerificacionDosPasosPorID(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromRequest(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, "id invalido")
		return
	}
	row := config.DB.QueryRow(`SELECT id_verificacion, id_usuario, tipo_metodo, codigo_otp, expiracion_otp, activo, fecha_creacion, fecha_modificacion FROM "VerificacionDosPasos" WHERE id_verificacion = $1`, id)
	item, err := scanVerificacionDosPasos(row.Scan)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			writeError(w, http.StatusNotFound, "verificacion_dos_pasos no encontrado")
			return
		}
		writeError(w, http.StatusInternalServerError, "error al consultar verificacion_dos_pasos")
		return
	}
	writeJSON(w, http.StatusOK, item)
}

func CrearVerificacionDosPasos(w http.ResponseWriter, r *http.Request) {
	var item models.VerificacionDosPasos
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
	row := config.DB.QueryRow(`INSERT INTO "VerificacionDosPasos" (id_usuario, tipo_metodo, codigo_otp, expiracion_otp, activo) VALUES ($1, $2, $3, $4, $5) RETURNING id_verificacion, id_usuario, tipo_metodo, codigo_otp, expiracion_otp, activo, fecha_creacion, fecha_modificacion`,
		item.IdUsuario, item.TipoMetodo, item.CodigoOtp, item.ExpiracionOtp, item.Activo)
	item, err := scanVerificacionDosPasos(row.Scan)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "error al crear verificacion_dos_pasos")
		return
	}
	writeJSON(w, http.StatusCreated, item)
}

func ActualizarVerificacionDosPasos(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromRequest(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, "id invalido")
		return
	}
	var item models.VerificacionDosPasos
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
	row := config.DB.QueryRow(`UPDATE "VerificacionDosPasos" SET id_usuario = $1, tipo_metodo = $2, codigo_otp = $3, expiracion_otp = $4, activo = $5 WHERE id_verificacion = $6 RETURNING id_verificacion, id_usuario, tipo_metodo, codigo_otp, expiracion_otp, activo, fecha_creacion, fecha_modificacion`,
		item.IdUsuario, item.TipoMetodo, item.CodigoOtp, item.ExpiracionOtp, item.Activo, id)
	item, err = scanVerificacionDosPasos(row.Scan)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			writeError(w, http.StatusNotFound, "verificacion_dos_pasos no encontrado")
			return
		}
		writeError(w, http.StatusInternalServerError, "error al actualizar verificacion_dos_pasos")
		return
	}
	writeJSON(w, http.StatusOK, item)
}

func EliminarVerificacionDosPasos(w http.ResponseWriter, r *http.Request) {
	eliminarGenericoUsuarios(w, r, `"VerificacionDosPasos"`, "id_verificacion", "verificacion_dos_pasos")
}

func scanVerificacionDosPasos(scan func(dest ...any) error) (models.VerificacionDosPasos, error) {
	var item models.VerificacionDosPasos
	var codigoOtp sql.NullString
	var expiracionOtp sql.NullTime
	err := scan(&item.IdVerificacion, &item.IdUsuario, &item.TipoMetodo, &codigoOtp, &expiracionOtp, &item.Activo, &item.FechaCreacion, &item.FechaModificacion)
	if err != nil {
		return models.VerificacionDosPasos{}, err
	}
	item.CodigoOtp = nullStringToPointer(codigoOtp)
	item.ExpiracionOtp = nullTimeToPointer(expiracionOtp)
	return item, nil
}
