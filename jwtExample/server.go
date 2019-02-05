package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

type user struct {
	Name     string
	LastName string
	password string
}

var mySigninKey []byte

func apiRoute(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var us user
	err := decoder.Decode(&us)
	if err != nil {
		fmt.Println("There was an error")
		return
	}
	var text = "User is " + us.Name + " with last name " + us.LastName
	fmt.Println(text)
	fmt.Fprintf(w, text)
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	tokenString, err := GenerateJWT()
	if err != nil {
		fmt.Fprintf(w, "There was an error while creating the token")
		return
	}
	fmt.Fprintf(w, tokenString)
}

// GenerateJWT Generates jwt token
func GenerateJWT() (string, error) {

	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		Issuer:    "Luis",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(mySigninKey)
	if err != nil {
		fmt.Println("There was an error ", err.Error())
		fmt.Errorf("Something went wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}

func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Token"] != nil {
			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}
				return mySigninKey, nil
			})

			if err != nil {
				fmt.Fprintf(w, err.Error())
			}

			if token.Valid {
				endpoint(w, r)
			}

		} else {
			fmt.Fprintf(w, "Not authorized")
		}
	})
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Super secret information")
}

func requestHandler() {
	http.HandleFunc("/api", apiRoute)
	http.HandleFunc("/login", handleLogin)
	http.Handle("/home", isAuthorized(homePage))
	log.Fatal(http.ListenAndServe(":4000", nil))
}

func main() {
	// Load env variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file " + err.Error())
	}
	mySigninKey = []byte(os.Getenv("SECRET_JWT_KEY"))
	// mt.Println("My key is " + mySigninKey)
	requestHandler()
}
