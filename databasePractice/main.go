package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "root@/golangPractice")
	if err != nil {
		log.Fatal("Not able to connect ", err.Error())
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
	http.ListenAndServe(":8080", nil)
}
