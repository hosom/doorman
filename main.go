package main

import (
	"log"
	"net/http"
	"time"

	"github.com/hosom/doorman/api"
)

func main() {

	apiHandler := api.NewAPI(1 * time.Hour)

	servMux := http.NewServeMux()
	servMux.Handle("/blocklist", apiHandler)

	serv := &http.Server{
		Addr:         ":8000",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  30 * time.Second,
		Handler:      servMux,
	}

	if err := serv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
