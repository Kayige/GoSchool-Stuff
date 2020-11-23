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

func InsertRecord(db *sql.DB, ID int, FN string, LN string, Age int) {
	query := fmt.Sprintf("INSERT INTO Persons VALUES (%d, '%s', '%s', %d)",
		ID, FN, LN, Age)
	_, err := db.Query(query)
	if err != nil {
		panic(err.Error())

	}

}

func EditRecord(db *sql.DB, ID int, FN string, LN string, Age int) {
	query := fmt.Sprintf("Update Persons Set FirstName='%s', LastName='%s', Age='%d', WHERE ID=%d", FN, LN, ID, Age)
	_, err := db.Exec(query)
	if err != nil {
		panic(err.Error())
	}

}

func DeleteRecord(db *sql.DB, ID int) {
	query := fmt.Sprintf(
		"DELETE FROM Persons WHERE ID='%d'", ID)
	_, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}

}
