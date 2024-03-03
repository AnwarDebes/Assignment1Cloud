package restfull_Server

import (
	"assignment1/controller"
	"net/http"
)

func MainPage(r http.ResponseWriter, w *http.Request) {
	output := "Welcome to the main page!" + "This is the main page for the Library Stats API. " +
		"Use the endpoint '/librarystats/v1/bookcount/' to access the book count information. " +
		"Use the endpoint '/librarystats/v1/readership/' to access the readership information. " +
		"Use the endpoint '/librarystats/v1/status/' to access the status information. "
	controller.JsonEncodeResponse(r, output)
}
