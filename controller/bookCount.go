package controller

import (
	"assignment1/model"
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

// BookCounterHandlerMethod handles the /librarystats/v1/bookcount/ endpoint.
// It routes the request based on the method and query parameters.
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

/*
// getBookCountFromAPI handles GET requests to /librarystats/v1/bookcount/ endpoint
// by fetching book counts based on the specified language(s).
func getBookCountFromAPI(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("language")
	languages := strings.Split(query, ",")

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
		authorsCount, bookCount := GetResultOfCount(results)
		book := model.BookCount{
			Language: lang,
			Count:    bookCount,
			Authors:  authorsCount,
			Fraction: float64(bookCount) / float64(model.TOTAL_BOOKS_IN_Library),
		}
		JsonEncodeResponse(w, book)
	}
}
*/

func getBookCountFromAPI(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("language")
	languages := strings.Split(query, ",")

	var bookCounts []model.BookCount // Step 1: Declare a slice to hold the book counts

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
		if results == nil { // Check for nil results to handle errors gracefully
			continue
		}
		authorsCount, bookCount := GetResultOfCount(results)
		book := model.BookCount{
			Language: lang,
			Count:    bookCount, // Ensure this matches the field name in your model
			Authors:  authorsCount,
			Fraction: float64(bookCount) / float64(model.TOTAL_BOOKS_IN_Library),
		}

		bookCounts = append(bookCounts, book) // Step 2: Append the book count object to the slice
	}

	// Step 3: Encode the entire slice as JSON and write it to the response
	JsonEncodeResponse(w, bookCounts)
}

// getResultOfCount parses the API response to count the number of authors and books.
func GetResultOfCount(res []map[string]interface{}) (int, int) {
	authorsCount := make(map[string]interface{})
	bookCount := 0

	for _, book := range res {
		if authors, ok := book["authors"].([]interface{}); ok && len(authors) > 0 {
			for _, newAuthor := range authors {
				if author, ok := newAuthor.(map[string]interface{}); ok {
					if authorName, nameOk := author["name"].(string); nameOk && authorName != "Unknown" { //TODO handle unknown author
						// Trim whitespaces and convert to lowercase
						authorName = strings.TrimSpace(strings.ToLower(authorName))
						// Add to unique authors map
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

// mainPageForBookCounter provides information on how to use the /librarystats/v1/bookcount/ endpoint.
func mainPageForBookCounter(w http.ResponseWriter) {
	output := "Here is the book counter main page where you can count the number of the books of a certain language as well as the author." +
		"Here is an explanation in how to use the service:\n" + "First the" + model.Book_Count + " take you to book counter main page." +
		"Secondly the " + model.Book_Language + " provide you the service of choosing language such as count the dutch books and the authors for example:" +
		model.Book_Language_Dutch + " And if you want to count books and authors for several languages at the same time, you can just add that by" +
		" typing a comma and the next two letter language code such like that" + model.Book_Language_Dutch_French + " the URL here provide you a " +
		" the count of books in both dutch and french along with authors"

	JsonEncodeResponse(w, output)

}

// jsonEncodeResponse encodes the given response object as JSON and writes it to the response writer.
func JsonEncodeResponse(w http.ResponseWriter, responseObject interface{}) {
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(responseObject); err != nil {
		http.Error(w, "Error encoding JSON response.", http.StatusInternalServerError)
		log.Println("Error encoding JSON:", err)
	}
}

func GetResultFromAPI(w http.ResponseWriter, lang string) []map[string]interface{} {
	var allResults []map[string]interface{}

	// Start with the initial URL
	nextURL := model.GUTENDEX + lang

	// Keep looping as long as there's a next URL to go to
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
		// Ensure response body is closed

		// Decode the JSON response
		pageResults, nextPage := model.JsonDecodeRequest(res)
		allResults = append(allResults, pageResults...)

		// Update the nextURL with the nextPage URL
		nextURL = nextPage
		res.Body.Close()
	}
	return allResults
}
