package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprintln(w, r.URL.String(), "hi iremnur")

	})
	fmt.Println("i am running")
	err := http.ListenAndServe(":8080", nil)
	fmt.Println(err)
}
