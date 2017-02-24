// My Shopping List
package main

import (
	"html/template"
	"net/http"
)

//content of my shopping list
type ShoppingList struct {
	Task string
	Done bool
}

func main() {
	tmpl := template.Must(template.ParseFiles("ShoppingList/list.html")) // the folder name has to be defined for html
	list := []ShoppingList{
		{"Sun Screen", true},
		{"New flip flops", true},
		{"Better googles", true},
		{"New Bikinis", false},
		{"New Suitcase", false},
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, struct{ Vacation []ShoppingList }{list})
	})

	http.ListenAndServe(":8080", nil)
}
