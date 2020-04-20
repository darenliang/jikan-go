package jikan

import "fmt"

// Character struct
type Character struct {
	MalID           int            `json:"mal_id"`
	URL             string         `json:"url"`
	Name            string         `json:"name"`
	NameKanji       string         `json:"name_kanji"`
	Nicknames       []interface{}  `json:"nicknames"`
	About           string         `json:"about"`
	MemberFavorites int            `json:"member_favorites"`
	ImageURL        string         `json:"image_url"`
	Animeography    []malRoleStaff `json:"animeography"`
	Mangaography    []malRoleStaff `json:"mangaography"`
	VoiceActors     []struct {
		MalID    int    `json:"mal_id"`
		Name     string `json:"name"`
		URL      string `json:"url"`
		ImageURL string `json:"image_url"`
		Language string `json:"language"`
	} `json:"voice_actors"`
}

// CharacterPictures struct
type CharacterPictures = AnimePictures

// GetCharacter returns character
func GetCharacter(id int) (*Character, error) {
	res := &Character{}

	err := urlToStruct(fmt.Sprintf("/character/%d", id), res)

	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetCharacterPictures returns character pictures
func GetCharacterPictures(id int) (*CharacterPictures, error) {
	res := &CharacterPictures{}

	err := urlToStruct(fmt.Sprintf("/character/%d/pictures", id), res)

	if err != nil {
		return nil, err
	}

	return res, nil
}
