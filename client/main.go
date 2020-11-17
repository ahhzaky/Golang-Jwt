package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte("mysupersecretphrase")

func homePage(w http.ResponseWriter, r *http.Request) {
	validToken, err := GenerateJWT()
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	fmt.Fprintf(w, validToken)
}

func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user"] = "Elliot Forbes"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("Something went wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}

func handleRequest() {
	http.HandleFunc("/", homePage)

	log.Fatal(http.ListenAndServe(":9001", nil))
}

func main() {
	fmt.Println("simple rest api")

	handleRequest()
}

// https://www.youtube.com/watch?v=-Scg9INymBs  (804)
