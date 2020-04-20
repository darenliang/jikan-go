package jikan

import (
	"encoding/json"
	"net/url"
	"strings"
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

func processArray(arr []string, process func(string) string) (ret []string) {
	for _, s := range arr {
		ret = append(ret, process(s))
	}
	return
}

func escapeOption(opt string) string {
	idx := strings.Index(opt, "=")
	if idx < 0 || idx+1 == len(opt) {
		return opt
	}
	return opt[:idx+1] + url.QueryEscape(opt[idx+1:])
}
