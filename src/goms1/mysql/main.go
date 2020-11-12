package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// Person struct values
type Person struct {
	ID        int
	FirstName string
	LastName  string
	Age       int
}

func GetRecords(db *sql.DB) {
	results, err := db.Query("select * from my_db.Persons")
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		var person Person
		err := results.Scan(&person.ID, &person.FirstName, &person.LastName, &person.Age)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(person.ID, person.FirstName, person.LastName, person.Age)
	}
}

func main() {
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/my_db")

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Database Open")
	}
	GetRecords(db)
	defer db.Close()

}
