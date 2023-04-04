package controller

import (
	"fmt"
	"go-sqlite/pkg/database"
	"net/http"
)

func HandleFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hey!\n")
	database.Insert()
	fmt.Fprintf(w, "Persons details was inserted\n")

	person := database.Query()
	for i := range person{
	fmt.Fprintf(w, "%v\n\n", person[i])

	}
	

}
