package jikan

import (
	"fmt"
	"time"
)

type UserInfo struct {
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

func GetUserInfo(username string) (UserInfo, error) {
	user := UserInfo{}
	err := ParseResults(fmt.Sprintf("%s/user/%s", Endpoint, username), &user)
	return user, err
}


