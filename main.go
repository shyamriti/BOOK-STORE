package main

import (
	"fmt"
	"go-sqlite/pkg/controller"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	
	// Init the mux router
	Router := mux.NewRouter()

	// Route handles & endpoints

	// Get all data
	Router.HandleFunc("/person", controller.GetAllHandleFunc).Methods("GET")

	// Insert a data
	Router.HandleFunc("/person", controller.InsertHandleFunc).Methods("POST")

	// Update a specific data by the id
	Router.HandleFunc("/person/{id}", controller.UpdateHandleFunc).Methods("PUT")

	// delete a specific data by the id
	Router.HandleFunc("/person/{id}", controller.DeleteHandleFunc).Methods("DELETE")

	// serve the app
	fmt.Println("Server running at 8000")
	http.ListenAndServe(":8000", Router)
}
