package jikan

import "fmt"

type CharacterInfo struct {
	MalID              int      `json:"mal_id"`
	URL                string   `json:"url"`
	Name               string   `json:"name"`
	NameKanji          string   `json:"name_kanji"`
	Nicknames          []string `json:"nicknames"`
	About              string   `json:"about"`
	MemberFavorites    int      `json:"member_favorites"`
	ImageURL           string   `json:"image_url"`
	Animeography       []struct {
		MalID    int    `json:"mal_id"`
		Name     string `json:"name"`
		URL      string `json:"url"`
		ImageURL string `json:"image_url"`
		Role     string `json:"role"`
	} `json:"animeography"`
	Mangaography []struct {
		MalID    int    `json:"mal_id"`
		Name     string `json:"name"`
		URL      string `json:"url"`
		ImageURL string `json:"image_url"`
		Role     string `json:"role"`
	} `json:"mangaography"`
	VoiceActors []struct {
		MalID    int    `json:"mal_id"`
		Name     string `json:"name"`
		URL      string `json:"url"`
		ImageURL string `json:"image_url"`
		Language string `json:"language"`
	} `json:"voice_actors"`
}

func GetCharacterInfo(id int) (CharacterInfo, error) {
	character := CharacterInfo{}
	err := ParseResults(fmt.Sprintf("%s/character/%d", Endpoint, id), &character)
	return character, err
}

func GetCharacterPictures(id int) (Pictures, error) {
	return getPictures(fmt.Sprintf("%s/character/%d/pictures", Endpoint, id))
}
