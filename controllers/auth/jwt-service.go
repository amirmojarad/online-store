package auth

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

var signingKey = []byte(os.Getenv("SECRET_KEY"))

func GetJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["client"] = "online-supermarket"
	claims["aud"] = "supermarket.online"
	claims["iss"] = "jwtgo.io"
	claims["exp"] = time.Now().Add(time.Hour * 48).Unix()
	if tokenString, err := token.SignedString(signingKey); err != nil {
		log.Println("Something went wron in jwt-service.go: ", err)
		return "", err
	} else {
		return tokenString, err
	}
}
