package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type user struct {
	Name     string
	LastName string
}

func apiRoute(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var us user
	err := decoder.Decode(&us)
	if err != nil {
		fmt.Println("There was an error")
		return
	}
	var text = "User is " + us.Name + " with last name " + us.LastName
	fmt.Println(text)
	fmt.Fprintf(w, text)
}

func main() {
	http.HandleFunc("/api", apiRoute)
	http.ListenAndServe(":4000", nil)
}
