package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"strconv"
)

type viewData struct {
	randNumber, userNumber, help string
}

var randNumber int

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	tmpl.ExecuteTemplate(w, "index", nil)

}

func startHandler(w http.ResponseWriter, r *http.Request) {
	var data viewData
	if r.Method == "POST" {
		userNumber := r.FormValue("userNumber")
		fmt.Println("user num = " + userNumber)
		fmt.Println("computer num = " + moreZeros(strconv.Itoa(randNumber), 4))
		data.randNumber = moreZeros(strconv.Itoa(randNumber), 4)
		data.userNumber = userNumber
		data.help = ""
	} else {
		randNumber = rand.Intn(10000)
		data.randNumber = moreZeros(strconv.Itoa(randNumber), 4)
	}

	tmpl, err := template.ParseFiles("templates/start.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	tmpl.ExecuteTemplate(w, "start", data)
}

func moreZeros(str string, length int) string {
	for len(str) < length {
		str = "0" + str
	}
	return str
}

func main() {
	fmt.Println("listening on port :3000")

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./css"))))
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/start", startHandler)

	http.ListenAndServe(":3000", nil)
}
