package controller

import (
	"assignment1/model"
	"net/http"
	"strconv"
	"strings"
)

// mainPageForReadership provides a default response for the /librarystats/v1/readership/ endpoint when no specific language code is provided.
// It gives an introduction to the readership statistics feature, explaining how to use the endpoint to get information on readership
// by language, including instructions for limiting the number of results returned.
func mainPageForReadership(w http.ResponseWriter, r *http.Request) {
	output := "Welcome to the Readership Main Page! Here, you can discover the readership statistics for different languages, " +
		"including the country name, country ISO code (a short abbreviation for the country), and the total number of books and authors. " +
		"To start, navigate to " + model.ReadershipLink + " - you're already here! To find readership details for a specific language, enter the language's ISO code in the URL." +
		"For example, " + model.ReadershipExampleRequest + " provides readership information for Norwegian, including details about the country, ISO code, books, and authors. " +
		"Want to narrow down your search? You can limit the number of results by adding a 'limit' parameter in the URL. " +
		"For example, " + model.ReadershipExampleRequestWithLimit + " will show you the top 5 readership statistics for Norwegian."
	JsonEncodeResponse(w, output)
}

// ReaderCountHandlerMethod is responsible for handling requests to the /librarystats/v1/readership/ endpoint.
// It parses the request URL to extract the language code and optional limit for the number of results.
// Depending on the request method (GET is supported), it either retrieves readership statistics from the API or returns an error for unsupported methods.
func ReaderCountHandlerMethod(w http.ResponseWriter, r *http.Request) {
	path := strings.Trim(r.URL.Path, "/")
	pathParts := strings.Split(path, "/")

	if len(pathParts) < 4 {
		mainPageForReadership(w, r)
		return
	}

	languageCode := pathParts[3]

	// Parse query parameters for the optional "limit" parameter
	queryParams := r.URL.Query()
	limitParam := queryParams.Get("limit") // Get the limit parameter if it exists

	var limit int = -1 // Default no limit, indicated by -1
	var err error
	if limitParam != "" {
		limit, err = strconv.Atoi(limitParam) // Convert limit to int
		if err != nil {
			http.Error(w, "Invalid limit parameter", http.StatusBadRequest)
			return
		}
	}

	if r.Method == http.MethodGet {
		if languageCode == "" {
			mainPageForReadership(w, r)
			return
		}
		getReadershipFromAPI(w, languageCode, limit)
	} else {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

// getReadershipFromAPI queries external APIs to gather and compile readership statistics for a specified language code and optional result limit.
// It first retrieves country data associated with the language, then fetches population data for each country, and finally compiles
// the readership statistics into a slice of model.Readership objects, which is then encoded as a JSON response.
func getReadershipFromAPI(w http.ResponseWriter, languageCode string, limit int) {
	res := getResponseFromAPI(w, model.LANGUAGE_TO_COUNTRY+"language2countries/"+languageCode)

	// If limit is set and less than the length of res, reslice res to limit's length
	if limit > 0 && len(res) > limit {
		res = res[:limit]
	}

	counts := GetResultFromAPI(w, languageCode)
	books, authors := GetResultOfCount(counts)

	// Initialize a slice to hold all readershipData objects
	var readershipDataList []model.Readership

	for _, country := range res {
		nameOfCountry, _ := country["Official_Name"].(string)
		isoOfCountry, _ := country["ISO3166_1_Alpha_2"].(string)
		populationResponse := getResponseFromAPI(w, model.POPULATION_API+"v3.1/name/"+nameOfCountry)
		population, _ := populationResponse[0]["population"].(float64)
		readershipData := model.Readership{
			Country:    nameOfCountry,
			IsoCode:    isoOfCountry,
			Books:      books,
			Authors:    authors,
			Readership: population,
		}
		// Append the data to the slice
		readershipDataList = append(readershipDataList, readershipData)
	}

	// Encode the entire slice as a JSON response
	JsonEncodeResponse(w, readershipDataList)
}

// getResponseFromAPI sends a GET request to the specified external API path and decodes the JSON response into a slice of maps.
// It handles errors associated with request creation and execution, responding with appropriate HTTP error statuses.
func getResponseFromAPI(w http.ResponseWriter, path string) []map[string]interface{} {
	req, err := http.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		http.Error(w, "Error occurred during creating the new request to the external API", http.StatusBadRequest)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, "Error occurred during receiving the response from the external API", http.StatusBadRequest)
	}
	// Decode the JSON response
	results := model.Decode(res)
	return results
}
