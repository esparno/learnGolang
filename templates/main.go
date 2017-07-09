package main

import (
	"text/template"
	"log"
	"os"
)

func main() {
	tpl, err := template.ParseGlob("*.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", nil)
}
