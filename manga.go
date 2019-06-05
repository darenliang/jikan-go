package jikan

import (
	"fmt"
	"time"
)

type MangaInfo struct {
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
	Published     struct {
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
	} `json:"published"`
	Rank       int     `json:"rank"`
	Score      float64 `json:"score"`
	ScoredBy   int     `json:"scored_by"`
	Popularity int     `json:"popularity"`
	Members    int     `json:"members"`
	Favorites  int     `json:"favorites"`
	Synopsis   string  `json:"synopsis"`
	Background string  `json:"background"`
	Related    struct {
		SideStory []struct {
			MalID int    `json:"mal_id"`
			Type  string `json:"type"`
			Name  string `json:"name"`
			URL   string `json:"url"`
		} `json:"Side story"`
		Adaptation []struct {
			MalID int    `json:"mal_id"`
			Type  string `json:"type"`
			Name  string `json:"name"`
			URL   string `json:"url"`
		} `json:"Adaptation"`
	} `json:"related"`
	Genres []struct {
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
	Serializations []struct {
		MalID int    `json:"mal_id"`
		Type  string `json:"type"`
		Name  string `json:"name"`
		URL   string `json:"url"`
	} `json:"serializations"`
}

type MangaCharacters struct {
	Characters []struct {
		MalID    int    `json:"mal_id"`
		URL      string `json:"url"`
		ImageURL string `json:"image_url"`
		Name     string `json:"name"`
		Role     string `json:"role"`
	} `json:"characters"`
}

type MangaStats struct {
	Reading    int `json:"reading"`
	Completed  int `json:"completed"`
	OnHold     int `json:"on_hold"`
	Dropped    int `json:"dropped"`
	PlanToRead int `json:"plan_to_read"`
	Total      int `json:"total"`
	Scores     struct {
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
			Votes      int     `json:"votes"`
			Percentage float64 `json:"percentage"`
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

type MangaReviews struct {
	Reviews []struct {
		MalID        int       `json:"mal_id"`
		URL          string    `json:"url"`
		HelpfulCount int       `json:"helpful_count"`
		Date         time.Time `json:"date"`
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

type MangaUserupdates struct {
	Users []struct {
		Username      string    `json:"username"`
		URL           string    `json:"url"`
		ImageURL      string    `json:"image_url"`
		Score         int       `json:"score"`
		Status        string    `json:"status"`
		VolumesRead   int       `json:"volumes_read"`
		VolumesTotal  int       `json:"volumes_total"`
		ChaptersRead  int       `json:"chapters_read"`
		ChaptersTotal int       `json:"chapters_total"`
		Date          time.Time `json:"date"`
	} `json:"users"`
}

func GetMangaInfo(id int) (MangaInfo, error) {
	manga := MangaInfo{}
	err := ParseResults(fmt.Sprintf("%s/manga/%d", Endpoint, id), &manga)
	return manga, err
}

func GetMangaCharacters(id int) (MangaCharacters, error) {
	manga := MangaCharacters{}
	err := ParseResults(fmt.Sprintf("%s/manga/%d/characters", Endpoint, id), &manga)
	return manga, err
}

func GetMangaNews(id int) (News, error) {
	return getNews(fmt.Sprintf("%s/manga/%d/news", Endpoint, id))
}

func GetMangaPictures(id int) (Pictures, error) {
	return getPictures(fmt.Sprintf("%s/manga/%d/pictures", Endpoint, id))
}

func GetMangaStats(id int) (MangaStats, error) {
	manga := MangaStats{}
	err := ParseResults(fmt.Sprintf("%s/manga/%d/stats", Endpoint, id), &manga)
	return manga, err
}

func GetMangaForum(id int) (Forum, error) {
	return getForum(fmt.Sprintf("%s/manga/%d/forum", Endpoint, id))
}

func GetMangaMoreinfo(id int) (Moreinfo, error) {
	return getMoreinfo(fmt.Sprintf("%s/manga/%d/forum", Endpoint, id))
}

func GetMangaReviews(id int, page int) (MangaReviews, error) {
	manga := MangaReviews{}
	err := ParseResults(fmt.Sprintf("%s/manga/%d/reviews/%d", Endpoint, id, page), &manga)
	return manga, err
}

func GetMangaRecommendations(id int) (Recommendations, error) {
	return getRecommendations(fmt.Sprintf("%s/manga/%d/recommendations", Endpoint, id))
}

func GetMangaUserupdates(id int, page int) (MangaUserupdates, error) {
	manga := MangaUserupdates{}
	err := ParseResults(fmt.Sprintf("%s/manga/%d/userupdates/%d", Endpoint, id, page), &manga)
	return manga, err
}
