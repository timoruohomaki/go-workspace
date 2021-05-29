package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type sage struct {
	Name    string
	Contact string
}

func init() {
	tpl = template.Must(template.ParseFiles("007.thtml"))
}

func main() {

	// OPTION 1: WITH SLICE

	// sages := []string{"Billy", "Bob", "Joe", "Mark", "Jason"}

	// OPTION 2: WITH MAP

	// sages := map[string]string{
	//	"Houston":   "Billy",
	//	"San Diego": "Bob",
	//	"Charlotte": "Joe",
	//	"Boston":    "Mark",
	//	"Seattle":   "Jason",
	// }

	// OPTION 3: WITH STRUCT

	city := sage{
		Name:    "Houston",
		Contact: "Billy",
	}

	// err := tpl.Execute(os.Stdout, sages)
	err := tpl.Execute(os.Stdout, city)
	if err != nil {
		log.Fatalln(err)
	}
}
