package jikan

import "fmt"

// Schedule struct defines a producer
type Schedule struct {
	Day string // day of the week
}

// Get returns a map of a schedule as specified in the Schedule struct.
// Calls responses through the /schedule/ endpoint.
func (schedule Schedule) Get() (map[string]interface{}, error) {
	var result map[string]interface{}
	var err error
	result, err = getMapFromUrl(fmt.Sprintf("/schedule/%v", schedule.Day)), nil
	if _, ok := result["error"]; ok {
		result, err = nil, getResultError(result)
	}
	return result, err
}
