package jikan

import (
	"net/url"
	"testing"
)

func TestGetClubsById(t *testing.T) {
	if _, err := GetClubsById(1); err != nil {
		t.Error(err)
	}
}

func TestGetClubMembers(t *testing.T) {
	if _, err := GetClubMembers(1, 1); err != nil {
		t.Error(err)
	}
}

func TestGetClubStaff(t *testing.T) {
	if _, err := GetClubStaff(1); err != nil {
		t.Error(err)
	}
}

func TestGetClubRelations(t *testing.T) {
	if _, err := GetClubRelations(1); err != nil {
		t.Error(err)
	}
}

func TestGetClubsSearch(t *testing.T) {
	query := url.Values{}
	query.Set("q", "Cowboy Bebop")
	if _, err := GetClubsSearch(query); err != nil {
		t.Error(err)
	}
}
