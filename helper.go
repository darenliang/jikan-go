package jikan

import (
	"encoding/json"
)

func urlToStruct(url string, target interface{}) error {
	resp, err := Client.Get(Endpoint + url)

	if err != nil {
		return err
	}

	err = json.NewDecoder(resp.Body).Decode(target)

	errClose := resp.Body.Close()

	if errClose != nil {
		return errClose
	}

	if err != nil {
		return err
	}

	return nil
}
