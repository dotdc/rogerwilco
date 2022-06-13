package main

/*
    Author: David Calvert
	Purpose: Run the Roger Wilco dummy application!
*/

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var AppVersion = "unset"

type Page struct {
	Title     string
	Version   string
	UserAgent string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	values := Page{
		Title:     "Roger Wilco App",
		Version:   AppVersion,
		UserAgent: r.UserAgent(),
	}

	template, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		rc := http.StatusInternalServerError
		w.WriteHeader(rc)
		log.Println(r.URL.Path, rc, "[ERROR] parsing files:", err)
		return
	}

	terr := template.Execute(w, values)
	if terr != nil {
		returncode := http.StatusInternalServerError
		w.WriteHeader(returncode)
		log.Println(r.URL.Path, returncode, "[ERROR] executing template:", terr)
		return
	}

	log.Println(r.URL.Path, http.StatusOK, "[INFO] http request received")
}

// Starts the webserver and exposes the different features on port 8080
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", indexHandler)
	router.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	log.Fatal(http.ListenAndServe(":8080", router))
}
