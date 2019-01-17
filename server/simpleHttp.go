package main

import (
	"fmt"
	"net/http"
	"time"
)

func responder(respuesta http.ResponseWriter, solicitud *http.Request) {
	current_time := time.Now().Local()
	fmt.Fprintf(respuesta, "Hola, la fecha actual es "+current_time.Format("2006-01-02"))
}

func main() {
	http.HandleFunc("/", responder)
	http.ListenAndServe(":4000", nil)
}
