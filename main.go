package main

import (
	ascii "asciiws/func"
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(404)
		// tmplErr404.Execute(w, nil)
		ascii.Templates.ExecuteTemplate(w, "errors.html", http.StatusNotFound)
		return
	}
	// ascii.Templates
	// tmplHomePage.Execute(w, nil)
	ascii.Templates.ExecuteTemplate(w, "mainPage.html", nil)
}
func workPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ascii-art/" {
		w.WriteHeader(404)
		// tmplErr404.Execute(w, nil)
		ascii.Templates.ExecuteTemplate(w, "errors.html", http.StatusNotFound)
		return
	}

	if r.Method != "POST" {
		w.WriteHeader(405)
		// tmplErr405.Execute(w, nil)
		ascii.Templates.ExecuteTemplate(w, "errors.html", http.StatusMethodNotAllowed)
		return
	}
	ascii.Post(w, r)
}

func main() {
	if ascii.TmplError != nil {
		fmt.Println("Templates Error")
		return
	}
	ascii.Fun("picture")
	fmt.Println("Server connected ...")
	http.Handle("/media/", http.StripPrefix("/media/", http.FileServer(http.Dir("media"))))
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ascii-art/", workPage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
