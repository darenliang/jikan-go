package jikan

import (
	"testing"
)

func TestGetRecentAnimeRecommendations(t *testing.T) {
	if _, err := GetRecentAnimeRecommendations(); err != nil {
		t.Error(err)
	}
}

func TestGetRecentMangaRecommendations(t *testing.T) {
	if _, err := GetRecentMangaRecommendations(); err != nil {
		t.Error(err)
	}
}
