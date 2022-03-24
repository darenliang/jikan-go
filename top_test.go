package jikan

import (
	"testing"
)

func TestGetTopAnime(t *testing.T) {
	if _, err := GetTopAnime("", "", 1); err != nil {
		t.Error(err)
	}
}

func TestGetTopManga(t *testing.T) {
	if _, err := GetTopManga("", "", 1); err != nil {
		t.Error(err)
	}
}

func TestGetTopPeople(t *testing.T) {
	if _, err := GetTopPeople(1); err != nil {
		t.Error(err)
	}
}

func TestGetTopCharacters(t *testing.T) {
	if _, err := GetTopCharacters(1); err != nil {
		t.Error(err)
	}
}

func TestGetTopReviews(t *testing.T) {
	if _, err := GetTopReviews(1); err != nil {
		t.Error(err)
	}
}
