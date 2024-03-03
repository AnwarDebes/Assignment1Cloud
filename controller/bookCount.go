package controller

import (
	"assignment1/model"
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

// BookCounterHandlerMethod is the entry point for handling requests to the /librarystats/v1/bookcount/ endpoint.
// It determines the appropriate action based on the HTTP method and query parameters provided in the request.
// If a 'language' query parameter is present, it will attempt to process the request using getBookCountFromAPI for GET requests,
// otherwise, it will respond with an error for unsupported HTTP methods. If no 'language' parameter is provided, it directs
// the user to the main page with instructions on how to use the endpoint.
func BookCounterHandlerMethod(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("language") != "" {
		if r.Method == http.MethodGet {
			getBookCountFromAPI(w, r)
		} else {
			http.Error(w, "REST Method '"+r.Method+"' not supported. Only 'GET' is supported.", http.StatusNotImplemented)
		}
	} else {
		mainPageForBookCounter(w)
	}
}

// getBookCountFromAPI handles the retrieval of book counts by language from an external API.
// It extracts the 'language' query parameter, validates it, and then queries the API for each language specified.
// The results are aggregated into a slice of model.BookCount, which is then encoded as a JSON response.
func getBookCountFromAPI(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("language")
	languages := strings.Split(query, ",")

	var bookCounts []model.BookCount // Holds the aggregated results

	for _, language := range languages {
		if len(language) == 0 {
			http.Error(w, "No language code provided. Please provide a language code of two letters.", http.StatusBadRequest)
			return
		} else if len(language) != 2 {
			http.Error(w, "Invalid language code: '"+language+"'. Please provide a language code of two letters.", http.StatusBadRequest)
			return
		}
	}

	for _, lang := range languages {
		results := GetResultFromAPI(w, lang)
		if results == nil { // Handle API errors gracefully
			continue
		}
		authorsCount, bookCount := GetResultOfCount(results)
		bookCounts = append(bookCounts, model.BookCount{
			Language: lang,
			Count:    bookCount,
			Authors:  authorsCount,
			Fraction: float64(bookCount) / float64(model.TOTAL_BOOKS_IN_Library),
		})
	}

	JsonEncodeResponse(w, bookCounts) // Encode and send the response
}

// GetResultOfCount processes the raw API response to count unique authors and books.
// It iterates over each book in the response, counts the number of unique authors (ignoring duplicates and 'Unknown' authors),
// and increments the book count for each valid entry. The result is the total number of unique authors and the total number of books.
func GetResultOfCount(res []map[string]interface{}) (int, int) {
	authorsCount := make(map[string]interface{})
	bookCount := 0

	for _, book := range res {
		if authors, ok := book["authors"].([]interface{}); ok && len(authors) > 0 {
			for _, newAuthor := range authors {
				if author, ok := newAuthor.(map[string]interface{}); ok {
					if authorName, nameOk := author["name"].(string); nameOk && authorName != "Unknown" {
						authorName = strings.TrimSpace(strings.ToLower(authorName))
						if _, exist := authorsCount[authorName]; !exist {
							authorsCount[authorName] = nil
						}
					}
				}
			}
			bookCount++
		}
	}

	return len(authorsCount), bookCount
}

// mainPageForBookCounter provides a default response for the /librarystats/v1/bookcount/ endpoint when no language parameter is provided.
// It gives a brief description of the service and instructions on how to use the endpoint to get book and author counts by language.
func mainPageForBookCounter(w http.ResponseWriter) {
	output := "Here is the book counter main page where you can count the number of the books of a certain language as well as the author." +
		"Here is an explanation in how to use the service:\n" + "First the" + model.Book_Count + " take you to book counter main page." +
		"Secondly the " + model.Book_Language + " provide you the service of choosing language such as count the dutch books and the authors for example:" +
		model.Book_Language_Dutch + " And if you want to count books and authors for several languages at the same time, you can just add that by" +
		" typing a comma and the next two letter language code such like that" + model.Book_Language_Dutch_French + " the URL here provide you a " +
		" the count of books in both dutch and french along with authors"

	JsonEncodeResponse(w, output) // Encode and send the instructional response
}

// JsonEncodeResponse is a utility function that takes any Go data structure and encodes it into JSON format,
// writing the resulting JSON to the provided http.ResponseWriter. It sets the appropriate content type and handles any errors in encoding,
// sending a server error response if encoding fails.
func JsonEncodeResponse(w http.ResponseWriter, responseObject interface{}) {
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(responseObject); err != nil {
		http.Error(w, "Error encoding JSON response.", http.StatusInternalServerError)
		log.Println("Error encoding JSON:", err)
	}
}

// GetResultFromAPI queries an external API for book data based on the specified language code.
// It constructs and sends HTTP GET requests to the API, parsing and aggregating the paginated JSON responses until all pages have been processed.
// The function returns a slice of map[string]interface{} representing the aggregated results, or nil if an error occurs during API communication.
func GetResultFromAPI(w http.ResponseWriter, lang string) []map[string]interface{} {
	var allResults []map[string]interface{}

	nextURL := model.GUTENDEX + lang // Initial URL to start fetching from

	for nextURL != "" {
		req, err := http.NewRequest(http.MethodGet, nextURL, nil)
		if err != nil {
			http.Error(w, "Error occurred during creating the new request to the external API", http.StatusBadRequest)
			return nil
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			http.Error(w, "Error occurred during receiving the response from the external API", http.StatusBadRequest)
			return nil
		}

		pageResults, nextPage := model.JsonDecodeRequest(res) // Process the API response
		allResults = append(allResults, pageResults...)       // Aggregate results
		nextURL = nextPage                                    // Prepare for the next iteration
		res.Body.Close()
	}
	return allResults
}
