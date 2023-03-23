package routes

import (
	"github.com/C:/Users/user/Desktop/New folder/Golang/CR Tejavardhan Reddy/Demo_projects/DATABASE/package/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookIb}".controllers.DeleteBook).Methods("DELETE")
}
