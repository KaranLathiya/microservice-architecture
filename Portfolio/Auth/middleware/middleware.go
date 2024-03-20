package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"portfolio/Auth/jwt"
	"portfolio/response"
)

func Middleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		token := r.Header.Get("Authorization")
		fmt.Println(token)
		errMessage := jwt.VerifyToken(token)
		if !(errMessage.Code == 0 || errMessage.Message == "") {
			w.WriteHeader(errMessage.Code)
			errMessage, err := json.MarshalIndent(errMessage, "", "  ")
			if err != nil {
				response.MessageShow(400, "Error on marshaling", w)
				return
			}
			w.Write(errMessage)
			return
		}
		next.ServeHTTP(w, r)
	})
}
