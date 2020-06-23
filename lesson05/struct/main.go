package main

import (
	"log"
	"os"
	"text/template"
)

var tmpl *template.Template

type person struct {
	Name string
	Age  int
}

func init() {
	tmpl = template.Must(template.ParseGlob("./templates/*"))
}

func main() {

	p1 := person{Name: "ahmed",
		Age: 25,
	}
	err := tmpl.ExecuteTemplate(os.Stdout, "index.gohtml", p1)
	if err != nil {
		log.Fatalln(err)
	}

}
