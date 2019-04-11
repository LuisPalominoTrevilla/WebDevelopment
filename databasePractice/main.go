package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/sessions"
)

var db *sql.DB

// var caja *sessions.CookieStore
var caja *sessions.FilesystemStore

func init() {
	var err error
	db, err = sql.Open("mysql", "root@/golangPractice")
	if err != nil {
		log.Fatal("Not able to connect ", err.Error())
	}

	// caja = sessions.NewCookieStore([]byte("llave"))
	caja = sessions.NewFilesystemStore("./sesiones", []byte("llavesssfsdgf"))
}

func activar(w http.ResponseWriter, r *http.Request) {
	sesion, err := caja.Get(r, "otropues")
	if err != nil {
		log.Fatal("error: ", err.Error())
	}
	sesion.Values["otropues"] = 3
	err = sesion.Save(r, w)
	if err != nil {
		log.Fatal("error: ", err.Error())
	}
}

func revisar(w http.ResponseWriter, r *http.Request) {
	sesion, err := caja.Get(r, "nombre")
	if err != nil {
		log.Fatal("error: ", err.Error())
	}
	cok := sesion.Values["nombre"]
	if cok == nil {
		fmt.Fprint(w, "Aqui no puedes estar")
	} else {
		fmt.Fprint(w, cok)
	}
}

func basedatos(w http.ResponseWriter, r *http.Request) {
	filas, err := db.Query("SELECT * FROM movies;")

	if err != nil {
		log.Fatal("error en el servidor ", err.Error())
	}

	for filas.Next() {

		var (
			id    int
			title string
		)

		err := filas.Scan(&id, &title)

		if err != nil {
			log.Fatal("error en el servidor ", err.Error())
		}
		fmt.Printf("ID is %d, title is %s\n", id, title)
	}
}

func main() {
	http.HandleFunc("/db", basedatos)
	http.HandleFunc("/check", revisar)
	http.HandleFunc("/activate", activar)
	http.ListenAndServe(":8080", nil)
}
