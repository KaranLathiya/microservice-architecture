package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"portfolio/dal"
	"portfolio/model"
	"portfolio/response"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

var jwtKey []byte

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
	jwtKey = []byte(os.Getenv("JWTKEY"))
	db, _ := dal.Connect()
	defer db.Close()
	fmt.Println("Server started")
	// run()
	r := chi.NewRouter()
	r.Route("/user", func(r chi.Router) {
		r.Use(middleware)
		r.Post("/portfolio", portfolioList)
	})

	r.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(405)
		w.Write([]byte("wrong method"))
	})
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		w.Write([]byte("route does not exist"))
	})
	http.Handle("/", r)
	http.ListenAndServe(":8082", nil)
}

func middleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		fmt.Println(token)
		errMessage := verifyToken(token)
		if !(errMessage.Code == 0 || errMessage.Message == "") {
			w.WriteHeader(errMessage.Code)
			errMessage, _ := json.MarshalIndent(errMessage, "", "  ")
			w.Write(errMessage)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func portfolioList(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var inputForPortfolios model.InputForPortfolios
	bodyData, err := io.ReadAll(r.Body)
	if err != nil {
		response.MessageShow(500, "Internal server error", w)
		return
	}
	err = json.Unmarshal(bodyData, &inputForPortfolios)
	if err != nil {
		response.MessageShow(500, "Internal server error", w)
		return
	}
	expertPortfolioMap := make(map[string]model.PortfoliosOfExpert)
	for _, expertID := range inputForPortfolios.ExpertIDs {
		var portfoliosOfExpert model.PortfoliosOfExpert
		expertPortfolioMap[expertID] = portfoliosOfExpert
	}
	db := dal.GetDB()
	query := fmt.Sprintf(`WITH RankedPortfolios AS (
		SELECT id, name, created_by, image, created_at, ROW_NUMBER() OVER( partition  by created_by order by created_at desc) row_num
		FROM public.portfolio where created_by in ('%v') order by created_at desc )
		select id, name, created_by, image, created_at from RankedPortfolios where row_num <= %d;`, strings.Join(inputForPortfolios.ExpertIDs, "' , '"), inputForPortfolios.IncludeNumberOfPortfolios)
	query = sqlx.Rebind(sqlx.DOLLAR, query)
	rows, err := db.Query(query)
	if err != nil {
		databaseErrorMessage, databaseErrorCode := response.DatabaseErrorShow(err)
		response.MessageShow(databaseErrorCode, databaseErrorMessage, w)
		return
	}
	for rows.Next() {
		var portfolio model.Portfolio
		var expertID string
		err = rows.Scan(&portfolio.ID, &portfolio.Name, &expertID, &portfolio.Image, &portfolio.CreatedAt)
		if err != nil {
			databaseErrorMessage, databaseErrorCode := response.DatabaseErrorShow(err)
			response.MessageShow(databaseErrorCode, databaseErrorMessage, w)
			return
		}

		if expertPortfoliosMapValue, ok := expertPortfolioMap[expertID]; ok {
			expertPortfoliosMapValue.Portfolios = append(expertPortfoliosMapValue.Portfolios, portfolio)
			expertPortfolioMap[expertID] = expertPortfoliosMapValue
		}
		// else {
		// 	var portfoliosOfExpert model.PortfoliosOfExpert
		// 	portfoliosOfExpert.Portfolios = append(portfoliosOfExpert.Portfolios, portfolio)
		// 	expertPortfolioMap[expertID] = portfoliosOfExpert
		// }
	}
	query = fmt.Sprintf(`SELECT created_by, count(id) as total_portfolios FROM public.portfolio where created_by in ('%v')  group by created_by ;`, strings.Join(inputForPortfolios.ExpertIDs, "' , '"))
	query = sqlx.Rebind(sqlx.DOLLAR, query)
	rows, err = db.Query(query)
	if err != nil {
		databaseErrorMessage, databaseErrorCode := response.DatabaseErrorShow(err)
		response.MessageShow(databaseErrorCode, databaseErrorMessage, w)
		return
	}
	for rows.Next() {
		var totalPortfolios int
		var expertID string
		err = rows.Scan(&expertID, &totalPortfolios)
		if err != nil {
			databaseErrorMessage, databaseErrorCode := response.DatabaseErrorShow(err)
			response.MessageShow(databaseErrorCode, databaseErrorMessage, w)
			return
		}

		if expertPortfoliosMapValue, ok := expertPortfolioMap[expertID]; ok {
			expertPortfoliosMapValue.TotalPortfolios = int16(totalPortfolios)
			expertPortfolioMap[expertID] = expertPortfoliosMapValue
		}
		//  else {
		// 	var portfoliosOfExpert model.PortfoliosOfExpert
		// 	portfoliosOfExpert.TotalPortfolios = int16(totalPortfolios)
		// 	expertPortfolioMap[expertID] = portfoliosOfExpert
		// }
	}
	result, _ := json.MarshalIndent(expertPortfolioMap, "", "  ")
	w.Write(result)
}

func verifyToken(token string) model.Message {
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
