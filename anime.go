package jikan

import (
	"fmt"
	"time"
)

type AnimeInfo struct {
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
	Aired         struct {
		From time.Time `json:"from"`
		To   time.Time `json:"to"`
		Prop struct {
			From struct {
				Day   int `json:"day"`
				Month int `json:"month"`
				Year  int `json:"year"`
			} `json:"from"`
			To struct {
				Day   int `json:"day"`
				Month int `json:"month"`
				Year  int `json:"year"`
			} `json:"to"`
		} `json:"prop"`
		String string `json:"string"`
	} `json:"aired"`
	Duration   string  `json:"duration"`
	Rating     string  `json:"rating"`
	Score      float64 `json:"score"`
	ScoredBy   int     `json:"scored_by"`
	Rank       int     `json:"rank"`
	Popularity int     `json:"popularity"`
	Members    int     `json:"members"`
	Favorites  int     `json:"favorites"`
	Synopsis   string  `json:"synopsis"`
	Background string  `json:"background"`
	Premiered  string  `json:"premiered"`
	Broadcast  string  `json:"broadcast"`
	Related    struct {
		Adaptation []struct {
			MalID int    `json:"mal_id"`
			Type  string `json:"type"`
			Name  string `json:"name"`
			URL   string `json:"url"`
		} `json:"Adaptation"`
		SideStory []struct {
			MalID int    `json:"mal_id"`
			Type  string `json:"type"`
			Name  string `json:"name"`
			URL   string `json:"url"`
		} `json:"Side story"`
		SpinOff []struct {
			MalID int    `json:"mal_id"`
			Type  string `json:"type"`
			Name  string `json:"name"`
			URL   string `json:"url"`
		} `json:"Spin-off"`
		AlternativeVersion []struct {
			MalID int    `json:"mal_id"`
			Type  string `json:"type"`
			Name  string `json:"name"`
			URL   string `json:"url"`
		} `json:"Alternative version"`
		AlternativeSetting []struct {
			MalID int    `json:"mal_id"`
			Type  string `json:"type"`
			Name  string `json:"name"`
			URL   string `json:"url"`
		} `json:"Alternative setting"`
		Character []struct {
			MalID int    `json:"mal_id"`
			Type  string `json:"type"`
			Name  string `json:"name"`
			URL   string `json:"url"`
		} `json:"Character"`
		Prequel []struct {
			MalID int    `json:"mal_id"`
			Type  string `json:"type"`
			Name  string `json:"name"`
			URL   string `json:"url"`
		} `json:"Prequel"`
		Sequel []struct {
			MalID int    `json:"mal_id"`
			Type  string `json:"type"`
			Name  string `json:"name"`
			URL   string `json:"url"`
		} `json:"Sequel"`
		Summary []struct {
			MalID int    `json:"mal_id"`
			Type  string `json:"type"`
			Name  string `json:"name"`
			URL   string `json:"url"`
		} `json:"Summary"`
		Other []struct {
			MalID int    `json:"mal_id"`
			Type  string `json:"type"`
			Name  string `json:"name"`
			URL   string `json:"url"`
		} `json:"Other"`
	} `json:"related"`
	Producers []struct {
		MalID int    `json:"mal_id"`
		Type  string `json:"type"`
		Name  string `json:"name"`
		URL   string `json:"url"`
	} `json:"producers"`
	Licensors []struct {
		MalID int    `json:"mal_id"`
		Type  string `json:"type"`
		Name  string `json:"name"`
		URL   string `json:"url"`
	} `json:"licensors"`
	Studios []struct {
		MalID int    `json:"mal_id"`
		Type  string `json:"type"`
		Name  string `json:"name"`
		URL   string `json:"url"`
	} `json:"studios"`
	Genres []struct {
		MalID int    `json:"mal_id"`
		Type  string `json:"type"`
		Name  string `json:"name"`
		URL   string `json:"url"`
	} `json:"genres"`
	OpeningThemes []string `json:"opening_themes"`
	EndingThemes  []string `json:"ending_themes"`
}

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

