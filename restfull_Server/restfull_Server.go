package restfull_Server

import (
	"assignment1/controller"
	"assignment1/model"
	"log"
	"net/http"
)

// StartApplication initializes the application by setting up HTTP routes and starting the server.
// It begins by calling HandleUpTime to start tracking the uptime of the service.
// Then, it registers several HTTP handlers for different paths using http.HandleFunc.
// Each path is associated with a specific function in the controller package that handles requests for that route.
// Finally, it starts an HTTP server listening on port 8080 and logs any errors that occur.
func StartApplication() {
	controller.HandleUpTime() // Start tracking uptime as soon as the application starts

	// Setup HTTP routes
	http.HandleFunc(model.HOME_PAGE, MainPage)                                 // Handler for the main page
	http.HandleFunc(model.Book_Count, controller.BookCounterHandlerMethod)     // Handler for book count requests
	http.HandleFunc(model.ReadershipLink, controller.ReaderCountHandlerMethod) // Handler for readership statistics requests
	http.HandleFunc(model.Status_Endpoint, controller.StatusHandlerMethod)     // Handler for status information requests

	// Start the HTTP server on port 8080 and log any errors
	log.Fatal(http.ListenAndServe(":8080", nil))
}
