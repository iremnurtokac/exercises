// json.go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//Work information
type Work struct {
	MasterThesis string `json:"masterThesis"`
	Deadline1    string `json:"26 February 2017"`
	WorkshopPro  string `json:"workshopPro"`

	Deadline2 string `json:"27 February 2017"`
}

func main() {
	http.HandleFunc("/decode", func(w http.ResponseWriter, r *http.Request) {
		var work Work
		json.NewDecoder(r.Body).Decode(&work)

		fmt.Fprintf(w, "%s deadline: %s and %s deadline: %s", work.MasterThesis, work.Deadline1, work.WorkshopPro, work.Deadline2)
	})

	http.HandleFunc("/encode", func(w http.ResponseWriter, r *http.Request) {
		peter := Work{
			MasterThesis: "masterThesis",
			Deadline1:    "26 February 2017",
			WorkshopPro:  "workshopPro",
			Deadline2:    "27 February 2017",
		}

		json.NewEncoder(w).Encode(peter)
	})

	http.ListenAndServe(":8080", nil)
}
