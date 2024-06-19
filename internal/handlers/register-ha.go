package handlers

import (
	"bufio"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/panprogramadorgh/jsonwebtokenserver/internal/dbutils"
	"github.com/panprogramadorgh/jsonwebtokenserver/internal/utils"
)

func RegisterHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// Leer el cuerpo de la req
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
			}
		}
		var user utils.User
		if err := json.Unmarshal([]byte(body), &user); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			if _, err := w.Write([]byte("internal server error")); err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(err)
			return
		}
		if err := dbutils.PostUser(db, user); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			if _, err := w.Write([]byte("internal server error")); err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(err)
			return
		}
		resMap := map[string]interface{}{
			"message": "user created successfully",
		}
		resJson, err := json.Marshal(resMap)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			if _, err := w.Write([]byte("internal server error")); err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(err)
			return
		}
		if _, err := w.Write(resJson); err != nil {
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
			if _, err := w.Write(jsonRes); err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(err)
		}
	}
}
