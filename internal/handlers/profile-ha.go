package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/panprogramadorgh/goquickjwt/pkg/jwt"
	"github.com/panprogramadorgh/jsonwebtokenserver/internal/dbutils"
	"github.com/panprogramadorgh/jsonwebtokenserver/internal/utils"
)

func GetProfileHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	// Comprueba que el metodo de autenticacion sea el adecuado
	if r.Method != "GET" {
		w.WriteHeader(http.StatusNotFound)
		resMap := map[string]interface{}{
			"error": "not found",
		}
		resJson, err := json.Marshal(resMap)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			if _, err := w.Write([]byte("internal server error")); err != nil {
				fmt.Println(err)
				return
			}
			return
		}
		if _, err := w.Write(resJson); err != nil {
			fmt.Println(err)
			return
		}
		return
	}

	// Comprueba que la cabecera con el token sea valida
	token := strings.Trim(r.Header.Get("Authorization"), " ")
	if token == "" {
		w.WriteHeader(http.StatusUnauthorized)
		resMap := map[string]interface{}{
			"error": "unauthorized",
		}
		resJson, err := json.Marshal(resMap)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			if _, err := w.Write([]byte("internal server error")); err != nil {
				fmt.Println(err)
				return
			}
			return
		}
		if _, err := w.Write(resJson); err != nil {
			fmt.Println(err)
			return
		}
		return
	}

	// Obtiene los usuarios de la base de datos
	users, err := dbutils.GetUsers(db)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := w.Write([]byte("internal server error")); err != nil {
			fmt.Println(err)
			return
		}
		return
	}

	// Verifica el token y obtiene el payload
	p, err := jwt.VerifyWithHS256(utils.Secret, token)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		if _, err := w.Write([]byte(err.Error())); err != nil {
			fmt.Println(err)
			return
		}
		return
	}

	// Obtiene el nombre de usuario del payload del token
	username := p["Username"]
	if v, ok := username.(string); ok {
		userI := utils.Users(users).IndexOfUserPerUsername(v)
		user := users[userI]
		jsonUser, _ := json.Marshal(user)
		if _, err := w.Write(jsonUser); err != nil {
			fmt.Println(err)
			return
		}
		return
	}

	// Si el campo Username no es el que se espeara el usuario no esta autorizado
	w.WriteHeader(http.StatusUnauthorized)
	resMap := map[string]interface{}{
		"error": "unauthorized",
	}
	resJson, err := json.Marshal(resMap)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := w.Write([]byte("internal server error")); err != nil {
			fmt.Println(err)
			return
		}
		return
	}
	if _, err := w.Write(resJson); err != nil {
		fmt.Println(err)
		return
	}
}
