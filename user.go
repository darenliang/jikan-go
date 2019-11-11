package jikan

import (
	"fmt"
	"strings"
)

// User struct defines a user
type User struct {
	Username string // MyAnimeList username
	Request  int    // Request type
	Data     string // Specify data type

	// Use only when choosing history
	HistoryType string // Specify history type

	// Use only when choosing friends
	FriendsPage int // Specify friends page

	AnimeListFilter string // Anime List Filter
	MangaListFilter string // Manga List Filter

	// Use only when choosing animelist or mangalist (Only need to choose what you need)
	Search string // Search query
	Page   int    // Page number
	Sort   string // Sort

	// Only for AnimeListFilter
	OrderBy      string // Order by
	OrderBy2     string // Second order by
	AiredFrom    string // Aired from
	AiredTo      string // Aired to
	Producer     int    // MyAnimeList Producer ID
	Year         int    // Year
	Season       string // Season
	AiringStatus string // Airing status

	// Only for MangaListFilter
	PublishedFrom    string // Published from
	PublishedTo      string // Published to
	Magazine         int    // MyAnimeList Magazine ID
	PublishingStatus string // Publishing status
}

// Get returns a map of a user as specified in the User struct
// Calls responses through the /user/ endpoint
func (user User) Get() (map[string]interface{}, error) {
	var result map[string]interface{}
	var err error
	var query, advancedQuery strings.Builder
	query.WriteString(fmt.Sprintf("/user/%v/%v", user.Username, user.Request))
	if user.Data != "" {
		query.WriteString(fmt.Sprintf("/%v", user.Data))
		if user.Data == "history" && user.HistoryType != "" {
			query.WriteString(fmt.Sprintf("/%v", user.HistoryType))
		} else if user.Data == "friends" && user.FriendsPage != 0 {
			query.WriteString(fmt.Sprintf("/%v", user.FriendsPage))
		} else if user.Data == "animelist" && user.AnimeListFilter != "" {
			query.WriteString(fmt.Sprintf("/%v", user.AnimeListFilter))
			if user.Search != "" {
				advancedQuery.WriteString(fmt.Sprintf("search=%v&", user.Search))
			}
			if user.Page != 0 {
				advancedQuery.WriteString(fmt.Sprintf("page=%v&", user.Page))
			}
			if user.Sort != "" {
				advancedQuery.WriteString(fmt.Sprintf("sort=%v&", user.Sort))
			}
			if user.OrderBy != "" {
				advancedQuery.WriteString(fmt.Sprintf("order_by=%v&", user.OrderBy))
			}
			if user.OrderBy2 != "" {
				advancedQuery.WriteString(fmt.Sprintf("order_by2=%v&", user.OrderBy2))
			}
			if user.AiredFrom != "" {
				advancedQuery.WriteString(fmt.Sprintf("aired_from=%v&", user.AiredFrom))
			}
			if user.AiredTo != "" {
				advancedQuery.WriteString(fmt.Sprintf("aired_to=%v&", user.AiredTo))
			}
			if user.Producer != 0 {
				advancedQuery.WriteString(fmt.Sprintf("producer=%v&", user.Producer))
			}
			if user.Year != 0 {
				advancedQuery.WriteString(fmt.Sprintf("year=%v&", user.Year))
			}
			if user.Season != "" {
				advancedQuery.WriteString(fmt.Sprintf("season=%v&", user.Season))
			}
			if user.AiringStatus != "" {
				advancedQuery.WriteString(fmt.Sprintf("airing_status=%v&", user.AiringStatus))
			}
		} else if user.Data == "mangalist" && user.MangaListFilter != "" {
			query.WriteString(fmt.Sprintf("/%v", user.MangaListFilter))
			if user.Search != "" {
				advancedQuery.WriteString(fmt.Sprintf("search=%v&", user.Search))
			}
			if user.Page != 0 {
				advancedQuery.WriteString(fmt.Sprintf("page=%v&", user.Page))
			}
			if user.Sort != "" {
				advancedQuery.WriteString(fmt.Sprintf("sort=%v&", user.Sort))
			}
			if user.OrderBy != "" {
				advancedQuery.WriteString(fmt.Sprintf("order_by=%v&", user.OrderBy))
			}
			if user.OrderBy2 != "" {
				advancedQuery.WriteString(fmt.Sprintf("order_by2=%v&", user.OrderBy2))
			}
			if user.PublishedFrom != "" {
				advancedQuery.WriteString(fmt.Sprintf("published_from=%v&", user.PublishedFrom))
			}
			if user.PublishedTo != "" {
				advancedQuery.WriteString(fmt.Sprintf("published_to=%v&", user.PublishedTo))
			}
			if user.Magazine != 0 {
				advancedQuery.WriteString(fmt.Sprintf("magazine=%v&", user.Magazine))
			}
			if user.PublishingStatus != "" {
				advancedQuery.WriteString(fmt.Sprintf("airing_status=%v&", user.PublishingStatus))
			}
		}
	}
	result, err = getMapFromUrl(query.String()), nil
	if _, ok := result["error"]; ok {
		result, err = nil, getResultError(result)
	}
	return result, err
}
