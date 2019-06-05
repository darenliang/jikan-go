package jikan

import (
	"encoding/json"
	"github.com/aerogo/http/client"
	"time"
)

type News struct {
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

type Pictures struct {
	Pictures []struct {
		Large string `json:"large"`
		Small string `json:"small"`
	} `json:"pictures"`
}
type Forum struct {
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

type Moreinfo struct {
	Moreinfo string `json:"moreinfo"`
}

type Recommendations struct {
	Recommendations []struct {
		MalID               int    `json:"mal_id"`
		URL                 string `json:"url"`
		ImageURL            string `json:"image_url"`
		RecommendationURL   string `json:"recommendation_url"`
		Title               string `json:"title"`
		RecommendationCount int    `json:"recommendation_count"`
	} `json:"recommendations"`
}

func getNews(url string) (News, error) {
	result := News{}
	err := ParseResults(url, &result)
	return result, err
}

func getPictures(url string) (Pictures, error) {
	result := Pictures{}
	err := ParseResults(url, &result)
	return result, err
}

func getForum(url string) (Forum, error) {
	result := Forum{}
	err := ParseResults(url, &result)
	return result, err
}

func getMoreinfo(url string) (Moreinfo, error) {
	result := Moreinfo{}
	err := ParseResults(url, &result)
	return result, err
}

func getRecommendations(url string) (Recommendations, error) {
	result := Recommendations{}
	err := ParseResults(url, &result)
	return result, err
}

func ParseResults(url string, obj interface{}) error {
	response, err := client.Get(url).End()
	if err != nil {
		return err
	}
	jsonString := response.String()
	err = json.Unmarshal([]byte(jsonString), obj)
	return err
}
