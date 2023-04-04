package database

import (
	"database/sql"
	"fmt"
	"go-sqlite/pkg/models"
	"log"
)

var Db *sql.DB

func Connection() {
	var err error
	Db, err = sql.Open("sqlite3", "sqlite-database.db")

	if err != nil {
		fmt.Printf("err: %v\n", err)
		// log.Fatal(err)
	}

}

func Query() []models.Person {
	person := models.Person{}
	personArr := []models.Person{}

	rows, err := Db.Query("SELECT * FROM person")
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {

		err = rows.Scan(&person.Id, &person.Name, &person.Email, &person.PhoneNumber)

		if err != nil {
			log.Print(err)
		}
		personArr = append(personArr, person)
	}

	return personArr
}

func Insert() {

	insertingData := "INSERT INTO person (name, email, phone) VALUES ('shib', 'shib@gmail.com', '1234')"
	_, err := Db.Exec(insertingData)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

}
