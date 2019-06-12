package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type viewData struct {
	RandNumber  string
	UserNumber  string
	Help        string
	Attempts    int
	TypeRandNum string
	DisButton   string
}

var randNumber int
var data viewData

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		fmt.Println(err.Error())
	}
	tmpl.ExecuteTemplate(w, "index", nil)

}

func startHandler(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().Unix())
	tmpl, err := template.ParseFiles("templates/start.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	data.Help = ""
	data.TypeRandNum = "password"
	data.DisButton = ""

	if r.Method == "POST" {
		userNumber := r.FormValue("userNumber")
		data.RandNumber = moreZeros(strconv.Itoa(randNumber), 4)
		data.UserNumber = userNumber
		data.Attempts++
		if data.UserNumber == data.RandNumber {
			data.RandNumber = moreZeros(strconv.Itoa(randNumber), 4)
			data.TypeRandNum = "text"
			data.DisButton = "disabled"
			data.Help = "ПОБЕДА!"
			tmpl.ExecuteTemplate(w, "start", data)
			return
		}
		help := [4]string{"", "", "", ""}
		//create help
		userNumberS := strings.Split(data.UserNumber, "")
		randNumberS := strings.Split(data.RandNumber, "")
		for ind, value := range userNumberS {
			if value == randNumberS[ind] {
				help[ind] = "B"
			} else if strings.Count(data.RandNumber, value) != 0 {
				help[ind] = "K"
			} else {
				help[ind] = "  "
			}
		}
		for _, value := range help {
			data.Help += value
		}
	} else {
		randNumber = rand.Intn(10000)
		data.Attempts = 0
		data.RandNumber = moreZeros(strconv.Itoa(randNumber), 4)
		data.UserNumber = ""
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

	http.ListenAndServe(":80", nil)
}
