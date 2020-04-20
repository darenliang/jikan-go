// Package jikan provides Jikan API bindings for Go.
// Please note that all functions/structs are named in accordance to Jikan's endpoints/responses
package jikan

import (
	"net/http"
	"time"
)

// Client is a *http.Client
var Client *http.Client

func init() {
	Client = &http.Client{Timeout: ClientTimeout * time.Second}
}
