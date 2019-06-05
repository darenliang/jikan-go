package jikan

import (
	"fmt"
	"time"
)

type MagazineInfo struct {
	Meta struct {
		MalID int    `json:"mal_id"`
		Type  string `json:"type"`
		Name  string `json:"name"`
		URL   string `json:"url"`
	} `json:"meta"`
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
		Genres          []struct {
			MalID int    `json:"mal_id"`
			Type  string `json:"type"`
			Name  string `json:"name"`
			URL   string `json:"url"`
		} `json:"genres"`
		Authors []struct {
			MalID int    `json:"mal_id"`
			Type  string `json:"type"`
			Name  string `json:"name"`
			URL   string `json:"url"`
		} `json:"authors"`
		Score         float64  `json:"score"`
		Serialization []string `json:"serialization"`
	} `json:"manga"`
}

func GetMagazineInfo(id int, page int) (MagazineInfo, error) {
	magazine := MagazineInfo{}
	err := ParseResults(fmt.Sprintf("%s/magazine/%d/%d", Endpoint, id, page), &magazine)
	return magazine, err
}
