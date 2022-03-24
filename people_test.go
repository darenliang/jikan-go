package jikan

import (
	"net/url"
	"testing"
)

func TestGetPersonById(t *testing.T) {
	if _, err := GetPersonById(1); err != nil {
		t.Error(err)
	}
}
func TestGetPersonAnime(t *testing.T) {
	if _, err := GetPersonAnime(1); err != nil {
		t.Error(err)
	}
}
func TestGetPersonVoices(t *testing.T) {
	if _, err := GetPersonVoices(1); err != nil {
		t.Error(err)
	}
}
func TestGetPersonManga(t *testing.T) {
	if _, err := GetPersonManga(1); err != nil {
		t.Error(err)
	}
}

func TestGetPersonPictures(t *testing.T) {
	if _, err := GetPersonPictures(1); err != nil {
		t.Error(err)
	}
}

func TestGetPeopleSearch(t *testing.T) {
	query := url.Values{}
	query.Set("q", "Tomokazu Seki")
	if _, err := GetPeopleSearch(query); err != nil {
		t.Error(err)
	}
}
