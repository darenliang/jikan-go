package jikan

import (
	"fmt"
	"time"
)

// Manga struct
type Manga struct {
	MalID         int      `json:"mal_id"`
	URL           string   `json:"url"`
	Title         string   `json:"title"`
	TitleEnglish  string   `json:"title_english"`
	TitleSynonyms []string `json:"title_synonyms"`
	TitleJapanese string   `json:"title_japanese"`
	Status        string   `json:"status"`
	ImageURL      string   `json:"image_url"`
	Type          string   `json:"type"`
	Volumes       int      `json:"volumes"`
	Chapters      int      `json:"chapters"`
	Publishing    bool     `json:"publishing"`
	Published     malDates `json:"published"`
	Rank          int      `json:"rank"`
	Score         float64  `json:"score"`
	ScoredBy      int      `json:"scored_by"`
	Popularity    int      `json:"popularity"`
	Members       int      `json:"members"`
	Favorites     int      `json:"favorites"`
	Synopsis      string   `json:"synopsis"`
	Background    string   `json:"background"`
	Related       struct {
		SideStory  []malItem `json:"Side story"`
		Adaptation []malItem `json:"Adaptation"`
	} `json:"related"`
	Genres         []malItem `json:"genres"`
	Authors        []malItem `json:"authors"`
	Serializations []malItem `json:"serializations"`
}

// Manga characters struct
type MangaCharacters struct {
	Characters []malRoleStaff `json:"characters"`
}

// Manga news struct
type MangaNews = AnimeNews

// Manga pictures struct
type MangaPictures = AnimePictures

// Manga stats struct
type MangaStats struct {
	Reading    int       `json:"reading"`
	Completed  int       `json:"completed"`
	OnHold     int       `json:"on_hold"`
	Dropped    int       `json:"dropped"`
	PlanToRead int       `json:"plan_to_read"`
	Total      int       `json:"total"`
	Scores     malScores `json:"scores"`
}

// Manga forum struct
type MangaForum = AnimeForum

// Manga more info struct
type MangaMoreInfo = AnimeMoreInfo

// Manga reviews struct
type MangaReviews struct {
	Reviews []struct {
		MalID        int         `json:"mal_id"`
		URL          string      `json:"url"`
		Type         interface{} `json:"type"`
		HelpfulCount int         `json:"helpful_count"`
		Date         time.Time   `json:"date"`
		Reviewer     struct {
			URL          string `json:"url"`
			ImageURL     string `json:"image_url"`
			Username     string `json:"username"`
			ChaptersRead int    `json:"chapters_read"`
			Scores       struct {
				Overall   int `json:"overall"`
				Story     int `json:"story"`
				Art       int `json:"art"`
				Character int `json:"character"`
				Enjoyment int `json:"enjoyment"`
			} `json:"scores"`
		} `json:"reviewer"`
		Content string `json:"content"`
	} `json:"reviews"`
}

// Manga recommendations struct
type MangaRecommendations = AnimeRecommendations

// Manga user updates struct
type MangaUserUpdates struct {
	Users []struct {
		Username      string      `json:"username"`
		URL           string      `json:"url"`
		ImageURL      string      `json:"image_url"`
		Score         interface{} `json:"score"`
		Status        string      `json:"status"`
		VolumesRead   interface{} `json:"volumes_read"`
		VolumesTotal  int         `json:"volumes_total"`
		ChaptersRead  int         `json:"chapters_read"`
		ChaptersTotal int         `json:"chapters_total"`
		Date          time.Time   `json:"date"`
	} `json:"users"`
}

// GetManga returns manga
func GetManga(id int) (*Manga, error) {
	res := &Manga{}

	err := urlToStruct(fmt.Sprintf("/manga/%d", id), res)

	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetMangaCharacters returns manga characters
func GetMangaCharacters(id int) (*MangaCharacters, error) {
	res := &MangaCharacters{}

	err := urlToStruct(fmt.Sprintf("/manga/%d/characters", id), res)

	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetMangaNews returns manga news
func GetMangaNews(id int) (*MangaNews, error) {
	res := &MangaNews{}

	err := urlToStruct(fmt.Sprintf("/manga/%d/news", id), res)

	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetMangaPictures returns manga pictures
func GetMangaPictures(id int) (*MangaPictures, error) {
	res := &MangaPictures{}

	err := urlToStruct(fmt.Sprintf("/manga/%d/pictures", id), res)

	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetMangaStats returns manga stats
func GetMangaStats(id int) (*MangaStats, error) {
	res := &MangaStats{}

	err := urlToStruct(fmt.Sprintf("/manga/%d/stats", id), res)

	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetMangaForum returns manga forum
func GetMangaForum(id int) (*MangaForum, error) {
	res := &MangaForum{}

	err := urlToStruct(fmt.Sprintf("/manga/%d/forum", id), res)

	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetMangaMoreInfo returns manga more info
func GetMangaMoreInfo(id int) (*MangaMoreInfo, error) {
	res := &MangaMoreInfo{}

	err := urlToStruct(fmt.Sprintf("/manga/%d/moreinfo", id), res)

	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetMangaReviews returns manga reviews
func GetMangaReviews(id, page int) (*MangaReviews, error) {
	res := &MangaReviews{}

	err := urlToStruct(fmt.Sprintf("/manga/%d/reviews/%d", id, page), res)

	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetMangaRecommendations returns manga recommendations
func GetMangaRecommendations(id int) (*MangaRecommendations, error) {
	res := &MangaRecommendations{}

	err := urlToStruct(fmt.Sprintf("/manga/%d/recommendations", id), res)

	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetMangaUserUpdates returns manga user updates
func GetMangaUserUpdates(id, page int) (*MangaUserUpdates, error) {
	res := &MangaUserUpdates{}

	err := urlToStruct(fmt.Sprintf("/manga/%d/userupdates/%d", id, page), res)

	if err != nil {
		return nil, err
	}

	return res, nil
}
