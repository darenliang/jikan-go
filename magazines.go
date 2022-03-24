package jikan

import (
	"fmt"
)

// Magazines struct
type Magazines struct {
	Data       []MalItemCount `json:"data"`
	Pagination Pagination     `json:"pagination"`
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
