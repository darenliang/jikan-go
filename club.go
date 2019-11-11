package jikan

import "fmt"

// Club struct defines a club
type Club struct {
	ID      int    // MyAnimeList club ID
	Request string // request type
	Page    int    // page number (available upon request type activation)
}

// Get returns a map of a club as specified in the Club struct
// Calls responses through the /club/ endpoint
func (club Club) Get() (map[string]interface{}, error) {
	var result map[string]interface{}
	var err error
	switch club.Request {
	case
		"members":
		result, err = getMapFromUrl(fmt.Sprintf("/anime/%v/%v/%v", club.ID, club.Request, club.Page)), nil
	default:
		result, err = getMapFromUrl(fmt.Sprintf("/anime/%v", club.ID)), nil
	}
	if _, ok := result["error"]; ok {
		result, err = nil, getResultError(result)
	}
	return result, err
}
