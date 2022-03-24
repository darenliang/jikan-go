package jikan

import (
	"fmt"
	"net/url"
)

// Season struct
type Season struct {
	Data       []AnimeBase `json:"data"`
	Pagination Pagination  `json:"pagination"`
}

// GetSeason returns season
func GetSeason(year int, season string) (*Season, error) {
	res := &Season{}
	err := urlToStruct(fmt.Sprintf("/seasons/%d/%s", year, url.QueryEscape(season)), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetSeasonNow returns season now
func GetSeasonNow() (*Season, error) {
	res := &Season{}
	err := urlToStruct("/seasons/now", res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type SeasonsList struct {
	Data []struct {
		Year    int      `json:"year"`
		Seasons []string `json:"seasons"`
	} `json:"data"`
}

// GetSeasonsList returns seasons list
func GetSeasonsList() (*SeasonsList, error) {
	res := &SeasonsList{}
	err := urlToStruct("/seasons", res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetSeasonUpcoming returns season upcoming
func GetSeasonUpcoming() (*Season, error) {
	res := &Season{}
	err := urlToStruct("/seasons/upcoming", res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
