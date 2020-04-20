package jikan

import (
	"fmt"
	"net/url"
	"time"
)

// User struct for the /user endpoint
type User struct {
	UserID     int       `json:"user_id"`
	Username   string    `json:"username"`
	URL        string    `json:"url"`
	ImageURL   string    `json:"image_url"`
	LastOnline time.Time `json:"last_online"`
	Gender     string    `json:"gender"`
	Birthday   time.Time `json:"birthday"`
	Location   string    `json:"location"`
	Joined     time.Time `json:"joined"`
	AnimeStats struct {
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
	} `json:"anime_stats"`
	MangaStats struct {
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
	} `json:"manga_stats"`
	Favorites struct {
		Anime      []MalImgItem `json:"anime"`
		Manga      []MalImgItem `json:"manga"`
		Characters []MalImgItem `json:"characters"`
		People     []MalImgItem `json:"people"`
	} `json:"favorites"`
	About string `json:"about"`
}

// UserHistory struct for the /user/history endpoint
type UserHistory struct {
	History []struct {
		Meta      MalItem   `json:"meta"`
		Increment int       `json:"increment"`
		Date      time.Time `json:"date"`
	} `json:"history"`
}

// UserFriends struct for the /user/friends endpoint
type UserFriends struct {
	Friends []struct {
		URL          string    `json:"url"`
		Username     string    `json:"username"`
		ImageURL     string    `json:"image_url"`
		LastOnline   time.Time `json:"last_online"`
		FriendsSince time.Time `json:"friends_since"`
	} `json:"friends"`
}

// UserAnimeList struct for the /user/animelist endpoint
type UserAnimeList struct {
	Anime []struct {
		MalID           int       `json:"mal_id"`
		Title           string    `json:"title"`
		VideoURL        string    `json:"video_url"`
		URL             string    `json:"url"`
		ImageURL        string    `json:"image_url"`
		Type            string    `json:"type"`
		WatchingStatus  int       `json:"watching_status"`
		Score           int       `json:"score"`
		WatchedEpisodes int       `json:"watched_episodes"`
		TotalEpisodes   int       `json:"total_episodes"`
		AiringStatus    int       `json:"airing_status"`
		SeasonName      string    `json:"season_name"`
		SeasonYear      int       `json:"season_year"`
		HasEpisodeVideo bool      `json:"has_episode_video"`
		HasPromoVideo   bool      `json:"has_promo_video"`
		HasVideo        bool      `json:"has_video"`
		IsRewatching    bool      `json:"is_rewatching"`
		Tags            []string  `json:"tags"`
		Rating          string    `json:"rating"`
		StartDate       time.Time `json:"start_date"`
		EndDate         time.Time `json:"end_date"`
		WatchStartDate  time.Time `json:"watch_start_date"`
		WatchEndDate    time.Time `json:"watch_end_date"`
		Days            int       `json:"days"`
		Storage         string    `json:"storage"`
		Priority        string    `json:"priority"`
		AddedToList     bool      `json:"added_to_list"`
		Studios         []MalItem `json:"studios"`
		Licensors       []MalItem `json:"licensors"`
	} `json:"anime"`
}

// UserMangaList struct for the /mangalist endpoint
type UserMangaList struct {
	Manga []struct {
		MalID            int       `json:"mal_id"`
		Title            string    `json:"title"`
		URL              string    `json:"url"`
		ImageURL         string    `json:"image_url"`
		Type             string    `json:"type"`
		ReadingStatus    int       `json:"reading_status"`
		Score            int       `json:"score"`
		ReadChapters     int       `json:"read_chapters"`
		ReadVolumes      int       `json:"read_volumes"`
		TotalChapters    int       `json:"total_chapters"`
		TotalVolumes     int       `json:"total_volumes"`
		PublishingStatus int       `json:"publishing_status"`
		IsRereading      bool      `json:"is_rereading"`
		Tags             []string  `json:"tags"`
		StartDate        time.Time `json:"start_date"`
		EndDate          time.Time `json:"end_date"`
		ReadStartDate    time.Time `json:"read_start_date"`
		ReadEndDate      time.Time `json:"read_end_date"`
		Days             int       `json:"days"`
		Retail           string    `json:"retail"`
		Priority         string    `json:"priority"`
		AddedToList      bool      `json:"added_to_list"`
		Magazines        []MalItem `json:"magazines"`
	} `json:"manga"`
}

// GetUser returns user
func GetUser(username string) (*User, error) {
	res := &User{}

	err := urlToStruct(fmt.Sprintf("/user/%s",
		url.QueryEscape(username)), res)

	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetUserHistory returns user history
func GetUserHistory(username, opt string) (*UserHistory, error) {
	res := &UserHistory{}

	err := urlToStruct(fmt.Sprintf("/user/%s/history/%s",
		url.QueryEscape(username), url.QueryEscape(opt)), res)

	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetUserFriends returns user friends
func GetUserFriends(username string, page int) (*UserFriends, error) {
	res := &UserFriends{}

	err := urlToStruct(fmt.Sprintf("/user/%s/friends/%d",
		url.QueryEscape(username), page), res)

	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetUserAnimeList returns user anime list
func GetUserAnimeList(username, opt string, page int) (*UserAnimeList, error) {
	res := &UserAnimeList{}

	err := urlToStruct(fmt.Sprintf("/user/%s/animelist/%s/%d",
		url.QueryEscape(username), url.QueryEscape(opt), page), res)

	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetUserMangaList returns user manga list
func GetUserMangaList(username, opt string, page int) (*UserMangaList, error) {
	res := &UserMangaList{}

	err := urlToStruct(fmt.Sprintf("/user/%s/mangalist/%s/%d",
		url.QueryEscape(username), url.QueryEscape(opt), page), res)

	if err != nil {
		return nil, err
	}

	return res, nil
}
