package server

import (
	"IPchecker/src"
	"encoding/json"
	"net/http"
)

func Init() *http.ServeMux {
	r := http.NewServeMux()
	r.HandleFunc("/", RootHandler)
	return r
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	response := src.GetCountryName("8.8.8.8")
	js, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
