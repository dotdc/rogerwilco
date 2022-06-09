package main

/*
    Author: David Calvert
	Purpose: Run the Roger Wilco dummy application!
*/

import (
	"html/template"
	"log"
	"net/http"
)

var AppVersion = "unset"

type Page struct {
	Title   string
	Version string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	values := Page{
		Title:   "Roger Wilco App",
		Version: AppVersion,
	}

	template, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		log.Println("Error parsing files :", err)
	}

	terr := template.Execute(w, values)
	if terr != nil {
		log.Println("Error executing template :", terr)
		return
	}
}

// Starts the webserver and exposes the different features on port 8080
func main() {
	router := http.NewServeMux()
	router.HandleFunc("/", indexHandler)
	router.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	log.Fatal(http.ListenAndServe(":8080", router))
}
