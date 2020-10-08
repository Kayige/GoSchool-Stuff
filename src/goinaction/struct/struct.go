package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

// Student Struct with 3 variables
type Student struct {
	fname string
	lname string
}

func init() {
	tpl = template.Must(template.ParseFiles("template.gohtml"))
}
func main() {

	student1 := Student{
		fname: "Obi",
		lname: "Kenobi",
	}

	err := tpl.Execute(os.Stdout, student1)
	if err != nil {
		log.Fatalln(err)
	}
}
