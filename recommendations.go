package jikan

// RecentRecommendations struct
type RecentRecommendations struct {
	Data []struct {
		MalId   string        `json:"mal_id"`
		Entry   []EntryTitle3 `json:"entry"`
		Content string        `json:"content"`
		User    struct {
			Url      string `json:"url"`
			Username string `json:"username"`
		} `json:"user"`
	} `json:"data"`
	Pagination Pagination `json:"pagination"`
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
