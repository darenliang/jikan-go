package jikan

import "time"

// Pagination struct
type Pagination struct {
	LastVisiblePage int  `json:"last_visible_page"`
	HasNextPage     bool `json:"has_next_page"`
}

// MalItem struct
type MalItem struct {
	MalId int    `json:"mal_id"`
	Type  string `json:"type"`
	Name  string `json:"name"`
	Url   string `json:"url"`
}

// MalItemCount struct
type MalItemCount struct {
	MalId int    `json:"mal_id"`
	Name  string `json:"name"`
	Url   string `json:"url"`
	Count int    `json:"count"`
}

// UserItem struct
type UserItem struct {
	Username string  `json:"username"`
	Url      string  `json:"url"`
	Images   Images1 `json:"images"`
}

// DateRange struct
type DateRange struct {
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
		String string `json:"string"`
	} `json:"prop"`
}

// Images1 struct
type Images1 struct {
	Jpg struct {
		ImageUrl string `json:"image_url"`
	} `json:"jpg"`
	Webp struct {
		ImageUrl string `json:"image_url"`
	} `json:"webp"`
}

// Images2 struct
type Images2 struct {
	Jpg struct {
		ImageUrl      string `json:"image_url"`
		SmallImageUrl string `json:"small_image_url"`
	} `json:"jpg"`
	Webp struct {
		ImageUrl      string `json:"image_url"`
		SmallImageUrl string `json:"small_image_url"`
	} `json:"webp"`
}

// Images3 struct
type Images3 struct {
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
}

// Comment struct
type Comment struct {
	Url            string    `json:"url"`
	AuthorUsername string    `json:"author_username"`
	AuthorUrl      string    `json:"author_url"`
	Date           time.Time `json:"date"`
}

// ScoresShort struct
type ScoresShort struct {
	Score      float64 `json:"score"`
	Votes      int     `json:"votes"`
	Percentage float64 `json:"percentage"`
}

// ScoresAnime struct
type ScoresAnime struct {
	Overall   int `json:"overall"`
	Story     int `json:"story"`
	Animation int `json:"animation"`
	Sound     int `json:"sound"`
	Character int `json:"character"`
	Enjoyment int `json:"enjoyment"`
}

// ScoresManga struct
type ScoresManga struct {
	Overall   int `json:"overall"`
	Story     int `json:"story"`
	Art       int `json:"art"`
	Character int `json:"character"`
	Enjoyment int `json:"enjoyment"`
}

// ScoresLong struct
type ScoresLong struct {
	Overall   int `json:"overall"`
	Story     int `json:"story"`
	Art       int `json:"art"`
	Character int `json:"character"`
	Enjoyment int `json:"enjoyment"`
	Animation int `json:"animation"`
	Sound     int `json:"sound"`
}

// EntryName2 struct
type EntryName2 struct {
	MalId  int     `json:"mal_id"`
	Url    string  `json:"url"`
	Images Images2 `json:"images"`
	Name   string  `json:"title"`
}

// EntryTitle3 struct
type EntryTitle3 struct {
	MalId  int     `json:"mal_id"`
	Url    string  `json:"url"`
	Images Images3 `json:"images"`
	Title  string  `json:"title"`
}
