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

func ObtenerTokensRecuperacion(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query(`SELECT id_token, token, id_usuario, expiracion, usado, fecha_creacion FROM "TokenRecuperacion" ORDER BY id_token`)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "error al consultar token_recuperacion")
		return
	}
	defer rows.Close()
	items := []models.TokenRecuperacion{}
	for rows.Next() {
		var item models.TokenRecuperacion
		if err := rows.Scan(&item.IdToken, &item.Token, &item.IdUsuario, &item.Expiracion, &item.Usado, &item.FechaCreacion); err != nil {
			writeError(w, http.StatusInternalServerError, "error al leer token_recuperacion")
			return
		}
		items = append(items, item)
	}
	writeJSON(w, http.StatusOK, items)
}

func ObtenerTokenRecuperacionPorID(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromRequest(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, "id invalido")
		return
	}
	var item models.TokenRecuperacion
	err = config.DB.QueryRow(`SELECT id_token, token, id_usuario, expiracion, usado, fecha_creacion FROM "TokenRecuperacion" WHERE id_token = $1`, id).
		Scan(&item.IdToken, &item.Token, &item.IdUsuario, &item.Expiracion, &item.Usado, &item.FechaCreacion)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			writeError(w, http.StatusNotFound, "token_recuperacion no encontrado")
			return
		}
		writeError(w, http.StatusInternalServerError, "error al consultar token_recuperacion")
		return
	}
	writeJSON(w, http.StatusOK, item)
}

func CrearTokenRecuperacion(w http.ResponseWriter, r *http.Request) {
	var item models.TokenRecuperacion
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
	err := config.DB.QueryRow(`INSERT INTO "TokenRecuperacion" (token, id_usuario, expiracion, usado) VALUES ($1, $2, $3, $4) RETURNING id_token, token, id_usuario, expiracion, usado, fecha_creacion`,
		item.Token, item.IdUsuario, item.Expiracion, item.Usado).Scan(&item.IdToken, &item.Token, &item.IdUsuario, &item.Expiracion, &item.Usado, &item.FechaCreacion)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			writeError(w, http.StatusConflict, "token ya existe")
			return
		}
		writeError(w, http.StatusInternalServerError, "error al crear token_recuperacion")
		return
	}
	writeJSON(w, http.StatusCreated, item)
}

func ActualizarTokenRecuperacion(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromRequest(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, "id invalido")
		return
	}
	var item models.TokenRecuperacion
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
	err = config.DB.QueryRow(`UPDATE "TokenRecuperacion" SET token = $1, id_usuario = $2, expiracion = $3, usado = $4 WHERE id_token = $5 RETURNING id_token, token, id_usuario, expiracion, usado, fecha_creacion`,
		item.Token, item.IdUsuario, item.Expiracion, item.Usado, id).Scan(&item.IdToken, &item.Token, &item.IdUsuario, &item.Expiracion, &item.Usado, &item.FechaCreacion)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			writeError(w, http.StatusNotFound, "token_recuperacion no encontrado")
			return
		}
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			writeError(w, http.StatusConflict, "token ya existe")
			return
		}
		writeError(w, http.StatusInternalServerError, "error al actualizar token_recuperacion")
		return
	}
	writeJSON(w, http.StatusOK, item)
}

func EliminarTokenRecuperacion(w http.ResponseWriter, r *http.Request) {
	eliminarGenericoUsuarios(w, r, `"TokenRecuperacion"`, "id_token", "token_recuperacion")
}
