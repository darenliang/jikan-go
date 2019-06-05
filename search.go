package jikan

import (
	"fmt"
	"strings"
	"time"
)

type SearchAnimeInfo struct {
	Results []struct {
		MalID     int       `json:"mal_id"`
		URL       string    `json:"url"`
		ImageURL  string    `json:"image_url"`
		Title     string    `json:"title"`
		Airing    bool      `json:"airing"`
		Synopsis  string    `json:"synopsis"`
		Type      string    `json:"type"`
		Episodes  int       `json:"episodes"`
		Score     float64   `json:"score"`
		StartDate time.Time `json:"start_date"`
		EndDate   time.Time `json:"end_date"`
		Members   int       `json:"members"`
		Rated     string    `json:"rated"`
	} `json:"results"`
	LastPage int `json:"last_page"`
}

type SearchMangaInfo struct {
	Results []struct {
		MalID      int       `json:"mal_id"`
		URL        string    `json:"url"`
		ImageURL   string    `json:"image_url"`
		Title      string    `json:"title"`
		Publishing bool      `json:"publishing"`
		Synopsis   string    `json:"synopsis"`
		Type       string    `json:"type"`
		Chapters   int       `json:"chapters"`
		Volumes    int       `json:"volumes"`
		Score      float64   `json:"score"`
		StartDate  time.Time `json:"start_date"`
		EndDate    time.Time `json:"end_date"`
		Members    int       `json:"members"`
	} `json:"results"`
	LastPage int `json:"last_page"`
}

type SearchPersonInfo struct {
	Results []struct {
		MalID            int      `json:"mal_id"`
		URL              string   `json:"url"`
		ImageURL         string   `json:"image_url"`
		Name             string   `json:"name"`
		AlternativeNames []string `json:"alternative_names"`
	} `json:"results"`
	LastPage int `json:"last_page"`
}

type SearchCharacterInfo struct {
	Results []struct {
		MalID            int      `json:"mal_id"`
		URL              string   `json:"url"`
		ImageURL         string   `json:"image_url"`
		Name             string   `json:"name"`
		AlternativeNames []string `json:"alternative_names"`
		Anime            []struct {
			MalID int    `json:"mal_id"`
			Type  string `json:"type"`
			Name  string `json:"name"`
			URL   string `json:"url"`
		} `json:"anime"`
		Manga []struct {
			MalID int    `json:"mal_id"`
			Type  string `json:"type"`
			Name  string `json:"name"`
			URL   string `json:"url"`
		} `json:"manga"`
	} `json:"results"`
	LastPage int `json:"last_page"`
}

type SearchQuery struct {
	Query        string
	Page         int
	Type         string
	Status       string
	Rated        string
	Genre        []int
	Score        float64
	StartDate    time.Time
	EndDate      time.Time
	GenreExclude bool
	Limit        int
	OrderBy      string
	Sort         string
}

func NewSearchQuery(query string, page int) SearchQuery {
	return SearchQuery{Query: query, Page: page, Type: "", Status: "", Rated: "", Genre: []int{}, Score: 0, StartDate: time.Unix(0, 0), EndDate: time.Unix(0, 0), GenreExclude: false, Limit: -1, OrderBy: "", Sort: ""}
}

func BuildQueryString(search SearchQuery) string {
	queryString := ""
	if search.Type != "" {
		queryString += fmt.Sprintf("&type=%s", search.Type)
	}
	if search.Status != "" {
		queryString += fmt.Sprintf("&status=%s", search.Status)
	}
	if search.Rated != "" {
		queryString += fmt.Sprintf("&rated=%s", search.Rated)
	}
	if len(search.Genre) != 0 {
		queryString += "&genre=" + strings.Trim(strings.Join(strings.Fields(fmt.Sprint(search.Genre)), ","), "[]")
	}
	if search.Score != 0 {
		queryString += fmt.Sprintf("&score=%v", search.Score)
	}
	if search.StartDate != time.Unix(0, 0) {
		queryString += fmt.Sprintf("&start_date=%s", search.StartDate.Format("2006-01-02"))
	}
	if search.EndDate != time.Unix(0, 0) {
		queryString += fmt.Sprintf("&end_date=%s", search.EndDate.Format("2006-01-02"))
	}
	if search.GenreExclude != false {
		queryString += "&genre_exclude=1"
	}
	if search.Limit != -1 {
		queryString += fmt.Sprintf("&limit=%d", search.Limit)
	}
	if search.OrderBy != "" {
		queryString += fmt.Sprintf("&order_by=%s", search.OrderBy)
	}
	if search.Sort != "" {
		queryString += fmt.Sprintf("&sort=%s", search.Sort)
	}
	return queryString
}

func GetSearchAnimeInfo(search SearchQuery) (SearchAnimeInfo, error) {
	queryString := fmt.Sprintf("%s/search/anime/?q=%s&page=%d%s", Endpoint, search.Query, search.Page, BuildQueryString(search))
	result := SearchAnimeInfo{}
	err := ParseResults(queryString, &result)
	return result, err
}

func GetSearchMangaInfo(search SearchQuery) (SearchMangaInfo, error) {
	queryString := fmt.Sprintf("%s/search/manga/?q=%s&page=%d%s", Endpoint, search.Query, search.Page, BuildQueryString(search))
	result := SearchMangaInfo{}
	err := ParseResults(queryString, &result)
	return result, err
}

func GetSearchPersonInfo(search SearchQuery) (SearchPersonInfo, error) {
	queryString := fmt.Sprintf("%s/search/person/?q=%s&page=%d", Endpoint, search.Query, search.Page)
	result := SearchPersonInfo{}
	err := ParseResults(queryString, &result)
	return result, err
}

func GetSearchCharacterInfo(search SearchQuery) (SearchCharacterInfo, error) {
	queryString := fmt.Sprintf("%s/search/character/?q=%s&page=%d", Endpoint, search.Query, search.Page)
	result := SearchCharacterInfo{}
	err := ParseResults(queryString, &result)
	return result, err
}
