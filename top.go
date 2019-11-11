package jikan

import (
	"fmt"
	"strings"
)

// Top struct defines a top
type Top struct {
	Type    string // object type
	Page    int    // page number (optional)
	Subtype string // subtype (available upon page specification)
}

// Get returns a map of a top as specified in the Top struct
// Calls responses through the /top/ endpoint
func (top Top) Get() (map[string]interface{}, error) {
	var result map[string]interface{}
	var err error
	var query strings.Builder
	query.WriteString(fmt.Sprintf("/top/%v", top.Type))
	if top.Page != 0 {
		query.WriteString(fmt.Sprintf("/%v", top.Page))
		if top.Subtype != "" {
			query.WriteString(fmt.Sprintf("/%v", top.Subtype))
		}
	}
	result, err = getMapFromUrl(query.String()), nil
	if _, ok := result["error"]; ok {
		result, err = nil, getResultError(result)
	}
	return result, err
}
