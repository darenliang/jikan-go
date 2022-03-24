package jikan

import (
	"fmt"
	"net/url"
	"time"
)

// MangaBase struct
type MangaBase struct {
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
	Title         string   `json:"title"`
	TitleEnglish  string   `json:"title_english"`
	TitleJapanese string   `json:"title_japanese"`
	TitleSynonyms []string `json:"title_synonyms"`
	Type          string   `json:"type"`
	Chapters      int      `json:"chapters"`
	Volumes       int      `json:"volumes"`
	Status        string   `json:"status"`
	Publishing    bool     `json:"publishing"`
	Published     struct {
		From time.Time `json:"from"`
		To   time.Time `json:"to"`
		Prop struct {
			From struct {
				Day   int `json:"day"`
				Month int `json:"month"`
				Year  int `json:"year"`
			} `json:"from"`
			To struct {
				Day   int `json:"day"`
				Month int `json:"month"`
				Year  int `json:"year"`
			} `json:"to"`
			String string `json:"string"`
		} `json:"prop"`
	} `json:"published"`
	Score      float64 `json:"score"`
	ScoredBy   int     `json:"scored_by"`
	Rank       int     `json:"rank"`
	Popularity int     `json:"popularity"`
	Members    int     `json:"members"`
	Favorites  int     `json:"favorites"`
	Synopsis   string  `json:"synopsis"`
	Background string  `json:"background"`
	Authors    []struct {
		MalId int    `json:"mal_id"`
		Type  string `json:"type"`
		Name  string `json:"name"`
		Url   string `json:"url"`
	} `json:"authors"`
	Serializations []struct {
		MalId int    `json:"mal_id"`
		Type  string `json:"type"`
		Name  string `json:"name"`
		Url   string `json:"url"`
	} `json:"serializations"`
	Genres []struct {
		MalId int    `json:"mal_id"`
		Type  string `json:"type"`
		Name  string `json:"name"`
		Url   string `json:"url"`
	} `json:"genres"`
	ExplicitGenres []struct {
		MalId int    `json:"mal_id"`
		Type  string `json:"type"`
		Name  string `json:"name"`
		Url   string `json:"url"`
	} `json:"explicit_genres"`
	Themes []struct {
		MalId int    `json:"mal_id"`
		Type  string `json:"type"`
		Name  string `json:"name"`
		Url   string `json:"url"`
	} `json:"themes"`
	Demographics []struct {
		MalId int    `json:"mal_id"`
		Type  string `json:"type"`
		Name  string `json:"name"`
		Url   string `json:"url"`
	} `json:"demographics"`
}

// MangaById struct
type MangaById struct {
	Data MangaBase `json:"data"`
}

