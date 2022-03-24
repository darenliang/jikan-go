package jikan

// RecentAnimeReviews struct
type RecentAnimeReviews struct {
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
		Anime struct {
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
		} `json:"anime"`
		MalId           int    `json:"mal_id"`
		Url             string `json:"url"`
		Type            string `json:"type"`
		Votes           int    `json:"votes"`
		Date            string `json:"date"`
		Review          string `json:"review"`
		EpisodesWatched int    `json:"episodes_watched"`
		Scores          struct {
			Overall   int `json:"overall"`
			Story     int `json:"story"`
			Animation int `json:"animation"`
			Sound     int `json:"sound"`
			Character int `json:"character"`
			Enjoyment int `json:"enjoyment"`
		} `json:"scores"`
	} `json:"data"`
	Pagination struct {
		LastVisiblePage int  `json:"last_visible_page"`
		HasNextPage     bool `json:"has_next_page"`
	} `json:"pagination"`
}

// GetRecentAnimeReviews returns recent anime reviews
func GetRecentAnimeReviews() (*RecentAnimeReviews, error) {
	res := &RecentAnimeReviews{}
	err := urlToStruct("/reviews/anime", res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// RecentMangaReviews struct
type RecentMangaReviews struct {
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
		Manga struct {
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
		} `json:"manga"`
		MalId        int    `json:"mal_id"`
		Url          string `json:"url"`
		Type         string `json:"type"`
		Votes        int    `json:"votes"`
		Date         string `json:"date"`
		ChaptersRead int    `json:"chapters_read"`
		Review       string `json:"review"`
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

// GetRecentMangaReviews returns recent manga reviews
func GetRecentMangaReviews() (*RecentMangaReviews, error) {
	res := &RecentMangaReviews{}
	err := urlToStruct("/reviews/manga", res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
