package routes

import (
	"github.com/gorilla/mux"
	"full-crud-tutorial/src/controller"
)

var RegisterBookStoreRoutes = func (router *mux.Router)  {
	router.HandleFunc("/book",controller.CreateBook ).Methods("POST")
	router.HandleFunc("/books",controller.GetBooks).Methods("GET")
	router.HandleFunc("/books/{id}",controller.GetBookById).Methods("GET")
	router.HandleFunc("/books/{id}",controller.UpdateBooks).Methods("PUT")
	router.HandleFunc("/books/{id}",controller.DeleteBooks).Methods("DELETE")
}   