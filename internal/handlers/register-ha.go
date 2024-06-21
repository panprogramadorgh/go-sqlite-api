package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/panprogramadorgh/jsonwebtokenserver/internal/utils"
)

type RegisterHandler struct {
	DB *sql.DB
}

func (rh RegisterHandler) Handle(w http.ResponseWriter, r *http.Request) {
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
	users, err := utils.GetUsers(rh.DB)
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

	// Hashear password del usuario
	if err := userPay.HashPass(); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		utils.WriteJRes(w, map[string]any{
			"error": "internal server error",
		})
		return
	}

	// Subir usuario a la base de datos
	if err := utils.PostUser(rh.DB, userPay); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		utils.WriteJRes(w, map[string]interface{}{
			"error": "internal server error",
		})
		return
	}

	users.ReloadData(rh.DB)
	newUserIndex := users.IndexOf(userPay.Username)
	newUser := users[newUserIndex]

	utils.WriteJRes(w, map[string]interface{}{
		"message": "user created successfully",
		"user":    newUser,
	})
}
