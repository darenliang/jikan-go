package jikan

import (
	"fmt"
	"time"
)

// Anime struct
type Anime struct {
	MalID         int      `json:"mal_id"`
	URL           string   `json:"url"`
	ImageURL      string   `json:"image_url"`
	TrailerURL    string   `json:"trailer_url"`
	Title         string   `json:"title"`
	TitleEnglish  string   `json:"title_english"`
	TitleJapanese string   `json:"title_japanese"`
	TitleSynonyms []string `json:"title_synonyms"`
	Type          string   `json:"type"`
	Source        string   `json:"source"`
	Episodes      int      `json:"episodes"`
	Status        string   `json:"status"`
	Airing        bool     `json:"airing"`
	Aired         malDates `json:"aired"`
	Duration      string   `json:"duration"`
	Rating        string   `json:"rating"`
	Score         float64  `json:"score"`
	ScoredBy      int      `json:"scored_by"`
	Rank          int      `json:"rank"`
	Popularity    int      `json:"popularity"`
	Members       int      `json:"members"`
	Favorites     int      `json:"favorites"`
	Synopsis      string   `json:"synopsis"`
	Background    string   `json:"background"`
	Premiered     string   `json:"premiered"`
	Broadcast     string   `json:"broadcast"`
	Related       struct {
		Adaptation []malItem `json:"Adaptation"`
		SideStory  []malItem `json:"Side story"`
		Summary    []malItem `json:"Summary"`
	} `json:"related"`
	Producers     []malItem `json:"producers"`
	Licensors     []malItem `json:"licensors"`
	Studios       []malItem `json:"studios"`
	Genres        []malItem `json:"genres"`
	OpeningThemes []string  `json:"opening_themes"`
	EndingThemes  []string  `json:"ending_themes"`
}

// Anime character staff struct
type AnimeCharactersStaff struct {
	Characters []struct {
		MalID       int    `json:"mal_id"`
		URL         string `json:"url"`
		ImageURL    string `json:"image_url"`
		Name        string `json:"name"`
		Role        string `json:"role"`
		VoiceActors []struct {
			MalID    int    `json:"mal_id"`
			Name     string `json:"name"`
			URL      string `json:"url"`
			ImageURL string `json:"image_url"`
			Language string `json:"language"`
		} `json:"voice_actors"`
	} `json:"characters"`
	Staff []struct {
		MalID     int      `json:"mal_id"`
		URL       string   `json:"url"`
		Name      string   `json:"name"`
		ImageURL  string   `json:"image_url"`
		Positions []string `json:"positions"`
	} `json:"staff"`
}

// Anime episodes struct
type AnimeEpisodes struct {
	EpisodesLastPage int `json:"episodes_last_page"`
	Episodes         []struct {
		EpisodeID     int       `json:"episode_id"`
		Title         string    `json:"title"`
		TitleJapanese string    `json:"title_japanese"`
		TitleRomanji  string    `json:"title_romanji"`
		Aired         time.Time `json:"aired"`
		Filler        bool      `json:"filler"`
		Recap         bool      `json:"recap"`
		VideoURL      string    `json:"video_url"`
		ForumURL      string    `json:"forum_url"`
	} `json:"episodes"`
}

// Anime news struct
type AnimeNews struct {
	Articles []struct {
		URL        string    `json:"url"`
		Title      string    `json:"title"`
		Date       time.Time `json:"date"`
		AuthorName string    `json:"author_name"`
		AuthorURL  string    `json:"author_url"`
		ForumURL   string    `json:"forum_url"`
		ImageURL   string    `json:"image_url"`
		Comments   int       `json:"comments"`
		Intro      string    `json:"intro"`
	} `json:"articles"`
}

// Anime pictures struct
type AnimePictures struct {
	Pictures []struct {
		Large string `json:"large"`
		Small string `json:"small"`
	} `json:"pictures"`
}

// Anime videos struct
type AnimeVideos struct {
	Promo []struct {
		Title    string `json:"title"`
		ImageURL string `json:"image_url"`
		VideoURL string `json:"video_url"`
	} `json:"promo"`
	Episodes []struct {
		Title    string `json:"title"`
		Episode  string `json:"episode"`
		URL      string `json:"url"`
		ImageURL string `json:"image_url"`
	} `json:"episodes"`
}

// Anime stats struct
type AnimeStats struct {
	Watching    int       `json:"watching"`
	Completed   int       `json:"completed"`
	OnHold      int       `json:"on_hold"`
	Dropped     int       `json:"dropped"`
	PlanToWatch int       `json:"plan_to_watch"`
	Total       int       `json:"total"`
	Scores      malScores `json:"scores"`
}

// Anime forum struct
type AnimeForum struct {
	Topics []struct {
		TopicID    int       `json:"topic_id"`
		URL        string    `json:"url"`
		Title      string    `json:"title"`
		DatePosted time.Time `json:"date_posted"`
		AuthorName string    `json:"author_name"`
		AuthorURL  string    `json:"author_url"`
		Replies    int       `json:"replies"`
		LastPost   struct {
			URL        string    `json:"url"`
			AuthorName string    `json:"author_name"`
			AuthorURL  string    `json:"author_url"`
			DatePosted time.Time `json:"date_posted"`
		} `json:"last_post"`
	} `json:"topics"`
}

