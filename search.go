package jikan

import (
	"fmt"
	"strings"
)

// Search struct defines an search
type Search struct {
	Type string // object type

	// All below are optional
	Q            string  // search query
	Page         int     // page number
	TypeFormat   string  // media format
	Status       string  // status
	Rated        string  // age rating
	Genre        int     // MyAnimeList genre ID
	Score        float64 // score 0.0 - 10.0
	StartDate    string  // start date
	EndDate      string  // end date
	GenreExclude bool    // exclude genre filter
	Limit        int     // limit results
	OrderBy      string  // order by
	Sort         string  // sort
	Producer     int     // MyAnimeList producer ID
	Magazine     int     // MyAnimeList magazine ID
	Letter       rune    // starting letter
}

// Get returns a map of an search as specified in the Search struct
// Calls responses through the /search/ endpoint
func (search Search) Get() (map[string]interface{}, error) {
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
		result, err = nil, getResultError(result)
	}
	return result, err
}
