// todos.go
package main

import (
	"html/template"
	"net/http"
)

//content of the todo list
type Todo1 struct {
	Task string
	Done bool
}

func main() {
	tmpl := template.Must(template.ParseFiles("webExample4/todos.html")) // the folder name has to be defined for html
	todos := []Todo1{
		{"Learn Go", true},
		{"Read Go Web Examples", true},
		{"Create a web app in Go", false},
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, struct{ Todossomething []Todo1 }{todos})
	})

	http.ListenAndServe(":8080", nil)
}
