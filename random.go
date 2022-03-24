package jikan

// RandomAnime struct
type RandomAnime struct {
	Data AnimeBase `json:"data"`
}

// GetRandomAnime returns random anime
func GetRandomAnime() (*RandomAnime, error) {
	res := &RandomAnime{}
	err := urlToStruct("/random/anime", res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// RandomManga struct
type RandomManga struct {
	Data MangaBase `json:"data"`
}

// GetRandomManga returns random manga
func GetRandomManga() (*RandomManga, error) {
	res := &RandomManga{}
	err := urlToStruct("/random/manga", res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// RandomCharacters struct
type RandomCharacters struct {
	Data CharactersBase `json:"data"`
}

// GetRandomCharacters returns random characters
func GetRandomCharacters() (*RandomCharacters, error) {
	res := &RandomCharacters{}
	err := urlToStruct("/random/characters", res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// RandomPeople struct
type RandomPeople struct {
	Data PeopleBase `json:"data"`
}

// GetRandomPeople returns random people
func GetRandomPeople() (*RandomPeople, error) {
	res := &RandomPeople{}
	err := urlToStruct("/random/people", res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// RandomUsers struct
type RandomUsers struct {
	Data UsersBase `json:"data"`
}

// GetRandomUsers returns random users
func GetRandomUsers() (*RandomUsers, error) {
	res := &RandomUsers{}
	err := urlToStruct("/random/users", res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
