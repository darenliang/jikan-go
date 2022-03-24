package jikan

import (
	"fmt"
	"net/url"
	"time"
)

type AnimeBase struct {
	MalId   int     `json:"mal_id"`
	Url     string  `json:"url"`
	Images  Images3 `json:"images"`
	Trailer struct {
		YoutubeID string `json:"youtube_id"`
		Url       string `json:"url"`
		EmbedUrl  string `json:"embed_url"`
	} `json:"trailer"`
	Title         string    `json:"title"`
	TitleEnglish  string    `json:"title_english"`
	TitleJapanese string    `json:"title_japanese"`
	TitleSynonyms []string  `json:"title_synonyms"`
	Type          string    `json:"type"`
	Source        string    `json:"source"`
	Episodes      int       `json:"episodes"`
	Status        string    `json:"status"`
	Airing        bool      `json:"airing"`
	Aired         DateRange `json:"aired"`
	Duration      string    `json:"duration"`
	Rating        string    `json:"rating"`
	Score         float64   `json:"score"`
	ScoredBy      int       `json:"scored_by"`
	Rank          int       `json:"rank"`
	Popularity    int       `json:"popularity"`
	Members       int       `json:"members"`
	Favorites     int       `json:"favorites"`
	Synopsis      string    `json:"synopsis"`
	Background    string    `json:"background"`
	Season        string    `json:"season"`
	Year          int       `json:"year"`
	Broadcast     struct {
		Day      string `json:"day"`
		Time     string `json:"time"`
		Timezone string `json:"timezone"`
		String   string `json:"string"`
	} `json:"broadcast"`
	Producers      []MalItem `json:"producers"`
	Licensors      []MalItem `json:"licensors"`
	Studios        []MalItem `json:"studios"`
	Genres         []MalItem `json:"genres"`
	ExplicitGenres []MalItem `json:"explicit_genres"`
	Themes         []MalItem `json:"themes"`
	Demographics   []MalItem `json:"demographics"`
}

// AnimeById struct
type AnimeById struct {
	Data AnimeBase `json:"data"`
}

