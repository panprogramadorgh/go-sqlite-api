package handlers

import (
	"database/sql"
	"net/http"
)

type DBHandler func(db *sql.DB, w http.ResponseWriter, r *http.Request)

func UseDBHandler(db *sql.DB, h DBHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h(db, w, r)
	}
}
