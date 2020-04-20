package jikan

import (
	"fmt"
	"time"
)

type Season struct {
	SeasonName string `json:"season_name"`
	SeasonYear int    `json:"season_year"`
	Anime      []struct {
		MalID       int       `json:"mal_id"`
		URL         string    `json:"url"`
		Title       string    `json:"title"`
		ImageURL    string    `json:"image_url"`
		Synopsis    string    `json:"synopsis"`
		Type        string    `json:"type"`
		AiringStart time.Time `json:"airing_start"`
		Episodes    int       `json:"episodes"`
		Members     int       `json:"members"`
		Genres      []malItem `json:"genres"`
		Source      string    `json:"source"`
		Producers   []malItem `json:"producers"`
		Score       float64   `json:"score"`
		Licensors   []string  `json:"licensors"`
		R18         bool      `json:"r18"`
		Kids        bool      `json:"kids"`
		Continuing  bool      `json:"continuing"`
	} `json:"anime"`
}

// SeasonArchive struct
type SeasonArchive struct {
	Archive []struct {
		Year    int      `json:"year"`
		Seasons []string `json:"seasons"`
	} `json:"archive"`
}

// SeasonLater struct
type SeasonLater = Season

// GetSeason returns season
func GetSeason(year int, season string) (*Season, error) {
	res := &Season{}

	err := urlToStruct(fmt.Sprintf("/season/%d/%s", year, season), res)

	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetSeasonArchive returns season archive
func GetSeasonArchive() (*SeasonArchive, error) {
	res := &SeasonArchive{}

	err := urlToStruct("/season/archive", res)

	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetSeasonLater returns season later
func GetSeasonLater() (*SeasonLater, error) {
	res := &SeasonLater{}

	err := urlToStruct("/season/later", res)

	if err != nil {
		return nil, err
	}

	return res, nil
}
