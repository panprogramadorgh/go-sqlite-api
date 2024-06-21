package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/panprogramadorgh/jsonwebtokenserver/internal/dbutils"
	h "github.com/panprogramadorgh/jsonwebtokenserver/internal/handlers"
	m "github.com/panprogramadorgh/jsonwebtokenserver/internal/middlewares"
)

func main() {

	db, err := dbutils.ConnectDB()
	if err != nil {
		fmt.Println(err)
		return
	}

	http.HandleFunc("/users", m.CheckDBMid(db, h.UseDBHandler(db, h.UsersHandler)))

	http.HandleFunc("/login", m.CheckDBMid(db, h.UseDBHandler(db, h.LoginHandler)))

	http.HandleFunc("/profile", m.CheckDBMid(db, h.UseDBHandler(db, h.ProfileHandler)))

	http.HandleFunc("/register", m.CheckDBMid(db, h.UseDBHandler(db, h.RegisterHandler)))

	var p string = "3000"
	if len(os.Args) == 2 {
		p = os.Args[1]
	}
	fmt.Printf("server running on %s\n", p)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", p), nil); err != nil {
		log.Fatal(err)
	}
}
