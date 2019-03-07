package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

func ponerCookie(w http.ResponseWriter, r *http.Request) {
	oreo1 := http.Cookie{
		Name:    "idioma",
		Value:   "fr",
		Expires: time.Now().Add(time.Minute * 10),
	}
	http.SetCookie(w, &oreo1)
}

func obtenerCookie(w http.ResponseWriter, r *http.Request) {
	datos := r.Header["Cookie"]
	fmt.Fprint(w, datos)
}

func deleteCookie(w http.ResponseWriter, r *http.Request) {
	oreo1 := &http.Cookie{
		Name:    "idioma",
		Value:   "NADAAAAAA",
		Expires: time.Now().Add(time.Minute * -10),
	}
	http.SetCookie(w, oreo1)
}

func getHTML(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./index.html")
}

func getExercise2(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./exercise2.html")
}

func setCountry(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	country := r.PostFormValue("country")
	florentina := http.Cookie{
		Name:  "country",
		Value: country,
	}
	http.SetCookie(w, &florentina)
	http.Redirect(w, r, "http://localhost:8080/", 301)
}

func receiveIdiom(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	idioma := r.PostFormValue("idioma")
	fmt.Println("Received ", idioma)
	fmt.Fprintf(w, "mandaste "+idioma)
}

func mostrarPlantilla(w http.ResponseWriter, r *http.Request) {
	type product struct {
		ID          string
		Product     string
		Price       int
		IsFree      bool
		Description string
	}
	var products []product = []product{
		product{
			ID:          "!@#$%ˆ",
			Product:     "Pintura",
			Description: "Prisma",
			Price:       129,
			IsFree:      true,
		},
		product{
			ID:          "12345F",
			Product:     "Cemento",
			Description: "Cruz azul",
			Price:       150,
			IsFree:      false,
		},
		product{
			ID:          "7584DFT",
			Product:     "Jabón",
			Description: "ZOTE",
			Price:       300,
			IsFree:      false,
		},
	}
	p, _ := template.ParseFiles("plantilla.html")
	p.Execute(w, products)
}

func main() {
	http.HandleFunc("/poner", ponerCookie)
	http.HandleFunc("/obtener", obtenerCookie)
	http.HandleFunc("/delete", deleteCookie)
	http.HandleFunc("/setCountry", setCountry)
	http.HandleFunc("/sendLs", receiveIdiom)
	http.HandleFunc("/", getHTML)
	http.HandleFunc("/ex2", getExercise2)
	http.HandleFunc("/template", mostrarPlantilla)
	http.ListenAndServe(":8080", nil)
}
