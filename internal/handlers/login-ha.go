package handlers

import (
	"database/sql"
	"net/http"

	"github.com/panprogramadorgh/goquickjwt/pkg/jwt"
	"github.com/panprogramadorgh/jsonwebtokenserver/internal/dbutils"
	"github.com/panprogramadorgh/jsonwebtokenserver/internal/utils"
)

func LoginHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	// Comprobar el metodo de la solicitud
	if r.Method == "POST" {
		// Convertir el cuerpo (formato json) en tipo Credentials
		var credentials utils.Credentials
		if err := utils.ReadReqBody(r, &credentials); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			utils.WriteJRes(w, map[string]interface{}{
				"error": "internal server error",
			})
			return
		}

		// Obtener los usuarios de la base de datos
		users, err := dbutils.GetUsers(db)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			utils.WriteJRes(w, map[string]interface{}{
				"error": "internal server error",
			})
			return
		}

		// Autenticar usuario
		auth := users.Auth(credentials.Username, credentials.Password)

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

	} else {
		http.NotFound(w, r)
	}
}
