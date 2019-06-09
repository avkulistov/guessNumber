package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"strconv"
)

var randNumber = rand.Intn(10000)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	tmpl.ExecuteTemplate(w, "index", nil)

}

func startHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		numberForGuess := r.FormValue("numberForGuess")
		userNumber := r.FormValue("userNumber")
		fmt.Println("user num = " + userNumber)
		fmt.Println("computer num = " + numberForGuess)
		fmt.Println("random num = " + strconv.Itoa(randNumber))
	} else {
		randNumber = rand.Intn(10000)
	}

	tmpl, err := template.ParseFiles("templates/start.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	tmpl.ExecuteTemplate(w, "start", nil)
}

func main() {
	fmt.Println("listening on port :3000")

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./css"))))
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/start", startHandler)

	http.ListenAndServe(":3000", nil)
}
