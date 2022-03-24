package jikan

import "fmt"

// Genres struct
type Genres struct {
	Data []struct {
		MalId int    `json:"mal_id"`
		Name  string `json:"name"`
		Url   string `json:"url"`
		Count int    `json:"count"`
	} `json:"data"`
}

type GenreFilter string

const (
	GenresFilterGenres         GenreFilter = "genres"
	GenresFilterExplicitGenres GenreFilter = "explicit_genres"
	GenresFilterThemes         GenreFilter = "themes"
	GenresFilterDemographics   GenreFilter = "demographics"
)

// GetAnimeGenres returns anime genres
func GetAnimeGenres(page, limit int, filter GenreFilter) (*Genres, error) {
	res := &Genres{}
	req := fmt.Sprintf("/genres/anime?page=%d&limit=%d", page, limit)
	if filter != "" {
		req += fmt.Sprintf("&filter=%s", filter)
	}
	err := urlToStruct(req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetMangaGenres returns manga genres
func GetMangaGenres(page, limit int, filter GenreFilter) (*Genres, error) {
	res := &Genres{}
	req := fmt.Sprintf("/genres/manga?page=%d&limit=%d", page, limit)
	if filter != "" {
		req += fmt.Sprintf("&filter=%s", filter)
	}
	err := urlToStruct(req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
