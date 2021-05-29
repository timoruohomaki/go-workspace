package main

import (
	"log"
	"os"
	"text/template"
)

// PARSE ALL TEMPLATES AND CREATE OUTPUT FILE

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "004.thtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
