package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/aerogo/http/client"
)

type Anime struct {
	MalID         int           `json:"mal_id"`
	URL           string        `json:"url"`
	ImageURL      string        `json:"image_url"`
	TrailerURL    string        `json:"trailer_url"`
	Title         string        `json:"title"`
	TitleEnglish  string        `json:"title_english"`
	TitleJapanese string        `json:"title_japanese"`
	TitleSynonyms []interface{} `json:"title_synonyms"`
	Type          string        `json:"type"`
	Source        string        `json:"source"`
	Episodes      int           `json:"episodes"`
	Status        string        `json:"status"`
	Airing        bool          `json:"airing"`
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

func GetAnime(id string) (*Anime, error) {
	const Endpoint = "https://api.jikan.moe/v3"
	anime := &Anime{}
	response, err := client.Get(Endpoint + "/anime/" + id).End()

	if err != nil {
		return nil, err
	}
	jsonString := response.String()
	err = json.Unmarshal([]byte(jsonString), anime)
	return anime, err
}

func main() {
	anime, err := GetAnime("1")
	fmt.Println(anime, err)
}
