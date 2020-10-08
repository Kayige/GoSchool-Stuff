package main

import (
	"io"
	"net/http"
)

// GoMenu Function to serve string
func GoMenu(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Welcome to Go School")
}

// GoBasic Function to serve string
func GoBasic(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Welcome to Go Basic")
}

// GoAdvanced Function to serve string
func GoAdvanced(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Welcome to Go Advanced")
}
func main() {
	http.Handle("/", http.HandlerFunc(GoMenu))
	http.Handle("/basic", http.HandlerFunc(GoBasic))
	http.Handle("/advanced", http.HandlerFunc(GoAdvanced))
	http.ListenAndServe(":8080", nil)
}
