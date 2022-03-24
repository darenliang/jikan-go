package jikan

import (
	"net/url"
	"testing"
)

func TestGetMangaById(t *testing.T) {
	if _, err := GetMangaById(1); err != nil {
		t.Error(err)
	}
}
func TestGetMangaCharacters(t *testing.T) {
	if _, err := GetMangaCharacters(1); err != nil {
		t.Error(err)
	}
}
func TestGetMangaNews(t *testing.T) {
	if _, err := GetMangaNews(1, 1); err != nil {
		t.Error(err)
	}
}
func TestGetMangaTopics(t *testing.T) {
	if _, err := GetMangaForum(1, MangaForumFilterAll); err != nil {
		t.Error(err)
	}
}

func TestGetMangaPictures(t *testing.T) {
	if _, err := GetMangaPictures(1); err != nil {
		t.Error(err)
	}
}

func TestGetMangaStatistics(t *testing.T) {
	if _, err := GetMangaStatistics(1); err != nil {
		t.Error(err)
	}
}

func TestGetMangaMoreInfo(t *testing.T) {
	if _, err := GetMangaMoreInfo(1); err != nil {
		t.Error(err)
	}
}

func TestGetMangaRecommendations(t *testing.T) {
	if _, err := GetMangaRecommendations(1); err != nil {
		t.Error(err)
	}
}

func TestGetMangaUserUpdates(t *testing.T) {
	if _, err := GetMangaUserUpdates(1, 1); err != nil {
		t.Error(err)
	}
}

func TestGetMangaReviews(t *testing.T) {
	if _, err := GetMangaReviews(1, 1); err != nil {
		t.Error(err)
	}
}

func TestGetMangaRelations(t *testing.T) {
	if _, err := GetMangaRelations(1); err != nil {
		t.Error(err)
	}
}

func TestGetMangaExternal(t *testing.T) {
	if _, err := GetMangaExternal(1); err != nil {
		t.Error(err)
	}
}

func TestGetMangaSearch(t *testing.T) {
	query := url.Values{}
	query.Set("q", "Monster")
	if _, err := GetMangaSearch(query); err != nil {
		t.Error(err)
	}
}
