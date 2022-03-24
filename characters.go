package jikan

import (
	"fmt"
	"net/url"
)

// CharactersBase struct
type CharactersBase struct {
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
	Name      string   `json:"name"`
	NameKanji string   `json:"name_kanji"`
	Nicknames []string `json:"nicknames"`
	Favorites int      `json:"favorites"`
	About     string   `json:"about"`
}

// CharacterById struct
type CharacterById struct {
	Data CharactersBase `json:"data"`
}

// GetCharacterById returns character by id
func GetCharacterById(id int) (*CharacterById, error) {
	res := &CharacterById{}
	err := urlToStruct(fmt.Sprintf("/characters/%d", id), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// CharacterAnime struct
type CharacterAnime struct {
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
	} `json:"data"`
}

// GetCharacterAnime returns character anime
func GetCharacterAnime(id int) (*CharacterAnime, error) {
	res := &CharacterAnime{}
	err := urlToStruct(fmt.Sprintf("/characters/%d/anime", id), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// CharacterManga struct
type CharacterManga struct {
	Data []struct {
		Role  string `json:"role"`
		Manga struct {
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

// GetCharacterManga returns character manga
func GetCharacterManga(id int) (*CharacterManga, error) {
	res := &CharacterManga{}
	err := urlToStruct(fmt.Sprintf("/characters/%d/manga", id), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// CharacterVoiceActors struct
type CharacterVoiceActors struct {
	Data []struct {
		Language string `json:"language"`
		Person   struct {
			MalId  int    `json:"mal_id"`
			Url    string `json:"url"`
			Images struct {
				Jpg struct {
					ImageUrl string `json:"image_url"`
				} `json:"jpg"`
			} `json:"images"`
			Name string `json:"name"`
		} `json:"person"`
	} `json:"data"`
}

// GetCharacterVoiceActors returns character manga
func GetCharacterVoiceActors(id int) (*CharacterVoiceActors, error) {
	res := &CharacterVoiceActors{}
	err := urlToStruct(fmt.Sprintf("/characters/%d/voices", id), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type CharacterPictures struct {
	Data []struct {
		ImageUrl      string `json:"image_url"`
		LargeImageUrl string `json:"large_image_url"`
	} `json:"data"`
}

// GetCharacterPictures returns character manga
func GetCharacterPictures(id int) (*CharacterPictures, error) {
	res := &CharacterPictures{}
	err := urlToStruct(fmt.Sprintf("/characters/%d/pictures", id), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// CharactersSearch struct
type CharactersSearch struct {
	Data       []CharacterById `json:"data"`
	Pagination struct {
		LastVisiblePage int  `json:"last_visible_page"`
		HasNextPage     bool `json:"has_next_page"`
	} `json:"pagination"`
}

// GetCharactersSearch returns characters search
func GetCharactersSearch(query url.Values) (*CharactersSearch, error) {
	res := &CharactersSearch{}
	err := urlToStruct(fmt.Sprintf("/characters?%s", query.Encode()), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
