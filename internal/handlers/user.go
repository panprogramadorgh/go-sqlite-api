package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/panprogramadorgh/goquickjwt/pkg/jwt"
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

	// Comprobar que el usuario no tenga un username ya usado
	if user, err := utils.GetUser(rh.DB, userPay.Username); user != nil {
		w.WriteHeader(http.StatusBadRequest)
		utils.WriteJRes(w, map[string]interface{}{
			"error": "username already taken",
		})
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		utils.WriteJRes(w, map[string]interface{}{
			"error": "internal server error",
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

	newUser, err := utils.GetUser(rh.DB, userPay.Username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		utils.WriteJRes(w, map[string]interface{}{
			"error": "internal server error",
		})
		return
	}

	utils.WriteJRes(w, map[string]interface{}{
		"message": "user created successfully",
		"user":    newUser,
	})
}

type GetPofileHandler struct {
	DB *sql.DB
}

func (ph GetPofileHandler) Handle(w http.ResponseWriter, r *http.Request) {
	// Comprueba que la cabecera con el token sea valida
	token := utils.ReadReqHeader(r, "Authorization")
	if token == nil {
		w.WriteHeader(http.StatusUnauthorized)
		utils.WriteJRes(w, map[string]interface{}{
			"error": "unauthorized",
		})
		return
	}

	// Verifica el token y obtiene el payload
	p, err := jwt.VerifyWithHS256(utils.Secret, *token)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		utils.WriteJRes(w, map[string]interface{}{
			"error": "unauthorized",
		})
		return
	}

	// Obtiene el nombre de usuario del payload del token
	username := p["Username"]

	if v, ok := username.(string); !ok {
		// En caso de no obtener el resultado esperado del payload
		w.WriteHeader(http.StatusBadRequest)
		utils.WriteJRes(w, map[string]interface{}{
			"error": "token is not valid",
		})
	} else {
		user, err := utils.GetUser(ph.DB, v)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			utils.WriteJRes(w, map[string]interface{}{
				"error": "internal server error",
			})
			return
		}

		utils.WriteJRes(w, map[string]interface{}{
			"user": user,
		})
		return
	}

}

type LoginHandler struct {
	DB *sql.DB
}

func (lh LoginHandler) Handle(w http.ResponseWriter, r *http.Request) {
	// Convertir el cuerpo (formato json) en tipo Credentials
	var credentials utils.Credentials
	if err := utils.ReadReqBody(r, &credentials); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		utils.WriteJRes(w, map[string]interface{}{
			"error": "internal server error",
		})
		return
	}

	// Autenticar usuario
	auth := utils.Auth(lh.DB, credentials.Username, credentials.Password)

	if !auth {
		// Si el usuario no esta autenticado entonces se devuelve un mensaje de error
		w.WriteHeader(http.StatusUnauthorized)
		utils.WriteJRes(w, map[string]interface{}{
			"error": "invalid credentials",
		})
		return
	}

	// Si el usuario esta autenticado se genera un payload para el token
	p := jwt.Payload{
		"Username": credentials.Username,
	}

	// Se firma el token
	token, err := p.SignWithHS256(utils.Secret)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		utils.WriteJRes(w, map[string]interface{}{
			"error": "internal server error",
		})
		return
	}

	utils.WriteJRes(w, map[string]interface{}{
		"token": token,
	})
}

type DeleteUserHandler struct {
	DB *sql.DB
}

func (duh DeleteUserHandler) Handle(w http.ResponseWriter, r *http.Request) {
	token := utils.ReadReqHeader(r, "Authorization")
	if token == nil {
		w.WriteHeader(http.StatusUnauthorized)
		utils.WriteJRes(w, map[string]any{
			"error": "unauthorized",
		})
		return
	}
	p, err := jwt.VerifyWithHS256(utils.Secret, *token)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		utils.WriteJRes(w, map[string]any{
			"error": "unauthorized",
		})
		return
	}

	username := p["Username"]
	if v, ok := username.(string); username == nil || !ok {
		w.WriteHeader(http.StatusBadRequest)
		utils.WriteJRes(w, map[string]any{
			"error": "token is invalid",
		})
		return
	} else {
		// Elimina el usuario
		if err := utils.DeleteUser(duh.DB, v); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			utils.WriteJRes(w, map[string]any{
				"error": "internal server error",
			})
		}
		utils.WriteJRes(w, map[string]any{
			"message": "user deletion successful",
		})
	}
}
