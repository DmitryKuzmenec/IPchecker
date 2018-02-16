package server

import (
	"IPchecker/Sources/freegeoip"
	"IPchecker/Sources/nekudo"
	"IPchecker/db"
	"IPchecker/types"
	"encoding/json"
	"net/http"
	"strings"
)

var RequestChan chan string
var ResponseChan chan types.ResolverResponse

func init() {
	RequestChan = make(chan string, 1000)
	ResponseChan = make(chan types.ResolverResponse, 1000)

	//Providers
	freegeoip.Init(RequestChan, ResponseChan)
	nekudo.Init(RequestChan, ResponseChan)

}

func Init() *http.ServeMux {
	r := http.NewServeMux()
	r.HandleFunc("/", RootHandler)
	return r
}

func RootHandler(w http.ResponseWriter, r *http.Request) {

	ip := strings.Split(r.RemoteAddr, ":")[0]
	response := types.ResolverResponse{}

	db := db.Init()
	// Check cache
	name := db.Get(ip)
	if name != "" {
		response = types.ResolverResponse{IP: ip, CountryName: name, Source: "cache"}
	} else {
		RequestChan <- ip
		select {
		case response = <-ResponseChan:
		}
	}

	js, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
