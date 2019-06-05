package jikan

import (
	"fmt"
	"time"
)

type ScheduleInfo struct {
	Monday []struct {
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
		R18  bool `json:"r18"`
		Kids bool `json:"kids"`
	} `json:"monday"`
}

func GetScheduleInfo(day string) (ScheduleInfo, error) {
	schedule := ScheduleInfo{}
	err := ParseResults(fmt.Sprintf("%s/schedule/%s", Endpoint, day), &schedule)
	return schedule, err
}
