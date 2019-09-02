package jikan

import (
	"fmt"
	"strings"
)

// Search struct defines an search
type Search struct {
	Type string // Object type

	// All below are optional
	Q            string  // Search query
	Page         int     // Page number
	TypeFormat   string  // Media format
	Status       string  // Status
	Rated        string  // Age rating
	Genre        int     // MyAnimeList Genre ID
	Score        float64 // Score 0.0 - 10.0
	StartDate    string  // Start date
	EndDate      string  // End date
	GenreExclude bool    // Exclude genre filter
	Limit        int     // Limit results
	OrderBy      string  // Order by
	Sort         string  // Sort
	Producer     int     // MyAnimeList Producer ID
	Magazine     int     // MyAnimeList Magazine ID
	Letter       rune    // Starting letter
}

// GetSearch returns a map of an search as specified in the Search struct
func GetSearch(search Search) (map[string]interface{}, error) {
	var result map[string]interface{}
	var err error
	var query strings.Builder
	query.WriteString(fmt.Sprintf("/search/%v/?", search.Type))
	if search.Q != "" {
		query.WriteString(fmt.Sprintf("q=%v&", search.Q))
	}
	if search.Page != 0 {
		query.WriteString(fmt.Sprintf("page=%v&", search.Page))
	}
	if search.TypeFormat != "" {
		query.WriteString(fmt.Sprintf("type=%v&", search.TypeFormat))
	}
	if search.Status != "" {
		query.WriteString(fmt.Sprintf("status=%v&", search.Status))
	}
	if search.Rated != "" {
		query.WriteString(fmt.Sprintf("rated=%v&", search.Rated))
	}
	if search.Genre != 0 {
		query.WriteString(fmt.Sprintf("genre=%v&", search.Genre))
	}
	if search.Score != 0 {
		query.WriteString(fmt.Sprintf("score=%v&", search.Score))
	}
	if search.StartDate != "" {
		query.WriteString(fmt.Sprintf("start_date=%v&", search.StartDate))
	}
	if search.EndDate != "" {
		query.WriteString(fmt.Sprintf("end_date=%v&", search.EndDate))
	}
	if search.GenreExclude != false {
		query.WriteString(fmt.Sprintf("genre_exclude=%v&", 1))
	}
	if search.Limit != 0 {
		query.WriteString(fmt.Sprintf("limit=%v&", search.Limit))
	}
	if search.OrderBy != "" {
		query.WriteString(fmt.Sprintf("order_by=%v&", search.OrderBy))
	}
	if search.Sort != "" {
		query.WriteString(fmt.Sprintf("sort=%v&", search.Sort))
	}
	if search.Producer != 0 {
		query.WriteString(fmt.Sprintf("producer=%v&", search.Producer))
	}
	if search.Magazine != 0 {
		query.WriteString(fmt.Sprintf("magazine=%v&", search.Magazine))
	}
	if search.Letter != 0 {
		query.WriteString(fmt.Sprintf("letter=%v&", search.Letter))
	}
	result, err = getMapFromUrl(query.String()), nil
	if _, ok := result["error"]; ok {
		result, err = nil, fmt.Errorf("error %v, %v, %v", result["status"], result["message"], result["error"])
	}
	return result, err
}
