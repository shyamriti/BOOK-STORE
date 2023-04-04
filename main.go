package main

import (
	"fmt"
	"go-sqlite/pkg/controller"
	"go-sqlite/pkg/database"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	database.Connection()
	// name, email, phonenumber, address := database.Query()
	defer database.Db.Close()

	http.HandleFunc("/hello", controller.HandleFunc)

	fmt.Println("Server is running")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}

	// fmt.Fprintf("%s %s %d %s\n", name, email, phonenumber, address)
}
