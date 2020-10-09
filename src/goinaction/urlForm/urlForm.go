package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", form)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
func form(res http.ResponseWriter, req *http.Request) {
	value := req.FormValue("text")
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	//can try for using GET to see the different effect in url
	io.WriteString(res, `
	<form method="POST">
	<input type="text" name="text">
	<input type="submit"> </form>
	<br>`+value)
}
