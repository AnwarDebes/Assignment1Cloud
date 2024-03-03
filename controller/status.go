package controller

import (
	"assignment1/model"
	"net/http"
)

// StatusHandlerMethod routes requests to the status endpoint. It checks the request path and method,
// and either processes a GET request to return the status of various services or redirects to the status main page.
// It ensures that only GET requests are processed for the status endpoint and returns a Method Not Allowed error for other HTTP methods.
func StatusHandlerMethod(w http.ResponseWriter, r *http.Request) {
	// Avoiding index out of bounds by checking the exact match of the URL path with the status endpoint
	if r.URL.Path == model.Status_Endpoint {
		if r.Method == http.MethodGet {
			handleStatusGetRequest(w)
			return
		} else {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	} else {
		statusMainPage(w)
		return
	}
}

// getResponseFromAPIForStatus makes a GET request to the specified API and returns the HTTP status code.
// It handles errors that may occur during the creation and execution of the request,
// responding with appropriate HTTP error statuses.
func getResponseFromAPIForStatus(w http.ResponseWriter, path string) int {
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		http.Error(w, "error during creating request: "+err.Error(), http.StatusInternalServerError)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, "error during sending request: "+err.Error(), http.StatusInternalServerError)
	}
	return res.StatusCode
}

// handleStatusGetRequest handles the GET request for the status endpoint.
// It queries the status of several predefined APIs, aggregates their HTTP status codes,
// and includes additional service information such as version and uptime.
// The result is then encoded as a JSON response.
func handleStatusGetRequest(w http.ResponseWriter) {
	qutendexApiStatus := getResponseFromAPIForStatus(w, model.GUTENDEX)
	languageApiStatus := getResponseFromAPIForStatus(w, model.LANGUAGE_TO_COUNTRY)
	countriesApiStatus := getResponseFromAPIForStatus(w, model.POPULATION_API)

	status := model.Status{
		Qutendexapi:  qutendexApiStatus,
		Languageapi:  languageApiStatus,
		Countriesapi: countriesApiStatus,
		Version:      "v1",
		Uptime:       getUptime().String(), // Assuming getUptime() is a function that returns the system uptime
	}
	JsonEncodeResponse(w, status) // Encode and send the response as JSON
}

// statusMainPage provides a default response for the status endpoint when accessed without the proper path.
// It explains how to use the status endpoint to get the current status of dependent services,
// including instructions for checking the status for specific languages or combinations of languages.
func statusMainPage(w http.ResponseWriter) {
	output := "Welcome to the status main page! In this page, you can get the status of the different API requests. " +
		"Use the endpoint '" + model.Status_Endpoint + "' to access the status information directly through the API. " +
		"For example, making a GET request to '" + model.Book_Language_Dutch + "' will provide you with the current status " +
		"of all dependent services to the Dutch language, including their HTTP status codes and uptime." +
		" You can also use two languages at the same time by adding a comma and the next two letter language code such as '" +
		model.Book_Language_Dutch_French + "' to get the status of the services for both Dutch and French languages."
	JsonEncodeResponse(w, output)
}
