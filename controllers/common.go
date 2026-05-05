package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"USUARIOS/config"

	"github.com/gorilla/mux"
)

func getIDFromRequest(r *http.Request, key string) (int, error) {
	return strconv.Atoi(mux.Vars(r)[key])
}

func writeJSON(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(data)
}

func writeError(w http.ResponseWriter, statusCode int, message string) {
	writeJSON(w, statusCode, map[string]string{"error": message})
}

func nullStringToPointer(value sql.NullString) *string {
	if !value.Valid {
		return nil
	}
	text := value.String
	return &text
}

func nullTimeToPointer(value sql.NullTime) *time.Time {
	if !value.Valid {
		return nil
	}
	fecha := value.Time
	return &fecha
}

func existeRegistro(tabla string, columna string, id int) (bool, error) {
	query := "SELECT EXISTS (SELECT 1 FROM " + tabla + " WHERE " + columna + " = $1)"
	var existe bool
	err := config.DB.QueryRow(query, id).Scan(&existe)
	if err != nil {
		return false, err
	}
	return existe, nil
}
