package restfull_Server

import (
	"assignment1/controller"
	"assignment1/model"
	"log"
	"net/http"
)

func StartApplication() {
	controller.HandleUpTime()
	http.HandleFunc(model.HOME_PAGE, MainPage)
	http.HandleFunc(model.Book_Count, controller.BookCounterHandlerMethod)
	http.HandleFunc(model.ReadershipLink, controller.ReaderCountHandlerMethod)
	http.HandleFunc(model.Status_Endpoint, controller.StatusHandlerMethod) //model.Status
	log.Fatal(http.ListenAndServe(":8080", nil))

}
