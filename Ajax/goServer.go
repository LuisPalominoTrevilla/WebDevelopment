package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type Person struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Password string `json:"password"`
}

func responder(respuesta http.ResponseWriter, solicitud *http.Request) {
	currentTime := time.Now().Local()
	fmt.Fprintf(respuesta, "Hola, la fecha actual es "+currentTime.Format("2006-01-02"))
}

func jsongo(respuesta http.ResponseWriter, solicitud *http.Request) {
	// b := [5]string{"Hola", "Mundo", "Como", "Estas"}
	me := Person{
		Name:     "Luis",
		Age:      10,
		Password: "luis123",
	}
	// fmt.Fprintf(respuesta, b[3])
	json.NewEncoder(respuesta).Encode(me)
}

func showArray(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		superArray := [5]string{"Hola", "Mundo", "Como", "Estas"}

		num, _ := strconv.ParseInt(r.FormValue("number"), 10, 64)
		fmt.Fprint(w, superArray[num])
	}
}

func attendForm(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Fprintf(w, "Regresar datos")
	} else if r.Method == "POST" {
		r.ParseForm()
		fmt.Fprint(w, r.Form["libro"])
	}
}

func mostrarHTML(respuesta http.ResponseWriter, solicitud *http.Request) {
	fmt.Println("quibo")
	http.ServeFile(respuesta, solicitud, "ajax_file.html")
	fmt.Println("hola")
}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static", fs))

	http.HandleFunc("/request", responder)
	http.HandleFunc("/json", jsongo)
	http.HandleFunc("/login", attendForm)
	http.HandleFunc("/array", showArray)
	http.HandleFunc("/vue", mostrarHTML)
	http.ListenAndServe(":4000", nil)
}
