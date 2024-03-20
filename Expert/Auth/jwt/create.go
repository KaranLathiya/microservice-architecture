package jwt

import (
	"expert/model"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

var jwtKey []byte

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(fmt.Sprintf("Some error occured. Err: %s", err))
	}
	jwtKey = []byte(os.Getenv("JWTKEY"))
}

func CreateJWT() (string, error) {
	expirationTime := time.Now().Add(time.Minute * 5)

	claims := model.Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			Audience:  "Portfolio",
			Subject:   "Portfolios of expert",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
