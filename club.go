package jikan

import (
	"fmt"
	"time"
)

type ClubInfo struct {
	MalID         int       `json:"mal_id"`
	URL           string    `json:"url"`
	ImageURL      string    `json:"image_url"`
	Title         string    `json:"title"`
	MembersCount  int       `json:"members_count"`
	PicturesCount int       `json:"pictures_count"`
	Category      string    `json:"category"`
	Created       time.Time `json:"created"`
	Type          string    `json:"type"`
	Staff         []struct {
		MalID int    `json:"mal_id"`
		Type  string `json:"type"`
		Name  string `json:"name"`
		URL   string `json:"url"`
	} `json:"staff"`
	AnimeRelations []struct {
		MalID int    `json:"mal_id"`
		Type  string `json:"type"`
		Name  string `json:"name"`
		URL   string `json:"url"`
	} `json:"anime_relations"`
	MangaRelations []struct {
		MalID int    `json:"mal_id"`
		Type  string `json:"type"`
		Name  string `json:"name"`
		URL   string `json:"url"`
	} `json:"manga_relations"`
	CharacterRelations []struct {
		MalID int    `json:"mal_id"`
		Type  string `json:"type"`
		Name  string `json:"name"`
		URL   string `json:"url"`
	} `json:"character_relations"`
}

type ClubMembers struct {
	Members []struct {
		Username string `json:"username"`
		URL      string `json:"url"`
		ImageURL string `json:"image_url"`
	} `json:"members"`
}

func GetClubInfo(id int) (ClubInfo, error) {
	club := ClubInfo{}
	err := ParseResults(fmt.Sprintf("%s/club/%d", Endpoint, id), &club)
	return club, err
}

func GetClubMembers(id int, page int) (ClubMembers, error) {
	members := ClubMembers{}
	err := ParseResults(fmt.Sprintf("%s/magazine/%d/members/%d", Endpoint, id, page), &members)
	return members, err
}