// Anime more info struct
type AnimeMoreInfo struct {
	Moreinfo string `json:"moreinfo"`
}

// Anime reviews struct
type AnimeReviews struct {
	Reviews []struct {
		MalID        int       `json:"mal_id"`
		URL          string    `json:"url"`
		Type         string    `json:"type"`
		HelpfulCount int       `json:"helpful_count"`
		Date         time.Time `json:"date"`
		Reviewer     struct {
			URL          string `json:"url"`
			ImageURL     string `json:"image_url"`
			Username     string `json:"username"`
			EpisodesSeen int    `json:"episodes_seen"`
			Scores       struct {
				Overall   int `json:"overall"`
				Story     int `json:"story"`
				Animation int `json:"animation"`
				Sound     int `json:"sound"`
				Character int `json:"character"`
				Enjoyment int `json:"enjoyment"`
			} `json:"scores"`
		} `json:"reviewer"`
		Content string `json:"content"`
	} `json:"reviews"`
}

// Anime recommendations struct
type AnimeRecommendations struct {
	Recommendations []struct {
		MalID               int    `json:"mal_id"`
		URL                 string `json:"url"`
		ImageURL            string `json:"image_url"`
		RecommendationURL   string `json:"recommendation_url"`
		Title               string `json:"title"`
		RecommendationCount int    `json:"recommendation_count"`
	} `json:"recommendations"`
}

// Anime user updates struct
type AnimeUserUpdates struct {
	Users []struct {
		Username      string    `json:"username"`
		URL           string    `json:"url"`
		ImageURL      string    `json:"image_url"`
		Score         int       `json:"score"`
		Status        string    `json:"status"`
		EpisodesSeen  int       `json:"episodes_seen"`
		EpisodesTotal int       `json:"episodes_total"`
		Date          time.Time `json:"date"`
	} `json:"users"`
}

// GetAnime returns anime
func GetAnime(id int) (*Anime, error) {
	res := &Anime{}

	err := urlToStruct(fmt.Sprintf("/anime/%d", id), res)

	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetAnimeCharactersStaff returns anime character staff
func GetAnimeCharactersStaff(id int) (*AnimeCharactersStaff, error) {
	res := &AnimeCharactersStaff{}

	err := urlToStruct(fmt.Sprintf("/anime/%d/characters_staff", id), res)

	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetAnimeEpisodes returns anime episodes
func GetAnimeEpisodes(id, page int) (*AnimeEpisodes, error) {
	res := &AnimeEpisodes{}

	err := urlToStruct(fmt.Sprintf("/anime/%d/episodes/%d", id, page), res)

	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetAnimeNews returns anime news
func GetAnimeNews(id int) (*AnimeNews, error) {
	res := &AnimeNews{}

	err := urlToStruct(fmt.Sprintf("/anime/%d/news", id), res)

	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetAnimePictures returns anime pictures
func GetAnimePictures(id int) (*AnimePictures, error) {
	res := &AnimePictures{}

	err := urlToStruct(fmt.Sprintf("/anime/%d/pictures", id), res)

	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetAnimeVideos returns anime videos
func GetAnimeVideos(id int) (*AnimeVideos, error) {
	res := &AnimeVideos{}

	err := urlToStruct(fmt.Sprintf("/anime/%d/videos", id), res)

	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetAnimeStats returns anime stats
func GetAnimeStats(id int) (*AnimeStats, error) {
	res := &AnimeStats{}

	err := urlToStruct(fmt.Sprintf("/anime/%d/stats", id), res)

	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetAnimeForum returns anime forum
func GetAnimeForum(id int) (*AnimeForum, error) {
	res := &AnimeForum{}

	err := urlToStruct(fmt.Sprintf("/anime/%d/forum", id), res)

	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetAnimeMoreInfo returns anime more info
func GetAnimeMoreInfo(id int) (*AnimeMoreInfo, error) {
	res := &AnimeMoreInfo{}

	err := urlToStruct(fmt.Sprintf("/anime/%d/moreinfo", id), res)

	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetAnimeReviews returns anime reviews
func GetAnimeReviews(id, page int) (*AnimeReviews, error) {
	res := &AnimeReviews{}

	err := urlToStruct(fmt.Sprintf("/anime/%d/reviews/%d", id, page), res)

	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetAnimeRecommendations returns anime recommendations
func GetAnimeRecommendations(id int) (*AnimeRecommendations, error) {
	res := &AnimeRecommendations{}

	err := urlToStruct(fmt.Sprintf("/anime/%d/recommendations", id), res)

	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetAnimeUserUpdates returns anime user updates
func GetAnimeUserUpdates(id, page int) (*AnimeUserUpdates, error) {
	res := &AnimeUserUpdates{}

	err := urlToStruct(fmt.Sprintf("/anime/%d/userupdates/%d", id, page), res)

	if err != nil {
		return nil, err
	}

	return res, nil
}
