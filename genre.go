package jikan

import (
	"fmt"
	"strings"
)

// Genre struct defines a genre
type Genre struct {
	Type    string // Media type
	GenreID int    // MyAnimeList Genre ID
	Page    int    // Page number (Optional)
}

// Get returns a map of a genre as specified in the Genre struct
// Calls responses through the /genre/ endpoint
func (genre Genre) Get() (map[string]interface{}, error) {
	var result map[string]interface{}
	var err error
	var query strings.Builder
	query.WriteString(fmt.Sprintf("/genre/%v/%v", genre.Type, genre.GenreID))
	if genre.Page != 0 {
		query.WriteString(fmt.Sprintf("/%v", genre.Page))
	}
	result, err = getMapFromUrl(query.String()), nil
	if _, ok := result["error"]; ok {
		result, err = nil, getResultError(result)
	}
	return result, err
}
