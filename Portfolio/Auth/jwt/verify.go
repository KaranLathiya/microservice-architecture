package jwt

import (
	"fmt"
	"net/http"
	"os"
	"portfolio/model"

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

func VerifyToken(token string) model.Message {
	claims := &model.Claims{}
	var message model.Message
	tkn, err := jwt.ParseWithClaims(token, claims,
		func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			message.Code = http.StatusUnauthorized
			message.Message = err.Error()
			return message
		}
		message.Code = http.StatusBadRequest
		message.Message = err.Error()
		return message
	}

	if !tkn.Valid {
		message.Code = http.StatusUnauthorized
		message.Message = "token is invalid"
		return message
	}
	if claims.Audience != "Portfolio" && claims.Subject == "Portfolios of expert" {
		message.Code = http.StatusUnauthorized
		message.Message = "invalid token details"
		return message
	}
	return message
}
