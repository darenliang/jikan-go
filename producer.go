package jikan

import (
	"fmt"
	"strings"
)

// Producer struct defines a producer
type Producer struct {
	ProducerID int // MyAnimeList producer ID
	Page       int // page number (optional)
}

// Get returns a map of a producer as specified in the Producer struct.
// Calls responses through the /producer/ endpoint.
func (producer Producer) Get() (map[string]interface{}, error) {
	var result map[string]interface{}
	var err error
	var query strings.Builder
	query.WriteString(fmt.Sprintf("/producer/%v", producer.ProducerID))
	if producer.Page != 0 {
		query.WriteString(fmt.Sprintf("/%v", producer.Page))
	}
	result, err = getMapFromUrl(query.String()), nil
	if _, ok := result["error"]; ok {
		result, err = nil, getResultError(result)
	}
	return result, err
}
