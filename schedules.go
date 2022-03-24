package jikan

import "fmt"

// Schedules struct
type Schedules struct {
	Data       []AnimeBase `json:"data"`
	Pagination Pagination  `json:"pagination"`
}

type ScheduleFilter string

const (
	ScheduleFilterMonday    ScheduleFilter = "monday"
	ScheduleFilterTuesday   ScheduleFilter = "tuesday"
	ScheduleFilterWednesday ScheduleFilter = "wednesday"
	ScheduleFilterThursday  ScheduleFilter = "thursday"
	ScheduleFilterFriday    ScheduleFilter = "friday"
	ScheduleFilterUnknown   ScheduleFilter = "unknown"
	ScheduleFilterOther     ScheduleFilter = "other"
)

// GetSchedules returns schedules
func GetSchedules(filter ScheduleFilter) (*Schedules, error) {
	res := &Schedules{}
	req := "/schedules"
	if filter != "" {
		req += fmt.Sprintf("?filter=%s", filter)
	}
	err := urlToStruct(req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
