package jikan

import (
	"net/url"
	"testing"
)

func TestGetAnimeById(t *testing.T) {
	if _, err := GetAnimeById(1); err != nil {
		t.Error(err)
	}
}
func TestGetAnimeCharacters(t *testing.T) {
	if _, err := GetAnimeCharacters(1); err != nil {
		t.Error(err)
	}
}
func TestGetAnimeStaff(t *testing.T) {
	if _, err := GetAnimeStaff(1); err != nil {
		t.Error(err)
	}
}
func TestGetAnimeEpisodes(t *testing.T) {
	if _, err := GetAnimeEpisodes(1, 1); err != nil {
		t.Error(err)
	}
}

func TestGetAnimeEpisodeById(t *testing.T) {
	if _, err := GetAnimeEpisodeById(1, 1); err != nil {
		t.Error(err)
	}
}

func TestGetAnimeNews(t *testing.T) {
	if _, err := GetAnimeNews(1, 1); err != nil {
		t.Error(err)
	}
}

func TestGetAnimeForum(t *testing.T) {
	if _, err := GetAnimeForum(1, AnimeForumFilterAll); err != nil {
		t.Error(err)
	}
}

func TestGetAnimeVideos(t *testing.T) {
	if _, err := GetAnimeVideos(1); err != nil {
		t.Error(err)
	}
}

func TestGetAnimePictures(t *testing.T) {
	if _, err := GetAnimePictures(1); err != nil {
		t.Error(err)
	}
}

func TestGetAnimeStatistics(t *testing.T) {
	if _, err := GetAnimeStatistics(1); err != nil {
		t.Error(err)
	}
}

func TestGetAnimeMoreInfo(t *testing.T) {
	if _, err := GetAnimeMoreInfo(1); err != nil {
		t.Error(err)
	}
}

func TestGetAnimeRecommendations(t *testing.T) {
	if _, err := GetAnimeRecommendations(1); err != nil {
		t.Error(err)
	}
}

func TestGetAnimeUserUpdates(t *testing.T) {
	if _, err := GetAnimeUserUpdates(1, 1); err != nil {
		t.Error(err)
	}
}

func TestGetAnimeReviews(t *testing.T) {
	if _, err := GetAnimeReviews(1, 1); err != nil {
		t.Error(err)
	}
}

func TestGetAnimeRelations(t *testing.T) {
	if _, err := GetAnimeRelations(1); err != nil {
		t.Error(err)
	}
}

func TestGetAnimeThemes(t *testing.T) {
	if _, err := GetAnimeThemes(1); err != nil {
		t.Error(err)
	}
}

func TestGetAnimeExternal(t *testing.T) {
	if _, err := GetAnimeExternal(1); err != nil {
		t.Error(err)
	}
}

func TestGetAnimeSearch(t *testing.T) {
	query := url.Values{}
	query.Set("q", "Cowboy Bebop")
	if _, err := GetAnimeSearch(query); err != nil {
		t.Error(err)
	}
}
