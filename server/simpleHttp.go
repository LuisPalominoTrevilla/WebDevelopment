package main

import (
	"fmt"
	"net/http"
	"time"
)

func responder(respuesta http.ResponseWriter, solicitud *http.Request) {
	currentTime := time.Now().Local()
	fmt.Fprintf(respuesta, "Hola, la fecha actual es "+currentTime.Format("2006-01-02"))
}

func mostrarHTML(respuesta http.ResponseWriter, solicitud *http.Request) {
	fmt.Println("quibo")
	http.ServeFile(respuesta, solicitud, "vue.html")
	fmt.Println("hola")
}

func main() {
	http.HandleFunc("/request", responder)
	http.HandleFunc("/vue", mostrarHTML)
	http.ListenAndServe(":4000", nil)
}
