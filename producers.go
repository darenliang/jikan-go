package jikan

import "fmt"

// Producers struct
type Producers struct {
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

// GetProducers returns producers
func GetProducers(page int) (*Producers, error) {
	res := &Producers{}
	err := urlToStruct(fmt.Sprintf("/producers?page=%d", page), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
