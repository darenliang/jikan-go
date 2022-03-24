package jikan

import (
	"net/url"
	"testing"
)

func TestGetCharacterById(t *testing.T) {
	if _, err := GetCharacterById(1); err != nil {
		t.Error(err)
	}
}

func TestGetCharacterAnime(t *testing.T) {
	if _, err := GetCharacterAnime(1); err != nil {
		t.Error(err)
	}
}

func TestGetCharacterManga(t *testing.T) {
	if _, err := GetCharacterManga(1); err != nil {
		t.Error(err)
	}
}

func TestGetCharacterVoiceActors(t *testing.T) {
	if _, err := GetCharacterVoiceActors(1); err != nil {
		t.Error(err)
	}
}

func TestGetCharacterPictures(t *testing.T) {
	if _, err := GetCharacterPictures(1); err != nil {
		t.Error(err)
	}
}

func TestGetCharactersSearch(t *testing.T) {
	query := url.Values{}
	query.Set("q", "Spike Spiegel")
	if _, err := GetCharactersSearch(query); err != nil {
		t.Error(err)
	}
}
