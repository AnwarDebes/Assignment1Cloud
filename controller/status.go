package controller

import (
	"assignment1/model"
	"net/http"
)

/*
type Status struct {
	Qutendexapi  int    `json:"gutendexapi"`
	Languageapi  int    `json:"languageapi"`
	Countriesapi int    `json:"countriesapi"`
	Version      string `json:"version"`
	Uptime       string `json:"uptime"`
}

*/

func StatusHandlerMethod(w http.ResponseWriter, r *http.Request) {

	// avoiding index out of bounds
	if r.URL.Path == model.Status_Endpoint { //model.Status
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

// handle the request and response for status endpoint
func handleStatusGetRequest(w http.ResponseWriter) {
	qutendexApiStatus := getResponseFromAPIForStatus(w, model.GUTENDEX)
	languageApiStatus := getResponseFromAPIForStatus(w, model.LANGUAGE_TO_COUNTRY)
	countriesApiStatus := getResponseFromAPIForStatus(w, model.POPULATION_API)

	status := model.Status{Qutendexapi: qutendexApiStatus, Languageapi: languageApiStatus, Countriesapi: countriesApiStatus, Version: "v1", Uptime: getUptime().String()}
	// Encode JSON
	JsonEncodeResponse(w, status)
}

func statusMainPage(w http.ResponseWriter) {
	output := "Welcome to the status main page! In this page, you can get the status of the different API requests. " +
		"Use the endpoint '" + model.Status_Endpoint + "' to access the status information directly through the API. " +
		"For example, making a GET request to '" + model.Book_Language_Dutch + "' will provide you with the current status " +
		"of all dependent services to the Dutch language, including their HTTP status codes and uptime." +
		" You can also use two languages at the same time by adding a comma and the next two letter language code such as '" +
		model.Book_Language_Dutch_French + "' to get the status of the services for both Dutch and French languages."
	JsonEncodeResponse(w, output)

}
