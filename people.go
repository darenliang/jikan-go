package jikan

import (
	"fmt"
	"net/url"
	"time"
)

// PeopleBase struct
type PeopleBase struct {
	MalId      int    `json:"mal_id"`
	Url        string `json:"url"`
	WebsiteUrl string `json:"website_url"`
	Images     struct {
		Jpg struct {
			ImageUrl string `json:"image_url"`
		} `json:"jpg"`
	} `json:"images"`
	Name           string    `json:"name"`
	GivenName      string    `json:"given_name"`
	FamilyName     string    `json:"family_name"`
	AlternateNames []string  `json:"alternate_names"`
	Birthday       time.Time `json:"birthday"`
	Favorites      int       `json:"favorites"`
	About          string    `json:"about"`
}

// PersonById struct
type PersonById struct {
	Data PeopleBase `json:"data"`
}

// GetPersonById returns person by id
func GetPersonById(id int) (*PersonById, error) {
	res := &PersonById{}
	err := urlToStruct(fmt.Sprintf("/people/%d", id), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// PersonAnime struct
type PersonAnime struct {
	Data []struct {
		Position string `json:"position"`
		Anime    struct {
			MalId  int    `json:"mal_id"`
			Url    string `json:"url"`
			Images struct {
				Jpg struct {
					ImageUrl      string `json:"image_url"`
					SmallImageUrl string `json:"small_image_url"`
					LargeImageUrl string `json:"large_image_url"`
				} `json:"jpg"`
				Webp struct {
					ImageUrl      string `json:"image_url"`
					SmallImageUrl string `json:"small_image_url"`
					LargeImageUrl string `json:"large_image_url"`
				} `json:"webp"`
			} `json:"images"`
			Title string `json:"title"`
		} `json:"anime"`
	} `json:"data"`
}

// GetPersonAnime returns person anime
func GetPersonAnime(id int) (*PersonAnime, error) {
	res := &PersonAnime{}
	err := urlToStruct(fmt.Sprintf("/people/%d/anime", id), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// PersonVoices struct
type PersonVoices struct {
	Data []struct {
		Role  string `json:"role"`
		Anime struct {
			MalId  int    `json:"mal_id"`
			Url    string `json:"url"`
			Images struct {
				Jpg struct {
					ImageUrl      string `json:"image_url"`
					SmallImageUrl string `json:"small_image_url"`
					LargeImageUrl string `json:"large_image_url"`
				} `json:"jpg"`
				Webp struct {
					ImageUrl      string `json:"image_url"`
					SmallImageUrl string `json:"small_image_url"`
					LargeImageUrl string `json:"large_image_url"`
				} `json:"webp"`
			} `json:"images"`
			Title string `json:"title"`
		} `json:"anime"`
		Character struct {
			MalId  int    `json:"mal_id"`
			Url    string `json:"url"`
			Images struct {
				Jpg struct {
					ImageUrl      string `json:"image_url"`
					SmallImageUrl string `json:"small_image_url"`
				} `json:"jpg"`
				Webp struct {
					ImageUrl      string `json:"image_url"`
					SmallImageUrl string `json:"small_image_url"`
				} `json:"webp"`
			} `json:"images"`
			Name string `json:"name"`
		} `json:"character"`
	} `json:"data"`
}

// GetPersonVoices returns person voices
func GetPersonVoices(id int) (*PersonVoices, error) {
	res := &PersonVoices{}
	err := urlToStruct(fmt.Sprintf("/people/%d/voices", id), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// PersonManga struct
type PersonManga struct {
	Data []struct {
		Position string `json:"position"`
		Manga    struct {
			MalId  int    `json:"mal_id"`
			Url    string `json:"url"`
			Images struct {
				Jpg struct {
					ImageUrl      string `json:"image_url"`
					SmallImageUrl string `json:"small_image_url"`
					LargeImageUrl string `json:"large_image_url"`
				} `json:"jpg"`
				Webp struct {
					ImageUrl      string `json:"image_url"`
					SmallImageUrl string `json:"small_image_url"`
					LargeImageUrl string `json:"large_image_url"`
				} `json:"webp"`
			} `json:"images"`
			Title string `json:"title"`
		} `json:"manga"`
	} `json:"data"`
}

// GetPersonManga returns person manga
func GetPersonManga(id int) (*PersonManga, error) {
	res := &PersonManga{}
	err := urlToStruct(fmt.Sprintf("/people/%d/manga", id), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// PersonPictures struct
type PersonPictures struct {
	Data []struct {
		Jpg struct {
			ImageUrl string `json:"image_url"`
		} `json:"jpg"`
	} `json:"data"`
}

// GetPersonPictures returns person pictures
func GetPersonPictures(id int) (*PersonPictures, error) {
	res := &PersonPictures{}
	err := urlToStruct(fmt.Sprintf("/people/%d/pictures", id), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// PeopleSearch struct
type PeopleSearch struct {
	Data       []PeopleBase `json:"data"`
	Pagination struct {
		LastVisiblePage int  `json:"last_visible_page"`
		HasNextPage     bool `json:"has_next_page"`
	} `json:"pagination"`
}

// GetPeopleSearch returns people search
func GetPeopleSearch(query url.Values) (*PeopleSearch, error) {
	res := &PeopleSearch{}
	err := urlToStruct(fmt.Sprintf("/people?%s", query.Encode()), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
