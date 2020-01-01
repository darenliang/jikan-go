package jikan

import "fmt"

// Person struct defines a person
type Person struct {
	ID      int    // MyAnimeList person ID
	Request string // request type (optional)
}

// Get returns a map of a person as specified in the Person struct.
// Calls responses through the /person/ endpoint
func (person Person) Get() (map[string]interface{}, error) {
	var result map[string]interface{}
	var err error
	result, err = getMapFromUrl(fmt.Sprintf("/person/%v/%v", person.ID, person.Request)), nil
	if _, ok := result["error"]; ok {
		result, err = nil, getResultError(result)
	}
	return result, err
}