// GetAnimeById returns anime by id
func GetAnimeById(id int) (*AnimeById, error) {
	res := &AnimeById{}
	err := urlToStruct(fmt.Sprintf("/anime/%d", id), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// AnimeCharacters struct
type AnimeCharacters struct {
	Data []struct {
		Character struct {
			MalId  int     `json:"mal_id"`
			Url    string  `json:"url"`
			Images Images2 `json:"images"`
			Name   string  `json:"name"`
		} `json:"character"`
		Role        string `json:"role"`
		VoiceActors []struct {
			Person struct {
				MalId  int    `json:"mal_id"`
				Url    string `json:"url"`
				Images struct {
					Jpg struct {
						ImageUrl string `json:"image_url"`
					} `json:"jpg"`
				} `json:"images"`
				Name string `json:"name"`
			} `json:"person"`
			Language string `json:"language"`
		} `json:"voice_actors"`
	} `json:"data"`
}

// GetAnimeCharacters returns anime characters
func GetAnimeCharacters(id int) (*AnimeCharacters, error) {
	res := &AnimeCharacters{}
	err := urlToStruct(fmt.Sprintf("/anime/%d/characters", id), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// AnimeStaff struct
type AnimeStaff struct {
	Data []struct {
		Person struct {
			MalId  int    `json:"mal_id"`
			Url    string `json:"url"`
			Images struct {
				Jpg struct {
					ImageUrl string `json:"image_url"`
				} `json:"jpg"`
			} `json:"images"`
			Name string `json:"name"`
		} `json:"person"`
		Positions []string `json:"positions"`
	} `json:"data"`
}

// GetAnimeStaff returns anime staff
func GetAnimeStaff(id int) (*AnimeStaff, error) {
	res := &AnimeStaff{}
	err := urlToStruct(fmt.Sprintf("/anime/%d/staff", id), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// AnimeEpisodes struct
type AnimeEpisodes struct {
	Data []struct {
		MalId         int       `json:"mal_id"`
		Url           string    `json:"url"`
		Title         string    `json:"title"`
		TitleJapanese string    `json:"title_japanese"`
		TitleRomanji  string    `json:"title_romanji"`
		Duration      int       `json:"duration"`
		Aired         time.Time `json:"aired"`
		Filler        bool      `json:"filler"`
		Recap         bool      `json:"recap"`
		ForumUrl      string    `json:"forum_url"`
	} `json:"data"`
	Pagination Pagination `json:"pagination"`
}

// GetAnimeEpisodes returns anime episodes
func GetAnimeEpisodes(id, page int) (*AnimeEpisodes, error) {
	res := &AnimeEpisodes{}
	err := urlToStruct(fmt.Sprintf("/anime/%d/episodes?page=%d", id, page), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// AnimeEpisodeById struct
type AnimeEpisodeById struct {
	Data struct {
		MalId         int       `json:"mal_id"`
		Url           string    `json:"url"`
		Title         string    `json:"title"`
		TitleJapanese string    `json:"title_japanese"`
		TitleRomanji  string    `json:"title_romanji"`
		Duration      int       `json:"duration"`
		Aired         time.Time `json:"aired"`
		Filler        bool      `json:"filler"`
		Recap         bool      `json:"recap"`
		Synopsis      string    `json:"synopsis"`
	} `json:"data"`
}

// GetAnimeEpisodeById returns anime episodes
func GetAnimeEpisodeById(id, episode int) (*AnimeEpisodeById, error) {
	res := &AnimeEpisodeById{}
	err := urlToStruct(fmt.Sprintf("/anime/%d/episodes/%d", id, episode), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// AnimeNews struct
type AnimeNews struct {
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
	Pagination Pagination `json:"pagination"`
}

// GetAnimeNews returns anime news
func GetAnimeNews(id, page int) (*AnimeNews, error) {
	res := &AnimeNews{}
	err := urlToStruct(fmt.Sprintf("/anime/%d/news?page=%d", id, page), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// AnimeForum struct
type AnimeForum struct {
	Data []struct {
		MalId          int       `json:"mal_id"`
		Url            string    `json:"url"`
		Title          string    `json:"title"`
		Date           time.Time `json:"date"`
		AuthorUsername string    `json:"author_username"`
		AuthorUrl      string    `json:"author_url"`
		Comments       int       `json:"comments"`
		LastComment    Comment   `json:"last_comment"`
	} `json:"data"`
}

type AnimeForumFilter string

const (
	AnimeForumFilterAll     AnimeForumFilter = "all"
	AnimeForumFilterEpisode AnimeForumFilter = "episode"
	AnimeForumFilterOther   AnimeForumFilter = "other"
)

// GetAnimeForum returns anime forum
func GetAnimeForum(id int, filter AnimeForumFilter) (*AnimeForum, error) {
	res := &AnimeForum{}
	req := fmt.Sprintf("/anime/%d/forum", id)
	if filter != "" {
		req += fmt.Sprintf("?filter=%s", filter)
	}
	err := urlToStruct(req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// AnimeVideos struct
type AnimeVideos struct {
	Data struct {
		Promos []struct {
			Title   string `json:"title"`
			Trailer struct {
				YoutubeId string `json:"youtube_id"`
				Url       string `json:"url"`
				EmbedUrl  string `json:"embed_url"`
				Images    struct {
					DefaultImageUrl string `json:"default_image_url"`
					SmallImageUrl   string `json:"small_image_url"`
					MediumImageUrl  string `json:"medium_image_url"`
					LargeImageUrl   string `json:"large_image_url"`
					MaximumImageUrl string `json:"maximum_image_url"`
				} `json:"images"`
			} `json:"trailer"`
		} `json:"promos"`
		Episodes []struct {
			MalId   int    `json:"mal_id"`
			Url     string `json:"url"`
			Title   string `json:"title"`
			Episode string `json:"episode"`
			Images  struct {
				Jpg struct {
					ImageUrl string `json:"image_url"`
				} `json:"jpg"`
			} `json:"images"`
		} `json:"episodes"`
	} `json:"data"`
}

// GetAnimeVideos returns anime videos
func GetAnimeVideos(id int) (*AnimeVideos, error) {
	res := &AnimeVideos{}
	err := urlToStruct(fmt.Sprintf("/anime/%d/videos", id), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// AnimePictures struct
type AnimePictures struct {
	Data []struct {
		Images struct {
			Jpg struct {
				ImageUrl string `json:"image_url"`
			} `json:"jpg"`
		} `json:"images"`
	} `json:"data"`
}

// GetAnimePictures returns anime pictures
func GetAnimePictures(id int) (*AnimePictures, error) {
	res := &AnimePictures{}
	err := urlToStruct(fmt.Sprintf("/anime/%d/pictures", id), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// AnimeStatistics struct
type AnimeStatistics struct {
	Data struct {
		Watching    int           `json:"watching"`
		Completed   int           `json:"completed"`
		OnHold      int           `json:"on_hold"`
		Dropped     int           `json:"dropped"`
		PlanToWatch int           `json:"plan_to_watch"`
		Total       int           `json:"total"`
		Scores      []ScoresShort `json:"scores"`
	} `json:"data"`
}

// GetAnimeStatistics returns anime statistics
func GetAnimeStatistics(id int) (*AnimeStatistics, error) {
	res := &AnimeStatistics{}
	err := urlToStruct(fmt.Sprintf("/anime/%d/statistics", id), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// AnimeMoreInfo struct
type AnimeMoreInfo struct {
	Data struct {
		Moreinfo string `json:"moreinfo"`
	} `json:"data"`
}

// GetAnimeMoreInfo returns anime more info
func GetAnimeMoreInfo(id int) (*AnimeMoreInfo, error) {
	res := &AnimeMoreInfo{}
	err := urlToStruct(fmt.Sprintf("/anime/%d/moreinfo", id), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// AnimeRecommendations struct
type AnimeRecommendations struct {
	Data []struct {
		Entry EntryTitle3 `json:"entry"`
		Url   string      `json:"url"`
		Votes int         `json:"votes"`
	} `json:"data"`
}

// GetAnimeRecommendations returns anime recommendations
func GetAnimeRecommendations(id int) (*AnimeRecommendations, error) {
	res := &AnimeRecommendations{}
	err := urlToStruct(fmt.Sprintf("/anime/%d/recommendations", id), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// AnimeUserUpdates struct
type AnimeUserUpdates struct {
	Data []struct {
		User          UserItem `json:"user"`
		Score         float64  `json:"score"`
		Status        string   `json:"status"`
		EpisodesSeen  int      `json:"episodes_seen"`
		EpisodesTotal int      `json:"episodes_total"`
		Date          string   `json:"date"`
	} `json:"data"`
	Pagination Pagination `json:"pagination"`
}

// GetAnimeUserUpdates returns anime user updates
func GetAnimeUserUpdates(id, page int) (*AnimeUserUpdates, error) {
	res := &AnimeUserUpdates{}
	err := urlToStruct(fmt.Sprintf("/anime/%d/userupdates?page=%d", id, page), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// AnimeReviews struct
type AnimeReviews struct {
	Data []struct {
		User            UserItem  `json:"user"`
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
		} `json:"scores"`
	} `json:"data"`
	Pagination Pagination `json:"pagination"`
}

// GetAnimeReviews returns anime reviews
func GetAnimeReviews(id, page int) (*AnimeReviews, error) {
	res := &AnimeReviews{}
	err := urlToStruct(fmt.Sprintf("/anime/%d/reviews?page=%d", id, page), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// AnimeRelations struct
type AnimeRelations struct {
	Data []struct {
		Relation string    `json:"relation"`
		Entry    []MalItem `json:"entry"`
	} `json:"data"`
}

// GetAnimeRelations returns anime relations
func GetAnimeRelations(id int) (*AnimeRelations, error) {
	res := &AnimeRelations{}
	err := urlToStruct(fmt.Sprintf("/anime/%d/relations", id), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// AnimeThemes struct
type AnimeThemes struct {
	Data struct {
		Openings []string `json:"openings"`
		Endings  []string `json:"endings"`
	} `json:"data"`
}

// GetAnimeThemes returns anime themes
func GetAnimeThemes(id int) (*AnimeThemes, error) {
	res := &AnimeThemes{}
	err := urlToStruct(fmt.Sprintf("/anime/%d/themes", id), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// AnimeExternal struct
type AnimeExternal struct {
	Data []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"data"`
}

// GetAnimeExternal returns anime external
func GetAnimeExternal(id int) (*AnimeExternal, error) {
	res := &AnimeExternal{}
	err := urlToStruct(fmt.Sprintf("/anime/%d/external", id), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// AnimeSearch struct
type AnimeSearch struct {
	Data       []AnimeBase `json:"data"`
	Pagination Pagination  `json:"pagination"`
}

// GetAnimeSearch returns anime search
func GetAnimeSearch(query url.Values) (*AnimeSearch, error) {
	res := &AnimeSearch{}
	err := urlToStruct(fmt.Sprintf("/anime?%s", query.Encode()), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
