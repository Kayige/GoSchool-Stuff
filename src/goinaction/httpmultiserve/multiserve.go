package main

import (
	"io"
	"net/http"
)

// GoBasic Type
type GoBasic int

// GoAdvanced Type
type GoAdvanced int

// GoMenu Type
type GoMenu int

func (b GoBasic) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Welcome to Go Basic")
}

func (a GoAdvanced) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Welcome to Go Advanced")
}

func (c GoMenu) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Welcome to Go School")
}

func main() {
	var a GoAdvanced
	var b GoBasic
	var c GoMenu

	http.Handle("/", c)
	http.Handle("/advanced", a)
	http.Handle("/basic", b)
	http.ListenAndServe(":8080", nil)
}
