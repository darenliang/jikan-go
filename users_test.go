package jikan

import (
	"net/url"
	"testing"
)

func TestGetUsersSearch(t *testing.T) {
	query := url.Values{}
	query.Set("q", "Xinil")
	if _, err := GetUsersSearch(query); err != nil {
		t.Error(err)
	}
}

func TestGetUserByID(t *testing.T) {
	if _, err := GetUserByID(1); err != nil {
		t.Error(err)
	}
}

func TestGetUserProfile(t *testing.T) {
	if _, err := GetUserProfile("Xinil"); err != nil {
		t.Error(err)
	}
}

func TestGetUserStatistics(t *testing.T) {
	if _, err := GetUserStatistics("Xinil"); err != nil {
		t.Error(err)
	}
}

func TestGetUserFavorites(t *testing.T) {
	if _, err := GetUserFavorites("Xinil"); err != nil {
		t.Error(err)
	}
}

func TestGetUserUpdates(t *testing.T) {
	if _, err := GetUserUpdates("Xinil"); err != nil {
		t.Error(err)
	}
}

func TestGetUserAbout(t *testing.T) {
	if _, err := GetUserAbout("Xinil"); err != nil {
		t.Error(err)
	}
}

func TestGetUserHistory(t *testing.T) {
	if _, err := GetUserHistory("Xinil", UserHistoryFilterAnime); err != nil {
		t.Error(err)
	}
}

func TestGetUserFriends(t *testing.T) {
	if _, err := GetUserFriends("Xinil"); err != nil {
		t.Error(err)
	}
}

func TestGetUserAnimelist(t *testing.T) {
	if _, err := GetUserAnimelist("Xinil"); err != nil {
		t.Error(err)
	}
}

func TestGetUserMangalist(t *testing.T) {
	if _, err := GetUserMangalist("Xinil"); err != nil {
		t.Error(err)
	}
}

func TestGetUserReviews(t *testing.T) {
	if _, err := GetUserReviews("Xinil"); err != nil {
		t.Error(err)
	}
}

func TestGetUserRecommendations(t *testing.T) {
	if _, err := GetUserRecommendations("Xinil"); err != nil {
		t.Error(err)
	}
}

func TestGetUserClubs(t *testing.T) {
	if _, err := GetUserClubs("Xinil"); err != nil {
		t.Error(err)
	}
}
