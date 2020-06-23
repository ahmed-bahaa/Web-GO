package main

import (
	"log"
	"os"
	"text/template"
)

func main() {

	// Parsing the first
	tpl, err := template.ParseFiles("./tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	// Create html wtiter file
	htmlWriter, err := os.Create("index.html")
	if err != nil {
		log.Println("Couldn't create the file cause : ", err)
	}
	defer htmlWriter.Close()

	err = tpl.Execute(htmlWriter, nil)
	if err != nil {
		log.Fatalln(err)
	}

	// parsing the second
	tpl, err = template.ParseFiles("./tpl2.gohtml", "./tpl3.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	// Create html wtiter file 2
	htmlWriter2, err := os.Create("index1.html")
	if err != nil {
		log.Println("Couldn't create the file cause : ", err)
	}
	defer htmlWriter2.Close()

	// Create html wtiter file 3
	htmlWriter3, err := os.Create("index2.html")
	if err != nil {
		log.Println("Couldn't create the file cause : ", err)
	}
	defer htmlWriter3.Close()

	err = tpl.Execute(htmlWriter2, nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(htmlWriter3, "tpl3.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

}
