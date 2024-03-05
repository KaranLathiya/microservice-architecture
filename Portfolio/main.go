package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"portfolio/dal"
	"portfolio/model"
	"portfolio/response"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/jmoiron/sqlx"
)

var jwtKey = []byte("secret_key")

func main() {
	db, _ := dal.Connect()
	defer db.Close()
	fmt.Println("Server started")
	// run()
	http.HandleFunc("/user/portfolio", portfolioList)
	http.ListenAndServe(":8082", nil)
}

func portfolioList(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	token := r.Header.Get("Authorization")
	fmt.Println(token)
	var claims model.Claims
	claims, errMessage := verifyToken(token)
	if !(errMessage.Code == 0 || errMessage.Message == "") {
		w.WriteHeader(errMessage.Code)
		errMessage, _ := json.MarshalIndent(errMessage, "", "  ")
		w.Write(errMessage)
		return
	}

	expertPortfolioMap := make(map[string]model.PortfoliosOfExpert)
	for _, expertID := range claims.ExpertID {
		var portfoliosOfExpert model.PortfoliosOfExpert
		expertPortfolioMap[expertID] = portfoliosOfExpert
	}
	db := dal.GetDB()
	query := fmt.Sprintf(`WITH RankedPortfolios AS (
		SELECT id, name, created_by, image, created_at, ROW_NUMBER() OVER( partition  by created_by order by created_at desc) row_num
		FROM public.portfolio where created_by in ('%v') order by created_at desc )
		select id, name, created_by, image, created_at from RankedPortfolios where row_num <= %d;`, strings.Join(claims.ExpertID, "' , '"), claims.IncludeNumberOfPortfolios)
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
	query = fmt.Sprintf(`SELECT created_by, count(id) as total_portfolios FROM public.portfolio where created_by in ('%v')  group by created_by ;`, strings.Join(claims.ExpertID, "' , '"))
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

func verifyToken(token string) (model.Claims, model.Message) {
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
			return *claims, message
		}
		message.Code = http.StatusBadRequest
		message.Message = err.Error()
		return *claims, message
	}

	if !tkn.Valid {
		message.Code = http.StatusUnauthorized
		message.Message = "token is invalid"
		return *claims, message
	}
	return *claims, message
}

// func run() {
// 	fmt.Println("Sda")
// 	db := dal.GetDB()
// 	i := 86
// 	for i < 88 {
// 		val := "948489062412845057"
// 		name := "portfolio" + strconv.Itoa(i)
// 		image := "https://jwt.io/img/pic_logo.svg"
// 		var id string
// 		err := db.QueryRow("INSERT INTO public.portfolios ( name, created_by, image, created_at) VALUES( $1, $2, $3, $4) returning id;", name, val, image, currentTimeConvertToCurrentFormattedTime()).Scan(&id)
// 		if err != nil {
// 			fmt.Println(err)
// 			break
// 		}
// 		time.Sleep(1 * time.Second)
// 		i += 1
// 	}
// fmt.Println("done")
// }

// func currentTimeConvertToCurrentFormattedTime() string {
// 	// fmt.Println(time.Now().UTC())
// 	// fmt.Println(time.Now().Local().UTC())
// 	currentTime := time.Now().UTC()
// 	outputFormat := "2006-01-02 15:04:05-07:00"
// 	currentFormattedTime := currentTime.Format(outputFormat)
// 	// fmt.Println(currentFormattedTime)
// 	return currentFormattedTime
// }
