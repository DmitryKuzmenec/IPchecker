package server

import (
	"IPchecker/Sources/test"
	"IPchecker/Sources/test2"
	"IPchecker/Sources/test3"
	"IPchecker/types"
	"encoding/json"
	"net/http"
)

var RequestChan chan string
var ResponseChan chan types.ResolverResponse

func init() {
	RequestChan = make(chan string)
	ResponseChan = make(chan types.ResolverResponse)

	test.Init(RequestChan, ResponseChan)
	test2.Init(RequestChan, ResponseChan)
	test3.Init(RequestChan, ResponseChan)

}

func Init() *http.ServeMux {
	r := http.NewServeMux()
	r.HandleFunc("/", RootHandler)
	return r
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	RequestChan <- r.RemoteAddr
	select {
	case response := <-ResponseChan:
		js, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}
}
