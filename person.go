package jikan

import (
	"fmt"
	"time"
)

// Person struct
type Person struct {
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

// PersonPictures struct
type PersonPictures = AnimePictures

// GetPerson returns person
func GetPerson(id int) (*Person, error) {
	res := &Person{}

	err := urlToStruct(fmt.Sprintf("/person/%d", id), res)

	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetPersonPictures returns person pictures
func GetPersonPictures(id int) (*PersonPictures, error) {
	res := &PersonPictures{}

	err := urlToStruct(fmt.Sprintf("/person/%d/pictures", id), res)

	if err != nil {
		return nil, err
	}

	return res, nil
}
