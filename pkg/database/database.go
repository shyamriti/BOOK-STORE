package database

import (
	"database/sql"
	"fmt"
	"go-sqlite/pkg/models"
	"log"
)


var Db *sql.DB

var P models.Person

func Connection() *sql.DB {
	var err error
	Db, err = sql.Open("sqlite3", "sqlite-database.db")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	return Db
}

func ReadAll() []models.Person {
	db := Connection()
	P := models.Person{}
	PersonArr := []models.Person{}
	rows, err := db.Query("SELECT * FROM person")
	if err != nil {
		log.Print(err)
	}
	for rows.Next() {
		rows.Scan(&P.Id, &P.Name, &P.Email, &P.PhoneNumber)
		PersonArr = append(PersonArr, P)
	}
	return PersonArr
}

func Insert(person models.Person) models.Person {
	db := Connection()
	p := person
	_, err := db.Exec("INSERT INTO person (id, name, email, phone) VALUES (?, ?, ?, ?)", p.Id, p.Name, p.Email, p.PhoneNumber)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	return p
}

func Update(person models.Person, id int) {
	db := Connection()
	_, err := db.Exec("update person set name=?, email=?, phone=? where id=?", person.Name, person.Email, person.PhoneNumber, id)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

}

func Delete(id int) {
	db := Connection()
	defer db.Close()
	result, err := db.Exec("delete from person where id=?", id)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	result.RowsAffected()

}
