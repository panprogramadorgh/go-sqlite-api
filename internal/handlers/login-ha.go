package handlers

import (
	"bufio"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/panprogramadorgh/goquickjwt/pkg/jwt"
	"github.com/panprogramadorgh/jsonwebtokenserver/internal/dbutils"
	"github.com/panprogramadorgh/jsonwebtokenserver/internal/utils"
)

func PostLoginHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	// Comprobar el metodo de la solicitud
	if r.Method != "POST" {
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

	// Leer el cuerpo de la solicitud
	reader := bufio.NewReader(r.Body)
	body := ""
	for {
		line, err := reader.ReadString('\n')
		body += line
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			return
		}
	}

	// Convertir el cuerpo (formato json) en tipo Credentials
	var credentials utils.Credentials
	if err := json.Unmarshal([]byte(body), &credentials); err != nil {
		fmt.Println(err)
		return
	}

	// Obtener los usuarios de la base de datos
	users, err := dbutils.GetUsers(db)
	if err != nil {
		if _, err := w.Write([]byte(fmt.Sprint(err))); err != nil {
			fmt.Println(err)
			return
		}
		return
	}

	// Autenticar usuario
	auth := utils.Users(users).AuthUser(credentials.Username, credentials.Password)

	if auth {
		// Si el usuario esta autenticado se genera un payload para el token
		p := jwt.Payload{
			"Username": credentials.Username,
		}

		// Se firma el token
		token, err := p.SignWithHS256(utils.Secret)
		if err != nil {
			if _, err := w.Write([]byte(err.Error())); err != nil {
				fmt.Println(err)
				return
			}
			return
		}

		// Se devuelve una respuesta con el token
		resMap := map[string]interface{}{
			"token": token,
		}
		resJson, err := json.Marshal(resMap)
		if err != nil {
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
	// Si el usuario no esta autenticado entonces se devuelve un mensaje de error
	if _, err := w.Write([]byte("invalid credentials")); err != nil {
		fmt.Println(err)
		return
	}
}
