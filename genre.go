package jikan

import (
	"fmt"
	"time"
)

type GenreAnimeInfo struct {
	MalURL struct {
		MalID int    `json:"mal_id"`
		Type  string `json:"type"`
		Name  string `json:"name"`
		URL   string `json:"url"`
	} `json:"mal_url"`
	ItemCount int `json:"item_count"`
	Anime     []struct {
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
type GenreMangaInfo struct {
	MalURL struct {
		MalID int    `json:"mal_id"`
		Type  string `json:"type"`
		Name  string `json:"name"`
		URL   string `json:"url"`
	} `json:"mal_url"`
	ItemCount int `json:"item_count"`
	Manga     []struct {
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

func GetGenreAnimeInfo(id int, page int) (GenreAnimeInfo, error) {
	anime := GenreAnimeInfo{}
	err := ParseResults(fmt.Sprintf("%s/genre/anime/%d/%d", Endpoint, id, page), &anime)
	return anime, err
}

func GetGenreMangaInfo(id int, page int) (GenreMangaInfo, error) {
	manga := GenreMangaInfo{}
	err := ParseResults(fmt.Sprintf("%s/genre/manga/%d/%d", Endpoint, id, page), &manga)
	return manga, err
}
