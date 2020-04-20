package jikan

import (
	"fmt"
)

// Schedule struct
type Schedule struct {
	Monday    []malAnimeShort `json:"monday"`
	Tuesday   []malAnimeShort `json:"tuesday"`
	Wednesday []malAnimeShort `json:"wednesday"`
	Thursday  []malAnimeShort `json:"thursday"`
	Friday    []malAnimeShort `json:"friday"`
	Saturday  []malAnimeShort `json:"saturday"`
	Sunday    []malAnimeShort `json:"sunday"`
	Other     []malAnimeShort `json:"other"`
	Unknown   []malAnimeShort `json:"unknown"`
}

// GetSchedule returns schedule
func GetSchedule(day string) (*Schedule, error) {
	res := &Schedule{}

	err := urlToStruct(fmt.Sprintf("/schedule/%s", day), res)

	if err != nil {
		return nil, err
	}

	return res, nil
}
