package jikan

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
	"time"
)

// SearchQuery struct
type SearchQuery struct {
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

// Search struct
type Search struct {
	Results []struct {
		MalID     int       `json:"mal_id"`
		URL       string    `json:"url"`
		ImageURL  string    `json:"image_url"`
		Title     string    `json:"title"`
		Airing    bool      `json:"airing"`
		Synopsis  string    `json:"synopsis"`
		Type      string    `json:"type"`
		Episodes  int       `json:"episodes"`
		Score     float64   `json:"score"`
		StartDate time.Time `json:"start_date"`
		EndDate   time.Time `json:"end_date"`
		Members   int       `json:"members"`
		Rated     string    `json:"rated"`
	} `json:"results"`
	LastPage int `json:"last_page"`
}

// GetSearch returns search
func GetSearch(q SearchQuery) (*Search, error) {
	res := &Search{}

	if q.Type == "" {
		return nil, errors.New("search type not specified")
	}

	if q.Q == "" {
		return nil, errors.New("search query not specified")
	}

	var query strings.Builder
	query.WriteString(fmt.Sprintf("/search/%v/?", q.Type))
	if q.Q != "" {
		query.WriteString(fmt.Sprintf("q=%v&", url.QueryEscape(q.Q)))
	}
	if q.Page != 0 {
		query.WriteString(fmt.Sprintf("page=%v&", q.Page))
	}
	if q.TypeFormat != "" {
		query.WriteString(fmt.Sprintf("type=%v&", q.TypeFormat))
	}
	if q.Status != "" {
		query.WriteString(fmt.Sprintf("status=%v&", q.Status))
	}
	if q.Rated != "" {
		query.WriteString(fmt.Sprintf("rated=%v&", q.Rated))
	}
	if q.Genre != 0 {
		query.WriteString(fmt.Sprintf("genre=%v&", q.Genre))
	}
	if q.Score != 0 {
		query.WriteString(fmt.Sprintf("score=%v&", q.Score))
	}
	if q.StartDate != "" {
		query.WriteString(fmt.Sprintf("start_date=%v&", q.StartDate))
	}
	if q.EndDate != "" {
		query.WriteString(fmt.Sprintf("end_date=%v&", q.EndDate))
	}
	if q.GenreExclude != false {
		query.WriteString(fmt.Sprintf("genre_exclude=%v&", 1))
	}
	if q.Limit != 0 {
		query.WriteString(fmt.Sprintf("limit=%v&", q.Limit))
	}
	if q.OrderBy != "" {
		query.WriteString(fmt.Sprintf("order_by=%v&", q.OrderBy))
	}
	if q.Sort != "" {
		query.WriteString(fmt.Sprintf("sort=%v&", q.Sort))
	}
	if q.Producer != 0 {
		query.WriteString(fmt.Sprintf("producer=%v&", q.Producer))
	}
	if q.Magazine != 0 {
		query.WriteString(fmt.Sprintf("magazine=%v&", q.Magazine))
	}
	if q.Letter != 0 {
		query.WriteString(fmt.Sprintf("letter=%v&", q.Letter))
	}

	err := urlToStruct(query.String(), res)

	if err != nil {
		return nil, err
	}

	return res, nil
}
