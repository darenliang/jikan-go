package jikan

import "fmt"

// Anime struct defines an anime
type Anime struct {
	ID      int    // MyAnimeList Anime ID
	Request string // Request type (Optional)
	Page    int    // Page number (Available depending on request type)
}

// GetAnime returns a map of an anime as specified in the Anime struct
// Calls responses through the /anime/ endpoint
func GetAnime(anime Anime) (map[string]interface{}, error) {
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
		result, err = nil, fmt.Errorf("error %v, %v, %v", result["status"], result["message"], result["error"])
	}
	return result, err
}
