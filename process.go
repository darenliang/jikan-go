package jikan

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

// getResultError returns an error with information from result
func getResultError(result map[string]interface{}) error {
	return fmt.Errorf("error %v, %v, %v", result["status"], result["message"], result["error"])
}

// getJsonString returns a json string from a url
func getJsonString(rawURL string) string {
	var httpClient = &http.Client{Timeout: ClientTimeout * time.Second}
	escapedURL := url.QueryEscape(Endpoint + rawURL)
	finalURL, err := url.Parse(escapedURL)
	if err != nil {
		panic(err)
	}
	response, err := httpClient.Get(finalURL.Path)
	if err != nil {
		panic(err)
	}
	str, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	err = response.Body.Close()
	if err != nil {
		panic(err)
	}
	return string(str)
}

// getMapFromUrl returns a map from a json file url
func getMapFromUrl(url string) map[string]interface{} {
	jsonStr := getJsonString(url)
	jsonMap := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonStr), &jsonMap)
	if err != nil {
		panic(err)
	}
	return jsonMap
}

// PrettyPrint pretty prints an interface value
func PrettyPrint(val interface{}) {
	pretty, err := json.MarshalIndent(val, "", "  ")
	if err == nil {
		fmt.Println(string(pretty))
	}
	return
}
