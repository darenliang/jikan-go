package jikan

import (
	"fmt"
	"time"
)

type ProducerInfo struct {
	Meta struct {
		MalID int    `json:"mal_id"`
		Type  string `json:"type"`
		Name  string `json:"name"`
		URL   string `json:"url"`
	} `json:"meta"`
	Anime []struct {
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
		Score     float64  `json:"score"`
		Licensors []string `json:"licensors"`
		R18       bool     `json:"r18"`
		Kids      bool     `json:"kids"`
	} `json:"anime"`
}

func GetProducerInfo(id int, page int) (ProducerInfo, error) {
	producer := ProducerInfo{}
	err := ParseResults(fmt.Sprintf("%s/producer/%d/%d", Endpoint, id, page), &producer)
	return producer, err
}
