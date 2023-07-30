package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type ViewData struct {
	Title   string
	Message string
	Users   []string
}

func main() {

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		//http.ServeFile(w, r, "hello.html")
		name := r.URL.Query().Get("name")
		age := r.URL.Query().Get("age")
		data := ViewData{
			Title:   "World Cup " + name,
			Message: "FIFA will never regret it " + age,
			Users:   []string{"Tom", "Bob", "Sam"},
		}
		fmt.Println(age)
		fmt.Println(name)

		//fmt.Fprintf(w, "Имя: %s Возраст: %s", name, age)

		tmpl, _ := template.ParseFiles("templates/hello.html")
		tmpl.Execute(w, data)
	})
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "About Page")
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Index Page")
	})
	fmt.Println("Server is listening...")
	http.ListenAndServe(":8181", nil)
}
