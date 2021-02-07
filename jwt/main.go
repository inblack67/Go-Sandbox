package main

import (
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte("somesecret");

func getJWT() (string, error){
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["user"] = "ryuk"
	claims["exp"] = time.Now().Add(time.Minute * 60).Unix()

	tokenString, err := token.SignedString(jwtSecret)

	if err != nil{
		log.Fatal("err gen token", err)
		return "", err
	}

	return tokenString, nil
}

func main(){

	myToken, err := getJWT()

	if err != nil{
		log.Fatal("err gen token", err)
	}

	fmt.Println(myToken)

}