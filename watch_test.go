package jikan

import (
	"testing"
)

func TestGetWatchRecentEpisodes(t *testing.T) {
	if _, err := GetWatchRecentEpisodes(); err != nil {
		t.Error(err)
	}
}

func TestGetWatchPopularEpisodes(t *testing.T) {
	if _, err := GetWatchPopularEpisodes(); err != nil {
		t.Error(err)
	}
}

func TestGetWatchRecentPromos(t *testing.T) {
	if _, err := GetWatchRecentPromos(); err != nil {
		t.Error(err)
	}
}

func TestGetWatchPopularPromos(t *testing.T) {
	if _, err := GetWatchPopularPromos(); err != nil {
		t.Error(err)
	}
}