type AnimeStats struct {
	Watching    int `json:"watching"`
	Completed   int `json:"completed"`
	OnHold      int `json:"on_hold"`
	Dropped     int `json:"dropped"`
	PlanToWatch int `json:"plan_to_watch"`
	Total       int `json:"total"`
	Scores      struct {
		Num1 struct {
			Votes      int     `json:"votes"`
			Percentage float64 `json:"percentage"`
		} `json:"1"`
		Num2 struct {
			Votes      int     `json:"votes"`
			Percentage float64 `json:"percentage"`
		} `json:"2"`
		Num3 struct {
			Votes      int     `json:"votes"`
			Percentage float64 `json:"percentage"`
		} `json:"3"`
		Num4 struct {
			Votes      int     `json:"votes"`
			Percentage float64 `json:"percentage"`
		} `json:"4"`
		Num5 struct {
			Votes      int     `json:"votes"`
			Percentage float64 `json:"percentage"`
		} `json:"5"`
		Num6 struct {
			Votes      int     `json:"votes"`
			Percentage float64 `json:"percentage"`
		} `json:"6"`
		Num7 struct {
			Votes      int     `json:"votes"`
			Percentage float64 `json:"percentage"`
		} `json:"7"`
		Num8 struct {
			Votes      int `json:"votes"`
			Percentage int `json:"percentage"`
		} `json:"8"`
		Num9 struct {
			Votes      int     `json:"votes"`
			Percentage float64 `json:"percentage"`
		} `json:"9"`
		Num10 struct {
			Votes      int     `json:"votes"`
			Percentage float64 `json:"percentage"`
		} `json:"10"`
	} `json:"scores"`
}

type AnimeReviews struct {
	Reviews []struct {
		MalID        int       `json:"mal_id"`
		URL          string    `json:"url"`
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

type AnimeUserupdates struct {
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

func GetAnimeInfo(id int) (AnimeInfo, error) {
	anime := AnimeInfo{}
	err := ParseResults(fmt.Sprintf("%s/anime/%d", Endpoint, id), &anime)
	return anime, err
}

func GetAnimeCharactersStaff(id int) (AnimeCharactersStaff, error) {
	anime := AnimeCharactersStaff{}
	err := ParseResults(fmt.Sprintf("%s/anime/%d/characters_staff", Endpoint, id), &anime)
	return anime, err
}

func GetAnimeEpisodes(id int, page int) (AnimeEpisodes, error) {
	anime := AnimeEpisodes{}
	err := ParseResults(fmt.Sprintf("%s/anime/%d/episodes/%d", Endpoint, id, page), &anime)
	return anime, err
}

func GetAnimeNews(id int) (News, error) {
	return getNews(fmt.Sprintf("%s/anime/%d/news", Endpoint, id))
}

func GetAnimePictures(id int) (Pictures, error) {
	return getPictures(fmt.Sprintf("%s/anime/%d/pictures", Endpoint, id))
}

func GetAnimeVideos(id int) (AnimeVideos, error) {
	anime := AnimeVideos{}
	err := ParseResults(fmt.Sprintf("%s/anime/%d/videos", Endpoint, id), &anime)
	return anime, err
}

func GetAnimeStats(id int) (AnimeStats, error) {
	anime := AnimeStats{}
	err := ParseResults(fmt.Sprintf("%s/anime/%d/stats", Endpoint, id), &anime)
	return anime, err
}

func GetAnimeForum(id int) (Forum, error) {
	return getForum(fmt.Sprintf("%s/anime/%d/forum", Endpoint, id))
}

func GetAnimeMoreinfo(id int) (Moreinfo, error) {
	return getMoreinfo(fmt.Sprintf("%s/anime/%d/forum", Endpoint, id))
}

func GetAnimeReviews(id int, page int) (AnimeReviews, error) {
	anime := AnimeReviews{}
	err := ParseResults(fmt.Sprintf("%s/anime/%d/reviews/%d", Endpoint, id, page), &anime)
	return anime, err
}

func GetAnimeRecommendations(id int) (Recommendations, error) {
	return getRecommendations(fmt.Sprintf("%s/anime/%d/recommendations", Endpoint, id))
}

func GetAnimeUserupdates(id int, page int) (AnimeUserupdates, error) {
	anime := AnimeUserupdates{}
	err := ParseResults(fmt.Sprintf("%s/anime/%d/userupdates/%d", Endpoint, id, page), &anime)
	return anime, err
}
