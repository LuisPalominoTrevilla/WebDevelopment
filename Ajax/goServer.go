package main

import (
	"encoding/json"
	"fmt"
	"net/http"
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

func mostrarHTML(respuesta http.ResponseWriter, solicitud *http.Request) {
	fmt.Println("quibo")
	http.ServeFile(respuesta, solicitud, "ajax_file.html")
	fmt.Println("hola")
}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	http.HandleFunc("/request", responder)
	http.HandleFunc("/array", jsongo)
	http.HandleFunc("/vue", mostrarHTML)
	http.ListenAndServe(":4000", nil)
}
