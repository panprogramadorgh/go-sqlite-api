package handlers

import (
	"database/sql"
	"net/http"

	"github.com/panprogramadorgh/jsonwebtokenserver/internal/dbutils"
	"github.com/panprogramadorgh/jsonwebtokenserver/internal/utils"
)

func UsersHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// Obtener los usuarios de la base de datos
		users, err := dbutils.GetUsers(db)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			utils.WriteJRes(w, map[string]interface{}{
				"error": "internal server error",
			})
			return
		}
		utils.WriteJRes(w, map[string]interface{}{
			"useres": users,
		})
	} else if r.Method == "DELETE" {
		// Eliminar todos los usuarios
		err := dbutils.DeleteUsers(db)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			utils.WriteJRes(w, map[string]interface{}{
				"error": "internal server error",
			})
			return
		}

		utils.WriteJRes(w, map[string]interface{}{
			"message": "successful deletion",
		})
	} else {
		http.NotFound(w, r)
	}
}
