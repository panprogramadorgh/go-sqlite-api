package handlers

import (
	"database/sql"
	"net/http"

	"github.com/panprogramadorgh/goquickjwt/pkg/jwt"
	"github.com/panprogramadorgh/jsonwebtokenserver/internal/dbutils"
	"github.com/panprogramadorgh/jsonwebtokenserver/internal/utils"
)

func ProfileHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	// Comprueba que el metodo de autenticacion sea el adecuado
	if r.Method == "GET" {
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

		// Obtiene los usuarios de la base de datos
		users, err := dbutils.GetUsers(db)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			utils.WriteJRes(w, map[string]interface{}{
				"error": "unauthorized",
			})
			return
		}

		// Obtiene el nombre de usuario del payload del token
		username := p["Username"]

		if v, ok := username.(string); ok {
			userI := utils.Users(users).IndexOf(v)
			user := users[userI]
			utils.WriteJRes(w, map[string]interface{}{
				"user": user,
			})
			return
		}
		// En caso de no obtener el resultado esperado del payload
		w.WriteHeader(http.StatusUnauthorized)
		utils.WriteJRes(w, map[string]interface{}{
			"error": "unauthorized",
		})
	} else {
		http.NotFound(w, r)
	}
}
