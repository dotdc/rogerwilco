package main

/*
    Author: David Calvert
	Purpose: Run the Roger Wilco dummy application!
*/

import (

	"log"
	"net/http"
	"html/template"
)

type Page struct {
    Title string
    Version string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	values := Page{
        Title: "Roger Wilco App",
        Version: "0.1.0",
	}
	
	template, _ := template.ParseFiles("templates/index.html")
	err := template.Execute(w, values)
	if err != nil {
		log.Println("Error executing template :", err)
		return
	}
}

// Starts the webserver and exposes the different features on port 8080
func main() {
	router := http.NewServeMux()
	router.HandleFunc("/", indexHandler)
	router.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	log.Fatal(http.ListenAndServe(":8080", router))
}