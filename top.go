package jikan

import (
	"fmt"
	"net/url"
	"strconv"
	"time"
)

// TopAnime struct
type TopAnime struct {
	Data       []AnimeBase `json:"data"`
	Pagination struct {
		LastVisiblePage int  `json:"last_visible_page"`
		HasNextPage     bool `json:"has_next_page"`
	} `json:"pagination"`
}

type TopAnimeType string

const (
	TopAnimeTypeTv      TopAnimeType = "tv"
	TopAnimeTypeMovie   TopAnimeType = "movie"
	TopAnimeTypeOva     TopAnimeType = "ova"
	TopAnimeTypeSpecial TopAnimeType = "special"
	TopAnimeTypeOna     TopAnimeType = "ona"
	TopAnimeTypeMusic   TopAnimeType = "music"
)

type TopAnimeFilter string

const (
	TopAnimeFilterAiring       TopAnimeFilter = "airing"
	TopAnimeFilterUpcoming     TopAnimeFilter = "upcoming"
	TopAnimeFilterByPopularity TopAnimeFilter = "bypopularity"
	TopAnimeFilterFavorite     TopAnimeFilter = "favorite"
)

// GetTopAnime returns top anime
func GetTopAnime(subType TopAnimeType, filter TopAnimeFilter, page int) (*TopAnime, error) {
	res := &TopAnime{}
	query := url.Values{}
	query.Set("page", strconv.Itoa(page))
	if subType != "" {
		query.Set("type", string(subType))
	}
	if filter != "" {
		query.Set("filter", string(filter))
	}
	err := urlToStruct(fmt.Sprintf("/top/anime?%s", query.Encode()), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// TopManga struct
type TopManga struct {
	Data       []MangaBase `json:"data"`
	Pagination struct {
		LastVisiblePage int  `json:"last_visible_page"`
		HasNextPage     bool `json:"has_next_page"`
	} `json:"pagination"`
}

type TopMangaType string

const (
	TopMangaTypeManga      TopMangaType = "manga"
	TopMangaTypeNodel      TopMangaType = "novel"
	TopMangaTypeLightnovel TopMangaType = "lightnovel"
	TopMangaTypeOneshot    TopMangaType = "oneshot"
	TopMangaTypeDoujin     TopMangaType = "doujin"
	TopMangaTypeManhwa     TopMangaType = "manhwa"
	TopMangaTypeManhua     TopMangaType = "manhua"
)

type TopMangaFilter string

const (
	TopMangaFilterPublishing   TopMangaFilter = "publishing"
	TopMangaFilterUpcoming     TopMangaFilter = "upcoming"
	TopMangaFilterByPopularity TopMangaFilter = "bypopularity"
	TopMangaFilterFavorite     TopMangaFilter = "favorite"
)

// GetTopManga returns top manga
func GetTopManga(subType TopMangaType, filter TopMangaFilter, page int) (*TopManga, error) {
	res := &TopManga{}
	query := url.Values{}
	query.Set("page", strconv.Itoa(page))
	if subType != "" {
		query.Set("type", string(subType))
	}
	if filter != "" {
		query.Set("filter", string(filter))
	}
	err := urlToStruct(fmt.Sprintf("/top/manga?%s", query.Encode()), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// TopPeople struct
type TopPeople struct {
	Data       []PeopleBase `json:"data"`
	Pagination struct {
		LastVisiblePage int  `json:"last_visible_page"`
		HasNextPage     bool `json:"has_next_page"`
	} `json:"pagination"`
}

// GetTopPeople returns top people
func GetTopPeople(page int) (*TopPeople, error) {
	res := &TopPeople{}
	err := urlToStruct(fmt.Sprintf("/top/people?page=%d", page), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// TopCharacters struct
type TopCharacters struct {
	Data       []CharactersBase `json:"data"`
	Pagination struct {
		LastVisiblePage int  `json:"last_visible_page"`
		HasNextPage     bool `json:"has_next_page"`
	} `json:"pagination"`
}

// GetTopCharacters returns top characters
func GetTopCharacters(page int) (*TopCharacters, error) {
	res := &TopCharacters{}
	err := urlToStruct(fmt.Sprintf("/top/characters?page=%d", page), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// TopReviews struct
type TopReviews struct {
	Data []struct {
		MalId           int       `json:"mal_id"`
		Url             string    `json:"url"`
		Type            string    `json:"type"`
		Votes           int       `json:"votes"`
		Date            time.Time `json:"date"`
		Review          string    `json:"review"`
		EpisodesWatched int       `json:"episodes_watched"`
		Scores          struct {
			Overall   int `json:"overall"`
			Story     int `json:"story"`
			Animation int `json:"animation"`
			Sound     int `json:"sound"`
			Character int `json:"character"`
			Enjoyment int `json:"enjoyment"`
			Art       int `json:"art"`
		} `json:"scores"`
		Entry struct {
			MalId  int    `json:"mal_id"`
			Url    string `json:"url"`
			Images struct {
				Jpg struct {
					ImageUrl      string `json:"image_url"`
					SmallImageUrl string `json:"small_image_url"`
					LargeImageUrl string `json:"large_image_url"`
				} `json:"jpg"`
				Webp struct {
					ImageUrl      string `json:"image_url"`
					SmallImageUrl string `json:"small_image_url"`
					LargeImageUrl string `json:"large_image_url"`
				} `json:"webp"`
			} `json:"images"`
			Title string `json:"title"`
		} `json:"entry"`
		User struct {
			Url      string `json:"url"`
			Username string `json:"username"`
			Images   struct {
				Jpg struct {
					ImageUrl string `json:"image_url"`
				} `json:"jpg"`
				Webp struct {
					ImageUrl string `json:"image_url"`
				} `json:"webp"`
			} `json:"images"`
		} `json:"user"`
		ChaptersRead int `json:"chapters_read"`
	} `json:"data"`
	Pagination struct {
		LastVisiblePage int  `json:"last_visible_page"`
		HasNextPage     bool `json:"has_next_page"`
	} `json:"pagination"`
}

// GetTopReviews returns top reviews
func GetTopReviews(page int) (*TopReviews, error) {
	res := &TopReviews{}
	err := urlToStruct(fmt.Sprintf("/top/reviews?page=%d", page), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
