package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/hosom/go-ttlmap"
)

// message is an API message received from a client
type message struct {
	indicator string
	ttl       string
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

	m := a.blocklist.GetAll()
	var indicators []string
	for indicator := range m {
		indicators = append(indicators, indicator.(string))
	}

	response := strings.Join(indicators, "\n")

	fmt.Fprint(w, response)
}

func (a *API) post(w http.ResponseWriter, r *http.Request) {
	msg, _ := parseMessage(r)
	if msg != nil {
		// if the message parses properly, process it
		if msg.indicator != "" {
			ttl, err := time.ParseDuration(msg.ttl)
			if err != nil {
				ttl = a.ttl
			}

			a.blocklist.AddWithTTL(msg.indicator, msg, ttl)
		}
	} else {
		http.Error(w, "", http.StatusBadRequest)
	}
}
