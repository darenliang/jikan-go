package jikan

import (
	"fmt"
	"time"
)

type UserProfile struct {
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
		DaysRead     int `json:"days_read"`
		MeanScore    int `json:"mean_score"`
		Reading      int `json:"reading"`
		Completed    int `json:"completed"`
		OnHold       int `json:"on_hold"`
		Dropped      int `json:"dropped"`
		PlanToRead   int `json:"plan_to_read"`
		TotalEntries int `json:"total_entries"`
		Reread       int `json:"reread"`
		ChaptersRead int `json:"chapters_read"`
		VolumesRead  int `json:"volumes_read"`
	} `json:"manga_stats"`
	Favorites struct {
		Anime []struct {
			MalID    int    `json:"mal_id"`
			URL      string `json:"url"`
			ImageURL string `json:"image_url"`
			Name     string `json:"name"`
		} `json:"anime"`
		Manga []struct {
			MalID    int    `json:"mal_id"`
			URL      string `json:"url"`
			ImageURL string `json:"image_url"`
			Name     string `json:"name"`
		} `json:"manga"`
		Characters []struct {
			MalID    int    `json:"mal_id"`
			URL      string `json:"url"`
			ImageURL string `json:"image_url"`
			Name     string `json:"name"`
		} `json:"characters"`
		People []struct {
			MalID    int    `json:"mal_id"`
			URL      string `json:"url"`
			ImageURL string `json:"image_url"`
			Name     string `json:"name"`
		} `json:"people"`
	} `json:"favorites"`
	About string `json:"about"`
}

type UserHistory struct {
	History []struct {
		Meta struct {
			MalID int    `json:"mal_id"`
			Type  string `json:"type"`
			Name  string `json:"name"`
			URL   string `json:"url"`
		} `json:"meta"`
		Increment int       `json:"increment"`
		Date      time.Time `json:"date"`
	} `json:"history"`
}

type UserFriends struct {
	Friends []struct {
		URL          string    `json:"url"`
		Username     string    `json:"username"`
		ImageURL     string    `json:"image_url"`
		LastOnline   time.Time `json:"last_online"`
		FriendsSince time.Time `json:"friends_since"`
	} `json:"friends"`
}

type UserAnimelist struct {
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
		Studios         []struct {
			MalID int    `json:"mal_id"`
			Type  string `json:"type"`
			Name  string `json:"name"`
			URL   string `json:"url"`
		} `json:"studios"`
		Licensors []struct {
			MalID int    `json:"mal_id"`
			Type  string `json:"type"`
			Name  string `json:"name"`
			URL   string `json:"url"`
		} `json:"licensors"`
	} `json:"anime"`
}

type UserMangalist struct {
	RequestHash        string `json:"request_hash"`
	RequestCached      bool   `json:"request_cached"`
	RequestCacheExpiry int    `json:"request_cache_expiry"`
	Manga              []struct {
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
		Magazines        []struct {
			MalID int    `json:"mal_id"`
			Type  string `json:"type"`
			Name  string `json:"name"`
			URL   string `json:"url"`
		} `json:"magazines"`
	} `json:"manga"`
}

func GetUserProfile(username string) (UserProfile, error) {
	user := UserProfile{}
	err := ParseResults(fmt.Sprintf("%s/user/%s", Endpoint, username), &user)
	return user, err
}

func GetUserHistory(username string) (UserHistory, error) {
	user := UserHistory{}
	err := ParseResults(fmt.Sprintf("%s/user/%s/history", Endpoint, username), &user)
	return user, err
}

func GetUserFriends(username string) (UserFriends, error) {
	user := UserFriends{}
	err := ParseResults(fmt.Sprintf("%s/user/%s/friends", Endpoint, username), &user)
	return user, err
}

func GetUserAnimelist(username string) (UserAnimelist, error) {
	user := UserAnimelist{}
	err := ParseResults(fmt.Sprintf("%s/user/%s/animelist", Endpoint, username), &user)
	return user, err
}

func GetUserMangalist(username string) (UserMangalist, error) {
	user := UserMangalist{}
	err := ParseResults(fmt.Sprintf("%s/user/%s/mangalist", Endpoint, username), &user)
	return user, err
}
