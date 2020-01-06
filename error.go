package jikan

import "fmt"

// getResultError returns an error with information from jikan response.
// Note that getResultError does not whether or not an error has occured.
func getResultError(result map[string]interface{}) error {
	return fmt.Errorf("error %v, %v, %v", result["status"], result["message"], result["error"])
}
