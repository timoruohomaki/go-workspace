package main

import (
	"log"
	"os"
	"text/template"
)

// PARSE TEMPLATE AND CREATE FILE

func main() {
	tpl, err := template.ParseFiles("004.thtml")
	if err != nil {
		log.Fatalln(err)
	}

	nf, err := os.Create("004-out.html")
	if err != nil {
		log.Println("Error creating file", err)
	}

	defer nf.Close()

	err = tpl.Execute(nf, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
