package main

import (
	"net/http"
	"time"

	"github.com/hosom/doorman/api"
)

func main() {
	apiHandler := api.NewAPI(time.Hour * 1)
	http.Handle("/blocklist", apiHandler)
	http.ListenAndServe(":80", nil)
}
