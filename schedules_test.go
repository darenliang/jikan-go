package jikan

import (
	"testing"
)

func TestGetSchedules(t *testing.T) {
	if _, err := GetSchedules(ScheduleFilterMonday); err != nil {
		t.Error(err)
	}
}
