package model

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

/*
import (

	"encoding/json"
	"log"
	"net/http"

)

// jsonEncodeResponse encodes the given response object as JSON and writes it to the response writer.

	func jsonEncodeResponse(w http.ResponseWriter, responseObject interface{}) {
		encoder := json.NewEncoder(w)
		if err := encoder.Encode(responseObject); err != nil {
			http.Error(w, "Error encoding JSON response.", http.StatusInternalServerError)
			log.Println("Error encoding JSON:", err)
		}
	}

// jsonDecodeRequest decodes the JSON payload from an HTTP request and returns it as a slice of maps.

	func jsonDecodeRequest(r *http.Response) []map[string]interface{} {
		var response map[string]interface{}

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&response); err != nil {
			log.Printf("Error decoding JSON: %v", err)
			return nil // Return nil to indicate decoding failure
		}

		// Extract the 'results' field, if present, and convert to a slice of maps
		if results, found := response["results"].([]interface{}); found {
			var resultsSlice []map[string]interface{}
			for _, result := range results {
				if resultMap, isValid := result.(map[string]interface{}); isValid {
					resultsSlice = append(resultsSlice, resultMap)
				}
			}
			return resultsSlice
		}

		log.Println("JSON format unexpected: 'results' field missing or invalid")
		return nil // Return nil to indicate unexpected JSON format
	}
*/
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
