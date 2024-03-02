package main

import (
	"assignment1/controller"
	"assignment1/model"
	"log"
	"net/http"
)

func MainPage(r http.ResponseWriter, w *http.Request) {
	output := "Welcome to the main page! In this page, you can get the status of the different API requests. " +
		"Use the endpoint '/librarystats/v1/bookcount/' to access the status information directly through the API. " +
		"For example, making a GET request to '/librarystats/v1/bookcount/?language=de' will provide you with the current status " +
		"of all dependent services to the Dutch language, including their HTTP status codes and uptime." +
		" You can also use two languages at the same time by adding a comma and the next two letter language code such as '/librarystats/v1/bookcount/?language=de,fr' to get the status of the services for both Dutch and French languages."
	controller.JsonEncodeResponse(r, output)

}

func startApplication() {
	controller.HandleUpTime()
	http.HandleFunc(model.HOME_PAGE, MainPage)
	http.HandleFunc(model.Book_Count, controller.BookCounterHandlerMethod)
	http.HandleFunc(model.ReadershipLink, controller.ReaderCountHandlerMethod)
	http.HandleFunc(model.Status_Endpoint, controller.StatusHandlerMethod) //model.Status
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func main() {
	startApplication()
}
