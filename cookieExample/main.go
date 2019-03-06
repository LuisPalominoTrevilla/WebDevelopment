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
	chips, _ := r.Cookie("idioma")
	fmt.Fprint(w, chips)
	chips.MaxAge = -1
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
	p, _ := template.ParseFiles("plantilla.html")
	p.Execute(w, "Ya casi nos vamos")
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
