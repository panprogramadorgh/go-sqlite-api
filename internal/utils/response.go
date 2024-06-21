package utils

import (
	"encoding/json"
	"net/http"
)

func WriteJRes(w http.ResponseWriter, res map[string]any) {
	w.Header().Add("Content-Type", "text/json")
	jsonRes, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write(jsonRes)
}
