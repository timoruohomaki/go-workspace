package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("005.thtml"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "005.thtml", 42)
	if err != nil {
		log.Fatalln(err)
	}
}
