/*
Within GoMS1 folder, create a new folder called REST
Within the REST folder, create a new file named main.go

*/

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type courseInfo struct {
	Title string `json:"Title"`
}

var courses map[string]courseInfo

func validKey(r *http.Request) bool {
	v := r.URL.Query()
	if key, ok := v["key"]; ok {
		if key[0] == "2c78afaf-97da-4816-bbee-9ad239abb296" {
			return true
		} else {
			fmt.Println("Invalid Key")
			return false
		}
	} else {
		return false
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my REST api.")
}

func allcourses(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "List of all courses")
	kv := r.URL.Query()
	// k v pairs parsed through Query
	for k, v := range kv {
		fmt.Println(k, v)
	}
	json.NewEncoder(w).Encode(courses)
}

func course(w http.ResponseWriter, r *http.Request) {
	if !validKey(r) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("401 - Invalid Key."))
		return
	}
	params := mux.Vars(r)
	// fmt.Fprintf(w, "Details for course "+params["courseid"])
	// fmt.Fprintf(w, "\n")

	if r.Method == "GET" {
		if _, ok := courses[params["courseid"]]; ok {
			// course found!
			json.NewEncoder(w).Encode(courses[params["courseid"]])

		} else {
			// course not found!
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("404 - No course found!"))

		}

	}

	if r.Method == "DELETE" {
		if _, ok := courses[params["courseid"]]; ok {
			// if course exist delete
			delete(courses, params["courseid"])
			w.WriteHeader(http.StatusNoContent)
			w.Write([]byte("204 - Course is deleted"))

		} else {
			// course not found!
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("404 - No course found!"))

		}
	}

	if r.Header.Get("Content-type") == "application/json" {

		// Create new course
		if r.Method == "POST" {
			var newCourse courseInfo
			reqBody, err := ioutil.ReadAll(r.Body)

			if err == nil {
				json.Unmarshal(reqBody, &newCourse)
				if newCourse.Title == "" {
					w.WriteHeader(http.StatusUnprocessableEntity)
					w.Write([]byte("422 - Please supply course information in JSON format"))
					return
				}
				if _, ok := courses[params["courseid"]]; !ok {
					courses[params["courseid"]] = newCourse
					w.WriteHeader(http.StatusCreated)
					w.Write([]byte("201 - Course added: " + params["courseid"]))

				} else {
					w.WriteHeader(http.StatusConflict)
					w.Write([]byte("409 - Duplicate course ID"))
				}

			} else {
				w.WriteHeader(http.StatusUnprocessableEntity)
				w.Write([]byte("422 - Please supply course information in JSON format"))
			}

		}

		//Update course details
		if r.Method == "PUT" {
			var newCourse courseInfo
			reqBody, err := ioutil.ReadAll(r.Body)

			if err == nil {
				json.Unmarshal(reqBody, &newCourse)
				if newCourse.Title == "" {
					w.WriteHeader(http.StatusUnprocessableEntity)
					w.Write([]byte("422 - Sorry the course cannot be added"))
					return
				}

				if _, ok := courses[params["courseid"]]; !ok {
					courses[params["courseid"]] = newCourse
					w.WriteHeader(http.StatusCreated)
					w.Write([]byte("201 - Course is created"))
				} else {
					courses[params["courseid"]] = newCourse
					w.WriteHeader(http.StatusOK)
					w.Write([]byte("200 - Course is updated"))
				}
			} else {
				w.WriteHeader(http.StatusUnprocessableEntity)
				w.Write([]byte("422 - Sorry the course cannot be added/updated"))
			}
		}

	}

}

func main() {
	courses = make(map[string]courseInfo)
	router := mux.NewRouter()
	router.HandleFunc("/api/v1", home)
	router.HandleFunc("/api/v1/courses", allcourses)
	router.HandleFunc("/api/v1/courses/{courseid}", course).Methods("GET", "PUT", "POST", "DELETE")
	fmt.Println("Connecting at port:5001")
	log.Fatal(http.ListenAndServe(":5001", router))
}
