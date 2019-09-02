package jikan

import (
	"fmt"
	"strings"
)

// Meta struct defines a meta
type Meta struct {
	Request string
	Type    string
	Period  string
	Offset  int
}

// GetMeta returns a map of a meta as specified in the Meta struct
func GetMeta(meta Meta) (map[string]interface{}, error) {
	var result map[string]interface{}
	var err error
	var query strings.Builder
	query.WriteString(fmt.Sprintf("/meta/%v/%v/%v", meta.Request, meta.Type, meta.Period))
	if meta.Offset != 0 {
		query.WriteString(fmt.Sprintf("/%v", meta.Offset))
	}
	result, err = getMapFromUrl(query.String()), nil
	if _, ok := result["error"]; ok {
		result, err = nil, fmt.Errorf("error %v, %v, %v", result["status"], result["message"], result["error"])
	}
	return result, err
}
