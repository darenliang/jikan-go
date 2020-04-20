package jikan

import (
	"fmt"
)

// Magazine struct for the /magazine endpoint
type Magazine struct {
	Meta  MalItem        `json:"meta"`
	Manga []MalMangaDesc `json:"manga"`
}

// GetMagazine returns magazine
func GetMagazine(id, page int) (*Magazine, error) {
	res := &Magazine{}

	err := urlToStruct(fmt.Sprintf("/magazine/%d/%d", id, page), res)

	if err != nil {
		return nil, err
	}

	return res, nil
}
