package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
	h "github.com/panprogramadorgh/jsonwebtokenserver/internal/handlers"
	m "github.com/panprogramadorgh/jsonwebtokenserver/internal/middlewares"
	"github.com/panprogramadorgh/jsonwebtokenserver/internal/utils"
)

func main() {

	db, err := utils.ConnectDB("./database.db")
	if err != nil {
		fmt.Println(err)
		return
	}

	http.HandleFunc("GET /users", m.CheckDBMid(db, h.ImplementDBHandler(h.GetUsersHandler{DB: db})))

	http.HandleFunc("DELETE /users", m.CheckDBMid(db, h.ImplementDBHandler(h.DeleteUsersHandler{DB: db})))

	http.HandleFunc("POST /login", m.CheckDBMid(db, h.ImplementDBHandler(h.LoginHandler{DB: db})))

	http.HandleFunc("GET /profile", m.CheckDBMid(db, h.ImplementDBHandler(h.GetPofileHandler{DB: db})))

	http.HandleFunc("DELETE /profile", m.CheckDBMid(db, h.ImplementDBHandler(h.DeleteUserHandler{DB: db})))

	http.HandleFunc("POST /register", m.CheckDBMid(db, h.ImplementDBHandler(h.RegisterHandler{DB: db})))

	var p string = "3000"
	if len(os.Args) == 2 {
		p = os.Args[1]
	}
	fmt.Printf("server running on %s\n", p)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", p), nil); err != nil {
		log.Fatal(err)
	}
}