// GetMangaById returns manga by id
func GetMangaById(id int) (*MangaById, error) {
	res := &MangaById{}
	err := urlToStruct(fmt.Sprintf("/manga/%d", id), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// MangaCharacters struct
type MangaCharacters struct {
	Data []struct {
		Character struct {
			MalId  int    `json:"mal_id"`
			Url    string `json:"url"`
			Images struct {
				Jpg struct {
					ImageUrl      string `json:"image_url"`
					SmallImageUrl string `json:"small_image_url"`
				} `json:"jpg"`
				Webp struct {
					ImageUrl      string `json:"image_url"`
					SmallImageUrl string `json:"small_image_url"`
				} `json:"webp"`
			} `json:"images"`
			Name string `json:"name"`
		} `json:"character"`
		Role string `json:"role"`
	} `json:"data"`
}

// GetMangaCharacters returns manga characters
func GetMangaCharacters(id int) (*MangaCharacters, error) {
	res := &MangaCharacters{}
	err := urlToStruct(fmt.Sprintf("/manga/%d/characters", id), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// MangaNews struct
type MangaNews struct {
	Data []struct {
		MalId          int       `json:"mal_id"`
		Url            string    `json:"url"`
		Title          string    `json:"title"`
		Date           time.Time `json:"date"`
		AuthorUsername string    `json:"author_username"`
		AuthorUrl      string    `json:"author_url"`
		ForumUrl       string    `json:"forum_url"`
		Images         struct {
			Jpg struct {
				ImageUrl string `json:"image_url"`
			} `json:"jpg"`
		} `json:"images"`
		Comments int    `json:"comments"`
		Excerpt  string `json:"excerpt"`
	} `json:"data"`
	Pagination struct {
		LastVisiblePage int  `json:"last_visible_page"`
		HasNextPage     bool `json:"has_next_page"`
	} `json:"pagination"`
}

// GetMangaNews returns manga news
func GetMangaNews(id, page int) (*MangaNews, error) {
	res := &MangaNews{}
	err := urlToStruct(fmt.Sprintf("/manga/%d/news?page=%d", id, page), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// MangaForum struct
type MangaForum struct {
	Data []struct {
		MalId          int       `json:"mal_id"`
		Url            string    `json:"url"`
		Title          string    `json:"title"`
		Date           time.Time `json:"date"`
		AuthorUsername string    `json:"author_username"`
		AuthorUrl      string    `json:"author_url"`
		Comments       int       `json:"comments"`
		LastComment    struct {
			Url            string    `json:"url"`
			AuthorUsername string    `json:"author_username"`
			AuthorUrl      string    `json:"author_url"`
			Date           time.Time `json:"date"`
		} `json:"last_comment"`
	} `json:"data"`
}

type MangaForumFilter string

const (
	MangaForumFilterAll     MangaForumFilter = "all"
	MangaForumFilterEpisode MangaForumFilter = "episode"
	MangaForumFilterOther   MangaForumFilter = "other"
)

// GetMangaTopics returns manga forum
func GetMangaForum(id int, filter MangaForumFilter) (*MangaForum, error) {
	res := &MangaForum{}
	req := fmt.Sprintf("/manga/%d/forum", id)
	if filter != "" {
		req += fmt.Sprintf("?filter=%s", filter)
	}
	err := urlToStruct(req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// MangaPictures struct
type MangaPictures struct {
	Data []struct {
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
	} `json:"data"`
}

// GetMangaPictures returns manga pictures
func GetMangaPictures(id int) (*MangaPictures, error) {
	res := &MangaPictures{}
	err := urlToStruct(fmt.Sprintf("/manga/%d/pictures", id), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// MangaStatistics struct
type MangaStatistics struct {
	Data struct {
		Reading    int `json:"reading"`
		Completed  int `json:"completed"`
		OnHold     int `json:"on_hold"`
		Dropped    int `json:"dropped"`
		PlanToRead int `json:"plan_to_read"`
		Total      int `json:"total"`
		Scores     []struct {
			Score      float64 `json:"score"`
			Votes      int     `json:"votes"`
			Percentage float64 `json:"percentage"`
		} `json:"scores"`
	} `json:"data"`
}

// GetMangaStatistics returns manga statistics
func GetMangaStatistics(id int) (*MangaStatistics, error) {
	res := &MangaStatistics{}
	err := urlToStruct(fmt.Sprintf("/manga/%d/statistics", id), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type MangaMoreInfo struct {
	Data struct {
		Moreinfo string `json:"moreinfo"`
	} `json:"data"`
}

// GetMangaMoreInfo returns manga more info
func GetMangaMoreInfo(id int) (*MangaMoreInfo, error) {
	res := &MangaMoreInfo{}
	err := urlToStruct(fmt.Sprintf("/manga/%d/moreinfo", id), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// MangaRecommendations struct
type MangaRecommendations struct {
	Data []struct {
		Entry []struct {
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
		Url   string `json:"url"`
		Votes int    `json:"votes"`
	} `json:"data"`
}

// GetMangaRecommendations returns manga recommendations
func GetMangaRecommendations(id int) (*MangaRecommendations, error) {
	res := &MangaRecommendations{}
	err := urlToStruct(fmt.Sprintf("/manga/%d/recommendations", id), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// MangaUserUpdates struct
type MangaUserUpdates struct {
	Data []struct {
		User struct {
			Username string `json:"username"`
			Url      string `json:"url"`
			Images   struct {
				Jpg struct {
					ImageUrl string `json:"image_url"`
				} `json:"jpg"`
				Webp struct {
					ImageUrl string `json:"image_url"`
				} `json:"webp"`
			} `json:"images"`
		} `json:"user"`
		Score         float64   `json:"score"`
		Status        string    `json:"status"`
		VolumesRead   int       `json:"volumes_read"`
		VolumesTotal  int       `json:"volumes_total"`
		ChaptersRead  int       `json:"chapters_read"`
		ChaptersTotal int       `json:"chapters_total"`
		Date          time.Time `json:"date"`
	} `json:"data"`
	Pagination struct {
		LastVisiblePage int  `json:"last_visible_page"`
		HasNextPage     bool `json:"has_next_page"`
	} `json:"pagination"`
}

// GetMangaUserUpdates returns manga user updates
func GetMangaUserUpdates(id, page int) (*MangaUserUpdates, error) {
	res := &MangaUserUpdates{}
	err := urlToStruct(fmt.Sprintf("/manga/%d/userupdates?page=%d", id, page), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// MangaReviews struct
type MangaReviews struct {
	Data []struct {
		User struct {
			Username string `json:"username"`
			Url      string `json:"url"`
			Images   struct {
				Jpg struct {
					ImageUrl string `json:"image_url"`
				} `json:"jpg"`
				Webp struct {
					ImageUrl string `json:"image_url"`
				} `json:"webp"`
			} `json:"images"`
		} `json:"user"`
		MalId        int       `json:"mal_id"`
		Url          string    `json:"url"`
		Type         string    `json:"type"`
		Votes        int       `json:"votes"`
		Date         time.Time `json:"date"`
		ChaptersRead int       `json:"chapters_read"`
		Review       string    `json:"review"`
		Scores       struct {
			Overall   int `json:"overall"`
			Story     int `json:"story"`
			Art       int `json:"art"`
			Character int `json:"character"`
			Enjoyment int `json:"enjoyment"`
		} `json:"scores"`
	} `json:"data"`
	Pagination struct {
		LastVisiblePage int  `json:"last_visible_page"`
		HasNextPage     bool `json:"has_next_page"`
	} `json:"pagination"`
}

// GetMangaReviews returns manga reviews
func GetMangaReviews(id, page int) (*MangaReviews, error) {
	res := &MangaReviews{}
	err := urlToStruct(fmt.Sprintf("/manga/%d/reviews?page=%d", id, page), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// MangaRelations struct
type MangaRelations struct {
	Data []struct {
		Relation string `json:"relation"`
		Entry    []struct {
			MalId int    `json:"mal_id"`
			Type  string `json:"type"`
			Name  string `json:"name"`
			Url   string `json:"url"`
		} `json:"entry"`
	} `json:"data"`
}

// GetMangaRelations returns manga relations
func GetMangaRelations(id int) (*MangaRelations, error) {
	res := &MangaRelations{}
	err := urlToStruct(fmt.Sprintf("/manga/%d/relations", id), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// MangaExternal struct
type MangaExternal struct {
	Data []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"data"`
}

// GetMangaExternal returns manga external
func GetMangaExternal(id int) (*MangaExternal, error) {
	res := &MangaExternal{}
	err := urlToStruct(fmt.Sprintf("/manga/%d/external", id), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// MangaSearch struct
type MangaSearch struct {
	Data       []MangaBase `json:"data"`
	Pagination struct {
		LastVisiblePage int  `json:"last_visible_page"`
		HasNextPage     bool `json:"has_next_page"`
	} `json:"pagination"`
}

// GetMangaSearch returns manga search
func GetMangaSearch(query url.Values) (*MangaSearch, error) {
	res := &MangaSearch{}
	err := urlToStruct(fmt.Sprintf("/manga?%s", query.Encode()), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
