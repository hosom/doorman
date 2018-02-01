package main

import (
	"log"
	"net/http"
	"time"

	"github.com/hosom/doorman/api"
)

func main() {
	apiHandler := api.NewAPI(time.Hour * 1)
	http.Handle("/blocklist", apiHandler)
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
