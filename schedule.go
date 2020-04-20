package jikan

import (
	"fmt"
)

// Schedule struct
type Schedule struct {
	Monday    []MalAnimeDesc `json:"monday"`
	Tuesday   []MalAnimeDesc `json:"tuesday"`
	Wednesday []MalAnimeDesc `json:"wednesday"`
	Thursday  []MalAnimeDesc `json:"thursday"`
	Friday    []MalAnimeDesc `json:"friday"`
	Saturday  []MalAnimeDesc `json:"saturday"`
	Sunday    []MalAnimeDesc `json:"sunday"`
	Other     []MalAnimeDesc `json:"other"`
	Unknown   []MalAnimeDesc `json:"unknown"`
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
