package jikan

import (
	"fmt"
	"time"
)

// Magazine struct
type Magazine struct {
	Meta  MalItem `json:"meta"`
	Manga []struct {
		MalID           int       `json:"mal_id"`
		URL             string    `json:"url"`
		Title           string    `json:"title"`
		ImageURL        string    `json:"image_url"`
		Synopsis        string    `json:"synopsis"`
		Type            string    `json:"type"`
		PublishingStart time.Time `json:"publishing_start"`
		Volumes         int       `json:"volumes"`
		Members         int       `json:"members"`
		Genres          []MalItem `json:"genres"`
		Authors         []MalItem `json:"authors"`
		Score           float64   `json:"score"`
		Serialization   []string  `json:"serialization"`
	} `json:"manga"`
}

// GetMagazine returns magazine
func GetMagazine(id, page int) (*Magazine, error) {
	res := &Magazine{}

	err := urlToStruct(fmt.Sprintf("/magazine/%d/%d", id, page), res)

	if err != nil {
		return nil, err
	}

	return res, nil
}
