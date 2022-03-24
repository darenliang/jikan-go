package jikan

import (
	"testing"
)

func TestGetRandomAnime(t *testing.T) {
	if _, err := GetRandomAnime(); err != nil {
		t.Error(err)
	}
}

func TestGetRandomManga(t *testing.T) {
	if _, err := GetRandomManga(); err != nil {
		t.Error(err)
	}
}

func TestGetRandomCharacters(t *testing.T) {
	if _, err := GetRandomCharacters(); err != nil {
		t.Error(err)
	}
}

func TestGetRandomPeople(t *testing.T) {
	if _, err := GetRandomPeople(); err != nil {
		t.Error(err)
	}
}

func TestGetRandomUsers(t *testing.T) {
	if _, err := GetRandomUsers(); err != nil {
		t.Error(err)
	}
}
