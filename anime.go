package jikan

import "fmt"

// Anime struct defines an anime
type Anime struct {
	ID      int    // MyAnimeList anime ID
	Request string // request type (optional)
	Page    int    // page number (available depending on request type)
}

// Get returns a map of an anime as specified in the Anime struct
// Calls responses through the /anime/ endpoint
func (anime Anime) Get() (map[string]interface{}, error) {
	var result map[string]interface{}
	var err error
	switch anime.Request {
	case
		"episodes",
		"reviews",
		"userupdates":
		result, err = getMapFromUrl(fmt.Sprintf("/anime/%v/%v/%v", anime.ID, anime.Request, anime.Page)), nil
	default:
		result, err = getMapFromUrl(fmt.Sprintf("/anime/%v/%v", anime.ID, anime.Request)), nil
	}
	if _, ok := result["error"]; ok {
		result, err = nil, getResultError(result)
	}
	return result, err
}
