package handlers

import (
	"database/sql"
	"net/http"

	"github.com/panprogramadorgh/jsonwebtokenserver/internal/dbutils"
	"github.com/panprogramadorgh/jsonwebtokenserver/internal/utils"
)

func RegisterHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		// Convertir el cuerpo de la solicitud en un struct utils.User
		var userPay utils.UserPayload
		if err := utils.ReadReqBody(r, &userPay); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			utils.WriteJRes(w, map[string]interface{}{
				"error": "internal server error",
			})
			return
		}

		// Comprobacion de formato de los campos del cuerpo de la solicitud
		if len(userPay.Username) < 3 {
			w.WriteHeader(http.StatusBadRequest)
			utils.WriteJRes(w, map[string]interface{}{
				"error": "bad request",
			})
			return
		}
		if len(userPay.Password) < 5 {
			w.WriteHeader(http.StatusBadRequest)
			utils.WriteJRes(w, map[string]interface{}{
				"error": "bad request",
			})
			return
		}
		if len(userPay.Firstname) < 3 {
			w.WriteHeader(http.StatusBadRequest)
			utils.WriteJRes(w, map[string]interface{}{
				"error": "bad request",
			})
			return
		}
		if len(userPay.Lastname) < 3 {
			w.WriteHeader(http.StatusBadRequest)
			utils.WriteJRes(w, map[string]interface{}{
				"error": "bad request",
			})
			return
		}

		// Comprobar la existencia del username
		users, err := dbutils.GetUsers(db)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			utils.WriteJRes(w, map[string]interface{}{
				"error": "internal server error",
			})
			return
		}

		if users.IndexOf(userPay.Username) != -1 {
			w.WriteHeader(http.StatusBadRequest)
			utils.WriteJRes(w, map[string]interface{}{
				"error": "username already taken",
			})
			return
		}

		// Subir usuario a la base de datos
		if err := dbutils.PostUser(db, userPay); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			utils.WriteJRes(w, map[string]interface{}{
				"error": "internal server error",
			})
			return
		}

		newUserIndex := users.IndexOf(userPay.Username)
		if newUserIndex == -1 {
			w.WriteHeader(http.StatusInternalServerError)
			utils.WriteJRes(w, map[string]interface{}{
				"error": "internal server error",
			})
			return
		}
		newUser := users[newUserIndex]

		utils.WriteJRes(w, map[string]interface{}{
			"message": "user created successfully",
			"user":    newUser,
		})

	} else {
		http.NotFound(w, r)
	}
}
