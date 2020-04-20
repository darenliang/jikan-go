package jikan

import (
	"fmt"
)

// GenreAnime struct for the /genre/anime endpoint
type GenreAnime struct {
	MalURL    MalItem        `json:"mal_url"`
	ItemCount int            `json:"item_count"`
	Anime     []MalAnimeDesc `json:"anime"`
}

// GenreManga struct for the /genre/manga endpoint
type GenreManga struct {
	MalURL    MalItem        `json:"mal_url"`
	ItemCount int            `json:"item_count"`
	Manga     []MalMangaDesc `json:"manga"`
}

// GetGenreAnime returns genre anime
func GetGenreAnime(id, page int) (*GenreAnime, error) {
	res := &GenreAnime{}

	err := urlToStruct(fmt.Sprintf("/genre/anime/%d/%d", id, page), res)

	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetGenreManga returns genre manga
func GetGenreManga(id, page int) (*GenreManga, error) {
	res := &GenreManga{}

	err := urlToStruct(fmt.Sprintf("/genre/manga/%d/%d", id, page), res)

	if err != nil {
		return nil, err
	}

	return res, nil
}
