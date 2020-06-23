package main

import (
	"log"
	"os"
	"text/template"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	err := tmpl.Execute(os.Stdout, "Ahmed")
	if err != nil {
		log.Fatalln(err)
	}

}
