package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/panprogramadorgh/jsonwebtokenserver/internal/dbutils"
)

func UsersHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// Obtener los usuarios de la base de datos
		users, err := dbutils.GetUsers(db)
		if err != nil {
			if _, err := w.Write([]byte(fmt.Sprint(err))); err != nil {
				fmt.Println(err)
				return
			}
			return
		}

		// Convertir los usuarios en formato json
		if jsonUsers, err := json.Marshal(users); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			if _, err := w.Write([]byte("internal server error")); err != nil {
				fmt.Println(err)
				return
			}
		} else {
			if _, err := w.Write(jsonUsers); err != nil {
				fmt.Println(err)
				return
			}
		}
	} else if r.Method == "DELETE" {
		// Eliminar todos los usuarios
		if err := dbutils.DeleteUsers(db); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			if _, err := w.Write([]byte("internal server error")); err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(err)
		}
		mapRes := map[string]interface{}{
			"message": "successful deletion",
		}
		jsonRes, err := json.Marshal(mapRes)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			if _, err := w.Write([]byte("internal server error")); err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(err)
			return
		}
		if _, err := w.Write(jsonRes); err != nil {
			fmt.Println(err)
			return
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
		mapRes := map[string]interface{}{
			"error": "not found",
		}
		jsonRes, err := json.Marshal(mapRes)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			if _, err := w.Write([]byte("internal server error")); err != nil {
				fmt.Println(err)
				return
			}
			return
		}
		if _, err := w.Write(jsonRes); err != nil {
			fmt.Println(err)
			return
		}
	}
}
