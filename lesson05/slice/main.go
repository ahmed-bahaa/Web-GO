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
	names := []string{"beho", "bahabeho :P ", "test "}
	err := tmpl.ExecuteTemplate(os.Stdout, "index.gohtml", names)
	if err != nil {
		log.Fatalln(err)
	}

}
