package model

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// JsonDecodeRequest decodes the JSON payload from an HTTP response into a slice of maps and extracts a 'next' page URL if available.
// It is designed to handle responses where the data is wrapped in a 'results' field and possibly paginated with a 'next' field.
// This function is particularly useful for APIs that return lists of items and include pagination support.
func JsonDecodeRequest(r *http.Response) ([]map[string]interface{}, string) {
	var response map[string]interface{}
	decoder := json.NewDecoder(r.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("Error closing response body: %v", err)
		}
	}(r.Body)
	if err := decoder.Decode(&response); err != nil {
		log.Printf("Error decoding JSON: %v", err)
		return nil, "" // Return nil and empty string to indicate decoding failure
	}

	// Extract the 'results' field
	var resultsSlice []map[string]interface{}
	if results, found := response["results"].([]interface{}); found {
		for _, result := range results {
			if resultMap, isValid := result.(map[string]interface{}); isValid {
				resultsSlice = append(resultsSlice, resultMap)
			}
		}
	} else {
		log.Println("JSON format unexpected: 'results' field missing or invalid")
		return nil, "" // Return nil and empty string for unexpected JSON format
	}

	// Check for a 'next' URL field
	nextPage := ""
	if next, found := response["next"].(string); found && next != "" {
		nextPage = next
	}

	return resultsSlice, nextPage
}

// Decode is a more generic JSON decoding function that directly decodes an HTTP response body into a slice of maps.
// This function is suitable for simpler API responses that directly return a JSON array at the top level.
func Decode(r *http.Response) []map[string]interface{} {
	var results []map[string]interface{}
	decoder := json.NewDecoder(r.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("Error closing response body: %v", err)
		}
	}(r.Body)
	if err := decoder.Decode(&results); err != nil {
		log.Printf("Error decoding JSON (decode): %v", err)
		return nil // Return nil to indicate decoding failure
	}
	return results
}
