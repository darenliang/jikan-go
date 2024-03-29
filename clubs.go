package jikan

import (
	"fmt"
	"net/url"
	"time"
)

// ClubsBase struct
type ClubsBase struct {
	MalId  int    `json:"mal_id"`
	Name   string `json:"name"`
	Url    string `json:"url"`
	Images struct {
		Jpg struct {
			ImageUrl string `json:"image_url"`
		} `json:"jpg"`
	} `json:"images"`
	Members  int       `json:"members"`
	Category string    `json:"category"`
	Created  time.Time `json:"created"`
	Access   string    `json:"access"`
}

// ClubsById struct
type ClubsById struct {
	Data ClubsBase `json:"data"`
}

// GetClubsById returns club by id
func GetClubsById(id int) (*ClubsById, error) {
	res := &ClubsById{}
	err := urlToStruct(fmt.Sprintf("/clubs/%d", id), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type ClubMembers struct {
	Data       []UserItem `json:"data"`
	Pagination Pagination `json:"pagination"`
}

// GetClubMembers returns club members
func GetClubMembers(id, page int) (*ClubMembers, error) {
	res := &ClubMembers{}
	err := urlToStruct(fmt.Sprintf("/clubs/%d/members?page=%d", id, page), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// ClubStaff struct
type ClubStaff struct {
	Data []struct {
		Url      string `json:"url"`
		Username string `json:"username"`
	} `json:"data"`
}

// GetClubStaff returns club staff
func GetClubStaff(id int) (*ClubStaff, error) {
	res := &ClubStaff{}
	err := urlToStruct(fmt.Sprintf("/clubs/%d/staff", id), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// ClubRelations struct
type ClubRelations struct {
	Data struct {
		Anime      []MalItem `json:"anime"`
		Manga      []MalItem `json:"manga"`
		Characters []MalItem `json:"characters"`
	} `json:"data"`
}

// GetClubRelations returns club relations
func GetClubRelations(id int) (*ClubRelations, error) {
	res := &ClubRelations{}
	err := urlToStruct(fmt.Sprintf("/clubs/%d/relations", id), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// ClubsSearch struct
type ClubsSearch struct {
	Data       []ClubsBase `json:"data"`
	Pagination Pagination  `json:"pagination"`
}

// GetClubsSearch returns clubs search
func GetClubsSearch(query url.Values) (*ClubsSearch, error) {
	res := &ClubsSearch{}
	err := urlToStruct(fmt.Sprintf("/clubs?%s", query.Encode()), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
