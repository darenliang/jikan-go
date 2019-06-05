package jikan

import (
	"fmt"
	"time"
)

type SeasonInfo struct {
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
		Genres      []struct {
			MalID int    `json:"mal_id"`
			Type  string `json:"type"`
			Name  string `json:"name"`
			URL   string `json:"url"`
		} `json:"genres"`
		Source    string `json:"source"`
		Producers []struct {
			MalID int    `json:"mal_id"`
			Type  string `json:"type"`
			Name  string `json:"name"`
			URL   string `json:"url"`
		} `json:"producers"`
		Score     float64 `json:"score"`
		Licensors []struct {
			MalID int    `json:"mal_id"`
			Type  string `json:"type"`
			Name  string `json:"name"`
			URL   string `json:"url"`
		} `json:"licensors"`
		R18        bool `json:"r18"`
		Kids       bool `json:"kids"`
		Continuing bool `json:"continuing"`
	} `json:"anime"`
}

type SeasonArchive struct {
	Archive []struct {
		Year    int      `json:"year"`
		Seasons []string `json:"seasons"`
	} `json:"archive"`
}
type SeasonLater struct {
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
		Genres      []struct {
			MalID int    `json:"mal_id"`
			Type  string `json:"type"`
			Name  string `json:"name"`
			URL   string `json:"url"`
		} `json:"genres"`
		Source    string `json:"source"`
		Producers []struct {
			MalID int    `json:"mal_id"`
			Type  string `json:"type"`
			Name  string `json:"name"`
			URL   string `json:"url"`
		} `json:"producers"`
		Score     float64 `json:"score"`
		Licensors []struct {
			MalID int    `json:"mal_id"`
			Type  string `json:"type"`
			Name  string `json:"name"`
			URL   string `json:"url"`
		} `json:"licensors"`
		R18        bool `json:"r18"`
		Kids       bool `json:"kids"`
		Continuing bool `json:"continuing"`
	} `json:"anime"`
}

func GetSeasonInfo(season string, year int) (SeasonInfo, error) {
	result := SeasonInfo{}
	err := ParseResults(fmt.Sprintf("%s/season/%d/%s", Endpoint, year, season), &result)
	return result, err
}

func GetSeasonArchive() (SeasonArchive, error) {
	result := SeasonArchive{}
	err := ParseResults(fmt.Sprintf("%s/season/archive", Endpoint), &result)
	return result, err
}

func GetSeasonLater() (SeasonLater, error) {
	result := SeasonLater{}
	err := ParseResults(fmt.Sprintf("%s/season/later", Endpoint), &result)
	return result, err
}
