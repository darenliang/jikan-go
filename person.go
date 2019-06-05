package jikan

import (
	"fmt"
	"time"
)

type PersonInfo struct {
	MalID            int       `json:"mal_id"`
	URL              string    `json:"url"`
	ImageURL         string    `json:"image_url"`
	WebsiteURL       string    `json:"website_url"`
	Name             string    `json:"name"`
	GivenName        string    `json:"given_name"`
	FamilyName       string    `json:"family_name"`
	AlternateNames   []string  `json:"alternate_names"`
	Birthday         time.Time `json:"birthday"`
	MemberFavorites  int       `json:"member_favorites"`
	About            string    `json:"about"`
	VoiceActingRoles []struct {
		Role  string `json:"role"`
		Anime struct {
			MalID    int    `json:"mal_id"`
			URL      string `json:"url"`
			ImageURL string `json:"image_url"`
			Name     string `json:"name"`
		} `json:"anime"`
		Character struct {
			MalID    int    `json:"mal_id"`
			URL      string `json:"url"`
			ImageURL string `json:"image_url"`
			Name     string `json:"name"`
		} `json:"character"`
	} `json:"voice_acting_roles"`
	AnimeStaffPositions []struct {
		Position string `json:"position"`
		Anime    struct {
			MalID    int    `json:"mal_id"`
			URL      string `json:"url"`
			ImageURL string `json:"image_url"`
			Name     string `json:"name"`
		} `json:"anime"`
	} `json:"anime_staff_positions"`
	PublishedManga []struct {
		Position string `json:"position"`
		Manga    struct {
			MalID    int    `json:"mal_id"`
			URL      string `json:"url"`
			ImageURL string `json:"image_url"`
			Name     string `json:"name"`
		} `json:"manga"`
	} `json:"published_manga"`
}

func GetPersonInfo(id int) (PersonInfo, error) {
	person := PersonInfo{}
	err := ParseResults(fmt.Sprintf("%s/person/%d", Endpoint, id), &person)
	return person, err
}

func GetPersonPictures(id int) (Pictures, error) {
	return getPictures(fmt.Sprintf("%s/person/%d/pictures", Endpoint, id))
}
