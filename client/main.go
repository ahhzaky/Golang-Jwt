package main

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte("mysupersecretphrase")

func GeneateJWT() (string, error) {
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

func main() {
	fmt.Println("simple rest api")

	tokenString, err := GeneateJWT()
	if err != nil {
		fmt.Println("Something went wrong")
	}

	fmt.Println(tokenString)
}

// https://www.youtube.com/watch?v=-Scg9INymBs  (804)
