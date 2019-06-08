package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	tmpl.ExecuteTemplate(w, "index", nil)

}

func startHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/start.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	tmpl.ExecuteTemplate(w, "start", nil)

}

func main() {
	fmt.Println("listening on port :3000")

	//http.Handle("/css/", http.FileServer(http.Dir("css/")))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./css"))))
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/start", startHandler)

	http.ListenAndServe(":3000", nil)
}
