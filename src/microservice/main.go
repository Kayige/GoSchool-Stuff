package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// run sql db instance from db.go
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:32769)/my_db")

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Database Open")
	}
	// Test the DB functions from db.go
	// DeleteRecord(db, 2)
	// EditRecord(db, 2, "Taylor", "Swift", 23)
	// InsertRecord(db, 2, "Michael", "Jackson", 55)
	// GetRecords(db)
	defer db.Close()

	// REST API functions test
	courses = make(map[string]courseInfo)
	router := mux.NewRouter()
	router.HandleFunc("/api/v1", home)
	router.HandleFunc("/api/v1/courses", allcourses)
	router.HandleFunc("/api/v1/courses/{courseid}", course).Methods("GET", "PUT", "POST", "DELETE")
	fmt.Println("Connecting at port:5000")
	log.Fatal(http.ListenAndServe(":5000", router))

}
