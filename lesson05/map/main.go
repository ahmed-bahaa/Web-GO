package main

import (
	"log"
	"os"
	"text/template"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("./templates/*"))
}

func main() {
	names := map[string]string{
		"fname": "ahmed",
		"lname": "bahaa",
	}
	err := tmpl.ExecuteTemplate(os.Stdout, "index.gohtml", names)
	if err != nil {
		log.Fatalln(err)
	}

}
