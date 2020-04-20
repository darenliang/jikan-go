package jikan

import "time"

// MalScore struct
type MalScore struct {
	Votes      int     `json:"votes"`
	Percentage float64 `json:"percentage"`
}

// MalScores struct
type MalScores struct {
	Num1  MalScore `json:"1"`
	Num2  MalScore `json:"2"`
	Num3  MalScore `json:"3"`
	Num4  MalScore `json:"4"`
	Num5  MalScore `json:"5"`
	Num6  MalScore `json:"6"`
	Num7  MalScore `json:"7"`
	Num8  MalScore `json:"8"`
	Num9  MalScore `json:"9"`
	Num10 MalScore `json:"10"`
}

// MalDay struct
type MalDay struct {
	Day   int `json:"day"`
	Month int `json:"month"`
	Year  int `json:"year"`
}

// MalDates struct
type MalDates struct {
	From time.Time `json:"from"`
	To   time.Time `json:"to"`
	Prop struct {
		From MalDay `json:"from"`
		To   MalDay `json:"to"`
	} `json:"prop"`
	String string `json:"string"`
}

// MalItem struct
type MalItem struct {
	MalID int    `json:"mal_id"`
	Type  string `json:"type"`
	Name  string `json:"name"`
	URL   string `json:"url"`
}

// MalImgItem struct
type MalImgItem struct {
	MalID    int    `json:"mal_id"`
	URL      string `json:"url"`
	ImageURL string `json:"image_url"`
	Name     string `json:"name"`
}

// MalRoleStaff struct
type MalRoleStaff struct {
	MalID    int    `json:"mal_id"`
	Name     string `json:"name"`
	URL      string `json:"url"`
	ImageURL string `json:"image_url"`
	Role     string `json:"role"`
}

// MalAnimeDesc struct
type MalAnimeDesc struct {
	MalID       int       `json:"mal_id"`
	URL         string    `json:"url"`
	Title       string    `json:"title"`
	ImageURL    string    `json:"image_url"`
	Synopsis    string    `json:"synopsis"`
	Type        string    `json:"type"`
	AiringStart time.Time `json:"airing_start"`
	Episodes    int       `json:"episodes"`
	Members     int       `json:"members"`
	Genres      []MalItem `json:"genres"`
	Source      string    `json:"source"`
	Producers   []MalItem `json:"producers"`
	Score       float64   `json:"score"`
	Licensors   []string  `json:"licensors"`
	R18         bool      `json:"r18"`
	Kids        bool      `json:"kids"`
}

// MalMangaDesc struct
type MalMangaDesc struct {
	MalID           int       `json:"mal_id"`
	URL             string    `json:"url"`
	Title           string    `json:"title"`
	ImageURL        string    `json:"image_url"`
	Synopsis        string    `json:"synopsis"`
	Type            string    `json:"type"`
	PublishingStart time.Time `json:"publishing_start"`
	Volumes         int       `json:"volumes"`
	Members         int       `json:"members"`
	Genres          []MalItem `json:"genres"`
	Authors         []MalItem `json:"authors"`
	Score           float64   `json:"score"`
	Serialization   []string  `json:"serialization"`
}
