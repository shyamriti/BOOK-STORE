package main

import (
	"fmt"
	"go-sqlite/pkg/controller"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// database.Connection()

	// defer database.Db.Close()

	// http.HandleFunc("/getall", controller.GetAllHandleFunc)
	// http.HandleFunc("/getone", controller.GetOneHandleFunc)
	// http.HandleFunc("/insert", controller.InsertHandleFunc)
	// http.HandleFunc("/update", controller.UpdateHandleFunc)
	// http.HandleFunc("/delete", controller.DeleteHandleFunc)

	// fmt.Println("Server is running")
	// if err := http.ListenAndServe(":8000", nil); err != nil {
	// 	log.Fatal(err)
	// }

	// Init the mux router
	Router := mux.NewRouter()

	// Route handles & endpoints

	// Get all movies
	Router.HandleFunc("/person", controller.GetAllHandleFunc).Methods("GET")

	// Create a movie
	Router.HandleFunc("/person", controller.InsertHandleFunc).Methods("POST")

	// Delete a specific movie by the movieID
	Router.HandleFunc("/person/{id}", controller.UpdateHandleFunc).Methods("PUT")

	// Delete all movies
	Router.HandleFunc("/person/{id}", controller.DeleteHandleFunc).Methods("DELETE")

	// serve the app
	fmt.Println("Server running at 8000")
	http.ListenAndServe(":8000", Router)
}
