package jikan

import (
	"fmt"
	"strings"
)

// Producer struct defines a producer
type Producer struct {
	ProducerID int
	Page       int
}

// GetProducer returns a map of a producer as specified in the Producer struct
func GetProducer(producer Producer) (map[string]interface{}, error) {
	var result map[string]interface{}
	var err error
	var query strings.Builder
	query.WriteString(fmt.Sprintf("/producer/%v", producer.ProducerID))
	if producer.Page != 0 {
		query.WriteString(fmt.Sprintf("/%v", producer.Page))
	}
	result, err = getMapFromUrl(query.String()), nil
	if _, ok := result["error"]; ok {
		result, err = nil, fmt.Errorf("error %v, %v, %v", result["status"], result["message"], result["error"])
	}
	return result, err
}
