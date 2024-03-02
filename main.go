package main

import (
	"assignment1/controller"
	"assignment1/model"
	"log"
	"net/http"
)

func mainPageForUser(w http.ResponseWriter) {

}
func main() {
	controller.HandleUpTime()
	http.HandleFunc(model.Book_Count, controller.BookCounterHandlerMethod)
	http.HandleFunc(model.ReadershipLink, controller.ReaderCountHandlerMethod)
	http.HandleFunc(model.Status_Endpoint, controller.StatusHandlerMethod) //model.Status
	log.Fatal(http.ListenAndServe(":8080", nil))
}
