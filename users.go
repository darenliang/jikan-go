package jikan

import (
	"fmt"
	"net/url"
	"time"
)

// UsersBase struct
type UsersBase struct {
	MalId    int    `json:"mal_id"`
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
	LastOnline time.Time `json:"last_online"`
	Gender     string    `json:"gender"`
	Birthday   time.Time `json:"birthday"`
	Location   string    `json:"location"`
	Joined     time.Time `json:"joined"`
}

// UsersSearch struct
type UsersSearch struct {
	Data []struct {
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
		LastOnline time.Time `json:"last_online"`
	} `json:"data"`
	Pagination struct {
		LastVisiblePage int  `json:"last_visible_page"`
		HasNextPage     bool `json:"has_next_page"`
	} `json:"pagination"`
}

// GetUsersSearch returns users search
func GetUsersSearch(query url.Values) (*UsersSearch, error) {
	res := &UsersSearch{}
	err := urlToStruct(fmt.Sprintf("/users?%s", query.Encode()), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// UserByID struct
type UserByID struct {
	Data UsersBase `json:"data"`
}

// GetUserByID returns user by id
func GetUserByID(id int) (*UserByID, error) {
	res := &UserByID{}
	err := urlToStruct(fmt.Sprintf("/users/userbyid/%d", id), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// UserProfile struct
type UserProfile struct {
	MalId    int    `json:"mal_id"`
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
	LastOnline string `json:"last_online"`
	Gender     string `json:"gender"`
	Birthday   string `json:"birthday"`
	Location   string `json:"location"`
	Joined     string `json:"joined"`
}

// GetUserProfile returns user profile
func GetUserProfile(username string) (*UserProfile, error) {
	res := &UserProfile{}
	err := urlToStruct(fmt.Sprintf("/users/%s",
		url.QueryEscape(username)), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// UserStatistics struct
type UserStatistics struct {
	Data struct {
		Anime struct {
			DaysWatched     float64 `json:"days_watched"`
			MeanScore       float64 `json:"mean_score"`
			Watching        int     `json:"watching"`
			Completed       int     `json:"completed"`
			OnHold          int     `json:"on_hold"`
			Dropped         int     `json:"dropped"`
			PlanToWatch     int     `json:"plan_to_watch"`
			TotalEntries    int     `json:"total_entries"`
			Rewatched       int     `json:"rewatched"`
			EpisodesWatched int     `json:"episodes_watched"`
		} `json:"anime"`
		Manga struct {
			DaysRead     float64 `json:"days_read"`
			MeanScore    float64 `json:"mean_score"`
			Reading      int     `json:"reading"`
			Completed    int     `json:"completed"`
			OnHold       int     `json:"on_hold"`
			Dropped      int     `json:"dropped"`
			PlanToRead   int     `json:"plan_to_read"`
			TotalEntries int     `json:"total_entries"`
			Reread       int     `json:"reread"`
			ChaptersRead int     `json:"chapters_read"`
			VolumesRead  int     `json:"volumes_read"`
		} `json:"manga"`
	} `json:"data"`
}

// GetUserStatistics returns user statistics
func GetUserStatistics(username string) (*UserStatistics, error) {
	res := &UserStatistics{}
	err := urlToStruct(fmt.Sprintf("/users/%s/statistics",
		url.QueryEscape(username)), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// UserFavorites struct
type UserFavorites struct {
	Data struct {
		Anime []struct {
			Type      string `json:"type"`
			StartYear int    `json:"start_year"`
			MalId     int    `json:"mal_id"`
			Url       string `json:"url"`
			Images    struct {
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
		} `json:"anime"`
		Manga []struct {
			Type      string `json:"type"`
			StartYear int    `json:"start_year"`
			MalId     int    `json:"mal_id"`
			Url       string `json:"url"`
			Images    struct {
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
		} `json:"manga"`
		Characters []struct {
			Field1 struct {
				MalId int    `json:"mal_id"`
				Type  string `json:"type"`
				Title string `json:"title"`
				Url   string `json:"url"`
			} `json:""`
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
		} `json:"characters"`
		People []struct {
			Field1 struct {
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
			} `json:""`
		} `json:"people"`
	} `json:"data"`
}

// GetUserFavorites returns user favorites
func GetUserFavorites(username string) (*UserFavorites, error) {
	res := &UserFavorites{}
	err := urlToStruct(fmt.Sprintf("/users/%s/favorites",
		url.QueryEscape(username)), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// UserUpdates struct
type UserUpdates struct {
	Data struct {
		Anime []struct {
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
			Score         float64 `json:"score"`
			Status        string  `json:"status"`
			EpisodesSeen  int     `json:"episodes_seen"`
			EpisodesTotal int     `json:"episodes_total"`
			Date          string  `json:"date"`
		} `json:"anime"`
		Manga []struct {
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
			Score         float64 `json:"score"`
			Status        string  `json:"status"`
			ChaptersRead  int     `json:"chapters_read"`
			ChaptersTotal int     `json:"chapters_total"`
			VolumesRead   int     `json:"volumes_read"`
			VolumesTotal  int     `json:"volumes_total"`
			Date          string  `json:"date"`
		} `json:"manga"`
	} `json:"data"`
}

// GetUserUpdates returns user updates
func GetUserUpdates(username string) (*UserUpdates, error) {
	res := &UserUpdates{}
	err := urlToStruct(fmt.Sprintf("/users/%s/userupdates",
		url.QueryEscape(username)), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// UserAbout struct
type UserAbout struct {
	Data struct {
		About string `json:"about"`
	} `json:"data"`
}

// GetUserAbout returns user about
func GetUserAbout(username string) (*UserAbout, error) {
	res := &UserAbout{}
	err := urlToStruct(fmt.Sprintf("/users/%s/about",
		url.QueryEscape(username)), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// UserHistory struct
type UserHistory struct {
	Data []struct {
		Entry struct {
			MalId int    `json:"mal_id"`
			Type  string `json:"type"`
			Name  string `json:"name"`
			Url   string `json:"url"`
		} `json:"entry"`
		Increment int       `json:"increment"`
		Date      time.Time `json:"date"`
	} `json:"data"`
}

type UserHistoryFilter string

const (
	UserHistoryFilterAnime UserHistoryFilter = "anime"
	UserHistoryFilterManga UserHistoryFilter = "manga"
)

// GetUserHistory returns user history
func GetUserHistory(username string, filter UserHistoryFilter) (*UserHistory, error) {
	res := &UserHistory{}
	req := fmt.Sprintf("/users/%s/history", url.QueryEscape(username))
	if filter != "" {
		req += fmt.Sprintf("?type=%s", filter)
	}
	err := urlToStruct(req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// UserFriends struct
type UserFriends struct {
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
		LastOnline   string `json:"last_online"`
		FriendsSince string `json:"friends_since"`
	} `json:"data"`
	Pagination struct {
		LastVisiblePage int  `json:"last_visible_page"`
		HasNextPage     bool `json:"has_next_page"`
	} `json:"pagination"`
}

// GetUserFriends returns user friends
func GetUserFriends(username string) (*UserFriends, error) {
	res := &UserFriends{}
	err := urlToStruct(fmt.Sprintf("/users/%s/friends",
		url.QueryEscape(username)), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// UserAnimelist struct
// Deprecated: Anime lists will be discontinued from May 1st, 2022.
type UserAnimelist struct {
	Data []struct {
		WatchingStatus  int       `json:"watching_status"`
		Score           float64   `json:"score"`
		EpisodesWatched int       `json:"episodes_watched"`
		Tags            string    `json:"tags"`
		IsRewatching    bool      `json:"is_rewatching"`
		WatchStartDate  time.Time `json:"watch_start_date"`
		WatchEndDate    time.Time `json:"watch_end_date"`
		Days            int       `json:"days"`
		Storage         string    `json:"storage"`
		Priority        string    `json:"priority"`
		Anime           struct {
			MalId  int    `json:"mal_id"`
			Title  string `json:"title"`
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
			Type     string `json:"type"`
			Season   string `json:"season"`
			Year     int    `json:"year"`
			Episodes int    `json:"episodes"`
			Rating   string `json:"rating"`
			Status   string `json:"status"`
			Airing   bool   `json:"airing"`
			Aired    struct {
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
				} `json:"prop"`
				String string `json:"string"`
			} `json:"aired"`
			Studios []struct {
				MalId int    `json:"mal_id"`
				Type  string `json:"type"`
				Name  string `json:"name"`
				Url   string `json:"url"`
			} `json:"studios"`
			Licensors []struct {
				MalId int    `json:"mal_id"`
				Type  string `json:"type"`
				Name  string `json:"name"`
				Url   string `json:"url"`
			} `json:"licensors"`
			Genres []struct {
				MalId int    `json:"mal_id"`
				Type  string `json:"type"`
				Name  string `json:"name"`
				Url   string `json:"url"`
			} `json:"genres"`
			Demographics []struct {
				MalId int    `json:"mal_id"`
				Type  string `json:"type"`
				Name  string `json:"name"`
				Url   string `json:"url"`
			} `json:"demographics"`
		} `json:"anime"`
	} `json:"data"`
}

// GetUserAnimelist returns user animelist
// Deprecated: Anime lists will be discontinued from May 1st, 2022.
func GetUserAnimelist(username string) (*UserAnimelist, error) {
	res := &UserAnimelist{}
	err := urlToStruct(fmt.Sprintf("/users/%s/animelist",
		url.QueryEscape(username)), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// UserMangalist struct
// Deprecated: Manga lists will be discontinued from May 1st, 2022.
type UserMangalist struct {
	Data []struct {
		ReadingStatus int       `json:"reading_status"`
		Score         float64   `json:"score"`
		ChaptersRead  int       `json:"chapters_read"`
		VolumesRead   int       `json:"volumes_read"`
		Tags          string    `json:"tags"`
		IsRereading   bool      `json:"is_rereading"`
		ReadStartDate time.Time `json:"read_start_date"`
		ReadEndDate   time.Time `json:"read_end_date"`
		Days          int       `json:"days"`
		Retail        int       `json:"retail"`
		Priority      string    `json:"priority"`
		Manga         struct {
			MalId  int    `json:"mal_id"`
			Title  string `json:"title"`
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
			Type       string `json:"type"`
			Chapters   int    `json:"chapters"`
			Volumes    int    `json:"volumes"`
			Status     string `json:"status"`
			Publishing bool   `json:"publishing"`
			Published  struct {
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
				} `json:"prop"`
				String string `json:"string"`
			} `json:"published"`
			Magazines []struct {
				MalId int    `json:"mal_id"`
				Type  string `json:"type"`
				Name  string `json:"name"`
				Url   string `json:"url"`
			} `json:"magazines"`
			Genres []struct {
				MalId int    `json:"mal_id"`
				Type  string `json:"type"`
				Name  string `json:"name"`
				Url   string `json:"url"`
			} `json:"genres"`
			Demographics []struct {
				MalId int    `json:"mal_id"`
				Type  string `json:"type"`
				Name  string `json:"name"`
				Url   string `json:"url"`
			} `json:"demographics"`
		} `json:"manga"`
	} `json:"data"`
}

// GetUserMangalist returns user mangalist
// Deprecated: Manga lists will be discontinued from May 1st, 2022.
func GetUserMangalist(username string) (*UserMangalist, error) {
	res := &UserMangalist{}
	err := urlToStruct(fmt.Sprintf("/users/%s/mangalist",
		url.QueryEscape(username)), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// UserReviews struct
type UserReviews struct {
	Data []struct {
		MalId        int       `json:"mal_id"`
		Url          string    `json:"url"`
		Type         string    `json:"type"`
		Votes        int       `json:"votes"`
		Date         time.Time `json:"date"`
		Review       string    `json:"review"`
		ChaptersRead int       `json:"chapters_read"`
		Scores       struct {
			Overall   int `json:"overall"`
			Story     int `json:"story"`
			Art       int `json:"art"`
			Character int `json:"character"`
			Enjoyment int `json:"enjoyment"`
			Animation int `json:"animation"`
			Sound     int `json:"sound"`
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
		EpisodesWatched int `json:"episodes_watched"`
	} `json:"data"`
	Pagination struct {
		LastVisiblePage int  `json:"last_visible_page"`
		HasNextPage     bool `json:"has_next_page"`
	} `json:"pagination"`
}

// GetUserReviews returns user reviews
func GetUserReviews(username string) (*UserReviews, error) {
	res := &UserReviews{}
	err := urlToStruct(fmt.Sprintf("/users/%s/reviews",
		url.QueryEscape(username)), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// UserRecommendations struct
type UserRecommendations struct {
	Data []struct {
		MalId string `json:"mal_id"`
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
		Content string `json:"content"`
		User    struct {
			Data struct {
				Url      string `json:"url"`
				Username string `json:"username"`
			} `json:"data"`
		} `json:"user"`
	} `json:"data"`
	Pagination struct {
		LastVisiblePage int  `json:"last_visible_page"`
		HasNextPage     bool `json:"has_next_page"`
	} `json:"pagination"`
}

// GetUserRecommendations returns user recommendations
func GetUserRecommendations(username string) (*UserRecommendations, error) {
	res := &UserRecommendations{}
	err := urlToStruct(fmt.Sprintf("/users/%s/recommendations",
		url.QueryEscape(username)), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// UserClubs struct
type UserClubs struct {
	Data []struct {
		MalId int    `json:"mal_id"`
		Name  string `json:"name"`
		Url   string `json:"url"`
	} `json:"data"`
	Pagination struct {
		LastVisiblePage int  `json:"last_visible_page"`
		HasNextPage     bool `json:"has_next_page"`
	} `json:"pagination"`
}

// GetUserClubs returns user clubs
func GetUserClubs(username string) (*UserClubs, error) {
	res := &UserClubs{}
	err := urlToStruct(fmt.Sprintf("/users/%s/clubs",
		url.QueryEscape(username)), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
