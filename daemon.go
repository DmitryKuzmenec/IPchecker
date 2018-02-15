package main

import (
	"IPchecker/config"
	"IPchecker/server"
	"log"
	"net/http"
	"time"

	"github.com/spf13/viper"
)

func main() {
	config.Init()
	httpServer := &http.Server{
		Addr:         viper.GetString("addr"),
		Handler:      server.Init(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
