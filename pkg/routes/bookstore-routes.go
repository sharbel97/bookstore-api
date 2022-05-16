package routes

import (
	"github.com/gorilla/mux"
	"github.com/sharbel97/bookstore-api/pkg/controllers"
)

var RegisterBookstoreRoutes = func (router* mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/book/{bookID}", controllers.GetBooksByID).Methods("GET")
	router.HandleFunc("/book/{bookID}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookID}", controllers.DeleteBook).Methods("DELETE")
}