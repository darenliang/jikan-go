package jikan

import "fmt"

// Character struct defines a character
type Character struct {
	ID      int    // MyAnimeList Character ID
	Request string // Request type (Optional)
}

// Get returns a map of a character as specified in the Character struct
// Calls responses through the /character/ endpoint
func (character Character) Get() (map[string]interface{}, error) {
	var result map[string]interface{}
	var err error
	result, err = getMapFromUrl(fmt.Sprintf("/character/%v/%v", character.ID, character.Request)), nil
	if _, ok := result["error"]; ok {
		result, err = nil, getResultError(result)
	}
	return result, err
}
