package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", URLCall)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)

}

// URLCall Function to query form
func URLCall(res http.ResponseWriter, req *http.Request) {
	passValue := req.FormValue("query")
	fmt.Fprintln(res, "Value is: ", passValue)
}
