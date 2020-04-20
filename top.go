package jikan

import (
	"fmt"
	"net/url"
	"time"
)

// TopAnime struct for the /top/anime endpoint
type TopAnime struct {
	Top []struct {
		MalID     int     `json:"mal_id"`
		Rank      int     `json:"rank"`
		Title     string  `json:"title"`
		URL       string  `json:"url"`
		ImageURL  string  `json:"image_url"`
		Type      string  `json:"type"`
		Episodes  int     `json:"episodes"`
		StartDate string  `json:"start_date"`
		EndDate   string  `json:"end_date"`
		Members   int     `json:"members"`
		Score     float64 `json:"score"`
	} `json:"top"`
}

// TopManga struct for the /top/manga endpoint
type TopManga struct {
	Top []struct {
		MalID     int     `json:"mal_id"`
		Rank      int     `json:"rank"`
		Title     string  `json:"title"`
		URL       string  `json:"url"`
		Type      string  `json:"type"`
		Volumes   int     `json:"volumes"`
		StartDate string  `json:"start_date"`
		EndDate   string  `json:"end_date"`
		Members   int     `json:"members"`
		Score     float64 `json:"score"`
		ImageURL  string  `json:"image_url"`
	} `json:"top"`
}

// TopPeople struct for the /top/people endpoint
type TopPeople struct {
	Top []struct {
		MalID     int       `json:"mal_id"`
		Rank      int       `json:"rank"`
		Title     string    `json:"title"`
		URL       string    `json:"url"`
		NameKanji string    `json:"name_kanji"`
		Favorites int       `json:"favorites"`
		ImageURL  string    `json:"image_url"`
		Birthday  time.Time `json:"birthday"`
	} `json:"top"`
}

// TopCharacters struct for the /top/characters endpoint
type TopCharacters struct {
	Top []struct {
		MalID        int       `json:"mal_id"`
		Rank         int       `json:"rank"`
		Title        string    `json:"title"`
		URL          string    `json:"url"`
		NameKanji    string    `json:"name_kanji"`
		Animeography []MalItem `json:"animeography"`
		Mangaography []MalItem `json:"mangaography"`
		Favorites    int       `json:"favorites"`
		ImageURL     string    `json:"image_url"`
	} `json:"top"`
}

// GetTopAnime returns top anime
func GetTopAnime(subType string, page int) (*TopAnime, error) {
	res := &TopAnime{}

	err := urlToStruct(fmt.Sprintf("/top/anime/%d/%s",
		page, url.QueryEscape(subType)), res)

	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetTopManga returns top manga
func GetTopManga(subType string, page int) (*TopManga, error) {
	res := &TopManga{}

	err := urlToStruct(fmt.Sprintf("/top/manga/%d/%s",
		page, url.QueryEscape(subType)), res)

	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetTopPeople returns top people
func GetTopPeople(subType string, page int) (*TopPeople, error) {
	res := &TopPeople{}

	err := urlToStruct(fmt.Sprintf("/top/people/%d/%s",
		page, url.QueryEscape(subType)), res)

	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetTopCharacters returns top characters
func GetTopCharacters(subType string, page int) (*TopCharacters, error) {
	res := &TopCharacters{}

	err := urlToStruct(fmt.Sprintf("/top/characters/%d/%s",
		page, url.QueryEscape(subType)), res)

	if err != nil {
		return nil, err
	}

	return res, nil
}
