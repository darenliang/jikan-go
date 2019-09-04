package jikan

import "fmt"

// Club struct defines a club
type Club struct {
	ID      int    // MyAnimeList Club ID
	Request string // Request type
	Page    int    // Page number (Available upon request type activation)
}

// GetClub returns a map of a club as specified in the Club struct
// Calls responses through the /club/ endpoint
func GetClub(club Club) (map[string]interface{}, error) {
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
		result, err = nil, fmt.Errorf("error %v, %v, %v", result["status"], result["message"], result["error"])
	}
	return result, err
}
