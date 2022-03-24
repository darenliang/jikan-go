package jikan

import (
	"testing"
)

func TestGetAnimeGenres(t *testing.T) {
	if _, err := GetAnimeGenres(1, 1, GenresFilterGenres); err != nil {
		t.Error(err)
	}
}

func TestGetMangaGenres(t *testing.T) {
	if _, err := GetMangaGenres(1, 1, GenresFilterGenres); err != nil {
		t.Error(err)
	}
}
