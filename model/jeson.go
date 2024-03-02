package model

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func JsonDecodeRequest(r *http.Response) ([]map[string]interface{}, string) {
	var response map[string]interface{}
	decoder := json.NewDecoder(r.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("Error closing response bode. %v", err)
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

func Decode(r *http.Response) []map[string]interface{} {
	var results []map[string]interface{}
	decoder := json.NewDecoder(r.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("Error closing response bode. %v", err)
		}
	}(r.Body)
	if err := decoder.Decode(&results); err != nil {
		log.Printf("Error decoding JSON (decode): %v", err)
		return nil // Return nil and empty string to indicate decoding failure
	}
	return results
}
