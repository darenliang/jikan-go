package jikan

import "time"

// RecentAnimeReviews struct
type RecentAnimeReviews struct {
	Data []struct {
		User            UserItem    `json:"user"`
		Anime           EntryTitle3 `json:"anime"`
		MalId           int         `json:"mal_id"`
		Url             string      `json:"url"`
		Type            string      `json:"type"`
		Votes           int         `json:"votes"`
		Date            time.Time   `json:"date"`
		Review          string      `json:"review"`
		EpisodesWatched int         `json:"episodes_watched"`
		Scores          ScoresAnime `json:"scores"`
	} `json:"data"`
	Pagination Pagination `json:"pagination"`
}

// GetRecentAnimeReviews returns recent anime reviews
func GetRecentAnimeReviews() (*RecentAnimeReviews, error) {
	res := &RecentAnimeReviews{}
	err := urlToStruct("/reviews/anime", res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// RecentMangaReviews struct
type RecentMangaReviews struct {
	Data []struct {
		User         UserItem    `json:"user"`
		Manga        EntryTitle3 `json:"manga"`
		MalId        int         `json:"mal_id"`
		Url          string      `json:"url"`
		Type         string      `json:"type"`
		Votes        int         `json:"votes"`
		Date         time.Time   `json:"date"`
		ChaptersRead int         `json:"chapters_read"`
		Review       string      `json:"review"`
		Scores       ScoresManga `json:"scores"`
	} `json:"data"`
	Pagination Pagination `json:"pagination"`
}

// GetRecentMangaReviews returns recent manga reviews
func GetRecentMangaReviews() (*RecentMangaReviews, error) {
	res := &RecentMangaReviews{}
	err := urlToStruct("/reviews/manga", res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
