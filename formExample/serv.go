package main

import (
	"fmt"
	"net/http"
)

func showForm(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./form.html")
}

func attendForm(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Fprintf(w, "Regresar datos")
	} else if r.Method == "POST" {
		r.ParseForm()
		fmt.Fprint(w, r.Form["libro"])
	}
}

func main() {
	http.HandleFunc("/forma", showForm)
	http.HandleFunc("/login", attendForm)
	http.ListenAndServe(":4000", nil)
}
