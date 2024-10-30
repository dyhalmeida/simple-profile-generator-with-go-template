package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Profile struct {
	Name string
	Age  int64
	Bio  string
}

var profiles = []Profile{}
var templates *template.Template

func init() {
	funcMap := template.FuncMap{
		"ToUppercase": strings.ToUpper,
	}
	templates = template.Must(template.New("").Funcs(funcMap).ParseGlob("templates/*.html"))
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/new", formHandler)
	http.HandleFunc("/save", saveHandler)

	log.Println("Server started at http://localhost:3333")
	log.Fatal(http.ListenAndServe(":3333", nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "index.html", profiles)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "form.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		name := r.FormValue("name")
		age, err := strconv.ParseInt(r.FormValue("age"), 10, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		bio := r.FormValue("bio")

		profile := Profile{Name: name, Age: age, Bio: bio}
		profiles = append(profiles, profile)

	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
