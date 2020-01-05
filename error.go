package jikan

import "fmt"

// getResultError returns an error with information from result
func getResultError(result map[string]interface{}) error {
	return fmt.Errorf("error %v, %v, %v", result["status"], result["message"], result["error"])
}
