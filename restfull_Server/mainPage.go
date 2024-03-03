package restfull_Server

import (
	"assignment1/controller"
	"net/http"
)

// MainPage serves as the entry point for the Library Stats API, providing information about the available endpoints.
// It outlines the purpose of each endpoint and how to access them, guiding users on how to retrieve book count information,
// readership statistics, and the status of various services.
// This function leverages the JsonEncodeResponse function from the 'controller' package to ensure the output is formatted
// as a JSON response, making it consistent with the RESTful API design.
func MainPage(r http.ResponseWriter, w *http.Request) {
	// Construct the welcome message with instructions on how to use the API's endpoints
	output := "Welcome to the main page!" + "This is the main page for the Library Stats API. " +
		"Use the endpoint '/librarystats/v1/bookcount/' to access the book count information. " +
		"Use the endpoint '/librarystats/v1/readership/' to access the readership information. " +
		"Use the endpoint '/librarystats/v1/status/' to access the status information. "

	// Encode the welcome message as a JSON response and send it to the requester
	controller.JsonEncodeResponse(r, output)
}
