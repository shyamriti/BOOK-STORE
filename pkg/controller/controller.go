package controller

import (
	"encoding/json"
	"fmt"
	"go-sqlite/pkg/database"
	"go-sqlite/pkg/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetAllHandleFunc(w http.ResponseWriter, r *http.Request) {
	p := database.ReadAll()

	json.NewEncoder(w).Encode(p)
}

func InsertHandleFunc(w http.ResponseWriter, r *http.Request) {

	person := models.Person{}
	json.NewDecoder(r.Body).Decode(&person)
	p := database.Insert(person)
	fmt.Fprintf(w, "Persons details was inserted\n")
	json.NewEncoder(w).Encode(p)
}

func UpdateHandleFunc(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	person := models.Person{}
	json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	database.Update(person, id)
	json.NewEncoder(w).Encode("Record was updated")

}

func DeleteHandleFunc(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	database.Delete(id)
	json.NewEncoder(w).Encode("Record was deleted")

}
