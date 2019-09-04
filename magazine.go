package jikan

import (
	"fmt"
	"strings"
)

// Magazine struct defines a magazine
type Magazine struct {
	MagazineID int // MyAnimeList Magazine ID
	Page       int // Page number (Optional)
}

// GetMagazine returns a map of a magazine as specified in the Magazine struct
// Calls responses through the /magazine/ endpoint
func GetMagazine(magazine Magazine) (map[string]interface{}, error) {
	var result map[string]interface{}
	var err error
	var query strings.Builder
	query.WriteString(fmt.Sprintf("/magazine/%v", magazine.MagazineID))
	if magazine.Page != 0 {
		query.WriteString(fmt.Sprintf("/%v", magazine.Page))
	}
	result, err = getMapFromUrl(query.String()), nil
	if _, ok := result["error"]; ok {
		result, err = nil, fmt.Errorf("error %v, %v, %v", result["status"], result["message"], result["error"])
	}
	return result, err
}
