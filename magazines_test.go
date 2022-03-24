package jikan

import (
	"testing"
)

func TestGetMagazines(t *testing.T) {
	if _, err := GetMagazines(1); err != nil {
		t.Error(err)
	}
}
