package jikan

import (
	"fmt"
)

type Producer struct {
	Meta  malItem         `json:"meta"`
	Anime []malAnimeShort `json:"anime"`
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
