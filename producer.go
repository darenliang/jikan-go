package jikan

import (
	"fmt"
)

// Producer struct for the /producer endpoint
type Producer struct {
	Meta  MalItem        `json:"meta"`
	Anime []MalAnimeDesc `json:"anime"`
}

// GetProducer returns producer
func GetProducer(id, page int) (*Producer, error) {
	res := &Producer{}

	err := urlToStruct(fmt.Sprintf("/producer/%d/%d", id, page), res)

	if err != nil {
		return nil, err
	}

	return res, nil
}
