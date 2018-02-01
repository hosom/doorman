package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/hosom/go-ttlmap"
)

// message is an API message received from a client
type message struct {
	Indicator string
	Ttl       string
}

// parseMessage accepts an HTTP request and returns a message
// struct to be used as a helper for API requests
func parseMessage(r *http.Request) (*message, error) {
	decoder := json.NewDecoder(r.Body)

	var msg message
	err := decoder.Decode(&msg)
	if err != nil {
		return nil, err
	}

	return &msg, err
}

// API provides the HTTP handler for managing the doorman list
type API struct {
	blocklist *ttlmap.TTLMap
	ttl       time.Duration
}

// NewAPI creates an API handler
func NewAPI(ttl time.Duration) *API {
	return &API{ttlmap.NewTTLMap(ttl), ttl}
}

// ServeHTTP is a required method for handlers
func (a *API) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request.")
	switch r.Method {
	default:
		http.Error(w, "", http.StatusNotImplemented)
	case "GET":
		a.get(w, r)
	case "POST":
		a.post(w, r)
	}
}

func (a *API) get(w http.ResponseWriter, r *http.Request) {
	log.Println("Processing request to retrieve blocklist.")
	m := a.blocklist.GetAll()
	var indicators []string
	for indicator := range m {
		indicators = append(indicators, indicator.(string))
	}
	log.Printf("Found %d entries to return", len(indicators))
	response := strings.Join(indicators, "\n")

	fmt.Fprint(w, response)
}

func (a *API) post(w http.ResponseWriter, r *http.Request) {
	log.Println("Processing request to add to blocklist.")
	msg, _ := parseMessage(r)
	if msg != nil {
		log.Printf("Processing indicator: %s, using ttl: %s", msg.Indicator, msg.Ttl)
		// if the message parses properly, process it
		if msg.Indicator != "" {
			ttl, err := time.ParseDuration(msg.Ttl)
			if err != nil {
				ttl = a.ttl
			}

			a.blocklist.AddWithTTL(msg.Indicator, msg, ttl)
			fmt.Fprint(w, "OK")
		}
	} else {
		http.Error(w, "", http.StatusBadRequest)
	}
}
