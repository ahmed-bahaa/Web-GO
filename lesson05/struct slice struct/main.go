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

type animal struct {
	Name  string
	Kindd string
}

type animals struct {
	Persons []person
	Animals []animal
}

func init() {
	tmpl = template.Must(template.ParseGlob("./templates/*"))
}

func main() {

	// persons
	p1 := person{Name: "ahmed",
		Age: 25,
	}
	p2 := person{Name: "beho",
		Age: 26,
	}

	// animals
	a1 := animal{Name: "timor",
		Kindd: "cat",
	}
	a2 := animal{Name: "toti",
		Kindd: "dog",
	}

	ps := []person{p1, p2}
	as := []animal{a1, a2}

	ans := animals{ps, as}

	err := tmpl.ExecuteTemplate(os.Stdout, "index.gohtml", ans)
	if err != nil {
		log.Fatalln(err)
	}

}
