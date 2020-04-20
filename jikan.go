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
