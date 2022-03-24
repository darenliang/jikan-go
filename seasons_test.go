package jikan

import (
	"testing"
)

func TestGetSeason(t *testing.T) {
	if _, err := GetSeason(2022, "Winter"); err != nil {
		t.Error(err)
	}
}

func TestGetSeasonNow(t *testing.T) {
	if _, err := GetSeasonNow(); err != nil {
		t.Error(err)
	}
}

func TestGetSeasonList(t *testing.T) {
	if _, err := GetSeasonsList(); err != nil {
		t.Error(err)
	}
}

func TestGetSeasonUpcoming(t *testing.T) {
	if _, err := GetSeasonUpcoming(); err != nil {
		t.Error(err)
	}
}
