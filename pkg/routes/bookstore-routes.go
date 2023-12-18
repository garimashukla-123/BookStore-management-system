package routes

import (
	"github.com/garimashukla-123/BookStore-management-system/go-bookstore/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("Post")
	router.HandleFunc("/book/{bookId}", controllers.GetBokksById).Methods("Get")
	router.HandleFunc("/book/{bookId}", controllers.UpadateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("Delete")

}
