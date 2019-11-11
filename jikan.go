package jikan

// Mal is an interface that gets the specified Jikan json map
type Mal interface {
	Get() (map[string]interface{}, error) // get Jikan data as a map
}