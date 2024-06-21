package handlers

import (
	"net/http"
)

type DBHandler interface {
	Handle(w http.ResponseWriter, r *http.Request)
}

func ImplementDBHandler(h DBHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h.Handle(w, r)
	}
}
