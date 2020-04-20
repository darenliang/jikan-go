package jikan

import (
	"fmt"
	"net/url"
)

// Genre struct for the /genre endpoint
type Genre struct {
	MalURL    MalItem        `json:"mal_url"`
	ItemCount int            `json:"item_count"`
	Anime     []MalAnimeDesc `json:"anime"`
	Manga     []MalMangaDesc `json:"manga"`
}

// GetGenre returns genre
func GetGenre(mediaType string, id, page int) (*Genre, error) {
	res := &Genre{}

	err := urlToStruct(fmt.Sprintf("/genre/%s/%d/%d", url.QueryEscape(mediaType), id, page), res)

	if err != nil {
		return nil, err
	}

	return res, nil
}
