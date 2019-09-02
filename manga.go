package jikan

import "fmt"

// Manga struct defines a manga
type Manga struct {
	ID      int
	Request string
	Page    int
}

// GetManga returns a map of a manga as specified in the Manga struct
func GetManga(manga Manga) (map[string]interface{}, error) {
	var result map[string]interface{}
	var err error
	switch manga.Request {
	case
		"reviews",
		"userupdates":
		result, err = getMapFromUrl(fmt.Sprintf("/manga/%v/%v/%v", manga.ID, manga.Request, manga.Page)), nil
	default:
		result, err = getMapFromUrl(fmt.Sprintf("/manga/%v/%v", manga.ID, manga.Request)), nil
	}
	if _, ok := result["error"]; ok {
		result, err = nil, fmt.Errorf("error %v, %v, %v", result["status"], result["message"], result["error"])
	}
	return result, err
}
