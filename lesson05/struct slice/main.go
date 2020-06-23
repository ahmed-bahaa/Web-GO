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

	p2 := person{Name: "beho",
		Age: 26,
	}

	persons := []person{p1, p2}
	err := tmpl.ExecuteTemplate(os.Stdout, "index.gohtml", persons)
	if err != nil {
		log.Fatalln(err)
	}

}
