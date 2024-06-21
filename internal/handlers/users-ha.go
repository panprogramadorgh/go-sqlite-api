package handlers

import (
	"database/sql"
	"net/http"

	"github.com/panprogramadorgh/jsonwebtokenserver/internal/utils"
)

type GetUsersHandler struct {
	DB *sql.DB
}

func (uh GetUsersHandler) Handle(w http.ResponseWriter, r *http.Request) {
	// Obtener los usuarios de la base de datos
	users, err := utils.GetUsers(uh.DB)
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
}

type DeleteUsersHandler struct {
	DB *sql.DB
}

func (duh DeleteUsersHandler) Handle(w http.ResponseWriter, r *http.Request) {
	// Eliminar todos los usuarios
	err := utils.DeleteUsers(duh.DB)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		utils.WriteJRes(w, map[string]interface{}{
			"error": "internal server error",
		})
		return
	}
	// Escribir respuesta exitosa
	utils.WriteJRes(w, map[string]interface{}{
		"message": "successful deletion",
	})
}
