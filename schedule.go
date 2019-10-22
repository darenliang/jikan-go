package jikan

import "fmt"

// GetSchedule returns a map of a schedule depending on the day
// Calls responses through the /schedule/ endpoint
func GetSchedule(day string) (map[string]interface{}, error) {
	var result map[string]interface{}
	var err error
	result, err = getMapFromUrl(fmt.Sprintf("/schedule/%v", day)), nil
	if _, ok := result["error"]; ok {
		result, err = nil, getResultError(result)
	}
	return result, err
}
