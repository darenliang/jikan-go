package jikan

import (
	"fmt"
)

// Magazines struct
type Magazines struct {
	Data []struct {
		MalId int    `json:"mal_id"`
		Name  string `json:"name"`
		Url   string `json:"url"`
		Count int    `json:"count"`
	} `json:"data"`
	Pagination struct {
		LastVisiblePage int  `json:"last_visible_page"`
		HasNextPage     bool `json:"has_next_page"`
	} `json:"pagination"`
}

// GetMagazines returns magazines
func GetMagazines(page int) (*Magazines, error) {
	res := &Magazines{}
	err := urlToStruct(fmt.Sprintf("/magazines?page=%d", page), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
