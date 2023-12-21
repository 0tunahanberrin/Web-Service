package router

import (
	"web_service_ko/pkg/controller"

	"github.com/gorilla/mux"
)

var BookRoutes = func(r *mux.Router) {
	r.HandleFunc("/books", controller.ShowAllBookDetails).Methods("GET")
	r.HandleFunc("/book/{id}", controller.ShowBookDetail).Methods("GET")
	r.HandleFunc("/book", controller.CreateBookDetail).Methods("POST")
	r.HandleFunc("/book/{id}", controller.UpdateBookDetail).Methods("PUT")
	r.HandleFunc("/book/{id}", controller.DeleteBookDetail).Methods("DELETE")
}
