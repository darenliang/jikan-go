package jikan

import "fmt"

// Producers struct
type Producers struct {
	Data       []MalItemCount `json:"data"`
	Pagination Pagination     `json:"pagination"`
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
