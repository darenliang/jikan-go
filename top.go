package jikan

import (
	"fmt"
	"time"
)

type TopAnimeInfo struct {
	Top []struct {
		MalID     int     `json:"mal_id"`
		Rank      int     `json:"rank"`
		Title     string  `json:"title"`
		URL       string  `json:"url"`
		ImageURL  string  `json:"image_url"`
		Type      string  `json:"type"`
		Episodes  int     `json:"episodes"`
		StartDate string  `json:"start_date"`
		EndDate   string  `json:"end_date"`
		Members   int     `json:"members"`
		Score     float64 `json:"score"`
	} `json:"top"`
}

type TopMangaInfo struct {
	RequestHash        string `json:"request_hash"`
	RequestCached      bool   `json:"request_cached"`
	RequestCacheExpiry int    `json:"request_cache_expiry"`
	Top                []struct {
		MalID     int     `json:"mal_id"`
		Rank      int     `json:"rank"`
		Title     string  `json:"title"`
		URL       string  `json:"url"`
		Type      string  `json:"type"`
		Volumes   int     `json:"volumes"`
		StartDate string  `json:"start_date"`
		EndDate   string  `json:"end_date"`
		Members   int     `json:"members"`
		Score     float64 `json:"score"`
		ImageURL  string  `json:"image_url"`
	} `json:"top"`
}

type TopPeopleInfo struct {
	Top []struct {
		MalID     int       `json:"mal_id"`
		Rank      int       `json:"rank"`
		Title     string    `json:"title"`
		URL       string    `json:"url"`
		NameKanji string    `json:"name_kanji"`
		Favorites int       `json:"favorites"`
		ImageURL  string    `json:"image_url"`
		Birthday  time.Time `json:"birthday"`
	} `json:"top"`
}

type TopCharactersInfo struct {
	Top []struct {
		MalID        int    `json:"mal_id"`
		Rank         int    `json:"rank"`
		Title        string `json:"title"`
		URL          string `json:"url"`
		NameKanji    string `json:"name_kanji"`
		Animeography []struct {
			MalID int    `json:"mal_id"`
			Type  string `json:"type"`
			Name  string `json:"name"`
			URL   string `json:"url"`
		} `json:"animeography"`
		Mangaography []struct {
			MalID int    `json:"mal_id"`
			Type  string `json:"type"`
			Name  string `json:"name"`
			URL   string `json:"url"`
		} `json:"mangaography"`
		Favorites int    `json:"favorites"`
		ImageURL  string `json:"image_url"`
	} `json:"top"`
}

func GetTopAnimeInfo(subtype string, page int) (TopAnimeInfo, error) {
	anime := TopAnimeInfo{}
	err := ParseResults(fmt.Sprintf("%s/top/anime/%d/%s", Endpoint, page, subtype), &anime)
	return anime, err
}

func GetTopMangaInfo(subtype string, page int) (TopMangaInfo, error) {
	manga := TopMangaInfo{}
	err := ParseResults(fmt.Sprintf("%s/top/manga/%d/%s", Endpoint, page, subtype), &manga)
	return manga, err
}

func GetTopPeopleInfo(page int) (TopPeopleInfo, error) {
	people := TopPeopleInfo{}
	err := ParseResults(fmt.Sprintf("%s/top/people/%d", Endpoint, page), &people)
	return people, err
}

func GetTopCharactersInfo(page int) (TopCharactersInfo, error) {
	characters := TopCharactersInfo{}
	err := ParseResults(fmt.Sprintf("%s/top/characters/%d", Endpoint, page), &characters)
	return characters, err
}
