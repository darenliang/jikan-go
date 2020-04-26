package jikan

import (
	"fmt"
	"net/url"
	"time"
)

// SearchAnime struct for the /search/anime endpoint
type SearchAnime struct {
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

// SearchManga struct for the /search/manga endpoint
type SearchManga struct {
	Results []struct {
		MalID      int       `json:"mal_id"`
		URL        string    `json:"url"`
		ImageURL   string    `json:"image_url"`
		Title      string    `json:"title"`
		Publishing bool      `json:"publishing"`
		Synopsis   string    `json:"synopsis"`
		Type       string    `json:"type"`
		Chapters   int       `json:"chapters"`
		Volumes    int       `json:"volumes"`
		Score      float64   `json:"score"`
		StartDate  time.Time `json:"start_date"`
		EndDate    time.Time `json:"end_date"`
		Members    int       `json:"members"`
	} `json:"results"`
	LastPage int `json:"last_page"`
}

// SearchPeople struct for the /search/people endpoint
type SearchPeople struct {
	Results []struct {
		MalID            int      `json:"mal_id"`
		URL              string   `json:"url"`
		ImageURL         string   `json:"image_url"`
		Name             string   `json:"name"`
		AlternativeNames []string `json:"alternative_names"`
	} `json:"results"`
	LastPage int `json:"last_page"`
}

// SearchCharacter struct for the /search/character endpoint
type SearchCharacter struct {
	Results []struct {
		MalID            int       `json:"mal_id"`
		URL              string    `json:"url"`
		ImageURL         string    `json:"image_url"`
		Name             string    `json:"name"`
		AlternativeNames []string  `json:"alternative_names"`
		Anime            []MalItem `json:"anime"`
		Manga            []MalItem `json:"manga"`
	} `json:"results"`
	LastPage int `json:"last_page"`
}

// GetSearchAnime returns search anime
func GetSearchAnime(query url.Values) (*SearchAnime, error) {
	res := &SearchAnime{}

	err := urlToStruct(fmt.Sprintf("/search/anime?%s", query.Encode()), res)

	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetSearchManga returns search manga
func GetSearchManga(query url.Values) (*SearchManga, error) {
	res := &SearchManga{}

	err := urlToStruct(fmt.Sprintf("/search/manga?%s", query.Encode()), res)

	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetSearchPeople returns search people
func GetSearchPeople(query url.Values) (*SearchPeople, error) {
	res := &SearchPeople{}

	err := urlToStruct(fmt.Sprintf("/search/people?%s", query.Encode()), res)

	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetSearchCharacter returns search character
func GetSearchCharacter(query url.Values) (*SearchCharacter, error) {
	res := &SearchCharacter{}

	err := urlToStruct(fmt.Sprintf("/search/character?%s", query.Encode()), res)

	if err != nil {
		return nil, err
	}

	return res, nil
}
