package jikan

// WatchEpisodes struct
type WatchEpisodes struct {
	Data []struct {
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
		Episodes []struct {
			MalId   int    `json:"mal_id"`
			Url     string `json:"url"`
			Title   string `json:"title"`
			Premium bool   `json:"premium"`
		} `json:"episodes"`
		RegionLocked bool `json:"region_locked"`
	} `json:"data"`
	Pagination struct {
		LastVisiblePage int  `json:"last_visible_page"`
		HasNextPage     bool `json:"has_next_page"`
	} `json:"pagination"`
}

// GetWatchRecentEpisodes returns watch recent episodes
func GetWatchRecentEpisodes() (*WatchEpisodes, error) {
	res := &WatchEpisodes{}
	err := urlToStruct("/watch/episodes", res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetWatchPopularEpisodes returns watch popular episodes
func GetWatchPopularEpisodes() (*WatchEpisodes, error) {
	res := &WatchEpisodes{}
	err := urlToStruct("/watch/episodes/popular", res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// WatchPromos struct
type WatchPromos struct {
	Data []struct {
		Title string `json:"title"`
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
		Trailer struct {
			YoutubeId string `json:"youtube_id"`
			Url       string `json:"url"`
			EmbedUrl  string `json:"embed_url"`
			Images    struct {
				ImageUrl        string `json:"image_url"`
				SmallImageUrl   string `json:"small_image_url"`
				MediumImageUrl  string `json:"medium_image_url"`
				LargeImageUrl   string `json:"large_image_url"`
				MaximumImageUrl string `json:"maximum_image_url"`
			} `json:"images"`
		} `json:"trailer"`
	} `json:"data"`
	Pagination struct {
		LastVisiblePage int  `json:"last_visible_page"`
		HasNextPage     bool `json:"has_next_page"`
	} `json:"pagination"`
}

// GetWatchRecentPromos returns watch recent promos
func GetWatchRecentPromos() (*WatchPromos, error) {
	res := &WatchPromos{}
	err := urlToStruct("/watch/promos", res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetWatchPopularPromos returns watch popular promos
func GetWatchPopularPromos() (*WatchPromos, error) {
	res := &WatchPromos{}
	err := urlToStruct("/watch/promos/popular", res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
