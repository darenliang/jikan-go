package jikan

// RecentRecommendations struct
type RecentRecommendations struct {
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

// GetRecentAnimeRecommendations returns recent anime recommendations
func GetRecentAnimeRecommendations() (*RecentRecommendations, error) {
	res := &RecentRecommendations{}
	err := urlToStruct("/recommendations/anime", res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetRecentMangaRecommendations returns recent manga recommendations
func GetRecentMangaRecommendations() (*RecentRecommendations, error) {
	res := &RecentRecommendations{}
	err := urlToStruct("/recommendations/manga", res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
