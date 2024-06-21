package middlewares

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/panprogramadorgh/jsonwebtokenserver/internal/utils"
)

func CheckDBMid(db *sql.DB, next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		if err := utils.CheckDB(db); err != nil {
			w.Write([]byte("cannot connect to database"))
			fmt.Println(err)
			return
		}
		next.ServeHTTP(w, r)
	}

	// if err := utils.CheckDB(db); err != nil {
	// 	if _, err := w.Write([]byte("cannot connec to database")); err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}
	// 	fmt.Println(err)
	// 	return
	// }
	// n(w, r)
}
