package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Sonrisa struct {
	Tipo       string `json:"tipo"`
	Intensidad string `json:"intensidad"`
	HasTeeth   bool   `json:"hasTeeth"`
}

func entregarSonrisa(w http.ResponseWriter, r *http.Request) {
	type ArregloSonrisas struct {
		Sonrisas []*Sonrisa
	}
	deliverable := ArregloSonrisas{
		Sonrisas: []*Sonrisa{
			&Sonrisa{
				Tipo:       "ondular",
				Intensidad: "mucha",
				HasTeeth:   true,
			},
			&Sonrisa{
				Tipo:       "circular",
				Intensidad: "poca",
				HasTeeth:   false,
			},
		},
	}
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.Encode(deliverable)
}

func entregarPais(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println(vars["code"])
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/Smiles", entregarSonrisa)
	r.HandleFunc("/pais/{code:[A-Z][A-Z]}", entregarPais)
	http.ListenAndServe(":8080", r)
}
