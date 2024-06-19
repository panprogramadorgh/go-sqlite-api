package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/panprogramadorgh/jsonwebtokenserver/internal/dbutils"
	h "github.com/panprogramadorgh/jsonwebtokenserver/internal/handlers"
	mid "github.com/panprogramadorgh/jsonwebtokenserver/internal/middlewares"
)

func main() {

	db, err := dbutils.Connect()
	if err != nil {
		fmt.Println(err)
		return
	}

	http.HandleFunc("/users", mid.CheckDBMid(db, h.HandleDBHandler(db, h.UsersHandler)))

	http.HandleFunc("/login", mid.CheckDBMid(db, h.HandleDBHandler(db, h.PostLoginHandler)))

	http.HandleFunc("/profile", mid.CheckDBMid(db, h.HandleDBHandler(db, h.GetProfileHandler)))

	http.HandleFunc("/register", mid.CheckDBMid(db, h.HandleDBHandler(db, h.RegisterHandler)))

	var p string = "3000"
	if len(os.Args) == 2 {
		p = os.Args[1]
	}
	fmt.Printf("server running on %s\n", p)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", p), nil); err != nil {
		log.Fatal(err)
	}
}
