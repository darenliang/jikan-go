package jikan

import (
	"fmt"
	"time"
)

// Club struct
type Club struct {
	MalID              int       `json:"mal_id"`
	URL                string    `json:"url"`
	ImageURL           string    `json:"image_url"`
	Title              string    `json:"title"`
	MembersCount       int       `json:"members_count"`
	PicturesCount      int       `json:"pictures_count"`
	Category           string    `json:"category"`
	Created            time.Time `json:"created"`
	Type               string    `json:"type"`
	Staff              []malItem `json:"staff"`
	AnimeRelations     []malItem `json:"anime_relations"`
	MangaRelations     []malItem `json:"manga_relations"`
	CharacterRelations []malItem `json:"character_relations"`
}

// ClubMembers struct
type ClubMembers struct {
	Members []struct {
		Username string `json:"username"`
		URL      string `json:"url"`
		ImageURL string `json:"image_url"`
	} `json:"members"`
}

// GetClub returns club
func GetClub(id int) (*Club, error) {
	res := &Club{}

	err := urlToStruct(fmt.Sprintf("/club/%d", id), res)

	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetClubMembers returns club members
func GetClubMembers(id, page int) (*ClubMembers, error) {
	res := &ClubMembers{}

	err := urlToStruct(fmt.Sprintf("/club/%d/members/%d", id, page), res)

	if err != nil {
		return nil, err
	}

	return res, nil
}
