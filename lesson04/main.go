package main

import (
	"log"
	"os"
	"text/template"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	name := "behoo"
	err := tmpl.ExecuteTemplate(os.Stdout, "index.gohtml", name)
	if err != nil {
		log.Fatalln(err)
	}

}
