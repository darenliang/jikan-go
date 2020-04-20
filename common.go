package jikan

import "time"

// MalScores struct
type MalScores struct {
	Num1 struct {
		Votes      int     `json:"votes"`
		Percentage float64 `json:"percentage"`
	} `json:"1"`
	Num2 struct {
		Votes      int     `json:"votes"`
		Percentage float64 `json:"percentage"`
	} `json:"2"`
	Num3 struct {
		Votes      int     `json:"votes"`
		Percentage float64 `json:"percentage"`
	} `json:"3"`
	Num4 struct {
		Votes      int     `json:"votes"`
		Percentage float64 `json:"percentage"`
	} `json:"4"`
	Num5 struct {
		Votes      int     `json:"votes"`
		Percentage float64 `json:"percentage"`
	} `json:"5"`
	Num6 struct {
		Votes      int     `json:"votes"`
		Percentage float64 `json:"percentage"`
	} `json:"6"`
	Num7 struct {
		Votes      int     `json:"votes"`
		Percentage float64 `json:"percentage"`
	} `json:"7"`
	Num8 struct {
		Votes      int     `json:"votes"`
		Percentage float64 `json:"percentage"`
	} `json:"8"`
	Num9 struct {
		Votes      int     `json:"votes"`
		Percentage float64 `json:"percentage"`
	} `json:"9"`
	Num10 struct {
		Votes      int     `json:"votes"`
		Percentage float64 `json:"percentage"`
	} `json:"10"`
}

// MalDates struct
type MalDates struct {
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
