package routes

import {
	"github.com/gorilla/mux"
	"github.com/Garima/go-workspace/go-bookstore/pkg/controllers"

}

var RegisterBookStoreRoutes = func(router * mux.Router){
	router.HandleFunc("/book/",controllers.CreateBook).Methods("Post")
	router.HandleFunc("/book/{bookId}", controllers.GetBokksById).Methods("Get")
	router.HandleFunc("/book/{bookId}", controllers.UpadateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}",controllers.DeleteBook).Methods("Delete")

}