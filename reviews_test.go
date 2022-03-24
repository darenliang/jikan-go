package jikan

import (
	"testing"
)

func TestGetRecentAnimeReviews(t *testing.T) {
	if _, err := GetRecentAnimeReviews(); err != nil {
		t.Error(err)
	}
}

func TestGetRecentMangaReviews(t *testing.T) {
	if _, err := GetRecentMangaReviews(); err != nil {
		t.Error(err)
	}
}
