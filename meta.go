package jikan

import (
	"fmt"
	"strings"
)

// Meta struct defines a meta
type Meta struct {
	Request string // request type
	Type    string // object Type
	Period  string // request Period
	Offset  int    // page number (optional)
}

// Get returns a map of a meta as specified in the Meta struct
// Calls responses through the /meta/ endpoint
func (meta Meta) Get() (map[string]interface{}, error) {
	var result map[string]interface{}
	var err error
	var query strings.Builder
	query.WriteString(fmt.Sprintf("/meta/%v/%v/%v", meta.Request, meta.Type, meta.Period))
	if meta.Offset != 0 {
		query.WriteString(fmt.Sprintf("/%v", meta.Offset))
	}
	result, err = getMapFromUrl(query.String()), nil
	if _, ok := result["error"]; ok {
		result, err = nil, getResultError(result)
	}
	return result, err
}
