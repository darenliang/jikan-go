package jikan

import "fmt"

// Person struct defines a person
type Person struct {
	ID      int
	Request string
}

// GetPerson returns a map of a person as specified in the Person struct
func GetPerson(person Person) (map[string]interface{}, error) {
	var result map[string]interface{}
	var err error
	result, err = getMapFromUrl(fmt.Sprintf("/person/%v/%v", person.ID, person.Request)), nil
	if _, ok := result["error"]; ok {
		result, err = nil, fmt.Errorf("error %v, %v, %v", result["status"], result["message"], result["error"])
	}
	return result, err
}
