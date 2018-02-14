package server

import (
	"io"
	"net/http"
)

func Init() *http.ServeMux {
	r := http.NewServeMux()
	r.HandleFunc("/", RootHandler)
	return r
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "OK")
}
