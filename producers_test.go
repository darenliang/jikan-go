package jikan

import (
	"testing"
)

func TestGetProducers(t *testing.T) {
	if _, err := GetProducers(1); err != nil {
		t.Error(err)
	}
}
