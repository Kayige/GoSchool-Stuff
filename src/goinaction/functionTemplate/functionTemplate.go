package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

// Student Struct with 2 variables
type student struct {
	Name  string
	Class string
}

// Course Struct with 2 vars
type totalScore struct {
	Module1 int
	Module2 int
}

var funcMap = template.FuncMap{
	"studentName":  getStudentName,
	"studentClass": getStudentClass,
	"score":        calScore,
}

func init() {
	tpl = template.Must(template.New("").Funcs(funcMap).ParseFiles("template.gohtml"))
}

func getStudentName(s student) string {
	return s.Name
}

func getStudentClass(s student) string {
	return s.Class
}

func calScore(ts totalScore) int {
	score := ts.Module1 + ts.Module2
	return score
}
func main() {

	student1 := student{
		Name:  "Obi",
		Class: "Kenobi",
	}

	student2 := student{
		Name:  "Obi",
		Class: "Kenobi",
	}

	score1 := totalScore{
		Module1: 50,
		Module2: 100,
	}

	score2 := totalScore{
		Module1: 20,
		Module2: 50,
	}
	students := []student{student1, student2}
	scores := []totalScore{score1, score2}

	data := struct {
		StudentData []student
		ScoreData   []totalScore
	}{
		students,
		scores,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "template.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}
}
