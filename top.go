package jikan

import (
	"fmt"
	"strings"
)

// TopQuery struct
type TopQuery struct {
	Type string // object type

	Page    int    // page number (optional)
	Subtype string // subtype (available upon page specification)
}

// Top struct
type Top struct {
	Top []struct {
		MalID     int     `json:"mal_id"`
		Rank      int     `json:"rank"`
		Title     string  `json:"title"`
		URL       string  `json:"url"`
		ImageURL  string  `json:"image_url"`
		Type      string  `json:"type"`
		Episodes  int     `json:"episodes"`
		StartDate string  `json:"start_date"`
		EndDate   string  `json:"end_date"`
		Members   int     `json:"members"`
		Score     float64 `json:"score"`
	} `json:"top"`
}

// GetTop returns top
func GetTop(mediaType string, page int, subType string) (*Top, error) {
	res := &Top{}

	var query strings.Builder
	query.WriteString(fmt.Sprintf("/top/%v", mediaType))
	if page != 0 {
		query.WriteString(fmt.Sprintf("/%v", page))
		if subType != "" {
			query.WriteString(fmt.Sprintf("/%v", subType))
		}
	}

	err := urlToStruct(query.String(), res)

	if err != nil {
		return nil, err
	}

	return res, nil
}
