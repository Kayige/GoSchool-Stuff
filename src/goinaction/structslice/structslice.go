package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

// Student Struct with 2 variables
type student struct {
	FirstName string
	LastName  string
}

// Course Struct with 2 vars
type course struct {
	Name string
	Size int
}

func init() {
	tpl = template.Must(template.ParseFiles("templateVar.gohtml"))
}
func main() {

	student1 := student{
		FirstName: "Obi",
		LastName:  "Kenobi",
	}

	student2 := student{
		FirstName: "Obi",
		LastName:  "Kenobi",
	}

	course1 := course{
		Name: "Padawan",
		Size: 100,
	}

	course2 := course{
		Name: "Jedi",
		Size: 50,
	}
	students := []student{student1, student2}
	courses := []course{course1, course2}

	data := struct {
		StudentNames []student
		CourseNames  []course
	}{
		students,
		courses,
	}
	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}
